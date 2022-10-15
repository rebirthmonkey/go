# net/http

## 请求/响应信息

### 通用

HTTP1.1 中，请求和响应都由以下 4 部分组成，两者之间格式的区别是开始行不同。

1. 开始行：在请求中叫请求行，在响应中叫状态行。
   - 请求行：`请求方法 URI 协议/版本`，如 `GET "http://images.chingli.com/past/logo.gif HTTP/1.1`
   - 响应行： `协议版本 状态代码 状态描述`，如 `HTTP/1.1 200 OK`
2. 头：零或多行，包含一些额外信息来说明浏览器、服务器以及后续正文的一些信息。
3. 空行。
4. 正文：客户端提交或服务器返回的信息，请求和响应中都可以没有此部分。

开始行和头的各行必须以 `<CR><LF>` 作为结尾。空行内必须只有 `<CR><LF>` 而无其他空格。在 HTTP/1.1 协议中，开始行和头都是以 ASCII 编码的纯文本，所有的请求头，除 `Host` 外，都是可选的。

### Request

HTTP 请求由客户端发来，Web 程序要做的首先就是分析这些请求，并用 Go 语言中响应的数据对象来表示。在 `net/http` 包中，用 `Request` 结构体表示 HTTP 请求：

```go
type Request struct {
    Method string
    URL *url.URL
    Proto      string // "HTTP/1.0"
    ProtoMajor int    // 1
    ProtoMinor int    // 0
    Header Header
    Body io.ReadCloser
    ContentLength int64
    TransferEncoding []string
    Close bool
    Host string
    Form url.Values
    PostForm url.Values
    MultipartForm *multipart.Form
    Trailer Header
    RemoteAddr string
    RequestURI string
    TLS *tls.ConnectionState
    Cancel <-chan struct{}
}
```

### Response

当收到并理解了请求之后，就需要根据相应的处理逻辑构建响应。`net/http` 包中，用 `Response` 结构体表示响应：

```go
type Response struct {
    Status     string // e.g. "200 OK"
    StatusCode int    // e.g. 200
    Proto      string // e.g. "HTTP/1.0"
    ProtoMajor int    // e.g. 1
    ProtoMinor int    // e.g. 0
    Header Header
    Body io.ReadCloser
    ContentLength int64
    TransferEncoding []string
    Close bool
    Trailer Header
    Request *Request
    TLS *tls.ConnectionState
}
```

## Handler

编写 Web 程序的主要工作就是编写各种实现该 `Handler` 接口的类型，并在该类型的 `ServeHTTP` 方法中编写服务器响应逻辑（path <--> handler mapping）。编写的 Web 服务器程序可能主要就是由各种各样的 `fooHandler` 、`barHandler` 构成，Handler 接口就成为 `net/http` 包中最重要的东西。可以说，每个 `Handler` 接口的实现就是一个小的 Web 服务器。

## Server

### 结构

```go
type Server struct {
    Addr           string        // TCP address to listen on, ":http" if empty
    Handler        Handler       // handler to invoke, http.DefaultServeMux if nil
    ReadTimeout    time.Duration // maximum duration before timing out read of the request
    WriteTimeout   time.Duration // maximum duration before timing out write of the response
    MaxHeaderBytes int           // maximum size of request headers, DefaultMaxHeaderBytes if 0
    TLSConfig      *tls.Config   // optional TLS config, used by ListenAndServeTLS
    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
    ConnState func(net.Conn, ConnState)
    ErrorLog *log.Logger
    disableKeepAlives int32     // accessed atomically.
    nextProtoOnce     sync.Once // guards initialization of TLSNextProto in Serve
    nextProtoErr      error
}

func (srv *Server) ListenAndServe() error
func (srv *Server) Serve(l net.Listener) error
func (srv *Server) SetKeepAlivesEnabled(v bool)
```

### ListenAndServe()

ListenAndServe 在 TCP 网络地址 srv.Addr 上监听接入连接，并通过 Serve 方法处理连接。连接被接受后，使 TCP 保持连接。如果 srv.Addr 为空，则默认使用 ":http"。ListenAndServe 返回的 error 始终不为 nil。

### Serve()

Serve 在 net.Listener 类型的 l 上接受接入连接，为每个连接创建一个新的服务 goroutine。该 goroutine 读请求并调用 srv.Handler 以进行响应。同 ListenAndServe 一样，Serve 返回的 error 也一直不为 nil。

## 路由

### ServeMux

ServeMux 是 net/http 包自带的固定 HTTP 请求多路复用器（路由器），它包含一个映射列表，每个列表项主要将特定的 URL 模式与特定的 Handler 对应。因为其匹配能力不强，所以不适用于动态路由。

```go
type ServeMux struct {
    mu    sync.RWMutex
    m     map[string]muxEntry
    hosts bool // whether any patterns contain hostnames
}

func (mux *ServeMux) Handle(pattern string, handler Handler)
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

## Lab

- [Handler](10_handler.go)

```bash
go run 10_handler.go &
sleep 3
curl http://127.0.0.1:8080
ps aux | grep -i 10_handler |  grep -v grep | awk {'print $2'} | xargs kill -9
```

- [Handler](12_handler.go)

```bash
go run 12_handler.go &
sleep 3
curl http://127.0.0.1:8080/view/abc
ps aux | grep -i 12_handler |  grep -v grep | awk {'print $2'} | xargs kill -9
```

> 对于`/view/<NAME>`路径，该程序会展示`<NAME>.txt`文件的内容，如果路径下不存在`<NAME>.txt`，将会报错。

- [Handler with Interface](14_handler-interface.go)

```bash
go run 14_handler-interface.go &
sleep 3
curl http://127.0.0.1:8080/view/abc
ps aux | grep -i 14_handler-interface |  grep -v grep | awk {'print $2'} | xargs kill -9
```

- [Handler as Reverse Proxy](20_reverse-proxy.go)

```bash
go run 20_reverse-proxy.go &
sleep 3
curl http://127.0.0.1:8080
ps aux | grep -i 20_reverse-proxy |  grep -v grep | awk {'print $2'} | xargs kill -9
```

## Ref

1. [理解 Go 语言 Web 编程](https://chingli.com/coding/understanding-go-web-app.html)
