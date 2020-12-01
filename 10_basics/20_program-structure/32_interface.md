# Interface

## 简介

Go 语言中的接口是一种内置的类型，它定义了一组方法的签名。

## 声明

在接口中我们只能定义方法签名，不能包含成员变量。

```go
type error interface {
	Error() string
}
```



## 实现

```go
type RPCError struct {
	Code    int64
	Message string
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

```



