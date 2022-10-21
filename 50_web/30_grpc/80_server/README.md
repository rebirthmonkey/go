# apiserver GRPC Server

## 开发

### 创建 PB 文件

```bash
cd apiserver/user/controller/grpc/v1
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --experimental_allow_proto3_optional \
    user.proto
```

## 运行

- server.pem `/Users/ruan/workspace/go/50_web/30_grpc/80_server/configs/cert/server.pem`
- server.key `/Users/ruan/workspace/go/50_web/30_grpc/80_server/configs/cert/server.key`


```shell
cat configs/config.yaml | sed "s#{CERT-FILE}#$(pwd)/configs/cert/server.pem#g" | sed "s#{PRIVATE-KEY-FILE}#$(pwd)/configs/cert/server.key#g"  > configs/config-out.yaml
go run cmd/apiserver.go -c configs/config-out.yaml
```

## Test

```shell
cat configs/config.yaml | sed "s#{CERT-FILE}#$(pwd)/configs/cert/server.pem#g" | sed "s#{PRIVATE-KEY-FILE}#$(pwd)/configs/cert/server.key#g"  > configs/config-out.yaml
go run cmd/apiserver.go -c configs/config-out.yaml &
# 新建终端，并在新建的终端继续
go run test/grpc/user_client.go
curl -X GET http://127.0.0.1:8080/v1/users
curl -X POST -H "Content-Type: application/json" \
     -d '{"metadata":{"name":"user99", "password":"admin"},"description":"admin user"}' \
     http://127.0.0.1:8080/v1/users # create a new user user99
curl -X GET http://127.0.0.1:8080/v1/users
go run test/grpc/user_client.go
```

或者在同一个终端窗口中执行

```bash
cat configs/config.yaml | sed "s#{CERT-FILE}#$(pwd)/configs/cert/server.pem#g" | sed "s#{PRIVATE-KEY-FILE}#$(pwd)/configs/cert/server.key#g"  > configs/config-out.yaml
go run cmd/apiserver.go -c configs/config-out.yaml &
sleep 10
go run test/grpc/user_client.go
curl -X GET http://127.0.0.1:8080/v1/users
curl -X POST -H "Content-Type: application/json" \
     -d '{"metadata":{"name":"user99", "password":"admin"},"description":"admin user"}' \
     http://127.0.0.1:8080/v1/users # create a new user user99
curl -X GET http://127.0.0.1:8080/v1/users
go run test/grpc/user_client.go
```

