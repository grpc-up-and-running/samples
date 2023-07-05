module github.com/grpc-up-and-running/samples/ch06/token-based-authentication/go/server

go 1.15

require (
	github.com/google/uuid v1.3.0
	google.golang.org/grpc v1.53.0
	productinfo/server v0.0.0-20200901064603-1f9de1e3efd9
)

replace productinfo/server => github.com/grpc-up-and-running/samples/ch02/productinfo/go/server v0.0.0-20200901064603-1f9de1e3efd9
