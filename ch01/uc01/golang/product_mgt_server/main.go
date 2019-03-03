// Go to ${Advanced-gRPC}/samples/ch01/uc01
// Optional: Execute protoc --go_out=plugins=grpc:golang/product_mgt product_mgt.proto
// Execute go get -v github.com/advanced-grpc/samples/ch01/uc01/golang/product_mgt
// Excute go run golang/product_mgt_server/main.go

package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/advanced-grpc/samples/ch01/uc01/golang/product_mgt"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) AddProduct(ctx context.Context, in *pb.ProductRequest) (*pb.ProductResponse, error) {
	return &pb.ProductResponse{ProductID: "1234567"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductMgtServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
