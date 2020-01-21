# Chapter 6: Secured gRPC

## ``ProductInfo`` Service and Client 

- Online retail scenario has a `` ProductInfo`` microservice which is responsible for managing the products and their
 information. The consumer of that service can add, retrieve products via that service. 

- ``ProductInfo`` service and the consumer of that service are implemented in both ``Go`` and ``Java`` languages.

- This use case shows how you can implement both ``ProductInfo`` service and its consumer.

## Service Definition 

```proto
package ecommerce;

service ProductInfo {
    rpc addProduct(Product) returns (ProductID);
    rpc getProduct(ProductID) returns (Product);
}

message Product {
    string id = 1;
    string name = 2;
    string description = 3;
    float price = 4;
}

message ProductID {
    string value = 1;
}
```

## Implementation

- Secure Channel [[Go]](secure-channel/go/README.md) [[Java]](secure-channel/java/README.md)
- Mutual TLS Channel [[Go]](mutual-tls-channel/go/README.md) [[Java]](mutual-tls-channel/java/README.md)
- Basic Authentication [[Go]](basic-authentication/go/README.md) [[Java]](basic-authentication/java/README.md) 
- Token Based Authentication [[Go]](token-based-authentication/go/README.md) [[Java]](token-based-authentication/java/README.md)
