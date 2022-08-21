# App Gin 最佳实践

## 简介

本例子为构建一个 Gin Web App 的最佳实践，具体分为以下步骤：

- options.go：导入参数配置文件
- config.go：把参数配置转换为 app 的配置信息
- server.go：
  - 把配置信息赋值给 server 的状态实例 GINServer
  - 启动 GINServer 实例
- router.go：Gin 中的多级路由器
- middleware：Gin 用到的中间件
- handler.go：Gin 应用真正的逻辑代码实现


## 运行


```shell
$ go run main.go options.go config.go server.go router.go middleware.go handler.go 
```

```shell
# 创建产品
$ curl -XPOST -H"Content-Type: application/json" -d'{"username":"colin","name":"iphone12","category":"phone","price":8000,"description":"cannot afford"}' http://127.0.0.1:8080/v1/products
{"username":"colin","name":"iphone12","category":"phone","price":8000,"description":"cannot afford","createdAt":"2021-06-20T11:17:03.818065988+08:00"}

# 获取产品信息
$ curl -XGET http://127.0.0.1:8080/v1/products/iphone12
{"username":"colin","name":"iphone12","category":"phone","price":8000,"description":"cannot afford","createdAt":"2021-06-20T11:17:03.818065988+08:00"}
```


## 二次开发

- 添加 handler 或 middleware
- 配置 router.go