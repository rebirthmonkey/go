# 错误码

针对大部分 web service，很多都是对外暴露 RESTful API 接口，内部系统通信采用 RPC 协议。因为 RESTful API 接口有一些天生的优势，比如规范、调试友好、易懂，所以通常作为直接面向用户的通信规范。

既然是直接面向用户，那么首先就要求消息返回格式是规范的。其次，如果接口报错，还要能给用户提供一些有用的报错信息，通常需要包含错误 Code（用来唯一定位一次错误）和 Message（用来展示出错的信息），这就需要设计一套规范的、科学的错误码。

## HTTP状态码

Go net/http 包提供了 60 个 HTTP 码，大致分为如下 5 类：

- 1XX - 指示信息：表示请求已接收，继续处理。
- 2XX - 请求成功：表示成功处理了请求的状态代码。
- 3XX - 请求被重定向：表示要完成请求需要进一步操作，通常这些状态代码用来重定向。
- 4XX - 客户端错误：表示请求可能出错，妨碍了服务器的处理。通常是客户端出错，需要客户端做进一步的处理。
- 5XX - 服务器错误：表示服务器在尝试处理请求时发生内部错误。

这里建议 HTTP 状态码不要太多，基本上只需要这 6 个

- 200：表示请求成功执行。
- 400：表示客户端请求出问题。
  - 401：表示认证失败。
  - 403：表示授权失败。
  - 404：表示资源找不到。
- 500：表示服务端出问题。

## 错误码

Go 应用如果出错，需要用户或开发者感知到这些错误，并且能够提供一些有效的错误信息进行排障，这就需要对错误进行处理。除了 HTTP 的状态码，还需要有额外的业务错误码。因为 HTTP 状态码有限，并且是只跟 HTTP 层相关的。Go 的 net/http 中包含的 60 个 HTTP 状态码，基本都是跟 HTTP 请求相关的错误码。在一个大型系统中，这些错误码完全不够用，而且这些错误码跟业务没有任何关联，满足不了业务的需求，所以需要有业务错误码。请求出错时，可以通过 HTTP 状态码直接感知到请求出错。需要返回详细出错信息时，通常需要返回 3 类信息：业务错误码、错误信息和参考文档（可选）。具体与 HTTP 状态码的配合方式为：返回 HTTP 404、500 状态码，并在 Body 中返回详细的业务错误码和错误信息。这种方式既能通过 HTTP 状态码使客户端方便地知道请求出错，又可以根据返回的信息知道哪里出错，以及如何解决问题。同时，返回了机器友好的业务业务码，可以在有需要时让程序进一步判断处理。

在实际开发中，引入业务错误码有下面几个好处：

- 可以非常方便地定位问题和定位代码行（看到错误码知道什么意思、搜索错误码可以定位到错误码所在行、某个错误类型的唯一标识）。
- 错误码包含一定的信息，通过错误码可以判断出错误级别、错误模块和具体错误信息。
- Go 中的 HTTP 服务器开发都是引用 net/http 包，该包中只有 60 个错误码，基本都是跟 HTTP 请求相关的错误码，在一个大型系统中，这些错误码完全不够用，而且这些错误码跟业务没有任何关联，满足不了业务的需求。引入业务错误码，则可以解决这些问题。
- 业务开发过程中，可能需要判断错误是哪种类型，以便做相应的逻辑处理，通过定制的错误可以很容易做到这点，例如：

```go
if err == code.ErrBind {
  ...
}
```

这里要注意，业务错误码可以是一个整数，也可以是一个整型字符串，还可以是一个字符型字符串，它是错误的唯一标识。

### 业务错误码

业务错误码是在业务开发过程中，用于判断错误是哪种类型，以便做相应的逻辑处理。一方面，可以根据需要自行扩展，另一方面也能够精准地定位到具体是哪个业务错误。同时，因为业务错误码通常是对计算机友好的 10 进制整数，基于业务错误码，计算机也可以很方便地进行一些分支处理。业务码也要有一定规则，可以通过业务码迅速定位出是哪类错误。

参考新浪的业务错误码设计，业务错误码需要用纯数字表示，不同部位代表不同的服务、不同的模块。业务错误码从 100101 开始，其中：

+ 10：服务，如对外 REST 接口
+ 01：模块
+ 01：模块自身的错误码序号，每个模块可以注册 100 个错误

#### 服务/模块说明

10 通用为所有服务都适用的错误，提高复用性、避免重复造轮子。例如：

| 服务 | 模块 | 说明(服务 - 模块)            |
| ---- | ---- | ---------------------------- |
| 10   | 00   | 通用 - 基本错误              |
| 10   | 01   | 通用 - 数据库类错误          |
| 10   | 02   | 通用 - 认证授权类错误        |
| 10   | 03   | 通用 - 加解码类错误          |
| 11   | 00   | apiserver服务 - 用户模块错误 |
| 11   | 01   | apiserver服务 - 密钥模块错误 |
| 11   | 02   | apiserver服务 - 策略模块错误 |
| 12   | 00   | 其他服务 - 认证模块错误      |

#### 具体错误码

具体错误码可参考：[错误码](40_error-code.md)，该错误码描述是通过程序自动生成的。

### 错误信息

错误描述包括：对外的错误描述和对内的错误描述两部分。

#### 对外错误描述

- 统一大写开头，结尾不要加`.`
- 要简洁，并能准确说明问题
- 错误说明应该是 `该怎么做` 而不是 `哪里错了`

在下例中，返回中 code 表示业务错误码，message 表示该错误的具体信息。每个错误同时也对应一个 HTTP 状态码，比如上述错误码对应了 HTTP 状态码 500（Internal Server  Error）。另外，在出错时，也返回了 reference 字段，该字段包含了可以解决这个错误的文档链接地址。

```json
{
  "code": 100101,
  "message": "Database error",
  "reference": "https://github.com/rebirthmonkey/docs/guide/zh-CN/faq/xxx"
}
```

#### 对内错误描述

- 告诉用户他们可以做什么，而不是告诉他们不能做什么。
- 当声明一个需求时，用 must 而不是 should。例如，must be greater than 0、must match regex '[a-z]+'。
- 当声明一个格式不对时，用 must not。例如，must not contain。
- 当声明一个动作时用 may not。例如，may not be specified when otherField is empty、only name may be specified。
- 引用文字字符串值时，请在单引号中指示文字。例如，ust not contain '..'。
- 当引用另一个字段名称时，请在反引号中指定该名称。例如，must be greater than request。
- 指定不等时，请使用单词而不是符号。例如，must be less than 256、must be greater than or equal to 0 (不要用 larger than、bigger than、more than、higher than)。
- 指定数字范围时，请尽可能使用包含范围。
- 建议 Go 1.13 以上，error 生成方式为 fmt.Errorf("module xxx: %w", err)。
- 错误描述用小写字母开头，结尾不要加标点符号。
- 错误信息是直接暴露给用户的，不能包含敏感信息

## 操作规范

### 命名规范

- 建议告诉用户他们可以做什么，而不是告诉他们不能做什么。
- 当声明一个需求时，用 must 而不是 should。例如，must be greater than 0、must match regex  ‘[a-z]+’。
- 当声明一个格式不对时，用 must not。例如，must not contain。
- 当声明一个动作时用 may not。例如，may not be specified when otherField is empty、only name may be specified。
- 引用文字字符串值时，请在单引号中指示文字。例如，ust not contain  ‘…’。
- 当引用另一个字段名称时，请在反引号中指定该名称。例如，must be greater than  request。
- 指定不等时，请使用单词而不是符号。例如，must be less than 256、must be greater than  or equal to 0 (不要用 larger than、bigger than、more than、higher  than)。
- 指定数字范围时，请尽可能使用包含范围。
- **错误命名用小写字母开头，结尾不要加标点符号**，例如：

```go
  // bad
  errors.New("Redis connection failed")
  errors.New("redis connection failed.")

  // good
  errors.New("redis connection failed")
```

### 生成规范

- 建议 Go 1.13 以上，error 生成方式为 `fmt.Errorf("module  xxx: %w", err)`。
- 自建 Error：如果是自身代码产生的错误，建议在错误最开始处使用 errors.WithCode() 创建一个 withCode 类型的错误。

### 处理规范

Error 处理是为业务逻辑提供非主场景的处理，为了代码今后的可扩展性，建议对每个函数都有 Error 输出。不需要处理时，直接传递给调用函数。

#### 打印日志

- 自建 Error：在处理自建 Error 逻辑时，在上层觉得有需要的地方调用 log 包打印错误，通过 log 包的 caller 功能，可以定位到 log 语句的位置，也能定位到错误发生的位置。

- 第三方库 Error：如果是调用第三方库或接口产生的错误，建议在错误产生的最原始位置打印日志 `log.Errorf("message: %v", err)`打印错误信息，直接返回该位置。

```go
if err := os.Chdir("/root"); err != nil {
    log.Errorf("change dir failed: %v", err)
}
```

#### 函数格式

- error 作为函数的值返回且有多个返回值的时候，error 必须是最后一个参数。
- error 作为函数的值返回，必须对 error 进行处理，或将返回值赋值给明确忽略。对于defer xx.Close() 可以不用显式处理。

```go
func load() error {
  // normal code
}

// bad
load()

// good
 _ = load()
```

#### 处理逻辑

- 尽早进行错误处理，并尽早返回，减少嵌套。

```go
// bad
if err != nil {
  // error code
} else {
  // normal code
}

// good
if err != nil {
  // error handling
  return err
}
// normal code
```

- 错误要单独判断，不与其他逻辑组合判断。

```go
// bad
v, err := foo()
if err != nil || v  == nil {
  // error handling
  return err
}

// good
v, err := foo()
if err != nil {
  // error handling
  return err
}

if v == nil {
  // error handling
  return errors.New("invalid value v")
}
```

- 如果需要在 if 之外使用函数调用的结果，则应采用下面的方式。

```go
// bad
if v, err := foo(); err != nil {
  // error handling
}

// good
v, err := foo()
if err != nil {
  // error handling
}
```

- 上层在处理底层返回的错误时，可以根据需要使用 Wrap() 函数基于该错误封装新的错误信息。
- 只有对会产生错误的函数输出 error。如果判断该函数不会产生错误，则不返回 error。
- 如果函数本身不产生 error，但其调用的第三方函数返回 error，则该函数需要处理并传递 error。

#### Panic使用

- 只在 main 包中使用，只有当程序完全不可运行时使用 panic，例如无法打开文件、无法连接 DB、配置文件错误等导致程序无法正常运行。
- 使用 log.Fatal 来记录错误，这样就可以由 log 来结束程序，或将 panic 抛出的异常记录到日志文件中，方便排查问题。
- 在业务逻辑/包中禁止使用 panic，否则否则在调用该包时会出现莫名的 panic。包内建议采用 error 而不是 panic 来传递错误。

## errors 包

处理错误也离不开错误包，业界有很多优秀的、开源的错误包可供选择，如 Go 标准库自带的 errors 包、github.com/pkg/errors包。但是这些包目前还不支持错误码，很难满足生产级应用的需求。所以，在实际开发中，有必要开发出适合自己错误码设计的错误包。

### 功能需求

设计一个错误包时，需要包含以下功能：

- 支持错误堆栈：
- 支持不同打印格式：如%+v、%v、%s等格式，可以根据需要打印不同丰富度的错误信息
  - %s/%v：只打印可向用户展示的 External Message
  - %-v：打印调用栈中最后一个错误的详细信息
  - %+v：打印调用栈中所有错误的详细信息
  - %#-v：同 %-v，只是以 JSON 的格式输出
  - %#+v：同 %+v，只是以 JSON 的格式输出

- 支持 Wrap/Unwrap 功能：也就是在已有的错误上，追加一些新的信息，如 errors.Wrap(err, "open file failed") 。Wrap 通常用在调用函数中，调用函数可以基于被调函数报错时的错误 Wrap 一些自己的信息、丰富报错信息，方便后期的错误定位。
- 有 is 方法：在实际开发中，经常需要判断某个 error 是否是指定的 error。因为有了 wrapping  error，这样判断就会有问题。因为根本不知道返回的 err 是不是一个嵌套的 error，嵌套了几层。这种情况下，错误包就需要提供 Is 函数：func Is(err, target error) bool当 err 和 target 是同一个，或 err 是一个 wrapping error 时，如果 target 也包含在这个嵌套 error  链中，返回 true，否则返回 fasle。
- 把 error 转为另外一个 error：
- 非格式化创建和格式化创建：

```go
errors.New("file not found") // 非格式化创建
errors.Errorf("file %s not found", "iam-apiserver") // 格式化创建
```

### 实现

通过在文件 errors.go 中增加新的 withCode 结构体，来引入一种新的错误类型，该错误类型可以记录错误码、错误堆栈、错误原因和具体的错误信息。

```go
type withCode struct {
    err   error // error 错误
    code  int // 业务错误码
    cause error // cause error
    *stack // 错误堆栈
}
```

### 使用

```go
package main

import (
    "fmt"

    "github.com/rebirthmonkey/pkg/errors"
    code "github.com/marmotedu/sample-code"
)

func main() {
    if err := getUser(); err != nil {
        fmt.Printf("%+v\n", err)
        // do some business process based on the error type
        if errors.IsCode(err, code.ErrEncodingFailed) {
            fmt.Println("this is a ErrEncodingFailed error")
        }

        // we can also find the cause error
        fmt.Println(errors.Cause(err))
    }
}

func getUser() error {
    if err := queryDatabase(); err != nil {
        return errors.WrapC(err, code.ErrEncodingFailed, "get user 'wukong' failed.")
    }

    return nil
}

func queryDatabase() error {
    return errors.WithCode(code.ErrDatabase, "user 'wukong' not found.")
}
```

上述代码中，通过 WithCode 函数来创建新的 withCode 类型的错误。通过 WrapC 来将一个 error 封装成一个 withCode 类型的错误。通过 IsCode 来判断一个 error 链中是否包含指定的 code。

可能会问，这些错误信息中的 100101 业务错误码，还有 Database error 这种对外展示的报错信息等等，是从哪里获取的？首先，withCode 中包含了 int 类型的业务错误码，例如100101。其次，当使用 github.com/rebirthmonkey/pkg/errors 包时，需要调用 Register 或 MustRegister，将一个 Coder 注册到 github.com/rebithmonkey/pkg/errors 开辟的内存中，数据结构为：

```go
var codes = map[int]Coder{}
```

Coder 是一个接口，定义为：

```go
type Coder interface {
  // HTTP status that should be used for the associated error code.
  HTTPStatus() int

  // External (user) facing error text.
  String() string

  // Reference returns the detail documents for user.
  Reference() string

  // Code returns the code of the coder
  Code() int
}
```


这样 withCode 的 Format 方法，就能够通过 withCode 中的code 字段获取到对应的 Coder，并通过 Coder 提供的HTTPStatus、String、Reference、Code 函数，来获取 withCode 中 code 的详细信息，最后格式化打印。

这里有 2 个注册函数：Register 和 MustRegister，二者唯一区别是：当重复定义同一个业务错误码时，MustRegister 会 panic，这样可以防止后面注册的错误覆盖掉之前注册的错误。在实际开发中，建议使用 MustRegister。

XXX() 和 MustXXX() 的函数命名方式，是一种 Go 代码设计技巧，在 Go 代码中经常使用，例如 Go 标准库中 regexp 包提供的 Compile 和 MustCompile 函数。和 XXX 相比，MustXXX 会在某种情况不满足时 panic。因此使用 MustXXX 的开发者看到函数名就会有一个心理预期：使用不当，会造成程序 panic。

最后，建议在实际的生产环境中，可以使用 JSON 格式打印日志，JSON 格式的日志可以非常方便的供日志系统解析。开发者可以根据需要，选择 `%#-v` 或 `%#+v` 两种格式。

## Lab

- [pkg/errors应用](example1.go)：通过 wrap 层层包裹错误

```bash
go run example1.go
echo $?
```

- [pkg/errors应用](example2.go)：`%+v` 表明直接展示堆栈

```bash
go run example2.go
```

- [rebirthmonkey/pkg/errors自定义错误](example3.go)：自定义错误，并逐层 Wrap。除了 HTTP 状态码，还添加了业务错误码，并且在 Wrap 时可以添加业务错误码。通过自定义错误包 errorcode 首先定义了多个业务错误码，并且通过 registry() 函数注册，再在 example 中使用注册的业务错误码堆栈排错。注意，这边的业务错误码需要在 errcode 包中定义，并且**显示地 register() 到 pkg/errors 包中**。

```bash
go run example3.go
```

- [rebirthmonkey/go/pkg/errors第三方调用错误](example4.go)：在错误处直接打印日志。

```bash
go run example4.go
```

- [rebirthmonkey/go/pkg/errors Gin Server](example5.go)：通过 gin/util 包解析错误并返回。同时，`log.L(ctx).Info()` 会把 context 中的信息传给 log，用于输出更详细信息。

```shell
go run example5.go
# 新建终端，并在新建的终端继续
curl -X GET http://127.0.0.1:8080/ping
```

或者在同一个终端窗口中执行

```bash
go run example5.go &
sleep 5
curl -X GET http://127.0.0.1:8080/ping
```

- [rebirthmonkey/go/pkg/panic Gin Server](example6.go)：展示 log.Panic() 后会中断程序，停止后续操作（包括return）

```shell
go run example6.go 
# 新建终端，并在新建的终端继续
curl -X GET http://127.0.0.1:8080/ping
```

或者在同一个终端窗口中执行

```bash
go run example6.go &
sleep 10
curl -X GET http://127.0.0.1:8080/ping
```

## apiserver

原先示例中有许多错误输出都没处理，而且大部分的错误场景也没有处理，并且也没有加入业务错误码并返回客户端。在新的[示例](80_server/README.md)中，逐步将这些都予以纠正。
