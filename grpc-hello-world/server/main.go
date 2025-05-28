package main

import (
	"context"
	"log"
	"net"
	"sync/atomic"
	"time"

	pb "grpc-hello-world/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
	requestCount uint64
	startTime    time.Time
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	start := time.Now()
	requestNum := atomic.AddUint64(&s.requestCount, 1)

	// Process the request
	response := &pb.HelloReply{Message: "Hello " + req.GetName()}

	// Calculate metrics
	latency := time.Since(start)
	totalTime := time.Since(s.startTime)
	throughput := float64(requestNum) / totalTime.Seconds()

	log.Printf("Request #%d - Latency: %v, Throughput: %.2f req/sec, Name: %q",
		requestNum, latency, throughput, req.GetName())

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := &server{
		startTime: time.Now(),
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, s)

	log.Printf("Server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
