package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	ecpb "google.golang.org/grpc/examples/features/proto/echo"
	"google.golang.org/grpc/status"
)

var (
	addrs = []string{":50051", ":50052"}
)

type ecServer struct {
	addr string
}

func (s *ecServer) UnaryEcho(ctx context.Context, req *ecpb.EchoRequest) (*ecpb.EchoResponse, error) {
	return &ecpb.EchoResponse{Message: fmt.Sprintf("%s (from %s)", req.Message, s.addr)}, nil
}
func (s *ecServer) ServerStreamingEcho(*ecpb.EchoRequest, ecpb.Echo_ServerStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}
func (s *ecServer) ClientStreamingEcho(ecpb.Echo_ClientStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}
func (s *ecServer) BidirectionalStreamingEcho(ecpb.Echo_BidirectionalStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}

func startServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ecpb.RegisterEchoServer(s, &ecServer{addr: addr})
	log.Printf("serving on %s\n", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	var wg sync.WaitGroup
	for _, addr := range addrs {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			startServer(addr)
		}(addr)
	}
	wg.Wait()
}