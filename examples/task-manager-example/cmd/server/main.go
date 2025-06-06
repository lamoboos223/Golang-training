package main

import (
	"net"
	"net/http"
	"task-manager-example/internal/api"
	"task-manager-example/internal/api/grpcserver"
	"task-manager-example/internal/api/handlers"
	"task-manager-example/internal/api/proto"
	"task-manager-example/internal/config"
	"task-manager-example/internal/factory"
	"task-manager-example/internal/service"
	"task-manager-example/pkg/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Initialize logger
	log := logger.New()

	// Load configuration
	cfg := config.NewConfig()

	// Initialize dependencies
	repo, err := factory.NewRepository(cfg)
	if err != nil {
		log.Fatal("Failed to create repository",
			zap.String("error", err.Error()),
			zap.String("repository_type", cfg.RepositoryType))
	}

	taskService := service.NewTaskService(repo)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Setup HTTP router
	router := api.SetupRouter(taskHandler)

	// Setup gRPC server
	grpcServer := grpc.NewServer()
	grpcTaskServer := grpcserver.NewServer(taskService)
	proto.RegisterTaskServiceServer(grpcServer, grpcTaskServer)
	reflection.Register(grpcServer)

	// Start HTTP server
	go func() {
		log.Info("HTTP server starting",
			zap.String("port", "8080"),
			zap.String("repository_type", cfg.RepositoryType))

		if err := http.ListenAndServe(":8080", router); err != nil {
			log.Fatal("HTTP server failed to start",
				zap.String("error", err.Error()))
		}
	}()

	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", ":9090")
		if err != nil {
			log.Fatal("Failed to listen for gRPC",
				zap.String("error", err.Error()))
		}

		log.Info("gRPC server starting",
			zap.String("port", "9090"))

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("gRPC server failed to start",
				zap.String("error", err.Error()))
		}
	}()

	// Wait forever because we are using go routines to start the http and grpc servers
	select {}
}
