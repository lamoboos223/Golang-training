package main

import (
	"net/http"
	"task-manager-example/internal/api"
	"task-manager-example/internal/api/handlers"
	"task-manager-example/internal/config"
	"task-manager-example/internal/factory"
	"task-manager-example/internal/service"
	"task-manager-example/pkg/logger"

	"go.uber.org/zap"
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

	// Setup router
	router := api.SetupRouter(taskHandler)

	// Start server
	log.Info("Server starting",
		zap.String("port", "8080"),
		zap.String("repository_type", cfg.RepositoryType))

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server failed to start",
			zap.String("error", err.Error()))
	}
}
