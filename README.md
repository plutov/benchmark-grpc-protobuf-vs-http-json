### gRPC or REST?

This repository contains 2 equal APIs: gRPC and REST. The goal is to run benchmarks for 2 approaches and compare them. APIs have 1 endpoint to create user, containing validation of request. Request, validation and response are the same in 2 packages, so we're benchmarking only mechanism itself. Benchmarks also include response parsing.

### Run tests

```
glide i
go run grpc/main.go &
go run rest/main.go &
go test -bench=.
```

### Results

Tested in Go 1.9.
```
BenchmarkGRPC-4   	   10000	    206301 ns/op
BenchmarkREST-4   	    3000	    497311 ns/op
```

**3** times faster!

### gRPC

Proto definition is described in `grpc/proto/api.proto` and Go bindings built with:

```
protoc --go_out=plugins=grpc:. grpc/proto/api.proto
```
