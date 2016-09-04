# gRPC Experiment

* An experiment with gRPC http://www.grpc.io/
* useful for learning gRPC https://github.com/kelseyhightower/grpc-hello-service

## Compiling and Running

### Service Definitions and Protobufs

* Services and messages are defined in `protobuf/...`
* If needed then recompile the Go libraries from the protobuf definitions.
* Prerequisites from http://www.grpc.io/docs/quickstart/go.html

```
protoc --proto_path=protobuf/field/ --go_out=plugins=grpc:field protobuf/field/*
protoc --proto_path=protobuf/data/ --go_out=plugins=grpc:data protobuf/data/*
```

### Server

* Simple token based auth.
* Interceptor for metrics.
* Auth requires end-to-end TLS encryption.
* Reads TLS cert from `certs/server.crt` and `cert/server.key` otherwise a self signed TLS cert is generated on the fly.
* A self signed cert can be generated using:

```
openssl genrsa -out server.key 2048
openssl req -new -x509 -key server.key -out server.crt -days 3650
```

* It's possible to handle gRPC and HTTP in the same server see https://coreos.com/blog/gRPC-protobufs-swagger.html

#### Integration tests for the server:

```
cd mtr-api
export $(cat env.list | grep = | xargs); go test -v
```

#### Build and run the server:

```
cd mtr-api
export $(cat env.list | grep = | xargs); go build && ./mtr-api
```

#### Docker

* Build an image with `./build.sh mtr-api`
* Run the container.  Self signed TLS certificates will be created on the fly.

```
docker run --env-file mtr-api/env.list -p 8443:8443 quay.io/gclitheroe/mtr-api:latest
```

* Mount a volume with TLS certificates into the container.

```
docker run --env-file mtr-api/env.list -p 8443:8443 -v /work/certs:/certs quay.io/gclitheroe/mtr-api:latest
```

### Client

* Build and run the client.
* Server will log messages.
* Client reconnects after server restarts.
 
```
cd mtr-client
export $(cat env.list | grep = | xargs); go build && ./mtr-client
``` 
