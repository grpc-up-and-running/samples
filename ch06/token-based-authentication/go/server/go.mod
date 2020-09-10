module github.com/grpc-up-and-running/samples/ch06/token-based-authentication/go/server

go 1.15

require (
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.2
	google.golang.org/grpc v1.32.0
	productinfo/server v0.0.0-20200901064603-1f9de1e3efd9
)

replace productinfo/server => github.com/grpc-up-and-running/samples/ch02/productinfo/go/server v0.0.0-20200901064603-1f9de1e3efd9
