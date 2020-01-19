# Chapter 2: Developing Product Info Service and Client 

- Online retail scenario has a `` ProductInfo`` micro service which is responsible for managing the products and their
 information. The consumer of that service can add, remove, retrieve products via that service. 

- ``ProductInfo`` service and the consumer of that service are implemented on both ``Go`` and ``Java`` languages.
- This use case shows how you can implement both ``ProductInfo`` service and its consumer.

-------------

# Chapter 2: Developing Product Info Service and Client 
(Introduction to the use case) 

## Service Definition and Code Generation 


## Running Service

### Go Implementation

In order to build, Go to ``Go`` module root directory location (productinfo/go/server) and execute the following
 shell command,
```
go build -i -v -o bin/server
```

In order to run, Go to ``Go`` module root directory location (productinfo/go/server) and execute the following
shell command,

```
./bin/server
```

### Java Implementation

In order to build gradle project, Go to ``Java`` project root directory location (productinfo/java/server) and execute
 the following shell command,
```
gradle build
```

In order to run, Go to ``Java`` project root directory location (productinfo/java/server) and execute the following
shell command,

```
java -jar build/libs/server.jar
```

## Running Client  

### Go Implementation 

In order to build, Go to ``Go`` module root directory location (productinfo/go/client) and execute the following
 shell command,
```
go build -i -v -o bin/client
```

In order to run, Go to ``Go`` module root directory location (productinfo/go/client) and execute the following
shell command,

```
./bin/client
```

### Java Implementation 

In order to build gradle project, Go to ``Java`` project root directory location (productinfo/java/client) and execute
 the following shell command,
```
gradle build
```

In order to run, Go to ``Java`` project root directory location (productinfo/java/client) and execute the following
shell command,

```
java -jar build/libs/client.jar
```
