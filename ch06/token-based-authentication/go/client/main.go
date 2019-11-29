// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc --go_out=plugins=grpc:golang/product_info product_info.proto
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/golang/product_info
// Execute go run go/client/main.go

package main

import (
	"context"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"log"
	"path/filepath"
	"time"

	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/grpc-up-and-running/samples/ch02/productinfo/go/product_info"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	hostname = "localhost"
	crtFile = filepath.Join("ch06", "secure-channel", "certs", "server.crt")
)

func main() {
	// Set up the credentials for the connection.
	perRPC := oauth.NewOauthAccess(fetchToken())

	creds, err := credentials.NewClientTLSFromFile(crtFile, hostname)
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(perRPC),
		// transport credentials.
		grpc.WithTransportCredentials(creds),
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

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
	log.Printf("Product: ", product.String())
}

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}
