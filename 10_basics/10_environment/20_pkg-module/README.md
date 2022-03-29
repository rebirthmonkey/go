# 包&Module

## 包

### 整体

- repo：一个 repo 一般对应一个 module，但一个 repo 也可以包含多个 module
- module：一个 module 包含多个 pkg，module 用 go.mod 来记录元信息
- pkg：每个 pkg 是一个目录下的多个源文件，往往采用 repo/module/ 的路径（github.com/bigwhite/gocmpp）作为 pkg 的导入路径（import path）

### 包结构

包指一个目录下的一组源文件，但这组源文件的包声明必须是一致的包名。所有语法可见性均定义在包这个级别，也就是一个目录/包下不同的源文件内的函数对于包来说同级，不会区分不同文件。

- 每个包开始于一个与目录同名的.go文件：如 http 包应该在 http/ 目录下的 http.go 文件中定义
- 不同的功能分布在不同的源文件中：如 message.go 包含 Request 和 Response 类型，负责 HTTP 序列化请求和响应；http.go 应该包含底层网络处理逻辑；client.go 包含 Client 类型，server.go 包含 Server 类型，实现 HTTP 业务逻辑、请求路由；

#### 导入路径 vs. 包名

包名可以与目录名不同，但尽量相同。

如在使用实时分布式消息平台 nsq 提供的 go client api 时，导入的路径为 `import “github.com/bitly/go-nsq”`。但在使用其提供的函数时，却直接用 nsq 做前缀包名 `q, _ := nsq.NewConsumer("write_test", "ch", config)`。因为 import 导入的是包在 $GOPATH 下的目录路径，并非包名，而在代码中使用的是真正的包名。【1】

#### 多源文件依赖

同一个目录下、相同包声明的多个源文件属于同一个包，当源文件间存在依赖关系时，需要在 go run/build 时指明多个文件。

- go run

```bash
cd samelib
go run main.go lib1.go -name="XXX" # 需要把涉及到的文件一起运行
```

- go build

```bash
cd samelib
go build # 会自动在同目录下寻找依赖函数/文件一起编译
./samelib -name="XXX"
```

### 作用域

名称的首字母为大写的变量或函数才可以被当前包外的代码引用（公有），否则它就只能被当前包内的其他代码引用（私有）。

- 小写：私有
- 大写：公有

### 分类

#### 来源

- 标准包：Go 自带的包
- 远程包：github 上的包，如：[github.com/user/hello](http://github.com/user/hello)

#### 类型

##### 库包

注意：以下放在在新版 Go 中已被 GO Module 机制替代，已无法运行。

```shell
cp -r wkpkg $GOPATH/src # 把 wkpkg 复制到 $GOPATH/src 目录下
go install -x wkpkg # build wkpgk.a 并归档文件到 $GOPATH/pkg/darwin_amd64 或 $GOROOT/.. 目录下，可直接被其他程序 import
```

##### 可执行包

注意：以下放在在新版 Go 中已被 GO Module 机制替代，已无法运行。

wkapp 依赖 wkpkg

```shell
cp -r wkapp $GOPATH/src # 把 wkapp 复制到 $GOPATH/src 目录下
# 在 GoLand 中 wkpkg 依赖无法解决，需要 `sudo ln -s $GOPATH/src/wkpkg $GOROOT/src/wkpkg`，因为 GoLand 包依赖搜索只寻找 GOROOT
go install -x wkapp # build 可执行文件 wkapp 并放在 $GOPATH/bin 目录下
```

## Module

### 原理

GO MODULE 的目的是确保在非 `$GOPATH/src` 目录下可以编译 go 源文件，并把结果输出到 $GOPATH 目录下。

在 module-aware mode 下，repo 的顶层目录下会放置一个 go.mod 文件，该文件定义把该 repo 定义为一个 Go module，而放置 go.mod 文件的目录被称为 module root 目录（通常一个 repo 对应一个 module，但不是必须的）。module root 目录以及其子目录下的所有 pkg 均归属于该 module，除了那些自身包含 go.mod 文件的子目录。

在 module-aware mode 下，go 编译器将不再在 $GOPATH/src 或 $GOPATH/pkg 下搜索目标程序依赖的第三方 pkg，而是主动将下载的依赖包缓存在 `$GOPATH/pkg/mod` 下。

#### go.mod

- 用来定义 module 名
- 确认 go 的版本
- 管理该 module 的所有依赖包及版本

执行 go build 后go.mod文件的内容为：

```bash
# cat go.mod
module wkmodule

go 1.17

require github.com/sirupsen/logrus v1.8.1

require golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
```

该 module 并没有直接依赖 d 包，因此在 d 的记录后面通过注释形式标记了 indirect，即非直接、传递依赖。

#### 导入路径

module 作为包的导入路径，一个 module 下可以包含多个包，所以可以通过不同层级的子目录来区分包。

如 `import "github.com/cnych/stardust/api/stringsx"`

- module名：stardust
- 内部路径：api
- 包名：stringsx

### 命令

#### 配置

Go 在 1.13 版本后是默认开启 MODULE, 代表当项目在 GOPATH/src 外且项目根目录有 go.mod 文件时，开启 go module。也就是说，如果不把代码放置在 GOPATH/src 下则默认使用 MODULE 管理.

```bash
go env -w GOPROXY=https://mirrors.cloud.tencent.com/go/,direct
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
go env -w GO111MODULE=on
```

#### 命令

- go mod init：创建 go.mod 文件
- go mod tidy：更新/细化 go.mod 文件，它会check go.mod 中所有依赖的 pkg 并下载到 `$GOPATH/pkg/mod/`下
- go mod downlowd（不需要）：下载所有依赖包到`$GOPATH/pkg/mod/` 下
- go mod vendor：将 `$GOPATH/pkg/mod/` 下的依赖包复制到本地 vendor/ 目录下

## Lab

### 新建mod/可执行-无本地依赖

所有依赖为外部包

```bash
cd wkmodule
go mod init wkmodule # 创建go.mod
go mod tidy # 检测该目录下所有引入的依赖，并放入 go.mod中，并自动下周所有依赖包至$GOPATH/pkg/mod
go run main.go
go install  # 安装wkmodule可执行文件到 $GOPATH/bin
```

- mod更新依赖：在 go.mod 所在目录

```bash
go mod tidy
```

- mod使用指定版本

```bash
go mod -require=bitbucket.org/bigwhite/c@v1.0.0 # 会更新go.mod reqire部分
```

#### 基于vendor编译（推荐）

可以不使用 $GOPATH/pkg/mod/ 而使用本地 vendor/ 来存放依赖包。Go Modules 默认会忽略 vendor/ 这个目录，但如果希望将依赖包都放入 vendor/，可以执行 `go mod vendor`。vender/ 目录会包含所有的依赖模块代码，并且会在该目录下面添加一个名为 modules.txt 的文件用来记录依赖包的一些信息。

```bash
go mod vendor # 将刚才下载至 $GOPATH/mod/pkg 的依赖包转移至该项目目录的 vendor 目录下
# go build -getmode=vendor main.go
go build -ldflags "-s -w" -a -installsuffix cgo -o wkmodule-local .
```

### 下载mod/可执行-无本地依赖

所有依赖为外部包，无本地依赖

```bash
cd wkapp
go mod tidy
go run main.go
```

### 下载mod/可执行-有本地依赖

改为本地依赖

````bash
cd wkapp2
go mod init wkapp # 创建go.mod
修改 go.mod，加入本地依赖
```
require (
	github.com/rebirthmonkey/wklib2 v0.0.0
)

replace github.com/rebirthmonkey/wklib2 => /Users/ruan/workspace/go/10_basics/10_environment/20_pkg-module/wklib2

```
go mod tidy 
go run main.go
````

### 下载mod/库-被依赖

如果下载的库是被其他程序 import，则：

- 公开：通过 import github.com/xxx/ 被导入
- 本地：通过 go.mod 中的 replace xxx=>yyy 被导入

## Ref

1. [理解Golang包导入](https://studygolang.com/articles/3189)
1. [亲测GO环境搭建，理解go build、go install、go get](https://blog.csdn.net/zhangliangzi/article/details/77914943)
1. [go module 基本使用](https://www.cnblogs.com/chnmig/p/11806609.html)
1. [一文搞懂 Go Modules 前世今生及入门使用](https://www.cnblogs.com/wongbingming/p/12941021.html)
1. [关于Go Modules，看这一篇文章就够了](https://zhuanlan.zhihu.com/p/105556877?utm_source=wechat_session)
1. [初窥Go module](https://tonybai.com/2018/07/15/hello-go-module/)
1. [Go 语言全新依赖管理系统 Go Modules 使用详解](https://www.toutiao.com/i6714564564194689543/?tt_from=weixin&utm_campaign=client_share&wxshare_count=1&timestamp=1597249873&app=news_article&utm_source=weixin&utm_medium=toutiao_ios&use_new_style=1&req_id=2020081300311301001405309209375D1B&group_id=6714564564194689543)
1. 

