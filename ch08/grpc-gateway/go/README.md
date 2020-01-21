## ``ProductInfo`` Service and Client - Go Implementation

## Building and Running Service

In order to build, Go to ``Go`` module root directory location (grpc-gateway/go/server) and execute the following
 shell command,
```
go build -i -v -o bin/server
```

In order to run, Go to ``Go`` module root directory location (grpc-gateway/go/server) and execute the following
shell command,

```
./bin/server
```

## Building and Running Client   

In order to build, Go to ``Go`` module root directory location (grpc-gateway/go/client) and execute the following
 shell command,
```
go build -i -v -o bin/client
```

In order to run, Go to ``Go`` module root directory location (grpc-gateway/go/client) and execute the following
shell command,

```
./bin/client
```

## Testing

* Add a new product to the ProductInfo service.

```
$ curl -X POST http://localhost:8081/v1/product -d '{"name": "Apple", "description": "iphone7", "price": 699}'

"38e13578-d91e-11e9-819f-6c96cfe0687d"
```

* Get the existing product using ProductID

```
$ curl http://localhost:8081/v1/product/38e13578-d91e-11e9-819f-6c96cfe0687d

{"id":"38e13578-d91e-11e9-819f-6c96cfe0687d","name":"Apple","description":"iphone7","price":
```

## Additional Information

### Generate Server and Client side code 
``` 
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. \
product_info.proto
```

### Update after changing the service definition
``` 
go get -u github.com/grpc-up-and-running/samples/ch08/grpc-gateway/go/pb
```

### Generate reverse proxy service code
```
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--grpc-gateway_out=logtostderr=true:. \
product_info.proto
```

### Update after changing the reverse proxy service definition
``` 
go get -u github.com/grpc-up-and-running/samples/ch08/grpc-gateway/go/gw
```

### Generate the swagger file correspond to reverse proxy service
```
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--swagger_out=logtostderr=true:. \
product_info.proto
```

