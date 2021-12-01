# 命令

## GOROOT

go 程序安装的地址，建议下载安装包安装

## build

拉取所有 import 的源代码，在当前目录下编译生成可执行/库文件。go build 会拉取所有 import 引用包的源码并重新编译，而不是直接使用 `$GOPATH/pkg` 里的编译文件。

如果在 `$GOROOT` 与`$GOPATH` 下没有找到 import 的源码，就会报错。

```shell
export GO111MODULE=off
go build hello.go  # 在本地验证编译，不会有文件生成
```

## install

拉取所有 import 的源代码，编译生成可执行/库文件

- 库文件放到 `$GOPATH/pkg/$GOOS_$GOARCH`目录下
- 可执行文件放到 `$GOPATH/bin`目录下

## run

## get

自动获取、构建和安装远程包

- git clone 到 $GOPATH/src
- go install

```markdown
go get [github.com/golang/example/hello](http://github.com/golang/example/hello)
```

## clean

``go clean -i # 删除对应的库文件或命令文件``

## test

在 _test.go 结尾的文件中写以 TestXXX 开头的函数

```shell
go test
```

注意：在 import 第三方包的时候，当源码和 .a 均已安装的情况下，编译器链接的是源码。
注意：在 import 标准库中的包（如fmt）时，也是必须要源码的。不过与自定义包不同的是，如果未找到源码，不会尝试重新编译标准包，而是在链接时链接已经编译好的`.a`文件。

