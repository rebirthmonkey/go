# gRPC

## 简介

`gRPC`是一款语言中立、平台中立、开源的远程过程调用系统。gRPC 的客户端和服务端可以在多种环境中运行和交互，例如用`java`写的服务端可以用 go 写客户端调用。各服务间使用高效的`protobuf`协议进行 RPC 调用，gRPC默认使用`protocol buffers`，这是 Google 开源的一套成熟的结构数据序列化机制。可以用 proto files 创建 gRPC 服务，用 message 类型来定义方法参数和返回类型。

### Install

```shell
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Example

### route-guide

#### 创建proto

创建 route_guide.pb.go 和 route_guide_grpc.pb.go

- route_guide.pb.go：contains all the protocol buffer code to populate, serialize, and retrieve request and response message types
- route_guide_grpc.pb.go：
  - An interface type (or *stub*) for clients to call with the methods defined in the `RouteGuide` service.
  - An interface type for servers to implement, also with the methods defined in the `RouteGuide` service.

```shell
protoc --go_out=plugins=grpc:./test/ ./test.proto

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    test/test.proto
```

#### 创建 server 端



#### 创建 client 端