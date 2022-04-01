# OS

## Arg

提取命令行输入的参数

- [命令行参数解析](02_os-args.go)

```shell
go run 02_os-args.go a b c d
```


## Flag

### Flag

提取命令行输入的flag：cmd -a xxx

- [命令行flag解析](20_flag/10_flag/example1.go)

```shell
go run example1.go -p 888
go run example2.go # 未输入参数，则采用默认
```

多flag

- [命令行多flag解析](20_flag/10_flag/example2.go)
- [命令行多flag解析](20_flag/10_flag/example3.go)

```shell
go run example2.go -d dd -l ll -w true
go run example2.go x y z # 未输入flag，则采用默认
```

### pflag

Go 的组件启动时需要多个参数来配置服务进程，像 kube-apiserver 就有多达 200 多个启动参数，而且这些参数的类型各不相同（如：string、int、ip 类型等），使用方式也不相同（如：--长选项、-短选项等），所以需要一个强大的命令行参数解析工具。虽然 Go 提供了一个标准库 Flag 包用来对命令行参数进行解析，但在大型项目中应用更广泛的是另外一个包：Pflag。Pflag  提供了很多强大的特性，非常适合用来构建大型项目，一些耳熟能详的开源项目都是用 Pflag 来进行命令行参数解析的，例如：Kubernetes、Istio、Helm、Docker、Etcd 等。

- [pflag解析](20_flag/20_pflag/example.go)

```shell
go run example.go -n nnn -a 88 -g female -o false
```

## Config

从配置文件中获取参数

- [配置文件解析](30_config-file/10_goconfig/example.go)

```shell
go run example1.go
```

## 子命令

为程序创建子命令

- [Cobra框架](40_sub-cmd/20_cobra/example.go)

```shell
go build -o wkctl .
./wkctl
./wkctl help
./wkctl version
./wkctl hello xxx
```

 



