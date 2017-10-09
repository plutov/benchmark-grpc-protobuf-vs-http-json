### gRPC+Protobuf or JSON over HTTP?

This repository contains 2 equal APIs: gRPC using Protobuf and JSON over HTTP. The goal is to run benchmarks for 2 approaches and compare them. APIs have 1 endpoint to create user, containing validation of request. Request, validation and response are the same in 2 packages, so we're benchmarking only mechanism itself. Benchmarks also include response parsing.

### Run tests

```
glide i
go test -bench=.
```

### Results

Tested in Go 1.9.
```
BenchmarkGRPCProtobuf-8   	   10000	    197919 ns/op
BenchmarkJSONHTTP-8       	    1000	   1720124 ns/op
```

gRPC+Protobuf is **10** times faster!

### gRPC

Proto definition is described in `grpc/proto/api.proto` and Go bindings built with:

```
protoc --go_out=plugins=grpc:. grpc/proto/api.proto
```
