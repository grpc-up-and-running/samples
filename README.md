
![gRPC Up and Running - Cover](./images/grpc-cover.png)


## Welcome
Welcome to the resource repository of the book "gRPC - Up and Running". 

All the samples of this repository require to have the accompanying book [gRPC Up and Running](https://www.amazon.com/gRPC-Running-Building-Applications-Kubernetes/dp/1492058335/). Each sample is based on a real-world use case and details of the use case can be found in the respective chapter of the book. 


## Don't have _gRPC Up and Running_ book yet? 

You can purchase it from following book sellers worldwide.  

* [Amazon](https://www.amazon.com/gRPC-Running-Building-Applications-Kubernetes/dp/1492058335/)
* [Barnes and Nobel](https://www.barnesandnoble.com/w/grpc-kasun-indrasiri/1132647211?ean=9781492058335#/) 

## Prerequisites

* Installing *Go*    
    Download the latest Go distribution file from [Go’s official download page](https://golang.org/dl/) and follow the steps to install Go language.
 
* Installing *Java*    
    Samples are tested in JDK 1.8 version. Recommended to install JDK 1.8 version. [Java official downland page](https://www.java.com/en/download/)

* Installing *Gradle*  
    [Gradle](https://gradle.org/) is required for builing and running all the Java samples. So, please install Gradle if you plan to try out Java samples of the book. 

* Installing *protoc*  
    [protoc](https://developers.google.com/protocol-buffers/docs/downloads) is required for generating code for your protocol buffer based gRPC service definitions (IDLs). Hence please install if you plan to use code generation. 

  
## Samples from the _gRPC Up and Running_ Book

- Chapter 01 - Introduction to gRPC
    - No code samples required for this chapter.

- Chapter 02 - Getting Started with gRPC
    - Getting Started with gRPC [[Go]](./ch02/productinfo/go) [[Java]](./ch02/productinfo/java) [[Python]](./ch02/productinfo/python) [[Node.js]](./ch02/productinfo/nodejs) [[Kotlin]](./ch02/productinfo/kotlin)
    
- Chapter 03 - gRPC Communication Patterns
    - Communication patterns [[Go]](./ch03/order-service/go) [[Java]](./ch03/order-service/java) 

- Chapter 04 - gRPC Under the Hood
    - No code samples required for this chapter. 

- Chapter 05 - gRPC Beyond the Basics 
    - Interceptors [[Go]](./ch05/interceptors/order-service/go) [[Java]](./ch05/interceptors/order-service/java) 
    - Deadline [[Go]](./ch05/deadlines/order-service/go) [[Java]](./ch05/deadlines/order-service/java) 
    - Cancellation [[Go]](./ch05/cancellation/order-service/go) [[Java]](./ch05/cancellation/order-service/java) 
    - Compression [[Go]](./ch05/compression/order-service/go) [[Java]](./ch05/compression/order-service/java) 
    - Keepalive [[Go]](./ch05/keepalive/order-service/go) [[Java]](./ch05/keepalive/order-service/java) 
    - Metadata [[Go]](./ch05/metadata/order-service/go) [[Java]](./ch05/metadata/order-service/java) 
    - Error Handling [[Go]](./ch05/error-handling/order-service/go) [[Java]](./ch05/error-handling/order-service/java) 
    - Load Balancing [[Go]](./ch05/loadbalancing/echo/go) [[Java]](./ch05/loadbalancing/echo/java) 
    - Multiplexing [[Go]](./ch05/multiplexing/order-service/go) [[Java]](./ch05/multiplexing/order-service/java)

- Chapter 06 - Secured gRPC
    - Secure Channel [[Go]](./ch06/secure-channel/go) [[Java]](./ch06/secure-channel/java) 
    - Mutual TLS Channel [[Go]](./ch06/mutual-tls-channel/go) [[Java]](./ch06/mutual-tls-channel/java) 
    - Basic Authentication [[Go]](./ch06/basic-authentication/go) [[Java]](./ch06/basic-authentication/java) 
    - Token Based Authentication [[Go]](./ch06/token-based-authentication/go) [[Java]](./ch06/token-based-authentication/java) 
    
- Chapter 07 - Running gRPC in Production
    - Continuous Integration [[Go]](./ch07/grpc-continous-integration/go)
    - Deploy in Docker [[Go]](./ch07/grpc-docker/go)
    - Deploy in Kubernetes [[Go]](./ch07/grpc-kubernetes/go)
    - OpenCensus metrics [[Go]](./ch07/grpc-opencensus/go) [[Java]](./ch07/grpc-opencensus/java)
    - OpenCensus tracing [[Go]](./ch07/grpc-opencensus-tracing/go) [[Java]](./ch07/grpc-opencensus-tracing/java)
    - OpenTracing [[Go]](./ch07/grpc-opentracing/go)
    - Prometheus [[Go]](./ch07/grpc-prometheus/go)
    
- Chapter 08 - The gRPC Ecosystem
    - gRPC Gateway [[Go]](./ch08/grpc-gateway/go) 
    - Server Reflection [[Go]](./ch08/server-reflection/go) [[Java]](./ch08/server-reflection/java) 
    
## Authors 
### Kasun Indrasiri
![Kasun](https://raw.githubusercontent.com/grpc-up-and-running/samples/master/images/kasun.jpg)

Kasun Indrasiri is an author and an evangelist of microservices and enterprise integration architecture with over ten years of experience in building distributed systems. He is the director of Integration Architecture at WSO2 and the product manager of the WSO2 Enterprise Integrator. He has authored [Microservices for Enterprise](https://www.amazon.com/Microservices-Enterprise-Designing-Developing-Deploying/dp/1484238575) (Apress, 2018) and [Beginning WSO2 ESB](https://www.amazon.com/Beginning-WSO2-ESB-Kasun-Indrasiri/dp/148422342X) (Apress, 2017) and has spoken at several conferences, including the O’Reilly Software Architecture Conference 2019 in San Jose and GOTO Con 2019 in Chicago, and WSO2 Conferences in San Francisco, London, and Barcelona. Kasun lives in San Jose, California, and has founded the “Silicon Valley Microservices, APIs and Integration” Meetup, which is one of the largest microservices meetups in the San Francisco Bay area..


### Danesh Kuruppu 
![Danesh](https://raw.githubusercontent.com/grpc-up-and-running/samples/master/images/danesh.jpg)

Danesh Kuruppu is an associate technical lead at WSO2 with expertise in microservices, messaging protocols, and service governance. Danesh has spearheaded the development
of gRPC language implementation and microservices framework.
