## ``ProductInfo`` Service and Client - Java Implementation

### Building and Running Service

In order to build gradle project, Go to ``Java`` project root directory location (samples) and execute
 the following shell command,
```
./gradlew :ch08:server-reflection:java:server:build
```

In order to run, Go to ``Java`` project root directory location (server-reflection/java/server) and execute the following
shell command,

```
java -jar build/libs/server.jar
```

## Testing

* List Services

```
$ ./grpc_cli ls localhost:50051

Output:
ecommerce.ProductInfo
grpc.reflection.v1alpha.ServerReflection
```

* List service details

```
$ ./grpc_cli ls localhost:50051 ecommerce.ProductInfo -l

Output:
package: ecommerce;
service ProductInfo {
rpc addProduct(ecommerce.Product) returns (google.protobuf.StringValue) {}
rpc getProduct(google.protobuf.StringValue) returns (ecommerce.Product) {}
}
```

* List method details

```
$ ./grpc_cli ls localhost:50051 ecommerce.ProductInfo.addProduct -l

Output:
rpc addProduct(ecommerce.Product) returns (google.protobuf.StringValue) {}
```

* Inspect message types
```
$ ./grpc_cli type localhost:50051 ecommerce.Product

Output:
message Product {
string id = 1[json_name = "id"];
string name = 2[json_name = "name"];
string description = 3[json_name = "description"];
float price = 4[json_name = "price"];
}
```

* Call remote methods
```
$ ./grpc_cli call localhost:50051 addProduct "name: 'Apple', description: 'iphone 11', price: 699"

Output:
connecting to localhost:50051
value: "d962db94-d907-11e9-b49b-6c96cfe0687d"
Rpc succeeded with OK status
```
