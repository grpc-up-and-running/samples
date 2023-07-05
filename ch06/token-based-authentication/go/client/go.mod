module github.com/grpc-up-and-running/samples/ch06/token-based-authentication/go/client

go 1.15

require (
	golang.org/x/oauth2 v0.4.0
	google.golang.org/grpc v1.53.0
	productinfo/client v0.0.0-20200901064603-1f9de1e3efd9
)

replace productinfo/client => github.com/grpc-up-and-running/samples/ch02/productinfo/go/client v0.0.0-20200901064603-1f9de1e3efd9
