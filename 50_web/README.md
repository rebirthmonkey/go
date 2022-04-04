# Web 框架

一个 Web  程序的编写往往要涉及更多的方面，目前各种各样的中间件能够完成一些任务。但许多时候，我们总是希望他人帮我们完成更多的事情，于是就产生了许多的 Web 框架。根据架构的不同，这些框架大致可分为两大类：

- 微架构型框架：其核心框架只提供很少的功能，而更多的功能则需要组合各种中间件来提供，因此这种框架也可称为混搭型框架。它相当灵活，但相对来说需要使用者在组合使用各种中间件时花费更大的力气。像 [Echo](https://github.com/labstack/echo)、[Goji](https://github.com/goji/goji)、[Gin](https://github.com/gin-gonic/gin) 等都属于微架构型框架。
- 全能型架构：它基本上提供了编写 Web 应用时需要的所有功能，因此更加重型，多数使用 MVC 架构模式设计。在使用这类框架时感觉更轻省，但其做事风格一般不同于 Go 语言惯用的风格，也较难弄明白这些框架是如何工作的。像 [Beego](http://beego.me/)、[Revel](http://revel.github.io/) 等就属于全能型架构。

对于究竟该选择微架构还是全能型架构，仍有较多的争议。像 [The Case for Go Web Frameworks](https://medium.com/@richardeng/the-case-for-go-web-frameworks-a791fcd79d47#.7qe9n08aw) 一文就力挺全能型架构，并且其副标题就是“Idiomatic Go is not a religion”，但该文也收到了较多的反对意见，见[这里](https://groups.google.com/forum/#!searchin/golang-nuts/framework/golang-nuts/vX086U_49Qo/KLXcyKwVil4J)和[这里](https://www.reddit.com/r/programming/comments/2jsrsq/the_case_for_go_web_frameworks_idiomatic_go_is/)。总体上来说，Go 语言社区已越来越偏向使用微架构型框架，当将来 `context` 包进入标准库后，`http.Handler` 本身就定义了较完善的中间件编写规范，这种使用微架构的趋势可能更加明显，并且各种微架构的实现方式有望进一步走向统一，这样其实 `http` 包就是一个具有庞大生态系统的微架构框架。

## 参数类型

HTTP 具有以下 5 种参数类型：

- 路径参数（path）：例如 `gin.Default().GET("/user/:name", nil)`， name 就是路径参数。
- 查询字符串参数（query）：例如 `/welcome?firstname=Wukong&lastname=Sun`，firstname 和 lastname 就是查询字符串参数。
- 表单参数（form）：例如 `curl -X POST -F 'username=colin' -F  'password=colin1234' http://rebirthmonkey.com/login`，username 和 password 就是表单参数。
- HTTP 头参数（header）：例如 `curl -X POST -H 'Content-Type:  application/json' -d '{"username":"colin","password":"colin1234"}'  http://mydomain.com/login`，Content-Type 就是 HTTP 头参数。
- 消息体参数（body）：例如 `curl -X POST -H 'Content-Type: application/json' -d  '{"username":"colin","password":"colin1234"}'  http://mydomain.com/login`，username 和 password 就是消息体参数。



## Web 服务的核心功能

<img src="figures/1a6d38450cdd0e115e505ab30113602e.jpg" alt="img" style="zoom: 33%;" />

### 路由匹配

Web 服务最核心的功能是路由匹配，其实就是根据（HTTP方法、请求路径）匹配到处理这个请求的函数，最终由该函数处理这次请求，并返回结果。一次 HTTP 请求经过路由匹配，最终将请求交由 Delete(c *gin.Context) 函数来处理。变量 c 中存放了这次请求的参数，在 Delete 函数中，可以进行参数解析、参数校验、逻辑处理，最终返回结果。

<img src="figures/1f5yydeffb32732e7d0e23a0a9cd369d.jpg" alt="img" style="zoom:33%;" />

### 路由分组

对于大型系统，可能会有很多个 API 接口，API 接口随着需求的更新迭代，可能会有多个版本，为了便于管理，需要对路由进行分组。

### 一进程多服务

有时候，需要在一个服务进程中，同时开启 HTTP 服务的 80 端口和 HTTPS 的 443 端口，这样就可以做到：对内的服务，访问 80 端口，简化服务访问复杂度；对外的服务，访问更为安全的 HTTPS 服务。显然，我们没必要为相同功能启动多个服务进程，所以这时候就需要 Web 服务能够支持一进程多服务的功能。

### 业务处理

开发 Web 服务最核心的诉求是：输入一些参数，校验通过后，进行业务逻辑处理，然后返回结果。所以 Web 服务还应该能够进行参数解析、参数校验、逻辑处理、返回结果。

### 中间件

在进行 HTTP 请求时，经常需要针对每一次请求都设置一些通用的操作，比如添加 Header、添加 RequestID、统计请求次数等，这就要求 Web 服务能够支持中间件特性。

### 认证

为了保证系统安全，对于每一个请求都需要进行认证。Web 服务中，通常有两种认证方式，一种是基于用户名和密码，一种是基于 Token。

### RequestID

为了方便定位和跟踪某一次请求，需要支持 RequestID，定位和跟踪 RequestID 主要是为了排障。

### 跨域

当前的软件架构很多采用了前后端分离的架构。在前后端分离的架构中，前端访问地址和后端访问地址往往是不同的，浏览器为了安全，会针对这种情况设置跨域请求，所以 Web 服务需要能够处理浏览器的跨域请求。

## net/http

- [net/http](10_net-http/README.md)

## Gin

- [Gin](20_gin/README.md)

## gRPC

- [gRPC](30_grpc/README.md)