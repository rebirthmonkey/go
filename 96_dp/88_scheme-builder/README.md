# SchemeBuilder

在 Go 的程序中，有时希望能在代码无入侵的情况下自动加载、安装某些某些插件。因为该设计模式最初来源于 k8s 的 SchemeBuilder，所以被称为 SchemeBuilder 设计模式。

SchemeBuilder 的本质是先把一组用于初始化插件的回调（callback）函数，也被称为 Builder 注册到一个统一的 BuilderScheme（注册表，也就是回调函数列表）中。然后在后续统一执行回调函数列表中的所有回调函数，从而实现插件的安装。因此，SchemeBuilder 分为：初始化注册表（BuilderScheme）、注册（Register）插件和安装（AddToManager）所有插件 3 步。

## 初始化注册表

- Builder：每个可被注册、安装的插件都被抽象为一个 Builder，它包含 Regisgter() 加载和 AddToManager() 安装 2 个方法。
- 初始化 BuilderScheme：统一且唯一的回调函数列表，因此必须在注册、安装插件前先完成初始化。初始化过程在 import mgr 包是自动通过 singleton 设计模式初始化 BuilderScheme 注册表。

## 注册插件

注册插件又可分为 init 和 Register 2 步：

- init：为了减少对代码的入侵，加载通过 import 插件包，触发该包的 init() 函数来触发注册。
- Register：同样为了减少对代码的入侵，在 Register 时会调用 registry 包的 Register() 函数，将回调函数注册到事先已经为 registry 包指定的 BuilderScheme 中。

## 安装所有插件

- AddToManager：在应用的主流程中，会通过 registry.AddToManager() 来调用已经加载的所有回调函数，从而实现将所有的插件安装到指定的对象中。

## 示例

具体示例见[这里](main.go)

```bash
go run main.go
```

