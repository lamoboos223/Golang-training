package main

import (
	"context"
	"log"
	"time"

	pb "grpc-hello-world/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Make 5 requests to demonstrate performance
	for i := 0; i < 5; i++ {
		start := time.Now()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		latency := time.Since(start)
		log.Printf("Request %d - Response: %s, Latency: %v", i+1, r.GetMessage(), latency)

		cancel()
		time.Sleep(100 * time.Millisecond) // Small delay between requests
	}
}
