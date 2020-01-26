## ``ProductInfo`` Service and Client - Python Implementation

## Building and Running Service

In order to build and run, Go to ``Python`` module root directory location (productinfo/python/server) and execute the following shell command,
```
python server.py
```

## Building and Running Client   

In order to build and run, Go to ``Python`` module root directory location (productinfo/python/client) and execute the following shell command,
```
python client.py
```

## Additional Information

### Generate Server and Client side code 
``` 
pip install grpcio
pip install grpcio-tools

python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. product_info.proto
``` 
