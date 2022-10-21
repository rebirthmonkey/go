# Strategy

## 简介

Strategy 模式定义一组 strategy，将所有 strategy 都封装成一个 interface，从而使它们之间可以互换使用。

在项目开发中，经常要根据不同的场景，采取不同的措施，也就是不同的算法。比如，假设需要对 a、b 这两个整数进行计算，根据条件的不同，需要执行不同的计算方式。可以把所有的操作都封装在同一个函数（算法）中，然后通过 if ... else ... 的形式来调用不同的计算方式，这种方式称之为硬编码。在实际应用中，随着功能和体验的不断增长，需要经常添加/修改算法，这样就需要不断修改已有代码，不仅会让这个函数越来越难维护，还可能因为修改带来一些 bug。所以为了解耦，需要使用 strategy 模式，定义一些独立的类来封装不同的算法，每一个类封装一个具体的算法（即 strategy）中。

### strategy

#### 声明 interface



#### 实现

##### struct

通过 struct 来实现 strategy



##### func

通过 funct 来实现 strategy



### operator

用于在不同的 strategy 之间 switch。



## Lab

- [Strategy](10_strategy/main.go)

```bash
cd 10_strategy
go run main.go
```

- [Auth](30_strategy/30_gin-auth/main.go)：Gin Auth 作为一个最佳实践范例

```shell
cd 30_gin-auth
go run example.go auth.go basic.go jwt.go
# 新建终端，并在新建的终端继续
basic=`echo -n 'admin:admin'|base64`
curl -XGET -H "Authorization: Basic ${basic}" http://127.0.0.1:8080/ping/basic/

curl -XPOST -H'Content-Type: application/json' -d'{"username":"admin","password":"admin"}' http://127.0.0.1:8080/login/jwt 
{"code":200,"expire":"2022-04-30T18:23:43+08:00","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTEzMTQyMjMsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY1MTMxMDYyM30.AugP8KBMBD7nOmEi03-JKBZ5v1Oo18MGVFE5HpgCS9I"}

jwt=`echo -n 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTEzMTQyMjMsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY1MTMxMDYyM30.AugP8KBMBD7nOmEi03-JKBZ5v1Oo18MGVFE5HpgCS9I'`

curl -XGET -H "Content-Type: application/json" -H "Authorization: Bearer ${jwt}"  http://127.0.0.1:8080/ping/jwt/
```

或者在同一个终端窗口中执行

```bash
cd 30_gin-auth
go run example.go auth.go basic.go jwt.go &
sleep 10
basic=`echo -n 'admin:admin'|base64`
curl -XGET -H "Authorization: Basic ${basic}" http://127.0.0.1:8080/ping/basic/

curl -XPOST -H'Content-Type: application/json' -d'{"username":"admin","password":"admin"}' http://127.0.0.1:8080/login/jwt 
{"code":200,"expire":"2022-04-30T18:23:43+08:00","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTEzMTQyMjMsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY1MTMxMDYyM30.AugP8KBMBD7nOmEi03-JKBZ5v1Oo18MGVFE5HpgCS9I"}

jwt=`echo -n 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTEzMTQyMjMsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY1MTMxMDYyM30.AugP8KBMBD7nOmEi03-JKBZ5v1Oo18MGVFE5HpgCS9I'`

curl -XGET -H "Content-Type: application/json" -H "Authorization: Bearer ${jwt}"  http://127.0.0.1:8080/ping/jwt/
```
