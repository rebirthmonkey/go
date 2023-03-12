# Builder

在 Go 的程序中，有时希望能在代码无入侵的情况下自动加载、安装某些某些插件。因为该设计模式最初来源于 k8s 的 SchemeBuilder，所以被称为 Builder 设计模式。

Builder 的本质是先把一组用于初始化插件的回调（callback）函数，通过 Register() 函数注册到 Builder（回调函数列表）中。然后通过 AddToXxx() 函数统一执行 Builder 回调函数列表中所有回调函数，从而实现插件的安装。

Builder 设计模式可以解决 2 个问题：

- 热插拔：在不修改主体代码的情况下，动态地加载插件。
- 延时调用：对于某些插件的初始化，可能会依赖其他组件（可能在初始化时还没创建）。通过 Builder 设计模式，可以在 AddToXX() 时才真正调用，到那时依赖组件已经创建完毕。

## 流程

Builder 分为：初始化回调函数列表（Builder）、注册回调函数（Register）和安装所有插件（AddToXxx） 3 步。

### 初始化Builder

Builder 为回调函数列表，它包含 Regisgter() 注册回调函数和 AddToXxx() 安装所有回调函数 2 个方法。

- 初始化 Builder：统一且唯一的回调函数列表，因此必须在注册、安装前先完成初始化。初始化过程在 import mgr 包是自动通过 singleton 设计模式在其 init() 函数中初始化。

### 注册回调函数

为了减少对代码的入侵，注册回调函数是通过加载 import 包来触发 init() ：

```go
func init() {
	Builder.Register(&Xxx{}, &XxxList{})
}
```

### 安装所有插件

一般在 main.go 文件的 init() 中，通过 AddToXxx() 来调用 Builder 中已经加载的所有回调函数，从而实现将所有的插件安装到指定的对象中。

## 示例

具体示例见[这里](main.go)

```bash
go run main.go
```

