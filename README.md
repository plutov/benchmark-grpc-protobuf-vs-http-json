### gRPC+Protobuf or JSON+HTTP?

This repository contains 2 equal APIs: gRPC using Protobuf and JSON over HTTP. The goal is to run benchmarks for 2 approaches and compare them. APIs have 1 endpoint to create user, containing validation of request. Request, validation and response are the same in 2 packages, so we're benchmarking only mechanism itself. Benchmarks also include response parsing.

### Run tests

Install dependencies first:

```
brew install glide
glide i
```

Run each command in different tab:

```
go run grpc-protobuf/main.go
go run http-json/main.go
go test -bench=.
```

### Results

Tested in Go 1.9.
```
BenchmarkGRPCProtobuf-8   	   10000	    197919 ns/op
BenchmarkJSONHTTP-8       	    1000	   1720124 ns/op
```

gRPC+Protobuf is **10** times faster!

### CPU usage comparison

Restart applications, then use profiling tool `pprof` during 30 sec when the client is talking to the server with these commands in different tabs:

```
go tool pprof http://localhost:6060/debug/pprof/profile
go tool pprof http://localhost:6061/debug/pprof/profile
```

Run tests to get client connections. Then in each `pprof` run `top` to see CPU usage.
My results show that Protobuf consumes less ressources, **30% less**.
