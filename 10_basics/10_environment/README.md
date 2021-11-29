# 包管理

## 简介
Golang使用包（package）这种语法元素来组织源码，所有语法可见性均定义在package这个级别。一个目录下不同 go 文件内的函数对于包来说同级，不会区分不同文件。比如在使用实时分布式消息平台 nsq 提供的 go client api 时，导入的路径如下：

   import “github.com/bitly/go-nsq”

但在使用其提供的函数时，却直接用 nsq 做前缀包名：

   q, _ := nsq.NewConsumer("write_test", "ch", config)

import 后面的最后一个元素应该是包在 $GOPATH 下的目录路径，并非包名。而在代码中使用的是真正的包名。【1】

## GOPATH Workspace
Go 的 workspace 被定义在 $GOPATH 工作目录下，其结构有三个子目录（需要自行创建）：
- src：存放源代码文件
- pkg：存放编译后的库文件
- bin：存放编译后的可执行文件 

![img](figures/2fdfb5620e072d864907870e61ae5f3c.png)

### Build/Install/Get
- build：在当前目录下编译生成可执行/库文件。go build 会拉取所有引用包的源码，重新编译，而不是直接使用 `$GOPATH/pkg` 里的编译文件。如果在 `$GOROOT` 与`$GOPATH` 下没有找到 import 的源码，就会报错。
- install：先执行 build 构建，然后进行链接操作，并且把结果文件搬运到指定目录。
  - 命令源码文件（package "main"且包含main方法）：会编译生成可执行文件到 $GOPATH/bin 目录下；
  - 库源码文件：会被编译到 `$GOPATH/pkg/$GOOS_$GOARCH` 目录下。 
- get：git clone 到 $GOPATH/src，然后 `go install` 
- clean: 
  - go clean -i：删除对应的库文件或命令文件

注意：在 import 第三方包的时候，当源码和 .a 均已安装的情况下，编译器链接的是源码。
注意：在 import 标准库中的包（如fmt）时，也是必须要源码的。不过与自定义包不同的是，如果未找到源码，不会尝试重新编译标准包，而是在链接时链接已经编译好的`.a`文件。

### Example

Build

- `export GO111MODULE=off`
- `go build hello.go`：在本地验证编译，不会有文件生成

Install

- `cp -r wkpkg $GOPATH/src`：把`wkpkg`复制到 $GOPATH/src 目录下
- `go install -x wkpkg`：build `wkpgk.a` 并归档文件到 $GOPATH/pkg/darwin_amd64 目录下，可以被其他程序 import
- `cp -r wkapp $GOPATH/src`: 把`wkapp`复制到 $GOPATH/src 目录下
- 在 GoLand 中 wkpkg 依赖无法解决，需要 `sudo ln -s $GOPATH/src/wkpkg $GOROOT/src/wkpkg`，因为GoLand包依赖搜索只寻找 GOROOT
- `go install -x wkapp`：build 可执行文件 `wkapp` 并放在 $GOPATH/bin 目录下


## GO Module
GO MODULE 确保在非 `$GOPATH/src` 目录下，可以编译 go 包，并把结果输出到 $GOPATH 目录下。

### Configuration
Go 在 1.13 版本后是默认开启 MODULE, 代表当项目在 GOPATH/src 外且项目根目录有 go.mod 文件时，开启 go module。也就是说，如果你不把代码放置在 GOPATH/src 下则默认使用 MODULE 管理.

```bash
go env -w GOPROXY=https://mirrors.cloud.tencent.com/go/,direct
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
go env -w GO111MODULE=on
```

### 步骤
```bash
cd wkmodule
go mod init wkmodule # 创建 go.mod
go mod tidy # tidy会检测该文件夹目录下所有引入的依赖，创建 go.sum 文件
go mod download # 将依赖全部下载至 $GOPATH 下
go install  # download pkg in $GOPATH/src and build wkmodule in $GOPATH/bin
go mod vendor # 将刚才下载至 GOPATH 下的依赖转移至该项目目录下的 vendor(自动新建) 文件夹下
```

更新依赖
```bash
go mod tidy
go mod download
go mod vendor
```

### 通过 $GOPATH 进行遍历
```bash
git clone XXX
cd XXX
go mod download
go build .
```

### 在本地目录进行编译（推荐）
不使用 $GOPATH/pkg 而使用本地目录 vendor/ 进行编译

```bash
go mod vendor
go build -ldflags "-s -w" -a -installsuffix cgo -o wkmodule-local .
```

## 库源代码
名称的首字母为大写的程序实体才可以被当前包外的代码引用，否则它就只能被当前包内的其他代码引用。

### Lib 在同一包中

同一个目录下、相同包名的文件属于同一个包

- go run
```bash
cd samelib
go run main.go lib1.go -name="XXX"
```

- go build
```bash
cd samelib
go build 
./samelib -name="XXX"
```

### Lib 在不同 Github 包 
```bash
cd gitlib
go mod init
go mod tidy
go run ./main.go
```

### Lib 在不同包的相对路径下
```bash
cd relatedlib/wkmodule3
go mod tidy
go mod download
go install
wkmodule3
```


## Ref

1. [理解Golang包导入](https://studygolang.com/articles/3189)
1. [亲测GO环境搭建，理解go build、go install、go get](https://blog.csdn.net/zhangliangzi/article/details/77914943)
1. [go module 基本使用](https://www.cnblogs.com/chnmig/p/11806609.html)
1. [一文搞懂 Go Modules 前世今生及入门使用](https://www.cnblogs.com/wongbingming/p/12941021.html)
1. [关于Go Modules，看这一篇文章就够了](https://zhuanlan.zhihu.com/p/105556877?utm_source=wechat_session)



