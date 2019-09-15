
# Multi stage build

# Build stage I : Go lang and Alpine Linux is only needed to build the program
FROM golang AS build

ENV location /go/src/github.com/grpc-up-and-running/samples/ch07/grpc-docker/go

WORKDIR ${location}/client

ADD ./client ${location}/client
ADD ./proto-gen ${location}/proto-gen

# Download all the dependencies
RUN go get -d ./...
# Install the package
RUN go install ./...


RUN CGO_ENABLED=0 go build -o /bin/grpc-productinfo-client



# Build stage II : Go binaries are self-contained executables.
FROM scratch
COPY --from=build /bin/grpc-productinfo-client /bin/grpc-productinfo-client


ENTRYPOINT ["/bin/grpc-productinfo-client"]
EXPOSE 50051