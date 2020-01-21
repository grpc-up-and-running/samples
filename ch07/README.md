# Chapter 7: Running gRPC in Production

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

- gRPC Continuous Integration  [[Go]](grpc-continous-integration/go/README.md)
- Deploy in Docker [[Go]](grpc-docker/go/README.md)
- Deploy in Kubernetes [[Go]](grpc-kubernetes/README.md)
- OpenCensus Metrics [[Go]](grpc-opencensus/go/README.md) [[Java]](grpc-opencensus/java/README.md)
- OpenCensus Tracing [[Go]](grpc-opencensus-tracing/go/README.md) [[Java]](grpc-opencensus-tracing/java/README.md)
- OpenTracing [[Go]](grpc-opentracing/go/README.md)
- Prometheus [[Go]](grpc-prometheus/go/README.md)