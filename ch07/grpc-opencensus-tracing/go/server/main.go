// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc -I proto proto/product_info.proto --go_out=plugins=grpc:go/product_info
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/go/product_info
// Execute go run go/server/main.go

package main

import (
	"context"
	"errors"
	"log"
	"net"

	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	pb "github.com/grpc-up-and-running/samples/ch07/grpc-prometheus/go/proto"
	"google.golang.org/grpc"
	"go.opencensus.io/trace"
	"contrib.go.opencensus.io/exporter/jaeger"
)

const (
	port = ":50051"
)

// server is used to implement ecommerce/product_info.
type server struct {
	productMap map[string]*pb.Product
}

// AddProduct implements ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*wrapper.StringValue, error) {
	ctx, span := trace.StartSpan(ctx, "ecommerce.AddProduct")
	defer span.End()
	out, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &wrapper.StringValue{Value: in.Id}, nil
}

// GetProduct implements ecommerce.GetProduct
func (s *server) GetProduct(ctx context.Context, in *wrapper.StringValue) (*pb.Product, error) {
	ctx, span := trace.StartSpan(ctx, "ecommerce.GetProduct")
	defer span.End()
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}

func main() {

	// Create a gRPC Server with stats handler.
	grpcServer := grpc.NewServer()
	pb.RegisterProductInfoServer(grpcServer, &server{})

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

    // initialize opencensus jaeger exporter
	initTracing()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


func initTracing() {
    // This is a demo app with low QPS. trace.AlwaysSample() is used here
    // to make sure traces are available for observation and analysis.
    // In a production environment or high QPS setup please use
    // trace.ProbabilitySampler set at the desired probability.
    trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
    agentEndpointURI := "localhost:6831"
    collectorEndpointURI := "http://localhost:14268/api/traces"
     exporter, err := jaeger.NewExporter(jaeger.Options{
             CollectorEndpoint: collectorEndpointURI,
             AgentEndpoint: agentEndpointURI,
             ServiceName: "product_info",

     })
     if err != nil {
        log.Fatal(err)
     }
     trace.RegisterExporter(exporter)

}
