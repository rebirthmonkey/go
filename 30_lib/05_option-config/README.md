# Option-Config

## 简介

一个命令的基本格式为：`App ARG --FLAG`

- ARG（argument）：代表非选项参数
- FLAG：代表选项参数（也叫标志）

### Arg

提取命令行输入的参数

- [命令行参数解析](10_arg/example.go)

```bash
cd 10_arg
go run example.go a b c d
```

## Flag

### Flag

#### Flag包

提取命令行输入的flag：cmd -a xxx

- [命令行flag解析](20_flag/10_flag/example1.go)

```bash
cd 20_flag/10_flag/
go run example1.go -p 888
go run example2.go # 未输入参数，则采用默认
```

多flag

- [命令行多flag解析](20_flag/10_flag/example2.go)
- [命令行多flag解析](20_flag/10_flag/example3.go)

```bash
cd 20_flag/10_flag/
go run example2.go -d dd -l ll -w=false
go run example3.go x y z # 未输入flag，则采用默认
```

#### pflag包

Go 的组件启动时需要多个参数来配置服务进程，像 kube-apiserver 就有多达 200 多个启动参数，而且这些参数的类型各不相同（如：string、int、ip 类型等），使用方式也不相同（如：--长选项、-短选项等），所以需要一个强大的命令行参数解析工具。虽然 Go 提供了一个标准库 Flag 包用来对命令行参数进行解析，但在大型项目中应用更广泛的是另外一个包 Pflag。Pflag 提供了很多强大的特性，非常适合用来构建大型项目，一些耳熟能详的开源项目都是用 Pflag 来进行命令行参数解析的，例如：Kubernetes、Istio、Helm、Docker、Etcd 等。

##### Flag

Pflag 可以对命令行参数进行处理，一个命令行参数在 Pflag 包中会解析为一个 Flag 类型的变量。Flag 是一个结构体，定义如下：

```go
type Flag struct {
  Name string // flag长选项的名称
  Shorthand string // flag短选项的名称，一个缩写的字符
  Usage string // flag的使用文本
  Value Value  // flag的值
  DefValue string // flag的默认值
  Changed bool // 记录flag的值是否有被设置过
  NoOptDefVal string // 当flag出现在命令行，但是没有指定选项值时的默认值
  Deprecated string // 记录该flag是否被放弃
  Hidden bool // 如果值为true，则从help/usage输出信息中隐藏该flag
  ShorthandDeprecated string // 如果flag的短选项被废弃，当使用flag的短选项时打印该信息
  Annotations map[string][]string // 给flag设置注解
}
```

##### FlagSet

FlagSet 是一组定义好的 Flag 的集合，几乎所有的 Pflag 操作都需要借助 FlagSet 提供的方法来完成。在实际开发中，可以使用 2 种方法来获取并使用 FlagSet：

- 调用 NewFlagSet() 创建一个 FlagSet。

```go
var version bool
flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
flagSet.BoolVar(&version, "version", true, "Print version information and quit.")
```

- 使用 Pflag 包定义的全局 FlagSet `CommandLine`：实际上 CommandLine 也是由  NewFlagSet 函数创建的。

```go
pflag.BoolVarP(&version, "version", "v", true, "Print version information and quit.")
```

##### 使用方法

###### 命令行参数

Pflag 支持以下 4 种命令行参数定义方式：

- 无短选项，并将值存储在指针中： `var name = pflag.String("name", "colin", "Input Your Name")`
- 有短选项，并将值存储在指针中：`var name = pflag.StringP("name", "n", "colin", "Input Your Name")`
- 无短选项，并将值绑定到变量：`pflag.StringVar(&name, "name", "colin", "Input Your Name")`
- 有短选项，并将值绑定到变量：`pflag.StringVarP(&name, "name", "n","colin", "Input Your Name")`

上面的函数命名是有规则的：

- 函数名带 Var 说明是将标志的值绑定到变量，否则是将标志的值存储在指针中。
- 函数名带 P 说明支持短选项，否则不支持短选项。

###### 使用 Get<Type> 获取参数值

可以使用 `Get<Type>` 来获取标志的值，<Type> 代表 Pflag 所支持的类型。例如：有一个 pflag.FlagSet，带有一个名为 flagname 的 int 类型的标志，可以使用 `GetInt()` 来获取 `int`值。需要注意 `flagname` 必须存在且必须是 `int`，如：

```go
i, err := flagset.GetInt("flagname")
```

###### 获取非选项参数

在定义完 flag 之后，可以调用 `pflag.Parse()` 来解析定义的 flag。解析后，可通过 `pflag.Args()`返回所有的非选项参数，通过 `pflag.Arg(i)` 返回第i个非选项参数。参数下标 0 到 `pflag.NArg() - 1`。

###### 认值设置 

创建一个 Flag 后，可以为这个 Flag 设置 pflag.NoOptDefVal。如果一个 Flag 具有 NoOptDefVal，并且该 Flag 在命令行上没有设置这个 Flag 的值，则该标志将设置为 NoOptDefVal 指定的值。如：

```go
var ip = pflag.IntP("flagname", "f", 1234, "help message")
pflag.Lookup("flagname").NoOptDefVal = "4321"
```

#### Lab

- [pflag解析](20_flag/20_pflag/example.go)

```bash
go run 20_flag/20_pflag/example.go -n nnn -a 88 -g female -o=false
```

- [获取非选项参数](20_flag/20_pflag/example2.go)

```bash
go run 20_flag/20_pflag/example2.go --flagname 222 arg1 arg2
```

- [FlagSet](20_flag/20_pflag/example3.go)

```bash
go run 20_flag/20_pflag/example3.go --version xxx
```

## Config

### Goconfig

从配置文件中获取参数

- [配置文件解析](30_config-file/10_goconfig/example.go)

```bash
cd 30_config-file/10_goconfig
go run example.go
```

### viper

几乎所有的后端服务都需要一些配置项来配置服务。一些小型的项目，配置不是很多，可以选择只通过命令行参数来传递配置。但是大型项目配置很多，通过命令行参数传递就变得很麻烦、不好维护。标准的解决方案是将这些配置信息保存在配置文件中，由程序启动时加载和解析。Go 中有很多包可以加载并解析配置文件，目前最受欢迎的就是 Viper 包。Viper 能够处理不同格式的配置文件，它可以从不同的位置读取配置，不同位置的配置具有不同的优先级，高优先级的配置会覆盖低优先级相同的配置。Viper 配置键不区分大小写。

[Viper](https://github.com/spf13/viper) 是国外大神 **spf13** 编写的开源配置解决方案，特性如下:

- 设置默认值
- 可以读取如下格式的配置文件：JSON、TOML、YAML、HCL
- 监控配置文件改动，并热加载配置文件
- 从环境变量读取配置
- 从远程配置中心读取配置（etcd/consul），并监控变动
- 从命令行 flag 读取配置
- 从缓存中读取配置
- 支持直接设置配置项的值

#### 加载配置文件

读入配置就是将配置文件整体读入到 Viper 中。Viper 可以读取配置文件来解析配置，支持JSON、TOML、YAML、YML、Properties、Props、Prop、HCL、Dotenv、Env 格式的配置文件。

Viper 支持设置多个配置文件搜索路径，需要注意添加搜索路径的顺序，Viper 会根据添加的路径顺序搜索配置文件，如果找到则停止搜索。如果通过搜索的方式查找配置文件，则需要注意 SetConfigName 设置的配置文件名是不带扩展名的，在搜索时 Viper 会在文件名之后追加文件扩展名，并尝试搜索所有支持的扩展类型。

Viper 同时支持在运行时让应用程序实时读取配置文件，也就是热加载配置。可以通过 WatchConfig 函数热加载配置。在调用 WatchConfig 函数之前，需要确保已经添加了配置文件的搜索路径。另外，还可以为 Viper 提供一个回调函数，以便在每次发生更改时运行，如：

```go
viper.WatchConfig()

viper.OnConfigChange(func(e fsnotify.Event) {
  // 配置文件发生变更之后会调用的回调函数
  fmt.Println("Config file changed:", e.Name)
})
```

不建议在实际开发中使用热加载功能，因为即使配置热加载了，程序中的代码也不一定会热加载。如：修改了服务监听端口，但是服务没有重启，这时候服务还是监听在老的端口上，会造成不一致。

Viper 配置读取顺序如下：

- 设置默认的配置文件名。
- 读取配置文件。
- 监听和重新读取配置文件。
- 从io.Reader读取配置。
- 从环境变量读取。
- 从命令行标志读取。
- 从远程Key/Value存储读取。

#### 读取配置项

在初始化配置文件后，读取配置只需要调用 `viper.GetString()`、`viper.GetInt()` 和 `viper.GetBool()` 等函数即可。Viper 提供了如下方法来读取配置：

- Get(key string) interface{}。
- Get<Type>(key string) <Type>。
- AllSettings() map[string]interface{}。
- IsSet(key string) : bool。

每一个 Get 方法在找不到值的时候都会返回零值。为了检查给定的键是否存在，可以使用 IsSet() 方法。<Type> 可以是 Viper 支持的类型，首字母大写：Bool、Float64、Int、IntSlice、String、StringMap、StringMapString、StringSlice、Time、Duration，如 GetInt()。

常见的读取配置方法有以下几种：

- 访问嵌套的键：Viper 也可以非常方便地读取多个层级的配置，比如这样一个 YAML 格式的配置：

```yaml
common:
  database:
    name: test
    host: 127.0.0.1
```

如果要读取 host 配置，执行 

```go
viper.GetString("common.database.host")
```

即可。大部分应用都采用 YAML 格式的配置文件，采用 YAML 格式，是因为 YAML 表达的格式更丰富，可读性更强。

- 序列化/反序列化：Viper 可以支持将所有或特定的值解析到结构体、map等，可以通过两个函数来实现：

```go
Unmarshal(rawVal interface{}) error
UnmarshalKey(key string, rawVal interface{}) error
```

下面是一个示例：

```go
type config struct {
  Port int
  Name string
  PathMap string `mapstructure:"path_map"`
}

var C config

err := viper.Unmarshal(&C)

if err != nil {
  t.Fatalf("unable to decode into struct, %v", err)
}
```

如果想要解析那些键本身就包含 . (默认的键分隔符）的配置，则需要修改分隔符：

```go
v := viper.NewWithOptions(viper.KeyDelimiter("::"))

v.SetDefault("chart::values", map[string]interface{}{
  "ingress": map[string]interface{}{
    "annotations": map[string]interface{}{
      "traefik.frontend.rule.type":         "PathPrefix",
      "traefik.ingress.kubernetes.io/ssl-redirect": "true",
    },
  },
})

type config struct {
  Chart struct{
    Values map[string]interface{}
  }
}

var C config
v.Unmarshal(&C)
```

Viper 在后台使用 `github.com/mitchellh/mapstructure` 来解析值，其默认情况下使用 mapstructure tags。当需要将 Viper 读取的配置反序列到我们定义的结构体变量中时，一定要使用 mapstructure tags。

#### Lab

- [viper example](30_config-file/20_viper/example.go)
- [viper with struct unmarshal](30_config-file/20_viper/example2.go)
- [viper 处理嵌套结构体](30_config-file/20_viper/example3.go)

```bash
cd 30_config-file/20_viper/
go run example.go
```

```bash
cd 30_config-file/20_viper/
go run example2.go
```

```bash
cd 30_config-file/20_viper/
go run example3.go
```

## 子命令

为程序创建子命令

### Cobra

Cobra 既是一个可以创建 CLI 应用程序的库，也是一个可以生成应用和命令文件的程序。有许多大型项目都是用 Cobra 来构建应用程序的，如 Kubernetes、Docker、etcd、Rkt、Hugo 等。Cobra 建立在 commands、arguments 和 flags 结构之上。一个好的应用程序应该是易懂的，用户可以清晰地知道如何去使用这个应用程序。应用程序通常遵循如下模式：`APPNAME COMMAND ARG --FLAG`，如：`git clone URL  --bare`

- clone 是一个命令，
- URL 是一个非选项参数，
- bare 是一个选项参数

#### 构建

Cobra 提供了两种方式来创建命令：Cobra 命令和 Cobra 库。Cobra 命令可以生成一个 Cobra 命令模板，而命令模板也是通过引用 Cobra 库来构建命令的。

#### flag

Cobra 可以跟 Pflag 结合使用，实现强大的 flag 功能，使用步骤如下：

- 使用持久化的 flag：flat 可以是“持久的”，这意味着该 flat 可用于它所分配的命令以及该命令下的每个子命令。可以在 rootCmd 上定义持久标志：

```go
rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
```

- 使用本地标志：也可以分配一个本地标志，本地标志只能在它所绑定的命令上使用。下例中 `--source` 标志只能在rootCmd上引用，而不能在 rootCmd 的子命令上引用。

```go
rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
```

- 将 flat 绑定到 Viper：可以将 flat 绑定到 Viper，这样就可以使用 `viper.Get()` 获取标志的值。

```go
var author string

func init() {
  rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
 viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
}
```

- 设置 flag 为必选：默认情况下，flat 是可选的，也可以设置 flat 为必选。当设置 flat 为必选，但是没有提供 flat 时，Cobra 会报错。示例如下：

```go
rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
rootCmd.MarkFlagRequired("region")
```

#### arg

在执行命令的过程中，经常会传入非选项参数 arguments，并且需要对这些 arg 进行验证。Cobra 提供了机制来对 arg 进行验证。可以使用 Command 的 Args 字段来验证非选项参数。Cobra 也内置了一些验证函数：

- NoArgs：如果存在任何非选项参数，该命令将报错。
- ArbitraryArgs：该命令将接受任何非选项参数。
- OnlyValidArgs：如果有任何非选项参数不在 Command 的 ValidArgs 字段中，该命令将报错。
- MinimumNArgs(int)：如果没有至少 N 个非选项参数，该命令将报错。
- MaximumNArgs(int)：如果有多于 N 个非选项参数，该命令将报错。
- ExactArgs(int)：如果非选项参数个数不为 N，该命令将报错。
- ExactValidArgs(int)：如果非选项参数的个数不为 N，或非选项参数不在 Command 的 ValidArgs 字段中，该命令将报错。
- RangeArgs(min, max)：如果非选项参数的个数不在 min 和 max 之间，该命令将报错。

使用预定义验证函数，示例如下：

```go
var cmd = &cobra.Command{
  Short: "hello",
  Args: cobra.MinimumNArgs(1), // 使用内置的验证函数
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hello, World!")
  },
}
```

当然也可以自定义验证函数，示例如下：

```go
var cmd = &cobra.Command{
  Short: "hello",
  // Args: cobra.MinimumNArgs(10), // 使用内置的验证函数
  Args: func(cmd *cobra.Command, args []string) error { // 自定义验证函数
    if len(args) < 1 {
      return errors.New("requires at least one arg")
    }
    if myapp.IsValidColor(args[0]) {
      return nil
    }
    return fmt.Errorf("invalid color specified: %s", args[0])
  },
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hello, World!")
  },
}
```

#### Lab

这里介绍如何使用 Cobra 库来创建命令。使用 Cobra 库创建命令如果要用 Cobra 库编码实现一个应用程序，需要首先创建一个空的 main.go 文件和一个 rootCmd  文件，之后可以根据需要添加其他命令。具体步骤如下：

- [Cobra框架1](40_cobra/example.go)
- [Cobra框架2](40_cobra/example2.go)

```bash
cd 40_cobra/
go build -o wkctl ./example.go
./wkctl
./wkctl help
./wkctl version
./wkctl hello xxx
go build -o wkctl ./example2.go
./wkctl
./wkctl help
./wkctl version
./wkctl hello xxx
```

## apiserver

- [apiserver 示例](80_option-config/README.md)

