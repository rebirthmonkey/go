# 安装

## Installation

### Ubuntu

Follow this [manuel](https://tecadmin.net/install-go-on-ubuntu/)

- `https://golang.org/doc/install`：从这里下载，建议安装 1.14.4 版本
- `sudo tar -xvf go1.10.1.linux-amd64.tar.gz`
- `sudo mv go /usr/local`
- `go version`: check installation
- `go env`

## 环境变量配置

```bash
go env
$GOROOT=/usr/local/go # GO安装目录，不可加上 /bin
$GOPATH=/Users/XXX/go # GO工作目录
$GOBIN=$GOPATH/bin # GO可执行文件目录
$PATH=$PATH:$GOBIN:$GOROOT/bin # 将GO可执行文件加入PATH中，使GO指令与我们编写的GO应用可以全局调用 
```

> 大部分Linux发行版的用户目录为`/home/XXX`，因此第二行`GOPATH=/Users/XXX/go`需要修改成`GOPATH=/home/XXX/go`

在MacOS下可以通过编辑`~/.zshrc` 文件添加以下四行内容配置GO环境变量

```bash
export GOROOT=/usr/local/go
export GOPATH=/Users/XXX/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN:$GOROOT/bin
```

> XXX 为用户名
> 可能需要替换`/usr/local/go`为安装目录，例如如果通过brew安装的go，则GOROOT应该是`/opt/homebrew/Cellar/go/<VERSION>/libexec`(M1)，`/usr/local/Cellar/go/<VERSION>/libexec`

通过 `sourc ~/.zshrc`更新当前终端下环境变量

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
