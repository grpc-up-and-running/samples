## GO Impl 



Generate Server and Client side code 

``` 
protoc -I order-service-gen/ order-service-gen/order_management.proto --go_out=plugins=grpc:order-service-gen

``` 

Update after changing the service definition
``` 
 go get -u github.com/grpc-up-and-running/samples/ch05/inteceptors/order-service/go/order-service-gen 
``` 