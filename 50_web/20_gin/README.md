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

## 核心概念

### Context

gin.Context，封装了 request 和 response

#### Request

##### Header

- Get：c.GetHeader()/c.Request.Header.Get(XRequestIDKey)
- Set：c.Request.Header.Set(XRequestIDKey, rid)


##### Param()/API 参数

在 URL 的路径中获取变量



##### Query()/ URL 参数

在 URL 路径后的 `?key1=value2&key2=value2` 中获取变量



##### FormPost()

POST 请求 Body 中的变量



##### Bind

###### ShouldBindJson()

将 Request 的 Body 绑定到指定的结构体变量

###### c.ShouldBindUri()

将 Request URL 中的 param 绑定到指定的结构体变量

###### c.ShouldBindQuery()



###### c.ShouldBind()：表单



###### c.ShouldBindHead()





#### Respons

##### Header

- Set：c.Writer.Header().Set(XRequestIDKey, rid)

##### Writer.

- Status()：返回的 HTTP 状态码



##### String()

返回的内容，可以是 String，也可以是 JSON

将 String 作为 Response 返回



##### JSON()

将 gin.H 结构体转化为 JSON 作为 Response 返回



#### Context Key-Value

- Get：c.Get(key)
- Set：c.Writer.Header().Set(XRequestIDKey, rid)



### 路由

#### Group

可以分为多个 Route Group 分别对待

#### Embedded

还可以多级嵌套



#### Redirect

重定向



## Middleware

### Next()

该函数会先交由后续 middleware、handler 处理完后，再继续处理剩下的代码。



## Lab

### Basics

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

- [Get with Parameter](13_get-param.go)

```shell
go run 13_get-param.go
curl http://127.0.0.1:8080/users/xxx 
```

- [Get with Path Parameter General](14_get-param.go)

```shell
go run 14_get-param.go
curl http://127.0.0.1:8080/users/xxx/
curl http://127.0.0.1:8080/users/xxx/yyy/zzz
```

- [Get with Query](16_get-query.go)

```shell
go run 16_get-query.go
curl http://127.0.0.1:8080/welcome
curl http://127.0.0.1:8080/welcome\?firstname\=中国
curl http://127.0.0.1:8080/welcome\?firstname\=中国\&lastname\=天朝
curl http://127.0.0.1:8080/welcome\?firstname\=\&lastname\=天朝
curl http://127.0.0.1:8080/welcome\?firstname\=%E4%B8%AD%E5%9B%BD
```

- [Post Body](21_post-form.go)

```shell
go run 21_post-form.go
curl -X POST http://127.0.0.1:8080/form -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&nick=wukong"
```

- [Post File](23_post-file.go)

```shell
go run 23_post-file.go
curl -X POST http://127.0.0.1:8080/upload \
	-F "file=@./23_post-file.go" \
	-H "Content-Type: multipart/form-data"
```

- [Post File](24_post-multi-file.go)

```shell
go run 24_post-multi-file.go
curl -X POST http://127.0.0.1:8080/upload \
  -F "file[]=@./23_post-file.go" \
  -F "file[]=@./24_post-multi-file.go" \
  -H "Content-Type: multipart/form-data"
```

- [Post JSON](27_bind-json.go)

```shell
go run 27_post-json.go
curl -X POST http://127.0.0.1:8080/login \
	-H "Content-Type:application/json" \
	-d '{"username": "ruan", "passwd": "123", "age": 21}'
```

- [Post JSON](28_bind-json.go)

```shell
go run 28_post-json.go
curl -X POST http://127.0.0.1:8080/login \
	-H "Content-Type:application/json" \
	-d '{"username": "ruan", "passwd": "123", "age": 21}'
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

- [Middleware1](41_middleware.go)

/before 会通过 middleware1

/after 会同时通过 middleware1 和 middleware2

```shell
curl http://127.0.0.1:8080/before
curl http://127.0.0.1:8080/after
```

- [Middleware2](43_middleware.go)

/before 会只通过 middleware3

/after 会只通过 middleware4，所以会报错 request4 不存在


```shell
curl http://127.0.0.1:8080/before
curl http://127.0.0.1:8080/after
```

- [Middleware with new Logger MW](45_middleware-logger.go)：自定义 logger middleware

Middleware c.next() 之前的代码会在调用前完成，之后的代码会在调用返回后完成。

```shell
curl http://127.0.0.1:8080/test
```

### Middleware

生产环境可用的 middleware

#### Basic Auth

- [Middleware with Basic Auth](46_middleware-basic-auth.go)：使用现有 Basic Auth Middleware

Basic Auth 必须采用浏览器登录 http://127.0.0.1:8080/auth



### Web Application

- [A Real Web Application with HTTP and HTTPS](80_app/80_bind-json.go)：通过 c.ShouldBindJSON 函数，将 Body 中的 JSON 格式数据解析到指定的 Struct 中，通过 c.JSON 函数返回 JSON 格式的数据。

主要做法是创建一个 struct，然后把 POST 的内容通过 c.ShouldBindJSON 添加到该 struct 的变量。
在给这个 struct 变量添加类似 Create()、Get() 等方法注册到 router 的 POST、GET 上。

```shell
# 创建产品
$ curl -XPOST -H"Content-Type: application/json" -d'{"username":"colin","name":"iphone12","category":"phone","price":8000,"description":"cannot afford"}' http://127.0.0.1:8080/v1/products
{"username":"colin","name":"iphone12","category":"phone","price":8000,"description":"cannot afford","createdAt":"2021-06-20T11:17:03.818065988+08:00"}

# 获取产品信息
$ curl -XGET http://127.0.0.1:8080/v1/products/iphone12
{"username":"colin","name":"iphone12","category":"phone","price":8000,"description":"cannot afford","createdAt":"2021-06-20T11:17:03.818065988+08:00"}
```

- [Middleware with Gin MW](80_app/82_bind-json-mw.go)：使用 Gin 现有的 middleware

### Advanced Topics

#### Cookie

Cookie 实际上就是服务器保存在浏览器上的一段信息。浏览器有了 Cookie 之后，每次向服务器发送请求时都会同时将该信息发送给服务器，服务器收到请求后，就可以根据该信息处理请求。Cookie 由服务器创建，并发送给浏览器，最终由浏览器保存。


- [Set and Get Cookie](90_advance/19_cookie.go)

```shell
go run 19_cookie.go
curl http://127.0.0.1:8080/login  # set the cookie for the browser
curl http://127.0.0.1:8080/home  # check the cookie value
```

#### Session

主要功能是：

- 简单的API：将其用作设置签名（以及可选的加密）cookie的简便方法。
- 内置的后端可将session存储在cookie或文件系统中。
- Flash消息：一直持续读取的session值。
- 切换session持久性（又称“记住我”）和设置其他属性的便捷方法。
- 旋转身份验证和加密密钥的机制。
- 每个请求有多个session，即使使用不同的后端也是如此。
- 自定义session后端的接口和基础结构：可以使用通用API检索并批量保存来自不同商店的session。



#### Async


- [Async](90_advanced/91_async.go)

```shell
curl http://127.0.0.1:8080/sync
curl http://127.0.0.1:8080/async
```

## Ref

1. [Golang 微框架 Gin 简介](https://www.jianshu.com/p/a31e4ee25305)
2. 



