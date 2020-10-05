# 包管理

## 简介
Golang使用包（package）这种语法元素来组织源码，所有语法可见性均定义在package这个级别。

一个目录下不同 go 文件内的函数对于包来说同级，不会区分不同文件。

比如在使用实时分布式消 息平台nsq提供的go client api时，导入的路径如下：

   import “github.com/bitly/go-nsq”

但在使用其提供的函数时，却直接用nsq做前缀包名：

   q, _ := nsq.NewConsumer("write_test", "ch", config)

import后面的最后一个元素应该是包在 $GOPATH 下的目录路径，并非包名。
而在代码中使用的是真正的包名。【1】

## GOPATH Workspace
Go的workspace被定义在$GOPATH工作目录下，其结构有三个子目录（需要自行创建）：
- src：存放源代码文件
- pkg：存放编译后的文件
- bin：存放编译后的可执行文件 

![img](figures/2fdfb5620e072d864907870e61ae5f3c.png)

### Install & Build & Get
- build: 在当前目录下编译生成可执行文件。go build指令会调用所有引用包的源码，重新编译，而不是直接使用 $GOPATH/pkg 里的编译后文件，如果在【$GOROOT】与【$GOPATH】下没有找到import引入包的项目源码，就会报错。
  - 库源码文件：那么操作后产生的结果文件只会存在于临时目录中。这里的构建的主要意义在于检查和验证。
  - 命令源码文件：那么操作的结果文件会被搬运到源码文件所在的目录中。
- install：安装操作会先执行构建，然后还会进行链接操作，并且把结果文件搬运到指定目录。
  - 如果为命令源码文件（package "main"且包含main方法）:会编译生成可执行文件到 $GOPATH/bin 目录下；
  - 如果为库源码文件：会被编译到 $GOPATH/pkg/$GOOS_$GOARCH 目录下。 
- get: git clone到 $GOPATH/src，然后 `go install` 

注意：**在import第三方包的时候，当源码和.a均已安装的情况下，编译器链接的是源码。**
注意：在import标准库中的包（如fmt）时，也是必须要源码的。不过与自定义包不同的是，如果未找到源码，不会尝试重新编译标准包，而是在链接时链接已经编译好的`.a`文件。

### Example
- `cp -r wkpkg $GOPATH/src`：把`wkpkg`复制到 $GOPATH/src 目录下
- `go install wkpkg`：build `wkpgk.a` 归档文件到 $GOPATH/pkg 目录下，可以被其他程序import
- `go build main.go`：在本地创建临时的无依赖关系的可执行文件
- `cp -r wkapp $GOPATH/src`: 把`wkapp`复制到 $GOPATH/src 目录下
- `go install wkapp`：build 可执行文件 `wkapp` 并放在 $GOPATH/bin 目录下


## GO Module
GO MODULE 是保证在非 $GOPATH/src 目录下，可以编译go包，并把结果输出到 $GOPATH 目录下。

### Configuration
Go在1.13版本后是默认开启 MODULE, 代表当项目在 GOPATH/src 外且项目根目录有 go.mod 文件时，开启 go module.
也就是说,如果你不把代码放置在 GOPATH/src 下则默认使用 MODULE 管理.

```bash
go env -w GOPROXY=https://mirrors.cloud.tencent.com/go/,direct
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
go env -w GO111MODULE=on
```

### 步骤
```bash
cd wkmodule
go mod init wkmodule # 创建 go.mod
go install  # download pkg in $GOPATH/src pkg and build wkmodule in $GOPATH/bin
go mod tidy # tidy会检测该文件夹目录下所有引入的依赖,写入 go.mod 文件
go mod download # 将依赖全部下载至 GOPATH 下
go mod vendor # 将刚才下载至 GOPATH 下的依赖转移至该项目根目录下的 vendor(自动新建) 文件夹下
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
不使用 $GOPATH 而使用本地目录 wkmodule 进行编译

```bash
go mod vendor
go build -ldflags "-s -w" -a -installsuffix cgo -o wkmodule-local .
```


## 源代码

- [命令源代码](cmd-code.md)
- [库源代码](archi-code.md)
- [测试源代码](test-code.md)



## Ref

1. [理解Golang包导入](https://studygolang.com/articles/3189)
1. [亲测GO环境搭建，理解go build、go install、go get](https://blog.csdn.net/zhangliangzi/article/details/77914943)
1. [go module 基本使用](https://www.cnblogs.com/chnmig/p/11806609.html)
1. [一文搞懂 Go Modules 前世今生及入门使用](https://www.cnblogs.com/wongbingming/p/12941021.html)


