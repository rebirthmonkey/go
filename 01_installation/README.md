# 安装

## Installation

### Ubuntu

Follow this [manuel](https://tecadmin.net/install-go-on-ubuntu/)

- `https://golang.org/doc/install`：从这里下载，建议安装 1.18.7 版本
- `sudo tar -xvf go1.10.1.linux-amd64.tar.gz`
- `sudo mv go /usr/local`
- `go version`: check installation
- `go env`

### MacOS

Follow this [manuel](https://golangdocs.com/install-go-mac-os)

- 先下载pkg格式的安装包
- 运行安装包，授予相应的权限

或者首先安装Homebrew，然后运行`brew install go`。

不管用何种形式，`go version`能够返回结果，就算是安装成功了

## 环境变量配置

主要是设置一下几个变量，如 GOROOT、GOPATH、GOBIN、PATH 等。

| **环境变量** | **含义**                                                     |
| ------------ | ------------------------------------------------------------ |
| GOROOT       | Go  语言编译工具、标准库等的安装路径                         |
| GOPATH       | Go  的工作目录，也就是编译后二进制文件的存放目录和 import 包时的搜索路径 |
| GO111MODULE  | 通过设置GO111MODULE值为 on、off、auto 来控制是否开启 Go Modules 特性。其中，**on** 代表开启 Go modules 特性，这会让 Go 编译器忽略$GOPATH 和 vendor 文件夹，只根据 go.mod 下载依赖。**off** 代表关闭 Go modules 特性，这会让 Go 编译器在$GOPATH 目录和 vendor目录来查找依赖关系，也就是继续使用“GOPATH 模式”。而 **auto** 在 Go1.14 和之后的版本中是默认值。当设置为 auto 后，源码在$GOPATH/src 下，并且没有包含 go.mod 则关闭 Go modules，其他情况下都开启 Go modules |
| GOPROXY      | Go  包下载代理服务器。众所周知的原因，在大陆的网络环境下是无法访问 golang.org 等 google 网站的，但在日常开发中使用的很多依赖包都要从 google 的服务器上下载。为了解决无法加载依赖的问题，需要设置一个代理服务器。以便我们能够使用 go  get 下载以来包 |
| GOPRIVATE    | 指定不走代理的 Go 包域名。go get 通过代理服务拉取私有仓库（内部仓库或托管站点的私有仓库），而代理服务无法访问私有仓库，会出现了 404 错误。go1.13 版本提供了一个方便的解决方案：GOPRIVATE 环境变量，通过该变量，可以使得指定的包不通过代理下载，而是直接下载 |
| GOSUMDB      | GOSUMDB 的值是一个 web 服务器，默认值是 sum.golang.org，该服务可以用来查询依赖包指定版本的哈希值，保证拉取到的模块版本数据未经过篡改 |

基本都是通过修改`~/.bashrc`或者`./zshrc`达成：

```shell
export GOVERSION=go1.17.2 # Go 版本设置
export GO_INSTALL_DIR=XXX # Go 安装目录
export GOROOT=XXX # GOROOT 设置
export GOPATH=$WORKSPACE/go # GOPATH 设置
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH # 将 Go 语言自带的和通过 go install 安装的二进制文件加入到 PATH 路径中
export GO111MODULE="on" # 开启 Go moudles 特性
export GOPROXY=https://goproxy.cn,direct # 安装 Go 模块时，代理服务器设置
export GOPRIVATE=
export GOSUMDB=off # 关闭校验 Go 依赖包的哈希值

```

运行`go env`命令可以查看和GO有关的环境变量的值。


### Ubuntu

```shell
export GOROOT=/usr/local/go
export GOPATH=/home/XXX/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN:$GOROOT/bin
```

> 大部分Linux发行版的用户目录为`/home/XXX`，XXX 为用户名

### MacOS

在Ubuntu/MacOS下可以通过编辑`~/.zshrc` 文件添加以下四行内容配置GO环境变量

```shell
export GOROOT=/usr/local/go
export GOPATH=/Users/XXX/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN:$GOROOT/bin
```

> XXX 为用户名
> 如果不是通过官方安装器安装的GO，可能需要替换`/usr/local/go`为具体的值。例如如果通过brew安装的go，则GOROOT应该是`/opt/homebrew/Cellar/go/<VERSION>/libexec`(M1)，`/usr/local/Cellar/go/<VERSION>/libexec`(Intel)

通过`source ~/.bashrc`或者`source ~/.zshrc`更新当前终端下环境变量

## hello-world

```bash
go run hello-world.go
go build hello-world.go
```

## GOPROXY

将如下两行配置加入终端配置文件，可以提高许多package的下载速度

```shell
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```

> `export GOPROXY=direct`可以指定包从源站下载
