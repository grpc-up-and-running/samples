package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/wrappers"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/grpc-up-and-running/samples/ch03/order-service/go/order_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

type server struct {
	orderMap map[string]*pb.Order
}

func (s *server) GetOrderStatus(ctx context.Context, in *wrapper.StringValue) (*wrapper.StringValue, error) {
	// value, exists := s.productMap[in.Value]

	value := "Order Processed!"
	// if exists {
	// 	return value, nil
	// }
	return &wrapper.StringValue{Value: value}, nil
	// return nil, errors.New("Product does not exist for the ID" + in.Value)
}

func (s *server) SearchOrder(ctx context.Context, req *wrappers.StringValue) (*pb.Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchOrder not implemented")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderInfoServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
