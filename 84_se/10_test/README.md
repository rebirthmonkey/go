# 测试

## testing包

### 命名规范

testing 包需要在编写测试文件、测试函数、测试变量时遵循一定的规范。这些规范有些来自于官方，有些则来自于社区。

- 测试文件的命名规范：Go 的测试文件名必须以 _test.go 结尾。例如，如果有一个名为 person.go 的文件，那它的测试文件必须命名为 person_test.go。这样做是因为，Go 需要区分哪些文件是测试文件。这些测试文件可以被 go test 命令行工具加载，用来测试编写的代码，但会被 Go 的构建程序忽略掉，因为 Go 程序的运行不需要这些测试代码。
- 包的命名规范：Go 的测试可以分为白盒测试和黑盒测试。
  - 白盒测试：将测试和生产代码放在同一个 Go 包中，这使我们可以同时测试 Go 包中可导出和不可导出的标识符。当编写的单元测试需要访问 Go 包中不可导出的变量、函数和方法时，就需要编写白盒测试用例。在白盒测试中，Go 的测试包名称需要跟被测试的包名保持一致，例如：person.go 定义了一个 person 包，则 person_test.go 的包名也要为 person，这也意味着 person.go 和 person_test.go 都要在同一个目录中。
  - 黑盒测试：将测试和生产代码放在不同的 Go 包中。这时，仅可以测试 Go 包的可导出标识符。这意味着测试包将无法访问生产代码中的任何内部函数、变量或常量。在黑盒测试中，Go 的测试包名称需要跟被测试的包名不同，但仍然可以存放在同一个目录下。比如，person.go 定义了一个 person 包，则 person_test.go 的包名需要跟 person 不同，通常命名为 person_test 。如果不是需要使用黑盒测试，在做单元测试时要尽量使用白盒测试。一方面，这是 go test 工具的默认行为；另一方面，使用白盒测试，可以测试和使用不可导出的标识符。
- 函数的命名规范：测试用例函数必须以 Test、Benchmark、Example 开头，例如 TestXxx、BenchmarkXxx、ExampleXxx，Xxx部分为任意字母数字的组合，首字母大写。这是由 Go 语言和 go test 工具来进行约束的，Xxx一般是需要测试的函数名。
- 变量的命名规范：Go 语言和 go test 没有对变量的命名做任何约束。但是，在编写单元测试用例时，还是有一些规范值得去遵守。单元测试用例通常会有一个实际的输出，在单元测试中，会将预期的输出跟实际的输出进行对比，来判断单元测试是否通过。为了清晰地表达函数的实际输出和预期输出，可以将这两类输出命名为expected/actual，或者 got/want。

## 单元测试

开发完一段代码后，第一个执行的测试就是单元测试。它可以保证代码是符合预期的，一些异常变动能够被及时感知到。进行单元测试，不仅需要编写单元测试用例，还需要确保代码是可测试的，以及具有一个高的单元测试覆盖率。

### go test

单元测试用例函数以 Test 开头，例如 TestXxx 或 Test_xxx（ Xxx 部分为任意字母数字组合，首字母大写）。函数参数必须是 *testing.T，可以使用该类型来记录错误或测试状态。可以调用 testing.T 的 Error 、Errorf 、FailNow  、Fatal 、FatalIf 方法，来说明测试不通过。调用 Log 、Logf 方法来记录测试信息。

<img src="figures/image-20220828110018192.png" alt="image-20220828110018192" style="zoom: 33%;" />

go test 命令自动搜集所有的测试文件，也就是格式为 *_test.go 的文件，从中提取全部测试函数并执行。go test 还支持下面三个参数。

- -v：显示所有测试函数的运行细节

```shell
$ go test -v
=== RUN   TestAbs
--- PASS: TestAbs (0.00s)
=== RUN   TestMax
--- PASS: TestMax (0.00s)
PASS
ok      github.com/marmotedu/gopractise-demo/31/test    0.002s
```

- -run < regexp>：指定要执行的测试函数

```shell
$ go test -v -run='TestA.*'
=== RUN   TestAbs
--- PASS: TestAbs (0.00s)
PASS
ok      github.com/marmotedu/gopractise-demo/31/test    0.001s
```

- -count N：指定执行测试函数的次数

```shell
$ go test -v -run='TestA.*' -count=2
=== RUN   TestAbs
--- PASS: TestAbs (0.00s)
=== RUN   TestAbs
--- PASS: TestAbs (0.00s)
PASS
ok      github.com/marmotedu/gopractise-demo/31/test    0.002s
```

#### TestMain()

有时候，在做测试的时候，可能会在测试之前做些准备工作，例如创建数据库连接等；在测试之后做些清理工作，例如关闭数据库连接、清理测试文件等。这时，可以在 _test.go 文件中添加 TestMain 函数，其入参为 *testing.M。

TestMain 是一个特殊的函数（相当于 main  函数），测试用例在执行时，会先执行 TestMain 函数，然后可以在 TestMain 中调用 m.Run() 函数执行普通的测试函数。在 m.Run() 函数前，可以编写准备逻辑，在 m.Run() 后面可以编写清理逻辑。

### Mock测试

一般来说，单元测试中是不允许有外部依赖的，那么也就是说，这些外部依赖都需要被模拟。在 Go 中，一般会借助各类 Mock  工具来模拟一些依赖。

#### GoMock

GoMock 是由 Golang 官方开发维护的测试框架，实现了较为完整的基于 interface 的 Mock  功能，能够与 Golang 内置的 testing 包良好集成，也能用于其他的测试环境中。GoMock 测试框架包含了 GoMock 包和  mockgen 工具两部分，其中 GoMock 包用来完成对象生命周期的管理，mockgen 工具用来生成 interface 对应的 Mock 类源文件。

gomock  支持以下输入参数匹配：

- gomock.Any()：可以用来表示任意的入参。
- gomock.Eq(value)：用来表示与 value  等价的值。
- gomock.Not(value)：用来表示非 value 以外的值。
- gomock.Nil()：用来表示 None 值。

如下例，GetBody 接收 2 个入参

```go
mockSpider.EXPECT().GetBody(gomock.Any(), gomock.Eq("admin")).Return("go1.8.3")
```

gomock 的返回值如下：

```go
func (c *Call) After(preReq *Call) *Call // After声明调用在preReq完成后执行
func (c *Call) AnyTimes() *Call // 允许调用次数为 0 次或更多次
func (c *Call) Do(f interface{}) *Call // 声明在匹配时要运行的操作
func (c *Call) MaxTimes(n int) *Call // 设置最大的调用次数为 n 次
func (c *Call) MinTimes(n int) *Call // 设置最小的调用次数为 n 次
func (c *Call) Return(rets ...interface{}) *Call //  // 声明模拟函数调用返回的值
func (c *Call) SetArg(n int, value interface{}) *Call // 声明使用指针设置第 n 个参数的值
func (c *Call) Times(n int) *Call // 设置调用次数为 n 次
```

#### 其他mock工具

- sqlmock：可以用来模拟数据库连接。数据库是项目中比较常见的依赖，在遇到数据库依赖时都可以用它。
- httpmock：可以用来 Mock HTTP 请求。
- bouk/monkey：猴子补丁，能够通过替换函数指针的方式来修改任意函数的实现。如果 golang/mock、sqlmock 和 httpmock 这几种方法都不能满足需求，可以尝试用猴子补丁的方式来 Mock 依赖。可以这么说，猴子补丁提供了单元测试 Mock 依赖的最终解决方案。

### 覆盖率

检查单元测试覆盖率

- 生成测试覆盖率数据 coverage.out

```shell
$ go test -coverprofile=coverage.out
```

-  分析覆盖率文件

```shell
$ go tool cover -func=coverage.out
```

- 生成HTML格式的分析文件

```shell
$ go tool cover -html=coverage.out -o coverage.html
```

### 自从生成

使用 gotests 工具自动生成单元测试代码，减少编写单元测试用例的工作量，从重复的劳动中解放出来。通过 gotests 可以自动生成单元测试文件及测试函数。

```shell
$ gotests -all -w .
```

## 示例测试

### go test

示例测试以 Example 开头，没有输入和返回参数，通常保存在 example_test.go 文件中。示例测试可能包含以 Output: 或者 Unordered output: 开头的注释，这些注释放在函数的结尾部分。Unordered output: 开头的注释会忽略输出行的顺序。

执行 go  test 命令时，会执行这些示例测试，并且 go test  会将示例测试输出到标准输出的内容，跟注释作对比。如果相等，则示例测试通过测试；如果不相等，则示例测试不通过测试。

```go
func ExampleMax() {
    fmt.Println(Max(1, 2))
    // Output: 2
}
```

#### 命名规范

示例测试需要遵循一些命名规范，因为只有这样，Godoc 才能将示例测试和包级别的标识符进行关联。

- 测试函数名：以 Example 开头，后面可以不跟任何字符串，也可以跟函数名、类型名或者类型_方法名，中间用下划线_连接。

## 接口测试

### 接口 Mock

- [golang/mock](https://github.com/golang/mock) 是官方提供的 Mock  框架。它实现了基于 interface 的 Mock 功能，能够与 Golang 内置的 testing 包做很好的集成，是最常用的 Mock  工具。golang/mock 提供了 mockgen 工具用来生成 interface 对应的 Mock  源文件。
- [sqlmock](https://github.com/DATA-DOG/go-sqlmock) 可以用来模拟数据库连接。数据库是项目中比较常见的依赖，在遇到数据库依赖时都可以用它。
- [httpmock](https://github.com/jarcoal/httpmock) 可以用来 Mock HTTP 请求。
- [bouk/monkey](https://github.com/bouk/monkey) 猴子补丁，能够通过替换函数指针的方式来修改任意函数的实现。如果 golang/mock、sqlmock 和 httpmock 这几种方法都不能满足我们的需求，我们可以尝试通过猴子补丁的方式来 Mock 依赖。可以这么说，猴子补丁提供了单元测试  Mock 依赖的最终解决方案。

## 性能测试

### go test

性能测试的用例函数必须以 Benchmark 开头，例如 BenchmarkXxx 或 Benchmark_Xxx（ Xxx  部分为任意字母数字组合，首字母大写）。函数参数必须是 *testing.B，函数内以 b.N 作为循环次数，其中 N 会在运行时动态调整，直到性能测试函数可以持续足够长的时间，以便能够可靠地计时。

go test 命令默认不会执行性能测试函数，需要通过指定参数 -bench 来运行性能测试函数。-bench 后可以跟正则表达式，选择需要执行的性能测试函数，例如 go test -bench=".*" 表示执行所有的压力测试函数。

- benchmem：输出内存分配统计。指定了-benchmem 参数后，执行结果中又多了两列： 0 B/op，表示每次执行分配了多少内存（字节），该值越小，说明代码内存占用越小；0  allocs/op，表示每次执行分配了多少次内存，该值越小，说明分配内存次数越少，意味着代码性能越高。
- benchtime：指定测试时间和循环执行次数（格式需要为 Nx，例如 100x）。
- cpu：指定 GOMAXPROCS。
- timeout：指定测试函数执行的超时时间。

## 代码可测试性

如果要对函数 A 进行测试，并且 A 中的所有代码均能够在单元测试环境下按预期被执行，那么函数 A 的代码块就是可测试的。

但在单元测试环境中：1/ 可能无法连接数据库；2/ 可能无法访问第三方服务。如果函数 A 依赖数据库连接、第三方服务，那么在单元测试环境下执行单元测试就会失败，函数就没法测试，函数是不可测的。解决方法是将依赖的数据库、第三方服务等抽象成接口，在被测代码中调用接口的方法，在测试时传入 mock 类型，从而将数据库、第三方服务等依赖从具体的被测函数中解耦出去。

<img src="figures/0cef423ec1a4f06f6f4715bd0b9f4497-20220324101751804.png" alt="img" style="zoom:50%;" />

为了提高代码的可测性，降低单元测试的复杂度，对 function 和 mock 的要求是：

- 要尽可能减少 function 中的依赖，让 function 只依赖必要的模块。编写一个功能单一、职责分明的函数，会有利于减少依赖。
- 依赖模块应该是易 Mock 的。

举个简单的例子：

- 不可测试代码：ListPosts 函数是不可测试的。因为 ListPosts 函数中调用了client.ListPosts() 方法，该方法依赖于一个 gRPC 连接。

```go
package post

import "google.golang.org/grpc"

type Post struct {
  Name    string
  Address string
}

func ListPosts(client *grpc.ClientConn) ([]*Post, error) {
  return client.ListPosts()
}
```

- 可测试代码：ListPosts 函数入参为 Service 接口类型，只要我们传入一个实现了 Service 接口类型的实例，ListPosts 函数即可成功运行。

```go
package main

type Post struct {
  Name    string
  Address string
}

type Service interface {
  ListPosts() ([]*Post, error)
}

func ListPosts(svc Service) ([]*Post, error) {
  return svc.ListPosts()
}
```

测试代码如下：

```go
package main

import "testing"

type fakeService struct {
}

func NewFakeService() Service {
  return &fakeService{}
}

func (s *fakeService) ListPosts() ([]*Post, error) {
  posts := make([]*Post, 0)
  posts = append(posts, &Post{
    Name:    "colin",
    Address: "Shenzhen",
  })
  posts = append(posts, &Post{
    Name:    "alex",
    Address: "Beijing",
  })
  return posts, nil
}

func TestListPosts(t *testing.T) {
  fake := NewFakeService()
  if _, err := ListPosts(fake); err != nil {
    t.Fatal("list posts failed")
  }
}
```

## Lab

### 单元测试

- [HelloWorld](10_unit/10_hello/equal_test.go)：

```bash
cd 10_unit/10_hello
go test
```

- [Struct-based](10_unit/11_struct-based/sqrt_test.go)：内部建立 struct list 统一测试

```bash
cd 10_unit/11_struct-based
go test
echo $?
```

- [Math Test](10_unit/12_math-test/math_test.go)：使用 TestMain() 函数做 test 前后准备

```bash
cd 10_unit/12_math-test/
go test
```

- `mock`：为接口 Spider 创建一个 mock

```shell
$ mockgen -destination spider/mock/mock_spider.go -package spider -source spider/spider.go
# 或者
$ cd spider
$ go generate # 通过 spider.go 的注释 //go:generate mockgen -destination ... 生成mock
$ cd .. 
$ go test
```

- [自动生成测试文件](10_unit/90_gotests/math_test.go)：通过 gotests 自动生成测试文件

```bash
cd 10_unit/90_gotests/
gotests -all -w .
# 在 math_test.go 中补全数据，添加测试内容
go test
```

### 示例测试

- [math示例测试](13_example/10_math/example_test.go)：示例测试只对比输出结果

```bash
cd 13_example/10_math/
go test
```

- [大型App示例测试](13_example/20_app/example_app_test.go)：

```bash
cd 13_example/20_app/
go test
```

### 性能测试

- [math性能测试](20_benchmark/12_math-test/math_test.go)：测试 math 函数性能

```bash
cd 20_benchmark/12_math-test/
go test -bench=".*"
```

## apierver 示例

在本章，会为 apiserver 示例添加额外的测试，具体包容“单元测试”和“接口测试”两部分。

### 单元测试

- 在 `Service` 内添加单元测试
- 运行单元测试

```shell
cd 80_server/apiserver/user/service/v1
go test
```

### 接口测试

- 添加接口测试脚本到[此处](80_server/test/api)
- 运行接口测试脚本

```bash
cd 80_server/
./test/api/test.sh api::test::user
```
