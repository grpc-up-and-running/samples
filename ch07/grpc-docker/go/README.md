## ``ProductInfo`` Service and Client - Go Implementation

### Building and Running Service

In order to build, Go to ``Go`` module root directory location (grpc-docker/go/server) and execute the following
 shell command,
```
go build -i -v -o bin/server
```

In order to run, Go to ``Go`` module root directory location (grpc-docker/go/server) and execute the following
shell command,

```
./bin/server
```

### Building and Running Client   

In order to build, Go to ``Go`` module root directory location (grpc-docker/go/client) and execute the following
 shell command,
```
go build -i -v -o bin/client
```

In order to run, Go to ``Go`` module root directory location (grpc-docker/go/client) and execute the following
shell command,

```
./bin/client
```

### Creating Docker Network 

``` 
    docker network create my-net
```


### Building Server 

``` 
    docker image build -t grpc-productinfo-server -f server/Dockerfile .
    docker run -it --network=my-net --name=productinfo --hostname=productinfo -p 50051:50051  grpc-productinfo-server
```

### Building Client 


``` 
    docker image build -t grpc-productinfo-client -f client/Dockerfile .
    docker run -it --network=my-net --hostname=client grpc-productinfo-client   
     
``` 

### Tagging and Pushing to a Docker Registry 

``` 
    docker image tag grpc-productinfo-server kasunindrasiri/grpc-productinfo-server
    docker image tag grpc-productinfo-client kasunindrasiri/grpc-productinfo-client
    docker image push kasunindrasiri/grpc-productinfo-server
    docker image push kasunindrasiri/grpc-productinfo-client

```





    