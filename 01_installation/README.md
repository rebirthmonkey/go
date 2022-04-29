# 安装

## Installation

### Ubuntu

Follow this [manuel](https://tecadmin.net/install-go-on-ubuntu/)

- `https://golang.org/doc/install`：从这里下载，建议安装 1.14.4 版本
- `sudo tar -xvf go1.10.1.linux-amd64.tar.gz`
- `sudo mv go /usr/local`
- `go version`: check 
- `go env`

## 环境变量配置
```bash
go env
$GOROOT=/usr/local/go # GO安装目录
$GOPATH=/Users/XXX/go # GO工作目录
$GOBIN=$GOPATH/bin # GO可执行文件目录
$PATH=$PATH:$GOBIN:$GOROOT/bin # 将GO可执行文件加入PATH中，使GO指令与我们编写的GO应用可以全局调用 
```
在MacOS下可以通过 `vim ~/.zshrc` 配置
```bash
export GOROOT=/usr/local/go
export GOPATH=/Users/XXX/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN:$GOROOT/bin
```
通过 `sourc ~/.zshrc`更新环境变量

## hello-world
```bash
go run hello-world.go
go build hello-world.go
```

