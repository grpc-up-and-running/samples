## ``ProductInfo`` Service and Client - Go Implementation

## Building and Running Service

In order to build, Go to ``Go`` module root directory location (secure-channel/go/server) and execute the following
 shell command,
```
go build -i -v -o bin/server
```

In order to run, Go to ``Go`` module root directory location (secure-channel/go/server) and execute the following
shell command,

```
./bin/server
```

## Building and Running Client   

In order to build, Go to ``Go`` module root directory location (secure-channel/go/client) and execute the following
 shell command,
```
go build -i -v -o bin/client
```

In order to run, Go to ``Go`` module root directory location (secure-channel/go/client) and execute the following
shell command,

```
./bin/client
```

## Additional Information

### Update after changing the service definition

```shell script 
go get -u github.com/grpc-up-and-running/samples/ch02/productinfo/go/product_info
```

### Generate Server key and certificate

* Generate Using [OpenSSL](../certs/README.md)