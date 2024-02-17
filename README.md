# grpc-demo-clone

```proto
// normal mod, one request, one response
rpc Add (TwoNum) returns (Response) {}
```

### How run

1. demo.pb.go

```proto
protoc --go_out=. demo.proto
```

2. demo_grpc.pb.go

```proto
protoc --go_grpc_out=. demo.proto
```

3. run

```bash
go run server/main.go
go run client/main.go
```
