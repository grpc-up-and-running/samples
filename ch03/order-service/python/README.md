## ``OrderManagement`` Service and Client - Python Implementation

## Prerequisites 
- Python 3.5 or higher
- pip version 9.0.1 or higher

If necessary, upgrade your version of pip:
```
python -m pip install --upgrade pip
```

Install gRPC and gRPC tools 

```
python -m pip install --upgrade pip
python -m pip install grpcio
python -m pip install grpcio-tools

```

## Code Generation 

Generate Python code by pointing to the .proto file. 

```
python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. order_management.proto

```


## Running the Service
Run the server with: 

```
python server.py
```

## Running the Client   

Run the client with: 
```
python server.py
```
