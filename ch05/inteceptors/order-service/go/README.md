## GO Impl 



Generate Server and Client side code 

``` 
protoc -I order_service ./order_service/order_management.proto --go_out=plugins=grpc:order_service order_service/order_service.proto
``` 

Update after changing the service definition
``` 
 go get -u github.com/grpc-up-and-running/samples/ch03/order-service/go/order_service 
``` 