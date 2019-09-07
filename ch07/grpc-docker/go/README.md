
Creating Docker Network 

``` 
    docker network create my-net
```


Building Server 

``` 
    docker image build -t grpc-productinfo-server -f server/Dockerfile .
    docker run -it --network=my-net --name=productinfo --hostname=productinfo -p 50051:50051  grpc-productinfo-server
```

Building Client 


``` 
    docker image build -t grpc-productinfo-client -f client/Dockerfile .
    docker run -it --network=my-net --hostname=client grpc-productinfo-client   
     
``` 

Tagging and Pushing to a Docker Registry 

``` 
    docker image tag grpc-productinfo-server kasunindrasiri/grpc-productinfo-server
    docker image tag grpc-productinfo-client kasunindrasiri/grpc-productinfo-client
    docker image push kasunindrasiri/grpc-productinfo-server
    docker image push kasunindrasiri/grpc-productinfo-client

```





    