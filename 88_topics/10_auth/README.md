# 认证

## 简介

### 认证 vs. 授权

认证是证明了你是谁，而授权是决定了你能做什么。

- 认证（Authentication，英文缩写  authn）：用来验证某个用户是否具有访问系统的权限。如果认证通过，该用户就可以访问系统，从而创建、修改、删除、查询平台支持的资源。
- 授权（Authorization，英文缩写 authz）：用来验证某个用户是否具有访问某个资源的权限。如果授权通过，该用户就能对资源做增删改查等操作。

例如，用户 james、colin、aaron 分别创建了 Product-A、Product-B、Product-C。现在用户 colin 通过用户名和密码（认证）成功登录到仓库系统中，但他尝试访问 Product-A、Product-C 失败，因为这两个产品不属于他（授权失败），但他可以成功访问自己创建的产品 Product-B（授权成功）。由此可见：**认证证明了你是谁，授权决定了你能做什么**。

<img src="figures/image-20220914155207130.png" alt="image-20220914155207130" style="zoom:50%;" />

## 方式

前常见的四种认证方法：Basic、Digest、OAuth、Bearer。

### Basic

Basic 是最简单的认证方式，它简单地将“用户名:密码”进行 base64 编码后，放到 HTTP Authorization Header 中。HTTP 请求到达后端服务后，后端服务会解析出 Authorization Header 中的 base64 字符串，解码获取用户名和密码，并将用户名和密码跟数据库中记录的值进行比较，如果匹配则认证通过。

通过 base64 编码，可以将密码以非明文的方式传输，增加一定的安全性。但 base64 不是加密技术，入侵者仍然可以截获 base64 字符串，并反编码获取用户名和密码。另外，即使 Basic 认证中密码被加密，入侵者仍可通过加密后的用户名和密码进行重放攻击。所以，Basic 认证虽然简单，但极不安全。使用 Basic 认证的唯一方式就是将它和 SSL 配合使用，来确保整个认证过程是安全的。

#### Lab

```shell
cd 10_basic
go run example.go auth.go basic.go
```

```shell
basic=`echo -n 'admin:admin'|base64`
curl -XGET -H "Authorization: Basic ${basic}" http://127.0.0.1:8080/ping
```

### Digest

Digest  是另一种 HTTP 认证协议，它与 Basic 兼容，但修复了基本认证的严重缺陷。Digest 具有如下特点：

- 绝不会用明文方式在网络上发送密码
- 可以有效防止恶意用户进行重放攻击
- 可以有选择地防止对报文内容的篡改

Digest 认证分以下四步：

- 客户端请求服务端的资源。
- 在客户端能够证明它知道密码从而确认其身份之前，服务端认证失败，返回 401 Unauthorized，并返回 WWW-Authenticate 头，里面包含认证需要的信息。
- 客户端根据 WWW-Authenticate 头中的信息，选择加密算法，并使用密码随机数 nonce，计算出 digest response，并再次请求服务端。
- 服务器将客户端提供的 digest 与服务器内部计算出的 digest 进行对比。如果匹配，说明客户端知道密码，认证通过，并返回一些与授权会话相关的附加信息，放在 Authorization-Info 中。

<img src="figures/image-20220914155242235.png" alt="image-20220914155242235" style="zoom:50%;" />

WWW-Authenticate 头中包含的信息见下表：

| **字段名**          | **说明**                                                 |
| ------------------- | -------------------------------------------------------- |
| username            | 用户名                                                   |
| realm               | 服务器返回的realm，一般是域名                            |
| method              | HTTP请求方法                                             |
| nonce               | 服务器发给客户端的随机字符串                             |
| nc(nonceCount)      | 请求的次数，用于标记、计数，防止重放攻击                 |
| cnonce(clinetNonce) | 客户端发送给服务器的随机字符串，用于客户端对服务器的认证 |
| qop                 | 保护质量参数，一般是auth或auth-int，这会影响摘要的算法   |
| uri                 | 请求的uri                                                |
| response            | 客户端根据算法算出的密码摘要值                           |

虽然使用 digest 可以避免密码以明文方式发送，一定程度上保护了密码的安全性，但是仅仅隐藏密码并不能保证请求是安全的。因为请求（包括 digest）仍然可以被截获，这样就可以重放给服务器，带来安全问题。为了防止重放攻击，服务器向客户端发送了密码随机数 nonce，nonce 每次请求都会变化。客户端会根据 nonce 生成 diget，这种方式可以使摘要随着随机数的变化而变化。服务端收到的密码摘要只对特定的随机数有效，而没有密码的话，攻击者就无法计算出正确的 digest，这样就可以防止重放攻击。digest 认证可以保护密码，比基本认证安全很多。但 digest 认证并不能保护内容，所以仍然要与 HTTPS 配合使用，来确保通信的安全。

### OAuth

OAuth（开放授权）是一个开放的授权标准，允许用户让第三方应用访问该用户在某一 Web 服务上存储的私密资源（例如照片、视频、音频等），而无需将用户名和密码提供给第三方应用。OAuth 目前的版本是 2.0  版，OAuth2.0  一共分为四种授权方式，分别为密码式、隐藏式、拼接式和授权码模式。

#### 密码式

密码式的授权方式，就是用户把用户名和密码直接告诉给第三方应用，然后第三方应用使用用户名和密码换取令牌。所以，使用此授权方式的前提是无法采用其他授权方式，并且用户高度信任某应用。认证流程如下：

- 网站 A 向用户发出获取用户名和密码的请求；
- 用户同意后，网站 A 凭借用户名和密码向网站 B 换取令牌；
- 网站 B 验证用户身份后，给出网站 A 令牌，网站 A 凭借令牌可以访问网站 B 对应权限的资源。

#### 隐藏式

这种方式适用于前端应用，认证流程如下：

- 网站 A 提供一个跳转到网站 B 的链接，用户点击后跳转至网站 B，并向用户请求授权；
- 用户登录网站 B，同意授权后，跳转回网站 A 指定的重定向 redirect_url 地址，并携带网站 B 返回的令牌，从而使用户在网站 B 的数据可以给网站 A 使用。

这个授权方式存在着“中间人攻击”的风险，因此只能用于一些安全性要求不高的场景，并且令牌的有效时间要非常短。

#### 凭借式

这种方式是在命令行中请求授权，适用于没有前端的命令行应用。认证流程如下：

- 应用 A 在命令行向应用 B 请求授权，此时应用 A 需要携带应用 B 提前颁发的 secretID 和 secretKey，其中 secretKey 出于安全性考虑，需在后端发送；
- 应用 B 接收到 secretID 和 secretKey，并进行身份验证，验证通过后返回给应用 A 令牌。

#### 授权码模式

这种方式就是第三方应用先提前申请一个授权码，然后再使用授权码来获取令牌。相对来说，这种方式安全性更高，前端传送授权码，后端存储令牌，与资源的通信都是在后端，可以避免令牌的泄露导致的安全问题。认证流程如下：

<img src="figures/image-20220914155315482.png" alt="image-20220914155315482" style="zoom:50%;" />

- 网站 A 提供一个跳转到网站 B 的链接 +redirect_url，用户点击后跳转至网站 B；
- 用户携带向网站 B 提前申请的 client_id，向网站 B 发起身份验证请求；
- 用户登录网站 B，通过验证，授予网站 A 权限，此时网站跳转回 redirect_url，其中会有网站 B 通过验证后的授权码附在该 url 后；
- 网站 A 携带授权码向网站 B 请求令牌，网站 B 验证授权码后，返回令牌即 access_token。

### Bearer

Bearer 认证，也被称为令牌认证，是一种 HTTP 身份验证方法。Bearer 认证的核心是 bearer token，它是一个加密字符串，通常由服务端根据密钥生成。客户端在请求服务端时，必须在请求头中包含 `Authorization: Bearer <token>`。服务端收到请求后，解析出 `<token>`，并校验 `<token>` 的合法性。如果校验通过，则认证通过。跟基本认证一样，Bearer 认证需要配合 HTTPS 一起使用，来保证认证安全性。

#### JWT

在典型的场景中，为了区分用户和保证安全，必须对 API 请求进行鉴权，但是不能要求每一个请求都进行登录操作。合理做法是，在第一次登录之后产生一个有一定有效期的 token，并将它存储在浏览器的 Cookie 或 LocalStorage 中。之后的请求都携带这个 token，请求到达服务器端后，服务器端用这个 token 对请求进行认证。在第一次登录之后，服务器会将这个 token 用文件、数据库或缓存服务器等方法存下来，用于之后请求中的比对。也可以采用更简单的方法：直接用密钥来签发 Token。这样可以省下额外的存储，也可以减少每一次请求时对数据库的查询压力。这种方法在业界的标准实现方式就是 JWT。

JWT 是 Bearer Token 的一个具体实现，由 JSON 数据格式组成，通过 Hash 散列算法生成一个字符串，该字符串可以用来进行授权和信息交换。使用 JWT Token 进行认证有很多优点，比如说无需在服务端存储用户数据，可以减轻服务端压力；而且采用 JSON 数据格式，比较易读。除此之外，使用 JWT  Token 还有跨语言、轻量级等优点。

用 JWT Token 进行认证具体可以分为 4 步：

- 客户端使用用户名和密码请求登录。
- 服务端收到请求后，会去验证用户名和密码。如果用户名和密码跟数据库记录不一致，则验证失败。如果一致则验证通过，服务端会签发一个 JWT Token 返回给客户端。
- 客户端收到回复后会将 Token 缓存起来，比如放在浏览器 Cookie 中或 LocalStorage 中，之后每次请求都会携带该 Token。
- 服务端收到请求后，会验证请求中的 Token，验证通过则进行业务逻辑处理，处理完后返回处理后的结果。

<img src="figures/image-20220914155333949.png" alt="image-20220914155333949" style="zoom:50%;" />

##### Header

JWT 由三部分组成，分别是 Header、Payload 和 Signature，它们之间用圆点.连接，例如：

```shell
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2NDI4NTY2MzcsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2MzUwODA2MzcsInN1YiI6ImFkbWluIn0.Shw27RKENE_2MVBq7-c8OmgYdF92UmdwS8xE-Fts2FM
```

<img src="figures/image-20220914155350215.png" alt="image-20220914155350215" style="zoom:50%;" />

JWT Token 的 Header 中包含两部分信息：一是 Token 的类型，二是 Token 所使用的加密算法。例如：

```shell
{
  "typ": "JWT",  
  "alg":  "HS256"
}
```

- typ：说明 Token 类型是 JWT。
- alg：说明 Token 的加密算法，这里是 HS256。

这里需要将 Header 进行 base64 编码：

```shell
$ echo -n '{"typ":"JWT","alg":"HS256"}'|base64
eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9
```

在某些场景下，可能还会有 kid 选项，用来标识一个密钥 ID，例如：

```shell
{    
  "alg": "HS256",
  "kid": "XhbY3aCrfjdYcP1OFJRu9xcno8JzSbUIvGE2",
  "typ": "JWT"
}
```

##### Payload

Payload 中携带 Token 的具体内容由 3 部分组成：JWT 标准中注册的声明（可选）、公共的声明、私有的声明。

标准中注册的声明：

| **字段名**           | **描述**                                                     |
| -------------------- | ------------------------------------------------------------ |
| iss(Issuer)          | JWT Token的签发者，其值应为大小写敏感的字符串或Uri           |
| sub(Subject)         | 主题，sub可以用来鉴别一个用户                                |
| exp(Expiration Time) | JWT Token过期时间                                            |
| aud(Audience)        | 接收JWT  Token的一方，其值应为大小写敏感的字符串或Uri，一般可以为特定的App、服务或模块。服务器端的安全策略在签发时和验证时，aud必须是一致的 |
| iat(Issued  At)      | JWT Token签发时间                                            |
| nbf(Not  Before)     | JWT Token生效时间                                            |
| Jti(JWT  ID)         | JWT Token ID，令牌的唯一标志符，通常用于一次性消费的Token    |

本例中的 payload 内容为：

```shell
{
  "aud": "iam.authz.marmotedu.com",
  "exp": 1604158987,
  "iat": 1604151787,
  "iss": "iamctl",
  "nbf": 1604151787
}
```

将 Payload 进行 base64 编码：

```shell
$ echo -n '{"aud":"iam.authz.marmotedu.com","exp":1604158987,"iat":1604151787,"iss":"iamctl","nbf":1604151787}'|base64
eyJhdWQiOiJpYW0uYXV0aHoubWFybW90ZWR1LmNvbSIsImV4cCI6MTYwNDE1ODk4NywiaWF0Ijox
NjA0MTUxNzg3LCJpc3MiOiJpYW1jdGwiLCJuYmYiOjE2MDQxNTE3ODd9
```

除此之外，还有公共的声明和私有的声明。公共的声明可以添加任何的需要的信息，一般添加用户的相关信息或其他业务需要的信息，注意不要添加敏感信息；私有声明是客户端和服务端所共同定义的声明，因为 base64 是对称解密的，所以一般不建议存放敏感信息。

##### Signature

Signature 是 Token 的签名部分，通过如下方式生成：将 Header 和 Payload 分别 base64 编码后，用 . 连接。然后再使用 Header 中声明的加密方式，利用 secretKey 对连接后的字符串进行加密，加密后的字符串即为最终的  Signature。secretKey 是密钥，保存在服务器中，一般通过配置文件来保存。这里要注意，密钥一定不能泄露。密钥泄露后，入侵者可以使用该密钥来签发 JWT Token，从而入侵系统。最后生成的 Token 如下：

```shell
eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJpYW0uYXV0aHoubWFybW90ZWR1LmNvbSIsImV4cCI6MTYwNDE1ODk4NywiaWF0IjoxNjA0MTUxNzg3LCJpc3MiOiJpYW1jdGwiLCJuYmYiOjE2MDQxNTE3ODd9.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ
```

签名后服务端会返回生成的 Token，客户端下次请求会携带该 Token。服务端收到 Token 后会解析出 header.payload，然后用相同的加密算法和密钥对 header.payload 再进行一次加密，得到 Signature。并且，对比加密后的 Signature 和收到的 Signature 是否相同，如果相同则验证通过，不相同则返回 HTTP 401 Unauthorized 错误。

#### Lab

```shell
cd 20_jwt
go run example.go auth.go jwt.go
```

- 通过用户名、密码 获取 JWT

```shell
curl -XPOST -H'Content-Type: application/json' -d'{"username":"admin","password":"admin"}' http://127.0.0.1:8080/login/jwt 
{"code":200,"expire":"2022-04-30T18:23:43+08:00","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTEzMTQyMjMsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY1MTMxMDYyM30.AugP8KBMBD7nOmEi03-JKBZ5v1Oo18MGVFE5HpgCS9I"}

jwt=`echo -n 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTEzMTQyMjMsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY1MTMxMDYyM30.AugP8KBMBD7nOmEi03-JKBZ5v1Oo18MGVFE5HpgCS9I'`
```

- 使用 JWT 认证、并获取 REST 资源

```shell
curl -XGET -H "Content-Type: application/json" -H "Authorization: Bearer ${jwt}"  http://127.0.0.1:8080/ping/
```



## Ref

1. [JWT Token](https://github.com/appleboy/gin-jwt)
