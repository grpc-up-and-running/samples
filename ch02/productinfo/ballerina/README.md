## ``ProductInfo`` Service and Client - Ballerina Implementation

## Building and Running Service

In order to build, Go to ``ballerina`` module root directory location (productinfo/ballerina/server) and execute the following
 shell command,
```
bal build
```

In order to run, Go to ``ballerina`` module root directory location (productinfo/ballerina/server) and execute the following
shell command,

```
java -jar target/bin/server.jar
```

## Building and Running Client   

In order to build, Go to ``ballerina`` module root directory location (productinfo/ballerina/client) and execute the following
 shell command,
```
bal build
```

In order to run, Go to ``ballerina`` module root directory location (productinfo/ballerina/client) and execute the following
shell command,

```
java -jar target/bin/client.jar
```

## Additional Information

### Generate Server and Client side code 
Pre-generated stub file is included in the ballerina project. If you need to generate the stub files please use the below
 command from the root directory(inside productinfo directory)
``` 
bal grpc --input proto/product_info.proto --output ballerina/server
``` 