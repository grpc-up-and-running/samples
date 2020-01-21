# Chapter 8: The gRPC Ecosystem

## ``ProductInfo`` Service and Client 

- Online retail scenario has a `` ProductInfo`` microservice which is responsible for managing the products and their
 information. The consumer of that service can add, retrieve products via that service. 

- ``ProductInfo`` service and the consumer of that service are implemented in both ``Go`` and ``Java`` languages.

- This use case shows how you can implement both ``ProductInfo`` service and its consumer.

## Service Definition 

```proto
package ecommerce;

service ProductInfo {
    rpc addProduct(Product) returns (google.protobuf.StringValue);
    rpc getProduct(google.protobuf.StringValue) returns (Product);
}

message Product {
    string id = 1;
    string name = 2;
    string description = 3;
    float price = 4;
}
```

## Implementation

- gRPC Gateway [[Go]](./grpc-gateway/go/README.md)
- Server Reflection [[Go]](./server-reflection/go/README.md) [[Java]](./server-reflection/java/README.md)
