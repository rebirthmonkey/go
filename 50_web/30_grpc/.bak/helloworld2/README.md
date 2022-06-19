# gRPC

## 步骤

### 更新 Proto

- 更新 `helloworld.proto` 文件

- 更新 `helloworld.pb.go`、`helloworld_grpc.pb.go`
```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```

- 更新 `server/main.go`：实现新的 rpc 函数

- 更新 `client/main.go`：使用新的 rpc 函数

### 运行

- `go run server/main.go`
- `go run client/main.go`


## Ref
1. [Quick start](https://grpc.io/docs/languages/go/quickstart/)

