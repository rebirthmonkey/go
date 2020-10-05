# 安装

## GO环境变量配置
```bash
go env
$GOROOT=/usr/local/go # GO安装目录
$GOPATH=/Users/XXX/go # GO工作目录
$GOBIN=$GOPATH/bin # GO可执行文件目录
# $PATH=$PATH:$GOBIN:$GOROOT/bin # 将GO可执行文件加入PATH中，使GO指令与我们编写的GO应用可以全局调用 
```
在MacOS下可以通过 `vim ~/.zshrc` 配置
```bash
export GOROOT=/usr/local/go
export GOPATH=/Users/ruanhe/go
export GOBIN=$GOPATH/bin
# export $PATH=$PATH:$GOBIN:$GOROOT/bin
```
通过 `sourc ~/.zshrc`更新环境变量


## Workspace
Go的workspace被定义在$GOPATH工作目录下，其结构有三个子目录（需要自行创建）：
- src：存放源代码文件
- pkg：存放编译后的文件
- bin：存放编译后的可执行文件 
