## GO Impl 



Generate Server and Client side code 

``` 
protoc -I order_service ./order_service/order_service.proto --go_out=plugins=grpc:order_service order_service/order_service.proto
``` 