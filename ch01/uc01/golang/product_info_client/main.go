// Go to ${Advanced-gRPC}/samples/ch01/uc01
// Optional: Execute protoc --go_out=plugins=grpc:golang/product_info product_info.proto
// Execute go get -v github.com/advanced-grpc/samples/ch01/uc01/golang/product_info
// Execute go run golang/product_info_client/main.go

package main

import (
	"context"
	"log"
	"time"

	pb "github.com/advanced-grpc/samples/ch01/uc01/golang/product_info"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

	// Contact the server and print out its response.
	name := "Sumsung S10"
	description := "Samsung Galaxy S10 is the latest smartphone, launched in February 2019"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: ", product.String())
}
