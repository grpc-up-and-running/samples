# Chapter 2: Developing Product Info Service and Client 

- Online retail scenario has a `` ProductInfo`` micro service which is responsible for managing the products and their
 information. The consumer of that service can add, remove, retrieve products via that service. 

- ``ProductInfo`` service is implemented on `` Go``. 
- One of the consumer of that service is written using ``Java``. 
- This use case shows you can implement both ``ProductInfo`` service and its consumer.

## Go Server and Client

### Go Server
To build gRPC Go server with Bazel, use the following command from sample directory
``` 
$ bazel build //ch02/productinfo/go/server:server
```

To spin up the Go microservice, run the following command from root directory.
``` 
$ bazel-bin/ch02/productinfo/go/server/darwin_amd64_stripped/server
```

### Go Client
To build gRPC Go Client with Bazel, use the following command from sample directory
``` 
$ bazel build //ch02/productinfo/go/client:client
```

To run gRPC Go Client and test microservice, use the following command from root directory.
``` 
$ bazel-bin/ch02/productinfo/go/client/darwin_amd64_stripped/client
```

## Java Server and Client

### Java Server
To build gRPC Java server with Bazel, use the following command from sample directory
``` 
$ bazel build //ch02/productinfo/java:server
```

To spin up the Java microservice, run the following command from root directory.
``` 
$ bazel-bin/ch02/productinfo/java/server
```

### Java Client

To build gRPC Java Client with Bazel, use the following command from sample directory
``` 
$ bazel build //ch02/productinfo/java:client
```

To run gRPC Java Client and test microservice, use the following command from root directory.
``` 
$ bazel-bin/ch02/productinfo/java/client
```
