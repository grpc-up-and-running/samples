## ``OrderManagement`` Service and Client - Go Implementation

## Building and Running Service

In order to build, Go to ``Go`` module root directory location (cancellation/order-service/go/server) and execute the following
 shell command,
```
go build -i -v -o bin/server
```

In order to run, Go to ``Go`` module root directory location (cancellation/order-service/go/server) and execute the following
shell command,

```
./bin/server
```

## Building and Running Client   

In order to build, Go to ``Go`` module root directory location (cancellation/order-service/go/client) and execute the following
 shell command,
```
go build -i -v -o bin/client
```

In order to run, Go to ``Go`` module root directory location (cancellation/order-service/go/client) and execute the following
shell command,

```
./bin/client
```

## Additional Information

### Generate Server and Client side code 
``` 
protoc -I order-service-gen/ order-service-gen/order_management.proto --go_out=plugins=grpc:order-service-gen
``` 

### Update after changing the service definition
``` 
go get -u github.com/grpc-up-and-running/samples/ch05/cancellation/order-service/go/order-service-gen
```
