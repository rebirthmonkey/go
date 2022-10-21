# gRPC

## 简介

`gRPC`是一款语言中立、平台中立、开源的远程过程调用系统。gRPC 的客户端和服务端可以在多种环境中运行和交互，例如用`java`写的服务端可以用 go 写客户端调用。各服务间使用高效的`protobuf`协议进行 RPC 调用，gRPC默认使用`protocol buffers`，这是 Google 开源的一套成熟的结构数据序列化机制。可以用 proto files 创建 gRPC 服务，用 message 类型来定义方法参数和返回类型。在 gRPC 中，客户端可以直接调用部署在不同机器上的 gRPC 服务所提供的方法，调用远端的 gRPC 方法就像调用本地的方法一样，非常简单方便，通过 gRPC 调用，我们可以非常容易地构建出一个分布式应用。

RPC 具有如下特性：

- 支持多种语言：如 Go、Java、C、C++、C#、Node.js、PHP、Python、Ruby 等。
- 基于 IDL（Interface Definition  Language）定义服务：预先定义好接口（接口的名字、传入参数和返回参数等），在服务端，gRPC 服务实现我们所定义的接口。在客户端，gRPC 存根提供了跟服务端相同的方法。通过 proto3 工具生成指定语言的数据结构、服务端接口以及客户端 Stub。通过这种方式也可以将服务端和客户端解耦，使客户端和服务端可以并行开发。
- 基于标准的 HTTP/2：支持双向流、消息头压缩、单 TCP 的多路复用、服务端推送等特性。
- 支持 Protobuf 和 JSON 序列化数据格式：Protobuf 是一种语言无关的高性能序列化框架，可以减少网络传输流量，提高通信效率。

### 服务方法

gRPC 支持定义 4 种类型的服务方法，分别是简单模式、服务端数据流模式、客户端数据流模式和双向数据流模式：

- 简单模式（Simple RPC）：是最简单的 gRPC 模式。客户端发起一次请求，服务端响应一个数据。定义格式为 `rpc SayHello (HelloRequest) returns  (HelloReply) {}`。
- 服务端数据流模式（Server-side streaming  RPC）：客户端发送一个请求，服务器返回数据流响应，客户端从流中读取数据直到为空。定义格式为 `rpc SayHello  (HelloRequest) returns (stream HelloReply) {}`。
- 客户端数据流模式（Client-side  streaming RPC）：客户端将消息以流的方式发送给服务器，服务器全部处理完成之后返回一次响应。定义格式为 `rpc SayHello  (stream HelloRequest) returns (HelloReply) {}`。
- 双向数据流模式（Bidirectional  streaming RPC）：客户端和服务端都可以向对方发送数据流，这个时候双方的数据可以同时互相发送，也就是可以实现实时交互 RPC。定义格式为 `rpc SayHello (stream HelloRequest) returns (stream  HelloReply) {}`。

### Install

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Protocol Buffers

Protocol Buffers（ProtocolBuffer/protobuf）是 Google 开发的一套对数据结构进行序列化的方法，可用作（数据）通信协议、数据存储格式等，也是一种更加灵活、高效的数据格式，与 XML、JSON 类似。它的传输性能非常好，所以常被用在一些对数据传输性能要求比较高的系统中，作为数据传输格式。Protocol Buffers 的主要特性有：

- 更快的数据传输速度：protobuf 在传输时，会将数据序列化为二进制数据，和 XML、JSON 的文本传输格式相比，可以节省大量的 IO 操作，从而提高数据传输速度。
- 跨平台多语言：protobuf 自带的编译工具 protoc 可以基于 protobuf 定义文件，编译出不同语言的客户端或服务端，供程序直接调用，因此可以满足多语言需求的场景。
- 具有非常好的扩展性和兼容性：可以更新已有的数据结构，而不破坏和影响原有的程序。
- 基于 IDL 文件定义服务：通过 proto3 工具生成指定语言的数据结构、服务端和客户端接口。

在 gRPC 的框架中，Protocol Buffers 主要有三个作用：

- 第一，可以用来定义数据结构：下面的代码定义了一个 SecretInfo 数据结构：

```go
// SecretInfo contains secret details.
message SecretInfo {
    string name = 1;
    string secret_id  = 2;
    string username   = 3;
    string secret_key = 4;
    int64 expires = 5;
    string description = 6;
    string created_at = 7;
    string updated_at = 8;
}
```

- 第二，可以用来定义服务接口：下面的代码定义了一个 Cache 服务，服务包含了 ListSecrets 和 ListPolicies 两个 API 接口。

```go
// Cache implements a cache rpc service.
service Cache{
  rpc ListSecrets(ListSecretsRequest) returns (ListSecretsResponse) {}
  rpc ListPolicies(ListPoliciesRequest) returns (ListPoliciesResponse) {}
}
```

- 第三，可以通过 protobuf 序列化和反序列化，提升传输效率。



## Example

### Hello-World

#### 定义 gRPC 服务接口文件

首先，需要定义我们的服务。进入 helloworld 目录，新建文件 helloworld.proto：

```go
syntax = "proto3";

package helloworld;

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```

本示例使用了简单模式，.proto 文件也包含了 Protocol Buffers 消息的定义，包括请求消息和返回消息。

#### 生成客户端和服务端接口

根据 .proto 服务定义生成 gRPC 客户端和服务器接口。可以使用 protoc 编译工具，并指定使用其 Go 语言插件来生成：

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld.proto

protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative ProductInfo.proto
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative ProductInfo.proto
    
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative \
  ProductInfo.proto
```

```shell
$ ls
helloworld.pb.go  helloworld.proto # 新增了一个 helloworld.pb.go 文件
```

#### 实现 gRPC 服务端

进入 ../greeter_server 目录，新建 main.go 文件。

在代码中实现了上一步根据服务定义生成的 Go 接口：

- 先定义了一个 Go 结构体 server，并为 server 结构体添加了`SayHello(context.Context,  pb.HelloRequest) (pb.HelloReply, error)`方法，也就是 server 是 GreeterServer  接口（位于 helloworld.pb.go 文件中）的一个实现。
- 在实现了 gRPC 服务所定义的方法之后，就可以通过 net.Listen(...) 指定监听客户端请求的端口；
- 通过 grpc.NewServer() 创建一个 gRPC Server 实例，并通过 `pb.RegisterGreeterServer(s, &server{})` 将该服务注册到 gRPC 框架中；
- 通过 s.Serve(lis) 启动 gRPC 服务。
- 创建完 main.go 文件后，在当前目录下执行 `go run main.go` 启动  gRPC 服务。

注意：如果 go-grpc 不同版本存在兼容性问题，需要手动调整 go.mod 中的版本号。

#### 实现 gRPC 客户端

进入 ../greeter-client 目录，新建 main.go 文件。

- 创建了一个 gRPC 连接，用来跟服务端进行通信。在创建连接时，可以指定不同的选项，用来控制创建连接的方式，如 `grpc.WithInsecure()`、`grpc.WithBlock()` 等，更多的选项可以参考 grpc 仓库下dialoptions.go文件中以 With 开头的函数。
- 连接建立起来之后，需要创建一个客户端 stub，用来执行 RPC 请求 `c :=  pb.NewGreeterClient(conn)`。
- 创建完成之后，就可以像调用本地函数一样，调用远程的方法了。例如 `r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})`。
  -  RPC 方便了，RPC 屏蔽了底层的网络通信细节，使得调用 RPC 就像调用本地方法一样方便。调用方式跟大家所熟知的调用类的方法一致：ClassName.ClassFuc(params)。
  - RPC 不需要打包和解包，RPC 调用的入参和返回的结果都是 Go 的结构体，不需要对传入参数进行打包操作，也不需要对返回参数进行解包操作，简化了调用步骤。
- 创建完 main.go 文件后，在当前目录下执行 `go run main.go`  发起 RPC 调用。

### route-guide

#### 创建proto

创建 route_guide.pb.go 和 route_guide_grpc.pb.go

- route_guide.pb.go：contains all the protocol buffer code to populate, serialize, and retrieve request and response message types
- route_guide_grpc.pb.go：
  - An interface type (or *stub*) for clients to call with the methods defined in the `RouteGuide` service.
  - An interface type for servers to implement, also with the methods defined in the `RouteGuide` service.

```bash
protoc --go_out=plugins=grpc:./test/ ./test.proto

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    test/test.proto
```

#### 创建 server 端



#### 创建 client 端