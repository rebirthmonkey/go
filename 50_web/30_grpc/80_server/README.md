# apiserver GRPC Server

## 开发

### 创建 PB 文件

```bash
cd apiserver/user/controller/grpc/v1
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --experimental_allow_proto3_optional \
    user.proto
ls
```

## 运行


```shell
go run cmd/apiserver.go -c configs/config-out.yaml
```

## Test

```shell
go run test/grpc/user_client.go
curl -X GET http://127.0.0.1:8080/v1/users
curl -X POST -H "Content-Type: application/json" \
     -d '{"metadata":{"name":"user99", "password":"admin"},"description":"admin user"}' \
     http://127.0.0.1:8080/v1/users # create a new user user99
curl -X GET http://127.0.0.1:8080/v1/users
go run test/grpc/user_client.go
```
