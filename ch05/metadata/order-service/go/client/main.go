package main

import (
	"context"
	"fmt"
	pb "github.com/grpc-up-and-running/samples/ch05/inteceptors/order-service/go/order-service-gen"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	// Setting up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrderManagementClient(conn)

	// originalCtx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	// defer cancel()


	// ****** Metadata : Creation *****

	md := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))

	mdCtx := metadata.NewOutgoingContext(context.Background(), md)

	// RPC using the context with new metadata.
	var header, trailer metadata.MD


	// Add Order
	order1 := pb.Order{Id: "101", Items:[]string{"iPhone XS", "Mac Book Pro"}, Destination:"San Jose, CA", Price:2300.00}
	res, addErr := client.AddOrder(mdCtx, &order1, grpc.Header(&header), grpc.Trailer(&trailer))

	if addErr != nil {
		got := status.Code(addErr)
		log.Printf("Error Occured -> addOrder : , %v:", got)
	} else {
		log.Print("AddOrder Response -> ", res.Value)
	}

	// Reading the headers
	if t, ok := header["timestamp"]; ok {
		log.Printf("timestamp from header:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("timestamp expected but doesn't exist in header")
	}
	if l, ok := header["location"]; ok {
		log.Printf("location from header:\n")
		for i, e := range l {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("location expected but doesn't exist in header")
	}




	// Get Order
	//retrievedOrder , err := client.GetOrder(originalCtx, &wrapper.StringValue{Value: "106"})
	//log.Print("GetOrder Response -> : ", retrievedOrder)


	// Search Order
	searchStream, _ := client.SearchOrders(mdCtx, &wrapper.StringValue{Value: "Google"})
	for {
		searchOrder, err := searchStream.Recv()
		if err == io.EOF {
			log.Print("EOF")
			break
		}

		if err == nil {
			log.Print("Search Result : ", searchOrder)
		}
	}


	// Update Orders
	updOrder1 := pb.Order{Id: "102", Items:[]string{"Google Pixel 3A", "Google Pixel Book"}, Destination:"Mountain View, CA", Price:1100.00}
	updOrder2 := pb.Order{Id: "103", Items:[]string{"Apple Watch S4", "Mac Book Pro", "iPad Pro"}, Destination:"San Jose, CA", Price:2800.00}
	updOrder3 := pb.Order{Id: "104", Items:[]string{"Google Home Mini", "Google Nest Hub", "iPad Mini"}, Destination:"Mountain View, CA", Price:2200.00}

	updateStream, _ := client.UpdateOrders(mdCtx)


	_ = updateStream.Send(&updOrder1)
	_ = updateStream.Send(&updOrder2)
	_ = updateStream.Send(&updOrder3)


	updateRes, _ := updateStream.CloseAndRecv()
	log.Printf("Update Orders Res : ", updateRes)


	//// Process Order
	//streamProcOrder, _ := client.ProcessOrders(originalCtx)
	//_ = streamProcOrder.Send(&wrapper.StringValue{Value:"102"})
	//_ = streamProcOrder.Send(&wrapper.StringValue{Value:"103"})
	//_ = streamProcOrder.Send(&wrapper.StringValue{Value:"104"})
	//
	//channel := make(chan bool, 1)
	//go asncClientBidirectionalRPC(streamProcOrder, channel)
	//time.Sleep(time.Millisecond * 1000)
	//
	//
	//_ = streamProcOrder.Send(&wrapper.StringValue{Value:"101"})
	//_ = streamProcOrder.CloseSend()
	//
	//<- channel
}






































