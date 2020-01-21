## ``OrderManagement`` Service and Client - Java Implementation

### Building and Running Service

In order to build gradle project, Go to ``Java`` project root directory location (samples) and execute
 the following shell command,
```
./gradlew :ch06:token-based-authentication:java:server:build
```

In order to run, Go to ``Java`` project root directory location (token-based-authentication/java/server) and execute the following
shell command,

```
java -jar build/libs/server.jar
```

### Building and Running Client

In order to build gradle project, Go to ``Java`` project root directory location (inside samples directory) and execute
 the following shell command,
```
./gradlew :ch06:token-based-authentication:java:client:build
```

In order to run, Go to ``Java`` project root directory location (token-based-authentication/java/client) and execute the following
shell command,

```
java -jar build/libs/client.jar
```

## Additional Information

### Generate Server key and certificate

* Generate Using [OpenSSL](../certs/README.md)