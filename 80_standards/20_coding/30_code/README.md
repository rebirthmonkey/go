# 代码

## 代码结构

### 分层 MVC

按层拆分和按功能拆分：

按层拆分 MVC：在 MVC 架构中，将服务中的不同组件按访问顺序，拆分成了 Model、View 和 Controller 三层。每层完成不同的功能：

- View（视图）：提供给用户的操作界面，用来处理数据的显示。
- Controller（控制器）：负责根据用户从 View 层输入的指令，选取 Model  层中的数据，然后对其进行相应的操作，产生最终结果。
- Model（模型）：应用程序中用于处理数据逻辑的部分。

<img src="figures/ed0c3dfyy52ac82539cb602eec9f0146.png" alt="img" style="zoom: 50%;" />

在 Go 项目中，按层拆分会带来很多问题。最大的问题是循环引用：相同功能可能在不同层被使用到，而这些功能又分散在不同的层中，很容易造成循环引用，所以不推荐。

### 按功能拆分

比如，一个订单系统，可以根据不同功能将其拆分成用户（user）、订单（order）和计费（billing）3 个模块，每一个模块提供独立的功能，功能更单一。

<img src="figures/0d65eb1363bf8055e209bc24d1d99ca5.png" alt="img" style="zoom:50%;" />

其目录结构如下：

```shell
$ tree --noreport -L 2 pkg
pkg
├── billing
├── order
│   └── order.go
└── user
```

相较于按层拆分，按功能拆分模块带来的好处也很好理解：不同模块，功能单一，可以实现高内聚低耦合。因为所有的功能只需要实现一次，引用逻辑清晰，会大大减少出现循环引用的概率。所以，有很多优秀的 Go 项目采用的都是按功能拆分的模块拆分方式，例如 Kubernetes、Docker、Helm、Prometheus 等。

## 代码规范

除了组织合理的代码结构这种方式外，编写高质量 Go 应用的另外一个行之有效的方法，是遵循 Go  语言代码规范来编写代码。

### 代码格式

- 运算符和操作数之间要留空格。
- 建议一行代码不超过 120 个字符，超过部分采用合适的换行方式换行。但也有些例外场景，例如 import 行、工具自动生成的代码、带 tag 的 struct 字段。
- 文件长度不能超过 800 行。
- 函数长度不能超过 80 行。
- import：
  - 不要使用相对路径引入包，例如 import …/util/net。
  - 包名称与导入路径的最后一个目录名不匹配时，或者多个相同包名冲突时，则必须使用导入别名。
  - 导入的包建议进行分组（标准包、第三方包、匿名包、内部包），匿名包的引用使用一个新的分组，并对匿名包引用进行说明。

```go
  import (
    // go 标准包
    "fmt"

    // 第三方包
      "github.com/jinzhu/gorm"
      "github.com/spf13/cobra"
      "github.com/spf13/viper"

    // 匿名包单独分组，并对匿名包引用进行说明
      // import mysql driver
      _ "github.com/jinzhu/gorm/dialects/mysql"

    // 内部包
      v1 "github.com/marmotedu/api/apiserver/v1"
      metav1 "github.com/marmotedu/apimachinery/pkg/meta/v1"
      "github.com/marmotedu/iam/pkg/cli/genericclioptions"
  )
```

### 命名规范

命名规范是代码规范中非常重要的一部分，一个统一的、短小的、精确的命名规范可以大大提高代码的可读性，也可以借此规避一些不必要的 Bug。

#### 包命名

- 包名必须和目录名一致，尽量采取有意义、简短的包名，不要和标准库冲突。
- 包名全部小写，没有大写或下划线，使用多级目录来划分层级。
- 项目名可以通过中划线来连接多个单词。
- 包名以及包所在的目录名，不要使用复数，例如，是net/url，而不是net/urls。
- 不要用 common、util、shared 或者 lib 这类宽泛的、无意义的包名。
- 包名要简单明了，例如 net、time、log。

#### 文件命名

- 文件名要简短有意义。
- 文件名应小写，并使用下划线分割单词。

#### 接口命名

接口命名的规则，基本和结构体命名规则保持一致

- 单个函数的接口名以 “er"”作为后缀（例如  Reader，Writer），有时候可能导致蹩脚的英文，但是没关系。
- 两个函数的接口名以两个函数名命名，例如  ReadWriter。
- 三个以上函数的接口名，类似于结构体名。

```go
  type Seeker interface {
      Seek(offset int64, whence int) (int64, error)
  }

  // ReadWriter is the interface that groups the basic Read and Write methods.
  type ReadWriter interface {
      Reader
      Writer
  }
```

#### 结构体命名

- 采用驼峰命名方式，首字母根据访问控制决定使用大写或小写，例如 MixedCaps 或 mixedCaps。
- 结构体名不应该是动词，应该是名词，比如 Node、NodeSpec。
- 避免使用 Data、Info 这类无意义的结构体名。
- 结构体的声明和初始化应采用多行

```go
// User 多行声明
type User struct {
    Name  string
    Email string
}

// 多行初始化
u := User{
    UserName: "colin",
    Email:    "colin404@foxmail.com",
}
```

#### 函数命名

- 函数名采用驼峰式，首字母根据访问控制决定使用大写或小写，例如：MixedCaps 或者 mixedCaps。
- 代码生成工具自动生成的代码  (如 xxxx.pb.go) 和为了对相关测试用例进行分组，而采用的下划线 (如  TestMyFunction_WhatIsBeingTested) 排除此规则。

#### 变量命名

- 变量名必须遵循驼峰式，首字母根据访问控制决定使用大写或小写。
- 在相对简单（对象数量少、针对性强）的环境中，可以将一些名称由完整单词简写为单个字母，例如：user 可以简写为 u；userID 可以简写 uid。
- 特有名词时，需要遵循以下规则：如果变量为私有，且特有名词为首个单词，则使用小写，如 apiClient。其他情况都应当使用该名词原有的写法，如 APIClient、repoID、UserID。
- 若变量类型为 bool 类型，则名称应以 Has，Is，Can 或 Allow 开头

```go
var hasConflict bool
var isExist bool
var canManage bool
var allowGitHook bool
```

- 局部变量应当尽可能短小，比如使用 buf 指代 buffer，使用 i 指代 index。
- 代码生成工具自动生成的代码可排除此规则（如 xxx.pb.go 里面的 Id）

#### 常量命名

- 常量名必须遵循驼峰式，首字母根据访问控制决定使用大写或小写。
- 如果是枚举类型的常量，需要先创建相应类型：

```go
// Code defines an error code type.
type Code int

// Internal errors.
const (
    // ErrUnknown - 0: An unknown error occurred.
    ErrUnknown Code = iota
    // ErrFatal - 1: An fatal error occurred.
    ErrFatal
)
```

#### Error的命名

- Error 类型应该写成 FooError 的形式。

```go
type ExitError struct {
  // ....
}
```

- Error 变量写成 ErrFoo 的形式。

```go
var ErrFormat = errors.New("unknown format")
```

### 注释规范

- 每个可导出的名字都要有注释，该注释对导出的变量、函数、结构体、接口等进行简要介绍。
- 全部使用单行注释，禁止使用多行注释。
- 和代码的规范一样，单行注释不要过长，禁止超过 120 字符，超过的请使用换行展示，尽量保持格式优雅。
- 注释必须是完整的句子，以需要注释的内容作为开头，句点作为结尾，格式为“// 名称 描述.”。

```go
// bad
// logs the flags in the flagset.
func PrintFlags(flags *pflag.FlagSet) {
  // normal code
}

// good
// PrintFlags logs the flags in the flagset.
func PrintFlags(flags *pflag.FlagSet) {
  // normal code
}
```

- 所有注释掉的代码在提交 code review 前都应该被删除，否则应该说明为什么不删除，并给出后续处理建议。
- 在多段注释之间可以使用空行分隔加以区分

```go
// Package superman implements methods for saving the world.
//
// Experience has shown that a small number of procedures can prove
// helpful when attempting to save the world.
package superman
```

#### 包注释

- 每个包都有且仅有一个包级别的注释。
- 包注释统一用 // 进行注释，格式为“// Package 包名 包描述”

```go
// Package genericclioptions contains flags which can be added to you command, bound, completed, and produce useful helper functions.
package genericclioptions
```

#### 结构体

- 注释每个需要导出的结构体或者接口都必须有注释说明，格式为“// 结构体名 结构体描述.”。
- 结构体内的可导出成员变量名，如果意义不明确，必须要给出注释，放在成员变量的前一行或同一行的末尾。

```go
// User represents a user restful resource. It is also used as gorm model.
type User struct {
    // Standard object's metadata.
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Nickname string `json:"nickname" gorm:"column:nickname"`
    Password string `json:"password" gorm:"column:password"`
    Email    string `json:"email" gorm:"column:email"`
    Phone    string `json:"phone" gorm:"column:phone"`
    IsAdmin  int    `json:"isAdmin,omitempty" gorm:"column:isAdmin"`
}
```

#### 方法注释

- 每个需要导出的函数或者方法都必须有注释，格式为“// 函数名 函数描述.”。

```go
// BeforeUpdate run before update database record.
func (p *Policy) BeforeUpdate() (err error) {
  // normal code
  return nil
}
```

#### 类型注释

- 每个需要导出的类型定义和类型别名都必须有注释说明，格式为“// 类型名 类型描述.”。

```go
// Code defines an error code type.
type Code int
```

#### 变量/常量注释

- 每个可导出的变量 / 常量都必须有注释说明，格式为“// 变量名 变量描述”

```go
// ErrSigningMethod defines invalid signing method error.
var ErrSigningMethod = errors.New("Invalid signing method")
```

- 出现大块常量或变量定义时，可在前面注释一个总的说明，然后在每一行常量的前一行或末尾详细注释该常量的定义

```go
// Code must start with 1xxxxx.    
const (                         
    // ErrSuccess - 200: OK.          
    ErrSuccess int = iota + 100001    
                                                   
    // ErrUnknown - 500: Internal server error.    
    ErrUnknown    

    // ErrBind - 400: Error occurred while binding the request body to the struct.    
    ErrBind    
                                                  
    // ErrValidation - 400: Validation failed.    
    ErrValidation 
)
```

### 声明/初始化/定义

- 当函数中需要使用到多个变量时，可以在函数开始处使用 var 声明。在函数外部声明必须使用 var ，不要采用 := ，容易踩到变量的作用域的问题。

```go
var (
  Width  int
  Height int
)
```

- 在初始化结构引用时，请使用 &T{} 代替 new(T)，以使其与结构体初始化一致。

```go
// bad
sptr := new(T)
sptr.Name = "bar"

// good
sptr := &T{Name: "bar"}
```

- struct 声明和初始化格式采用多行，定义如下。

```go
type User struct{
    Username  string
    Email     string
}

user := User{
  Username: "colin",
  Email: "colin404@foxmail.com",
}
```

- 相似的声明放在一组，同样适用于常量、变量和类型声明。
- 尽可能指定容器容量，以便为容器预先分配内存

```go
v := make(map[int]string, 4)
v := make([]string, 0, 4)
```

- 对于未导出的顶层常量和变量，使用 _ 作为前缀。

```go
// bad
const (
  defaultHost = "127.0.0.1"
  defaultPort = 8080
)

// good
const (
  _defaultHost = "127.0.0.1"
  _defaultPort = 8080
)
```

- 嵌入式类型（例如 mutex）应位于结构体内的字段列表的顶部，并且必须有一个空行将嵌入式字段与常规字段分隔开。

```go
// bad
type Client struct {
  version int
  http.Client
}

// good
type Client struct {
  http.Client

  version int
}
```

- type assertion 的单个返回值针对不正确的类型将产生 panic。请始终使用 “comma ok”的惯用法。

```go
// bad
t := n.(int)

// good
t, ok := n.(int)
if !ok {
  // error handling
}
// normal code
```

### 类型

#### 字符串

- 空字符串判断

```go
// bad
if s == "" {
    // normal code
}

// good
if len(s) == 0 {
    // normal code
}
```

- []byte/string 相等比较。

```go
// bad
var s1 []byte
var s2 []byte
...
bytes.Equal(s1, s2) == 0
bytes.Equal(s1, s2) != 0

// good
var s1 []byte
var s2 []byte
...
bytes.Compare(s1, s2) == 0
bytes.Compare(s1, s2) != 0
```

#### 切片

- 空 slice 判断，同样适用于 map、channel

```go
// bad
if len(slice) == 0 {
    // normal code
}

// good
if slice != nil && len(slice) == 0 {
    // normal code
}
```

- 声明 slice

```go
// bad
s := []string{}
s := make([]string, 0)

// good
var s []string
```

- slice 复制

```go
// bad
var b1, b2 []byte
for i, v := range b1 {
   b2[i] = v
}
for i := range b1 {
   b2[i] = b1[i]
}

// good
copy(b2, b1)
```

- slice 新增

```go
// bad
var a, b []int
for _, v := range a {
    b = append(b, v)
}

// good
var a, b []int
b = append(b, a...)
```

#### 结构体

- struct 初始化：struct 以多行格式初始化。

```go
type user struct {
  Id   int64
  Name string
}

u1 := user{100, "Colin"}

u2 := user{
    Id:   200,
    Name: "Lex",
}
```

### 控制结构

#### if

- if 接受初始化语句，约定如下方式建立局部变量。

```go
if err := loadConfig(); err != nil {
  // error handling
  return err
}
```

- if 对于 bool 类型的变量，应直接进行真假判断

```go
var isAllow bool
if isAllow {
  // normal code
}
```

#### for

- 采用短声明建立局部变量

```go
sum := 0
for i := 0; i < 10; i++ {
    sum += 1
}
```

- 不要在 for 循环里面使用 defer，defer 只有在函数退出时才会执行

```go
// bad
for file := range files {
  fd, err := os.Open(file)
  if err != nil {
    return err
  }
  defer fd.Close()
  // normal code
}

// good
for file := range files {
  func() {
    fd, err := os.Open(file)
    if err != nil {
      return err
    }
    defer fd.Close()
    // normal code
  }()
}
```

#### range

- 如果只需要第一项（key），就丢弃第二个。

```go
for key := range keys {
// normal code
}
```

- 如果只需要第二项，则把第一项置为下划线

```go
sum := 0
for _, value := range array {
    sum += value
}
```

#### switch

- 必须要有 default

```go
switch os := runtime.GOOS; os {
    case "linux":
        fmt.Println("Linux.")
    case "darwin":
        fmt.Println("OS X.")
    default:
        fmt.Printf("%s.\n", os)
}
```

#### goto

- 业务代码禁止使用 goto 。
- 框架或其他底层源码尽量不用。

### 函数

- 传入变量和返回变量以小写字母开头。
- 函数参数个数不能超过 5 个。
- 函数分组与顺序

  - 函数应按粗略的调用顺序排序。
  - 同一文件中的函数应按接收者分组。
- 尽量采用值传递，而非指针传递。
- 传入参数是 map、slice、chan、interface ，不要传递指针。

#### 函数参数

- 传入变量和返回变量都以小写字母开头。
- 尽量用值传递，非指针传递。
- 参数数量均不能超过 5 个。
- 多返回值最多返回三个，超过三个请使用 struct。

#### defer

- 当存在资源创建时，应紧跟 defer 释放资源（可以大胆使用 defer，defer 在 Go1.14 版本中，性能大幅提升，defer 的性能损耗即使在性能敏感型的业务中，也可以忽略）。
- 先判断是否错误，再 defer 释放资源

```go
rep, err := http.Get(url)
if err != nil {
    return err
}

defer resp.Body.Close()
```

#### 方法的接收器

- 推荐以类名第一个英文首字母的小写作为接收器的命名。
- 接收器的命名在函数超过 20 行的时候不要用单字符。
- 接收器的命名不能采用 me、this、self 这类易混淆名称。

#### 嵌套

- 嵌套深度不能超过 4 层。

#### 变量命名

- 变量声明尽量放在变量第一次使用的前面，遵循就近原则。
- 如果某数字出现超过两次，则禁止使用，改用一个常量代替

```go
// PI ...
const Prise = 3.14

func getAppleCost(n float64) float64 {
  return Prise * n
}

func getOrangeCost(n float64) float64 {
  return Prise * n
}
```

### 设置规范

#### GOPATH

- Go 1.11  之后，弱化了 GOPATH 规则，已有代码（很多库肯定是在 1.11 之前建立的）肯定符合这个规则，建议保留 GOPATH 规则，便于维护代码。
- 建议只使用一个 GOPATH，不建议使用多个 GOPATH。如果使用多个 GOPATH，编译生效的 bin 目录是在第一个  GOPATH 下。

#### 依赖管理

- Go 1.11 以上必须使用 Go Modules。
- 使用 Go Modules 作为依赖管理的项目时，不建议提交 vendor 目录。
- 使用 Go Modules 作为依赖管理的项目时，必须提交 go.sum 文件。

### 单元测试

- 单元测试文件名命名规范为 example_test.go。
- 每个重要的可导出函数都要编写测试用例。
- 因为单元测试文件内的函数都是不对外的，所以可导出的结构体、函数等可以不带注释。
- 如果存在 func (b *Bar) Foo ，单测函数可以为 func TestBar_Foo。

### 整体

- 尽量少用全局变量，而是通过参数传递，使每个函数都是“无状态”的。这样可以减少耦合，也方便分工和单元测试。
- 在编译时验证接口的符合性

```go
type LogHandler struct {
  h   http.Handler
  log *zap.Logger
}
var _ http.Handler = LogHandler{}
```

- 编译过程无法检查 interface{}  的转换，只能在运行时检查，小心引起 panic。
- 服务器处理请求时，应该创建一个 context，保存该请求的相关信息（如 requestID），并在函数调用链中传递。
- string 表示的是不可变的字符串变量，对 string 的修改是比较重的操作，基本上都需要重新申请内存。所以，如果没有特殊需要，需要修改时多使用 []byte。
- 优先使用 strconv 而不是 fmt。
- append 要小心自动分配内存，append 返回的可能是新分配的地址。如果要直接修改 map 的 value 值，则 value 只能是指针，否则要覆盖原来的值。
- map 在并发中需要加锁。

### 静态代码检查

golangci-lint

虽然 Go 提供了 go vet 和 go tool vet 来做静态代码检查，但它们检查的内容还不够全面。golangci-lint 是目前使用最多，也最受欢迎的静态代码检查工具。选择 golangci-lint，是因为它具有其他静态代码检查工具不具备的一些优点：

- 速度非常快：golangci-lint 是基于 gometalinter 开发的，但是平均速度要比 gometalinter 快 5 倍。golangci-lint  速度快的原因有三个：可以并行检查代码；可以复用 go build 缓存；会缓存分析结果。
- 可配置：支持 YAML 格式的配置文件，让检查更灵活、更可控。
- IDE 集成：可以集成进多个主流的 IDE，如 VS Code、GNU Emacs、Sublime  Text、Goland 等。
- linter 聚合器：1.41.1 版的 golangci-lint 集成了 76 个 linter，不需要再单独安装这 76 个 linter。并且 golangci-lint 还支持自定义 linter。
- 最小的误报数：golangci-lint 调整了所集成 linter 的默认设置，大幅度减少了误报。
- 良好的输出：输出的结果带有颜色、代码行号和 linter 标识，易于查看和定位。

#### Install

```shell
$ go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.1
$ golangci-lint version # 输出 golangci-lint 版本号，说明安装成功
$ golangci-lint -h # 查看其用法
```

#### run

run 命令执行 golangci-lint 对代码进行检查，是 golangci-lint 最为核心的一个命令。

- `golangci-lint run`：对当前目录及子目录下的所有 Go 文件进行静态代码检查
- `golangci-lint run dir1 dir2/... dir3/file1.go`：对指定的 Go 文件或者指定目录下的 Go 文件进行静态代码检查
- `golangci-lint -c .golangci.yaml ./...`：指定配置文件
- `golangci-lint run --no-config --disable-all -E errcheck ./...`：指定 linter。golangci-lint 默认会从当前目录一层层往上寻找配置文件名 .golangci.yaml、.golangci.toml、.golangci.json 直到根（/）目录。如果找到，就以找到的配置文件作为本次运行的配置文件，所以为了防止读取到未知的配置文件，可以用 --no-config 参数使 golangci-lint 不读取任何配置文件。
- `golangci-lint run --no-config -D godot,errcheck`：禁用某些 linter

#### cache 

cache 命令用来进行缓存控制，并打印缓存的信息。它包含两个子命令：

- clean 用来清除 cache，当觉得 cache 的内容异常，或 cache 占用空间过大时，可以通过 `golangci-lint cache clean 清除` cache。
- status 用来打印 cache 的状态，比如 cache 的存放目录和 cache 的大小，如：`golangci-lint cache status`

#### config

config 命令可以打印 golangci-lint 当前使用的配置文件路径，例如：`golangci-lint config path`

#### linters

可以打印出 golangci-lint 所支持的 linter，并将这些 linter 分成两类，分别是配置为启用的 linter 和配置为禁用的 linter，例如：`golangci-lint linters`

#### .golangci.yaml





## Ref

1. [Go编码规范](https://time.geekbang.org/column/article/385440)
2. [Uber Go 语言编码规范](https://github.com/xxjwxc/uber_go_guide_cn)
3. [Effective Go](https://go.dev/doc/effective_go)
4. [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)
5. [Style guideline for Go packages]()：包含了如何组织 Go 包、如何命名 Go 包、如何写 Go 包文档的一些建议。

