# CMD

commands 代表命令，arguments 代表非选项（flag）参数，flags 代表选项参数（也叫标志）。

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

#### FlagSet

FlagSet 是一组定义好的 Flag 的集合，几乎所有的 Pflag 操作都需要借助 FlagSet 提供的方法来完成。在实际开发中，可以使用两种方法来获取并使用 FlagSet：

- 调用 NewFlagSet 创建一个 FlagSet。

```go
var version bool
flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
flagSet.BoolVar(&version, "version", true, "Print version information and quit.")
```

- 使用 Pflag 包定义的全局 FlagSet：CommandLine，实际上 CommandLine 也是由  NewFlagSet 函数创建的。

```go
import (
    "github.com/spf13/pflag"
)
pflag.BoolVarP(&version, "version", "v", true, "Print version information and quit.")
```

#### 使用方法

命令行参数定义：

- 无短选项，并将值存储在指针中： `var name = pflag.String("name", "colin", "Input Your Name")`
- 有短选项，并将值存储在指针中：`var name = pflag.StringP("name", "n", "colin", "Input Your Name")`
- 无短选项，并将值绑定到变量：`pflag.StringVar(&name, "name", "colin", "Input Your Name")`
- 有短选项，并将值绑定到变量：`pflag.StringVarP(&name, "name", "n","colin", "Input Your Name")`

上面的函数命名是有规则的：

- 函数名带Var说明是将标志的值绑定到变量，否则是将标志的值存储在指针中。
- 函数名带P说明支持短选项，否则不支持短选项。


#### Lab

- [pflag解析](20_flag/20_pflag/example.go)

```shell
go run example.go -n nnn -a 88 -g female -o false
```

- [获取非选项参数](20_flag/20_pflag/example2.go)

```shell
go run example2.go --flagname 222 arg1 arg2
```

- [FlagSet](20_flag/20_pflag/example3.go)

```shell
go run example3.go --version xxx
```

## Config

### Goconfig

从配置文件中获取参数

- [配置文件解析](30_config-file/10_goconfig/example.go)

```shell
go run example1.go
```

### viper

几乎所有的后端服务都需要一些配置项来配置服务.一些小型的项目，配置不是很多，可以选择只通过命令行参数来传递配置。但是大型项目配置很多，通过命令行参数传递就变得很麻烦，不好维护。标准的解决方案是将这些配置信息保存在配置文件中，由程序启动时加载和解析。Go 中有很多包可以加载并解析配置文件，目前最受欢迎的就是 Viper 包。Viper 能够处理不同格式的配置文件，它可以从不同的位置读取配置，不同位置的配置具有不同的优先级，高优先级的配置会覆盖低优先级相同的配置。Viper 配置键不区分大小写。

#### 读入配置

读入配置就是将配置文件整体读入到 Viper 中，Viper 支持设置多个配置文件搜索路径，需要注意添加搜索路径的顺序，Viper 会根据添加的路径顺序搜索配置文件，如果找到则停止搜索。如果通过搜索的方式查找配置文件，则需要注意 SetConfigName 设置的配置文件名是不带扩展名的，在搜索时 Viper 会在文件名之后追加文件扩展名，并尝试搜索所有支持的扩展类型。

- [viper example](30_config-file/20_viper/example.go)
- [viper with struct unmarshal](30_config-file/20_viper/example2.go)
- [viper 处理嵌套结构体](30_config-file/20_viper/example3.go)

#### 使用 Pflag

Viper 支持 Pflag 包，能够绑定 key 到 Flag：

- 对于单个 flag，可以调用 BindPFlag() 进行绑定：`viper.BindPFlag("token", pflag.Lookup("token"))`
- 对于多个 flag，可以绑定一组 flagSet：`viper.BindPFlags(pflag.CommandLine)`

## 子命令

为程序创建子命令

### Cobra

Cobra 既是一个可以创建 CLI 应用程序的库，也是一个可以生成应用和命令文件的程序。有许多大型项目都是用 Cobra 来构建应用程序的，如 Kubernetes、Docker、etcd、Rkt、Hugo 等。Cobra 建立在 commands、arguments 和 flags 结构之上。一个好的应用程序应该是易懂的，用户可以清晰地知道如何去使用这个应用程序。应用程序通常遵循如下模式：`APPNAME COMMAND ARG --FLAG`，如：`git clone URL  --bare`

- clone 是一个命令，
- URL是一个非选项参数，
- bare是一个选项参数

#### 构建

Cobra 提供了两种方式来创建命令：Cobra 命令和 Cobra 库。Cobra 命令可以生成一个 Cobra 命令模板，而命令模板也是通过引用 Cobra 库来构建命令的。所以，这里直接介绍如何使用 Cobra 库来创建命令。使用 Cobra 库创建命令如果要用 Cobra 库编码实现一个应用程序，需要首先创建一个空的 main.go 文件和一个 rootCmd  文件，之后可以根据需要添加其他命令。具体步骤如下：

- [Cobra框架](40_sub-cmd/20_cobra/example.go)

```shell
go build -o wkctl .
./wkctl
./wkctl help
./wkctl version
./wkctl hello xxx
```

## XXX framework

构建的任何一个应用程序，至少要具备命令行程序、命令行参数解析和配置文件解析这 3 种功能：

- 命令行程序：用来启动一个应用。命令行程序需要实现诸如应用描述、help、参数校验等功能。还可以实现命令自动补全、打印命令行参数等高级功能。
- 命令行参数解析：用来在启动时指定应用程序的命令行参数，以控制应用的行为。
- 配置文件解析：用来解析不同格式的配置文件。

上述 3 类功能跟业务关系不大，可以抽象成一个统一的框架。





