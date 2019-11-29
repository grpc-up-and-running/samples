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
    "go.opencensus.io/trace"
    "contrib.go.opencensus.io/exporter/jaeger"
)

const (
	address = "localhost:50051"
)

func main() {
    initTracing()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address,
	        grpc.WithInsecure(),
	        )
	if err != nil {
		log.Fatalf("Can't connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewProductInfoClient(conn)
	for {
	    ctx, span := trace.StartSpan(context.Background(), "ecommerce.ProductInfoClient")
        // Contact the server and print out its response.
        name := "Sumsung S10"
        description := "Samsung Galaxy S10 is the latest smart phone, launched in February 2019"
        price := float32(700.0)
        r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
        if err != nil {
            span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
            log.Fatalf("Could not add product: %v", err)
        }
        log.Printf("Product ID: %s added successfully", r.Value)

        product, err := c.GetProduct(ctx, &wrapper.StringValue{Value: r.Value})
        if err != nil {
            span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
            log.Fatalf("Could not get product: %v", err)
        }
        log.Printf("Product: ", product.String())
        span.End()
        time.Sleep(3 * time.Second)

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
             ProjectID:    "product_info",

     })
     if err != nil {
        log.Fatal(err)
     }
     trace.RegisterExporter(exporter)

}
