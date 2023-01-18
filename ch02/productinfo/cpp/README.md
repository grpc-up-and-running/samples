# server
```shell
cd server
protoc -I ../../proto --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` ../../proto/product_info.proto
protoc -I ../../proto --cpp_out=. ../../proto/product_info.proto
```