package main

import (
	"context"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/grpc-up-and-running/samples/ch03/order-service/go/order_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
	orderBatchSize = 3
)

var orderMap = make(map[string]pb.Order)
var combinedShipmentMap = make(map[string]pb.CombinedShipment)

type server struct {
	orderMap map[string]*pb.Order
}

// Simple RPC
func (s *server) AddOrder(ctx context.Context, orderReq *pb.Order) (*wrappers.StringValue, error) {
	orderMap[orderReq.Id] = *orderReq
	return &wrapper.StringValue{Value: "Order Added: " + orderReq.Id}, nil
}

// Simple RPC
func (s *server) GetOrder(ctx context.Context, orderId *wrapper.StringValue) (*pb.Order, error) {
	ord := orderMap[orderId.Value]
	return &ord, nil
}

// Server-side Streaming RPC
func (s *server) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {

	for key, order := range orderMap {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(itemStr, searchQuery.Value) {
				// Send the matching orders in a stream
				stream.Send(&order)
				log.Print("Matching Order Found : " + key)
				time.Sleep(4000 * time.Millisecond)
				break
			}
		}
	}

	log.Print("EOF marked")

	return nil
}



// Client-side Streaming RPC
func (s *server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {

	ordersStr := "Updated Order IDs : "
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			// Finished reading the order stream.
			return stream.SendAndClose(&wrapper.StringValue{Value: "Orders processed " + ordersStr})
		}
		// Update order
		orderMap[order.Id] = *order

		log.Printf("Order ID ", order.Id, ": Updated")
		ordersStr += order.Id + ", "
	}
}


// Bi-directional Streaming RPC
func (s *server) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {

	for {
		orderId, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		i := 0
		for i <= orderBatchSize {
			shipment, found := combinedShipmentMap[orderMap[orderId.GetValue()].Destination]

			if found {
				shipment.OrderIDList[1]= orderMap[orderId]

			} else {
				orderList := make([]string, orderBatchSize)

				orderList[0] = orderMap[order.Id].Id
				comShip := pb.CombinedShipment{Id:"cmb" + order.Id, Status:"Processed!", OrderIDList:orderList}
				combinedShipmentMap[orderMap[order.Id].Destination] = comShip
			}
			i++
		}

		for _, comb  := range combinedShipmentMap {
			stream.Send(&comb)
		}




	}





	return nil
}

func main() {
	initSampleData()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initSampleData() {
	orderMap["102"] = pb.Order{Id: "102", Items:[]string{"Google Pixel 3A", "Mac Book Pro"}, Destination:"Mountain View, CA", Price:1800.00}
	orderMap["103"] = pb.Order{Id: "103", Items:[]string{"Apple Watch S4" }, Destination:"San Jose, CA", Price:400.00}
	orderMap["104"] = pb.Order{Id: "104", Items:[]string{"Google Home Mini", "Google Nest Hub" }, Destination:"Mountain View, CA", Price:400.00}
	orderMap["105"] = pb.Order{Id: "105", Items:[]string{"Amazon Echo"}, Destination:"San Jose, CA", Price:30.00}
	orderMap["106"] = pb.Order{Id: "106", Items:[]string{"Amazon Echo", "Apple iPhone XS"}, Destination:"Mountain View, CA", Price:30.00}
}
