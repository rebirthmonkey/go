# GIN

## 简介

Gin 是 Go 语言写的一个 web 框架，它具有运行速度快，分组的路由器，良好的崩溃捕获和错误处理，非常好的支持中间件和 JSON。总之在 Go语言开发领域是一款值得好好研究的 Web 框架，[开源网址](https://github.com/gin-gonic/gin)。

Gin 核心的路由功能是通过一个定制版的 HttpRouter 来实现的，具有很高的路由性能，其主要功能包括：

- 支持 HTTP 方法：GET、POST、PUT、PATCH、DELETE、OPTIONS。
- 支持不同的 HTTP 参数：路径参数（path）、查询字符串参数（query）、表单参数（form）、HTTP 头参数（header）、消息体参数（body）。
- 支持 HTTP 路由和路由分组。
- 支持 middleware 和自定义 middleware。
- 支持自定义 Log。
- 支持 binding 和 validation，支持自定义 validator。可以 bind 如下参数：query、path、body、header、form。
- 支持重定向。
- 支持 basic auth middleware。
- 支持自定义 HTTP 配置。
- 支持优雅关闭。
- 支持 HTTP2。
- 支持设置和获取 cookie。

Installation: `go get -u github.com/gin-gonic/gin`

## Bind

- c.ShouldBindUri()
- c.ShouldBindQuery()
- c.ShouldBind()：表单
- c.ShouldBindHead()
- c.ShouldBindJson()：消息体 body


## Lab

- [HelloWorld](01_hello-world.go)

```shell
go run 01_hello-world.go
curl http://127.0.0.1:8080
```

- [Get with JSON response](10_get-json.go)

```shell
go run 10_get-json.go
curl http://127.0.0.1:8080/ping
```

- [Get with JSON response 2](11_get-json.go)

```shell
go run 11_get-json.go
curl http://127.0.0.1:8080/ping
```

- [Get with Path Parameter](13_get-param-path.go)

```shell
go run 13_get-param-path.go
curl http://127.0.0.1:8080/users/xxx 
```

- [Get with Path Parameter General](14_get-param-path.go)

```shell
go run 14_get-param-path.go
curl http://127.0.0.1:8080/users/xxx/
curl http://127.0.0.1:8080/users/xxx/yyy/zzz
```

- [Get with Query Parameter](16_get-param-query.go)

```shell
go run 16_get-param-query.go
curl http://127.0.0.1:8080/welcome
curl http://127.0.0.1:8080/welcome\?firstname\=中国
curl http://127.0.0.1:8080/welcome\?firstname\=中国\&lastname\=天朝
curl http://127.0.0.1:8080/welcome\?firstname\=\&lastname\=天朝
curl http://127.0.0.1:8080/welcome\?firstname\=%E4%B8%AD%E5%9B%BD
```

- [Get Cookie](18_get-cookie.go)

```shell
go run 18_get-cookie.go
curl http://127.0.0.1:8080/auth/signin
```

- [Post Body](21_post-body.go)

```shell
go run 21_post-body.go
curl -X POST http://127.0.0.1:8080/form_post -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&nick=wukong"
```

- [Post File](23_post-file.go)

```shell
go run 23_post-file.go
curl -X POST http://localhost:8080/upload \
	-F "file=@./23_post-file.go" \
	-H "Content-Type: multipart/form-data"
```

- [Post File](23_post-file.go)

```shell
go run 24_post-multi-file.go
curl -X POST http://localhost:8080/upload \
  -F "file[]=@./23_post-file.go" \
  -F "file[]=@./24_post-multi-file.go" \
  -H "Content-Type: multipart/form-data"
```

- [Redirect](31_redirect.go)

```shell
curl http://127.0.0.1:8080/redirect/google
```

- [Multiple Routes](35_route-multi.go)

```shell
curl http://127.0.0.1:8080/v1/login
curl http://127.0.0.1:8080/v2/login
```

- [Embedded Routes](37_route-embedded.go)

```shell
curl http://127.0.0.1:8080/user/index
curl http://127.0.0.1:8080/user/login
curl http://127.0.0.1:8080/user/shop/index
```

- [Middleware1](41_mw.go)

```shell
curl http://127.0.0.1:8080/before
curl http://127.0.0.1:8080/after
```

- [Middleware2](43_mw.go)

```shell
curl http://127.0.0.1:8080/before
curl http://127.0.0.1:8080/after
```

- [Async](51_async.go)

```shell
curl http://127.0.0.1:8080/sync
curl http://127.0.0.1:8080/async
```

### A Real Web Application

- [A Real Web Application with HTTP and HTTPS](80_bind-json.go)：通过c.ShouldBindJSON函数，将 Body 中的 JSON 格式数据解析到指定的 Struct 中，通过c.JSON函数返回 JSON 格式的数据。

```shell
# 创建产品
$ curl -XPOST -H"Content-Type: application/json" -d'{"username":"colin","name":"iphone12","category":"phone","price":8000,"description":"cannot afford"}' http://127.0.0.1:8080/v1/products
{"username":"colin","name":"iphone12","category":"phone","price":8000,"description":"cannot afford","createdAt":"2021-06-20T11:17:03.818065988+08:00"}

# 获取产品信息
$ curl -XGET http://127.0.0.1:8080/v1/products/iphone12
{"username":"colin","name":"iphone12","category":"phone","price":8000,"description":"cannot afford","createdAt":"2021-06-20T11:17:03.818065988+08:00"}
```

- [Middleware with new Logger MW](82_mw-logger.go)：自定义 logger middleware

```shell
curl http://127.0.0.1:8080/test
```

- [Middleware with Gin MW](84_mw.go)：使用 Gin 现有的 middleware


## Ref

1. [Golang 微框架 Gin 简介](https://www.jianshu.com/p/a31e4ee25305)
2. 



