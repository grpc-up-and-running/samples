# Chapter 7: Examining metrics for product info service and client with Stackdriver stats exporter

export GOOGLE_APPLICATION_CREDENTIALS="/Users/daneshk/Advanced-gRPC/samples/ch07/grpc-opencensus/credentials/grpc-up-and-running-demo-452678b739e1.json"

## Obtaining service account credentials to publish metrics

* Create a service account GCP(Google Cloud Platform), generate key and download as JSON file to your computer.
* Provide credentials file(JSON file) as environment variable or pass the path to the service account key in code.

```
export GOOGLE_APPLICATION_CREDENTIALS="[PATH]"
For example:
export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/[FILE_NAME].json"
```

Refer https://cloud.google.com/docs/authentication/production#auth-cloud-implicit-java for more information


