# GRPC service

This repo contains the grpc generated code which could be used by client and server implementations.

## Code generation

### Generate hello service grpc code

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./helloservice/helloservice.proto
```