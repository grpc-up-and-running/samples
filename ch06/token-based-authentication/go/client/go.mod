module github.com/grpc-up-and-running/samples/ch06/token-based-authentication/go/client

go 1.15

require (
	github.com/golang/protobuf v1.4.2
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	google.golang.org/grpc v1.32.0
	productinfo/client v0.0.0-20200901064603-1f9de1e3efd9
)

replace productinfo/client => github.com/grpc-up-and-running/samples/ch02/productinfo/go/client v0.0.0-20200901064603-1f9de1e3efd9
