# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

ENV GOPATH /go

# Copy the local package files to the container's workspace.
ADD server /go/src/github.com/advanced-grpc/samples/ch01/productinfo/go/server
ADD product_info /go/src/github.com/advanced-grpc/samples/ch01/productinfo/go/product_info

WORKDIR /go/src/github.com/advanced-grpc/samples/ch01/productinfo/go/server

# Download all the dependencies
RUN go get -d ./...
# Install the package
RUN go install ./...

RUN  go build -i -o main.go
# Run the server command by default when the container starts.
ENTRYPOINT ["/go/bin/server"]

# Document that the service listens on port 50051.
EXPOSE 50051
