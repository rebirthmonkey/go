# Version

使用 Go 创建一个应用时，经常使用 `ldflags -X` 在构建期间将选项添加当应用内作为变量，而添加版本信息则是最常用的场景。

## Lab

- ldflags 注入功能：通过使用 `12_version/config.XXX` 将 build 时的参数注入应用。

```shell
go build -ldflags "-X 'github.com/rebirthmonkey/go/30_lib/12_version/config.Version=0.0.1' -X 'github.com/rebirthmonkey/go/30_lib/12_version/config.BuildTime=$(date)'" ./10_ldflags.go
./10_ldflags 
```

- version 注入功能：

```shell
go build -ldflags "-X 'github.com/rebirthmonkey/go/pkg/version.GitVersion=0.0.1' -X 'github.com/rebirthmonkey/go/pkg/version.BuildDate=$(date)'" ./20_version.go
./20_version 
```



