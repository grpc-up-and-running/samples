// Go to ${Advanced-gRPC}/samples/ch01/uc01
// Optional: Execute protoc --go_out=plugins=grpc:golang/product_mgt product_mgt.proto
// Execute go get -v github.com/advanced-grpc/samples/ch01/uc01/golang/product_mgt
// Excute go run golang/product_mgt_client/main.go

package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "github.com/advanced-grpc/samples/ch01/uc01/golang/product_mgt"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductMgtClient(conn)

	// Contact the server and print out its response.
	name := "Sumsung S10"
	description := "Samsung Galaxy S10 is the latest smartphone, launched in February 2019"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddProduct(ctx, &pb.ProductRequest{Name: name, Description: description})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Product ID: %s", r.ProductID)
}
