package main

import (
	"context"
	"io"
	"log"
	"net"
	"strings"

	"github.com/golang/protobuf/ptypes/wrappers"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/grpc-up-and-running/samples/ch03/order-service/go/order_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

var orderMap = make(map[string]pb.Order)

type server struct {
	orderMap map[string]*pb.Order
}

func (s *server) AddOrder(ctx context.Context, orderReq *pb.Order) (*wrappers.StringValue, error) {
	orderMap[orderReq.Id] = *orderReq
	return &wrapper.StringValue{Value: "Order Added: " + orderReq.Id}, nil
}

func (s *server) GetOrder(ctx context.Context, orderId *wrapper.StringValue) (*pb.Order, error) {
	ord := orderMap[orderId.Value]
	return &ord, nil
}

func (s *server) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {

	for key, order := range orderMap {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(itemStr, searchQuery.Value) {
				// Send the matching orders in a stream
				stream.Send(&order)
				log.Print("Matching Order Found : " + key)
				break
			}
		}
	}

	return nil
}

func (s *server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {

	ordersStr := "Updated Order IDs : "
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			// Finished reading order stream
			return stream.SendAndClose(&wrapper.StringValue{Value: "Orders processed " + ordersStr})
		}
		// Update order
		orderMap[order.Id] = *order

		log.Printf("Order ID ", order.Id, ": Updated")
		ordersStr += order.Id + ", "
		// ...
	}
}

func (s *server) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {
	order, _ := stream.Recv()
	orderList := []string{"100500", "100501"}
	comb := pb.CombinedShipment{Id: "123", OrderIDList: orderList, Status: "OK"}

	log.Printf("Order ID ", order.Id, " - Processed!")

	stream.Send(&comb)
	return nil
	//return status.Errorf(codes.Unimplemented, "method ProcessOrders not implemented")
}

func main() {
	initSampleData()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &server{})
	///	pb.RegisterOrderInfoServer(s, &server{})
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
