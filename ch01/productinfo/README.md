# Product Info Service and Client 

## Prerequisite

Install Bazel version 0.19.2

Installation Step: https://docs.bazel.build/versions/master/install-os-x.html#install-with-installer-mac-os-x


## Go Server and Client

### Go Server
Build Go Server

bazel build //ch01/productinfo/go/server:server

Run Go Server

bazel-bin/ch01/productinfo/go/server/darwin_amd64_stripped/server

### Go Client
Build Go Client

bazel build //ch01/productinfo/go/client:client

Run Go Client

bazel-bin/ch01/productinfo/go/client/darwin_amd64_stripped/client


## Java Server and Client

### Java Server
Build Java Server

bazel build //ch01/productinfo/java:server

Run Java Server

bazel-bin/ch01/productinfo/java/server

### Java Client

Build Java Client

bazel build //ch01/productinfo/java:client

Run Java Client

bazel-bin/ch01/productinfo/java/client

