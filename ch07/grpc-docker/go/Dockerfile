

# Multi stage build

# Build stage I : Go lang and Alpine Linux is only needed to build the program
FROM golang:1.11-alpine AS build

ENV location /go/src/github.com/grpc-up-and-running/samples/ch07/productinfo/go/server


WORKDIR ${location}
ADD server ${location}
ADD ../proto ${location}


# Download all the dependencies
RUN go get -d ./...
# Install the package
RUN go install ./...


RUN CGO_ENABLED=0 go build -o /bin/grpc-productinfo-server



# Build stage II : Go binaries are self-contained executables.
FROM scratch
COPY --from=build /bin/grpc-productinfo-server /bin/grpc-productinfo-server


ENTRYPOINT ["/bin/grpc-productinfo-server"]
EXPOSE 50051