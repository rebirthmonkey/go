# Gin-Swag

## Init
### Installation

```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

### Module Init

```shell
go mod tidy
go mod download
```

### Create Swag Docs
```bash
swag init -g cmd/main.go
```

!!! check cmd/main.go, `import _ "wukong/go/50_web/20_gin/90_swag/docs"`

### run & check

```bash
go run cmd/main.go &
sleep 3
curl http://127.0.0.1:8080/user/123/xxx
curl http://127.0.0.1:8080/docs/index.html
ps aux | grep -i cmd/main |  grep -v grep | awk {'print $2'} | xargs kill -9
```


## Grammar

### General
- // @title: Gin Swag API
- // @version: 1.0
- // @description: This is a Gin server.

- // @contact.name: API Support
- // @contact.url: http://www.xxx.io/support

- // @license.name: Apache 2.0
- // @license.url: http://www.apache.org/licenses/LICENSE-2.0.html

- // @host 127.0.0.1:8080
- // @BasePath /docs/index.html


### API
- // @Summary：接口概要说明
- // @Description：接口详细描述信息
- // @Tags：用户信息，swagger API分类标签, 同一个tag为一组
- // @accept json：浏览器可处理数据类型，浏览器默认发 Accept: */*
- // @Produce  json：设置返回数据的类型和编码
- // @Param id path int true "ID"：url参数：（name；参数类型[query(?id=),path(/123)]；数据类型；required；参数描述）
- // @Param name query string false "name"
- // @Success 200 {object} Res {"code":200,"data":null,"msg":""}：成功返回的数据结构，最后是示例
- // @Success 200 {object} model.User：另一个例子
- // @Failure 400 {object} Res {"code":200,"data":null,"msg":""}
- // @Router /test/{id} [get]：路由信息，一定要写上


