// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc -I proto proto/product_info.proto --go_out=plugins=grpc:go/product_info
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/golang/product_info
// Execute go run go/client/main.go

package main

import (
	"context"
	"github.com/daneshk/samples/ch07/grpc-opentracing/go/tracer"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	grpcopentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	pb "github.com/grpc-up-and-running/samples/ch07/grpc-opentracing/go/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {

	// Set up a connection to the server.
	conn, err := NewClientConn(address)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewProductInfoClient(conn)

	for {
		// Contact the server and print out its response.
		name := "Sumsung S10"
		description := "Samsung Galaxy S10 is the latest smart phone, launched in February 2019"
		price := float32(700.0)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
		if err != nil {
			log.Fatalf("Could not add product: %v", err)
		}
		log.Printf("Product ID: %s added successfully", r.Value)

		product, err := c.GetProduct(ctx, &wrapper.StringValue{Value: r.Value})
		if err != nil {
			log.Fatalf("Could not get product: %v", err)
		}
		log.Printf("Product: " + product.String())
		time.Sleep(3 * time.Second)
	}
}

func NewClientConn(address string) (*grpc.ClientConn, error) {
	// initialize jaegertracer
	jaegertracer, closer, err := tracer.NewTracer("product_mgt")
	if err != nil {
		return grpc.Dial(address, grpc.WithInsecure())
	}
	defer closer.Close()

	// initialize client with tracing interceptor using grpc client side chaining
	return grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithStreamInterceptor(
			grpcopentracing.StreamClientInterceptor(grpcopentracing.WithTracer(jaegertracer)),
		),
		grpc.WithUnaryInterceptor(
			grpcopentracing.UnaryClientInterceptor(grpcopentracing.WithTracer(jaegertracer)),
		),
	)
}
