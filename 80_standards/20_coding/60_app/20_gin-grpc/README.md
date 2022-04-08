# App Gin+GRPC 最佳实践

## 简介

本例子为构建一个 Gin+GRPC Web App 的最佳实践，具体分为以下步骤：

- main.go：程序入口
- options.go：导入参数配置文件
- config.go：把参数配置转换为 app 的通用配置信息
- Gin
  - gin-config.go: 把参数配置转化为 Gin 用到的配置信息
  - gin-server.go：
    - 把配置信息赋值给 server 的状态实例 GINServer
    - 启动 GINServer 实例
  - gin-router.go：Gin 中的多级路由器
  - gin-middleware：Gin 用到的中间件
  - gin-handler.go：Gin 应用真正的逻辑代码实现
- GRPC
  - grpc-config.go: 把参数配置转化为 Grpc 用到的配置信息
  - grpc-server.go: 启动 Grpc server
  - grpc-handler.go：GRPC 应用真正的逻辑代码实现，其数据结构定义在对应的 ProtoBuff 中
  - grpc-client.go：GRPC 客户端调用


## 运行
```shell
$ go run main.go options.go config.go \
  gin-config.go gin-server.go gin-router.go gin-middleware.go gin-handler.go \
  grpc-config.go grpc-server.go grpc-handler.go 
```

```shell
# 创建产品
curl -XPOST -H"Content-Type: application/json" -d'{"username":"colin","name":"iphone12","category":"phone","price":8000,"description":"cannot afford"}' http://127.0.0.1:8080/v1/products
{"username":"colin","name":"iphone12","category":"phone","price":8000,"description":"cannot afford","createdAt":"2021-06-20T11:17:03.818065988+08:00"}

# 获取产品信息
$ curl -XGET http://127.0.0.1:8080/v1/products/iphone12
{"username":"colin","name":"iphone12","category":"phone","price":8000,"description":"cannot afford","createdAt":"2021-06-20T11:17:03.818065988+08:00"}

$ go run grpc-client.go
```


### 二次开发

- 添加 handler 或 middleware
- 配置 router.go

