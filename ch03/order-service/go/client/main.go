package main

import (
	"context"
	"log"
	"time"

	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/grpc-up-and-running/samples/ch03/order-service/go/order_service"
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
	c := pb.NewOrderManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.GetOrderStatus(ctx, &wrapper.StringValue{Value: "Foo"})
	log.Printf("Order Status : ", res.Name)

	stream, _ := c.SearchOrders(ctx, &wrapper.StringValue{Value: "Foo"})
	order2, err2 := stream.Recv()
	if err2 != nil {
		log.Fatalf("Error ")
	}
	log.Printf("Search Order : ", order2.Description)

	order3, _ := stream.Recv()
	log.Printf("Search Order : ", order3.Description)

	// Update Orders
	// Send stream or 'Orders' to the service.
	streamU, _ := c.UpdateOrders(ctx)

	streamU.Send(order2)
	streamU.Send(order3)

	updateRes, _ := streamU.CloseAndRecv()
	log.Printf("Update response : ", updateRes)

	// Process Order
	order4 := pb.Order{Id: "100600", Name: "Order for prc"}
	streamProc, _ := c.ProcessOrders(ctx)
	streamProc.Send(&order4)

	combinedShipment, _ := streamProc.Recv()

	log.Printf("Combined shipment ", combinedShipment.OrderIDList)

}
