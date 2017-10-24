### gRPC+Protobuf or JSON+HTTP?

This repository contains 2 equal APIs: gRPC using Protobuf and JSON over HTTP. The goal is to run benchmarks for 2 approaches and compare them. APIs have 1 endpoint to create user, containing validation of request. Request, validation and response are the same in 2 packages, so we're benchmarking only mechanism itself. Benchmarks also include response parsing.

### Requirements

 - Go 1.9
 - [dep](https://github.com/golang/dep)

### Run tests

Install dependencies first:

```
dep ensure
```

Run benchmarks:
```
go test -bench=.
```

### Results

```
BenchmarkGRPCProtobuf-8   	   10000	    197919 ns/op
BenchmarkJSONHTTP-8       	    1000	   1720124 ns/op
```

gRPC+Protobuf is **10** times faster!

### CPU usage comparison

This will create an executable `benchmark-grpc-protobuf-vs-http-json.test` and the profile information will be stored in `grpcprotobuf.cpu` and `httpjson.cpu`:

```
go test -bench=BenchmarkGRPCProtobuf -cpuprofile=grpcprotobuf.cpu
go test -bench=BenchmarkHTTPJSON -cpuprofile=httpjson.cpu
```

Check CPU usage per approach using:

```
go tool pprof grpcprotobuf.cpu
go tool pprof httpjson.cpu
```

My results show that Protobuf consumes less ressources, around **30% less**.

### gRPC definition

```
```
protoc --go_out=plugins=grpc:. grpc-protobuf/proto/api.proto
```