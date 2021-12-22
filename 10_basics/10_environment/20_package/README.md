# 包

## 简介

Go 使用包（package）这种语法元素来组织源码，所有语法可见性均定义在 package 这个级别。一个目录下不同的源文件内的函数对于包来说同级，不会区分不同文件。

比如在使用实时分布式消息平台 nsq 提供的 go client api 时，导入的路径如下：

   import “github.com/bitly/go-nsq”

但在使用其提供的函数时，却直接用 nsq 做前缀包名：

   q, _ := nsq.NewConsumer("write_test", "ch", config)

import 后面的最后一个元素应该是包在 $GOPATH 下的目录路径，并非包名。而在代码中使用的是真正的包名。【1】

### 结构

- 包名：包名可以与目录名不同
- 目录
  - 源文件：指定相同包名

### 作用域

变量、函数的公有或私有

名称的首字母为大写的变量或程序才可以被当前包外的代码引用，否则它就只能被当前包内的其他代码引用。

- 小写：私有
- 大写：公有

### 类型

- 标准包：
- 远程包：

## Pkg/App

### 安装 Pkg
```shell
cp -r wkpkg $GOPATH/src # 把 wkpkg 复制到 $GOPATH/src 目录下
go install -x wkpkg # build wkpgk.a 并归档文件到 $GOPATH/pkg/darwin_amd64 目录或 $GOROOT/.. 下，可以被其他程序 import
```

### App 使用 Pkg

```shell
cp -r wkapp $GOPATH/src # 把 wkapp 复制到 $GOPATH/src 目录下
# 在 GoLand 中 wkpkg 依赖无法解决，需要 `sudo ln -s $GOPATH/src/wkpkg $GOROOT/src/wkpkg`，因为 GoLand 包依赖搜索只寻找 GOROOT
go install -x wkapp # build 可执行文件 wkapp 并放在 $GOPATH/bin 目录下
```


## 同包/多文件

同一个目录下、相同包名的多个文件属于同一个包。但在 go run/build 时需要多个文件

- go run

```bash
cd samelib
go run main.go lib1.go -name="XXX" # 需要把涉及到的文件一起编译
```

- go build

```bash
cd samelib
go build 
./samelib -name="XXX"
```

## Module

### 简介

#### 原理

GO MODULE 确保在非 `$GOPATH/src` 目录下可以编译 go 源文件，并把结果输出到 $GOPATH 目录下。

在 module-aware  mode 下，某repo的顶层目录下会放置一个 go.mod 文件，每个 go.mod 文件定义了一个 module，而放置 go.mod 文件的目录被称为 module root 目录（通常对应一个repo的root目录，但不是必须的）。module root 目录以及其子目录下的所有 pkg 均归属于该 module，除了那些自身包含 go.mod 文件的子目录。

在“module-aware mode”下，go 编译器将不再在 GOPATH 下面以及 vendor 下面搜索目标程序依赖的第三方 Go pkg。go compiler 没有去使用之前已经下载到 `$GOPATH/src` 或 pkg/ 下的包，而是主动将下载的依赖包缓存在 `$GOPATH/pkg/mod` 下面。执行go build后go.mod文件的内容为：

```bash
# cat go.mod
module hello

require (
    bitbucket.org/bigwhite/c v0.0.0-20180714063616-861b08fcd24b
    bitbucket.org/bigwhite/d v0.0.0-20180714005150-3e3f9af80a02 // indirect
)
```

 module 并没有直接依赖 d pkg，因此在 d 的记录后面通过注释形式标记了 indirect，即非直接依赖，也就是传递依赖。

#### 配置

Go 在 1.13 版本后是默认开启 MODULE, 代表当项目在 GOPATH/src 外且项目根目录有 go.mod 文件时，开启 go module。也就是说，如果你不把代码放置在 GOPATH/src 下则默认使用 MODULE 管理.

```bash
go env -w GOPROXY=https://mirrors.cloud.tencent.com/go/,direct
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
go env -w GO111MODULE=on
```

#### 初始化

```bash
cd wkmodule
go mod init wkmodule # 创建 go.mod
go mod tidy # 检测该文件夹目录下所有引入的依赖，并放入 go.mod 文件
go mod download # 将检测到的所有依赖下载至 $GOPATH 下
go install  # 安装 wkmodule 可执行文件到 $GOPATH/bin
```

#### 更新依赖

```bash
go mod tidy
go mod download
```

#### 指定 module 版本

```bash
go mod -require=bitbucket.org/bigwhite/c@v1.0.0 # 会更新 go.mod reqire 部分
```

#### 已有 module 安装

```bash
git clone XXX
cd XXX
go mod download
go build .
```

#### 基于 vendor 编译（推荐）

Go Modules 默认会忽略 vendor/ 这个目录，但如果希望将依赖放入 vendor/，可以执行 `go mod vendor`。vender/ 目录会包含所有的依赖模块代码，并且会在该目录下面添加一个名为modules.txt 的文件用来记录依赖包的一些信息。

不使用 $GOPATH/pkg 而使用本地 vendor/ 目录进行编译

```bash
go mod vendor # 将刚才下载至 $GOPATH/mod 的依赖转移至该项目目录的 vendor 目录下
go build -getmode=vendor main.go
go build -ldflags "-s -w" -a -installsuffix cgo -o wkmodule-local .
```
### 依赖私有远程包

```bash
cd gitlib
go mod init
go mod tidy
go run ./main.go
```

### 依赖私有本地包

包在相对路径下

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
1. [初窥Go module](https://tonybai.com/2018/07/15/hello-go-module/)
1. [Go 语言全新依赖管理系统 Go Modules 使用详解](https://www.toutiao.com/i6714564564194689543/?tt_from=weixin&utm_campaign=client_share&wxshare_count=1&timestamp=1597249873&app=news_article&utm_source=weixin&utm_medium=toutiao_ios&use_new_style=1&req_id=2020081300311301001405309209375D1B&group_id=6714564564194689543)
1. 

