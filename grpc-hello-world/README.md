# gRPC Hello World Example

This is a simple gRPC example in Go that demonstrates a basic client-server communication with performance metrics.

## Project Structure

```
grpc-hello-world/
├── proto/
│   └── hello.proto      # Protocol buffer definition
├── server/
│   └── main.go         # gRPC server implementation
├── client/
│   └── main.go         # gRPC client implementation
├── go.mod              # Go module file
└── README.md           # This file
```

## Prerequisites

- Go 1.21 or later
- Protocol Buffers compiler (protoc)
- Go plugins for protoc:
  - protoc-gen-go
  - protoc-gen-go-grpc

## Installation

1. Install Protocol Buffers compiler (protoc):

   ```bash
   # Using Chocolatey (Windows)
   choco install protoc

   # Or download from https://github.com/protocolbuffers/protobuf/releases
   ```

2. Install Go plugins:

   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

3. Generate Go code from proto file:

   ```bash
   protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/hello.proto
   ```

## Running the Example

1. Start the server:

   ```bash
   go run server/main.go
   ```

2. In a new terminal, run the client:

   ```bash
   go run client/main.go
   ```

## Features

- Simple gRPC service with a SayHello method
- Performance metrics logging:
  - Request latency
  - Throughput (requests per second)
  - Request count
- Client makes multiple requests to demonstrate performance
- HTTP/2 connection reuse

## Testing with Postman

To test the gRPC service using Postman:

1. Open Postman
2. Create a new gRPC request
3. Set the server URL to: `localhost:50051`
4. Select the `SayHello` method
5. In the request body, provide:

   ```json
   {
     "name": "World"
   }
   ```

## Performance Metrics

The server logs include:

- Request number
- Latency per request
- Current throughput
- Received name parameter

The client logs include:

- Request number
- Response message
- Total round-trip latency

## Dependencies

- google.golang.org/grpc v1.72.2
- google.golang.org/protobuf v1.36.5

## Notes

- The first request will be slower due to connection establishment
- Subsequent requests are faster due to HTTP/2 connection reuse
- The server runs on port 50051 by default
- The example includes proper error handling and connection management

## Testing the gRPC Server

Since gRPC uses HTTP/2 and Protocol Buffers, you cannot use regular curl. Here are the recommended ways to test the server:

### 1. Using grpcurl (CLI tool)

```bash
# Install grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# List available services
grpcurl -plaintext localhost:50051 list

# Make a request
grpcurl -plaintext -d '{"name": "World"}' localhost:50051 hello.Greeter/SayHello
```

### 2. Using Postman

1. Open Postman
2. Create a new gRPC request
3. Set the server URL to: `localhost:50051`
4. Select the `SayHello` method
5. In the request body, provide:

   ```json
   {
     "name": "World"
   }
   ```

### 3. Using Evans (CLI tool)

```bash
# Install Evans
go install github.com/ktr0731/evans@latest

# Run Evans
evans -r repl -p 50051

# Inside Evans:
package hello
service Greeter
call SayHello
```

### 4. Using BloomRPC (GUI tool)

1. Download BloomRPC from: <https://github.com/uw-labs/bloomrpc>
2. Import your proto file
3. Connect to `localhost:50051`
4. Select the `SayHello` method
5. Provide the request body and send
