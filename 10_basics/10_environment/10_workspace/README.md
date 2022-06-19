# Workspace

Go 的 workspace 被定义在 $GOPATH 工作目录下，其结构有三个子目录：

## GOPATH

指定 workspace 的位置，默认为 $HOME/go

Go 的编译模式为：

- go get 获取的代码会放在 `$GOPATH/src` 下面
- go build 如果依赖第三方 pkg，会在 `$GOROOT/src` 和 `$GOPATH/src` 下面按照 import path 去搜索 pkg

## src/

go 默认用于存放源文件：

- src/ 下面可以包含多个 git repo

## bin/ 

$GOBIN，一般为 $GOPATH/bin，用于存放编译后的可执行文件。

往往会将 ￥GOBIN 目录添加到 $PATH：``export PATH=$PATH:$(go env GOPATH)/bin``

## pkg/

存放编译后的库文件

<img src="figures/2fdfb5620e072d864907870e61ae5f3c.png" alt="img" style="zoom:50%;" />