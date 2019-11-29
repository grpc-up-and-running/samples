// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc --go_out=plugins=grpc:golang/product_info product_info.proto
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/go/product_info
// Execute go run go/server/main.go

package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	pb "github.com/grpc-up-and-running/samples/ch02/productinfo/go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"
)

// server is used to implement ecommerce/product_info.
type server struct {
	productMap map[string]*pb.Product
}

// AddProduct implements ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*wrapper.StringValue, error) {
	out, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &wrapper.StringValue{Value: in.Id}, nil
}

// GetProduct implements ecommerce.GetProduct
func (s *server) GetProduct(ctx context.Context, in *wrapper.StringValue) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, nil
	}
	return nil, errors.New("Product does not exist for the ID" + in.Value)
}

var (
	port = ":50051"
    crtFile = filepath.Join("ch06", "mutual-tls-channel", "certs", "server.crt")
    keyFile = filepath.Join("ch06", "mutual-tls-channel", "certs", "server.key")
    caFile = filepath.Join("ch06", "mutual-tls-channel", "certs", "ca.crt")
)

func main() {
	certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}

	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append client certs")
	}

	opts := []grpc.ServerOption{
		// Enable TLS for all incoming connections.
		grpc.Creds(    // Create the TLS credentials
			credentials.NewTLS(&tls.Config {
				ClientAuth:   tls.RequireAndVerifyClientCert,
				Certificates: []tls.Certificate{certificate},
				ClientCAs:    certPool,
				},
			)),
	}

	s := grpc.NewServer(opts...)
	pb.RegisterProductInfoServer(s, &server{})
	// Register reflection service on gRPC server.
	//reflection.Register(s)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
