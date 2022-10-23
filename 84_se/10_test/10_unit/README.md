# 单元测试

开发完一段代码后，第一个执行的测试就是单元测试。它可以保证代码是符合预期的，一些异常变动能够被及时感知到。进行单元测试，不仅需要编写单元测试用例，还需要确保代码是可测试的，以及具有一个高的单元测试覆盖率。

## 可测试代码

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

## Mock

- [golang/mock](https://github.com/golang/mock) 是官方提供的 Mock  框架。它实现了基于 interface 的 Mock 功能，能够与 Golang 内置的 testing 包做很好的集成，是最常用的 Mock  工具。golang/mock 提供了 mockgen 工具用来生成 interface 对应的 Mock  源文件。
- [sqlmock](https://github.com/DATA-DOG/go-sqlmock) 可以用来模拟数据库连接。数据库是项目中比较常见的依赖，在遇到数据库依赖时都可以用它。
- [httpmock](https://github.com/jarcoal/httpmock) 可以用来 Mock HTTP 请求。
- [bouk/monkey](https://github.com/bouk/monkey) 猴子补丁，能够通过替换函数指针的方式来修改任意函数的实现。如果 golang/mock、sqlmock 和 httpmock 这几种方法都不能满足我们的需求，我们可以尝试通过猴子补丁的方式来 Mock 依赖。可以这么说，猴子补丁提供了单元测试  Mock 依赖的最终解决方案。

## 覆盖率

- 使用 gotests 工具自动生成单元测试代码，减少编写单元测试用例的工作量，从重复的劳动中解放出来。
- 定期检查单元测试覆盖率，可以通过以下方法来检查：

```bash
go test -race -cover  -coverprofile=./coverage.out -timeout=10m -short -v ./...
go tool cover -func ./coverage.out
```

## Lab

- [单元测试](10_hello/equal_test.go)：单元测试

```bash
cd 10_hello/
go test
```

- [单元测试](12_math-test/math_test.go)：单元测试

```bash
cd 12_math-test
go test
```

- [自动生成测试文件](90_gotests/math_test.go)：单元测试

```bash
cd 90_gotests
gotests -all -w .
# 在 math_test.go 中补全数据，添加测试内容
go test
```
