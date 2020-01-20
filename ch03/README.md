# Chapter 3: gRPC Communication Patterns

## ``OrderManagement`` Service and Client 

- Online retail scenario has a `` OrderManagement`` microservice which is responsible for managing the orders and
 their information. The consumer of that service can add, retrieve, search and update order via that service. 

- ``OrderManagement`` service and the consumer of that service are implemented in both ``Go`` and ``Java`` languages.

- This sample shows how you can implement both ``OrderManagement`` service and its consumer.

## Service Definition 

```proto
package ecommerce;

service OrderManagement {
    rpc addOrder(Order) returns (google.protobuf.StringValue);
    rpc getOrder(google.protobuf.StringValue) returns (Order);
    rpc searchOrders(google.protobuf.StringValue) returns (stream Order);
    rpc updateOrders(stream Order) returns (google.protobuf.StringValue);
    rpc processOrders(stream google.protobuf.StringValue) returns (stream CombinedShipment);
}

message Order {
    string id = 1;
    repeated string items = 2;
    string description = 3;
    float price = 4;
    string destination = 5;
}

message CombinedShipment {
    string id = 1;
    string status = 2;
    repeated Order ordersList = 3;
}
```

## Implementation

- [Go](./order-service/go/README.md)
- [Java](./order-service/java/README.md)



