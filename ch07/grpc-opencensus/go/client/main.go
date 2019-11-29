// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc -I proto proto/product_info.proto --go_out=plugins=grpc:go/product_info
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/golang/product_info
// Execute go run go/client/main.go

package main

import (
	"context"
	"log"
	"time"

	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/grpc-up-and-running/samples/ch07/grpc-prometheus/go/proto"
	"google.golang.org/grpc"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/examples/exporter"
)

const (
	address = "localhost:50051"
)

func main() {
	// Register stats and trace exporters to export
	// the collected data.
	view.RegisterExporter(&exporter.PrintExporter{})

    // Register the view to collect gRPC client stats.
    if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
        log.Fatal(err)
    }

	// Set up a connection to the server.
	conn, err := grpc.Dial(address,
			grpc.WithStatsHandler(&ocgrpc.ClientHandler{}),
	        grpc.WithInsecure(),
	        )
	if err != nil {
		log.Fatalf("Can't connect: %v", err)
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
        log.Printf("Product: ", product.String())
        time.Sleep(3 * time.Second)
    }
}
