# interface-struct

## 简介

Go 是面向接口语言，不同对象间的交互都必须通过方法调用。对比与 OOP 面向对象，Go 的 interface、interface-var、struct、struct-var 类似于 OOP 的类定义声明、调用声明、实现和实例化：

- 只要定义了 interface 就等于“类的定义声明”
- interface var 则是“类的使用声明”
- struct 可以理解为“类的实现”
- struct var 则是“该类的实例化”

如下图，可以在函数中对 interface 的 var 对象进行操作声明，在具体运行时会 map 到对应的 struct 的实例（var）上。

![image-20220504163654074](figures/image-20220504163654074.png)



### 定义声明：interface

interface 定义了对象的对外的行为，是外部操作对象的抓手。



### 调用声明

在函数中通过操作 interface 的 var 来声明如何使用 interface 以及实现 interface 的 struct。



### 实现：struct+方法

#### 结构

struct 分为显性和隐性 2 部分：

- 显性：数据结构
- 隐性：方法，用于实现的 interface

#### 多态

struct 为 interface 提供不同的实现方式，一个 interface 可以由多个不同的 struct 来实现。通过 struct 内不同的数据结构及业务逻辑来实现相同的方法，从而实现 interface 的多态。

#### 继承

在 Go 中，继承通过 struct 的 embed 来实现。当 struct A embed 另一个 struct B 时，A 将继承 B 的所有数据及方法，并作为 A 自身的一部分。

在继承中，如果 A 对 B 的某个方法不满意，可以重新定义，从而覆盖 B 的该方法。



### 实例化

可以创建多个 struct 的实例来使用 interface 的方法。



## 最佳实践

见 gRPC productinfo_grpc.go 文件

- 包：对应一个对象的容器
- ProductInfoClient interface：大写，定义对外的“类”。
- productInfoClient struct：小写，该 struct 对 interface 的实现，此处定义 struct 内部的数据结。针对之前提及的通过方法调用交互的原则，所以一般 struct 都是 private 的，不对外公开。而方法（interface 的实现）则是 public 的，被外部调用。
- (c *productInfoClient) AddProduct() 等方法：大写，该 struct 对 interface 的某个方法的具体实现
- NewProductInfoClient()：因为 struct（类的定义）是 private，所以只能通过该函数来创建类的对象。
  - 如果某些类的对象是唯一的，则需要在 pkg 中先创建对象 `var p = &productInfoClient{cc} `，然后 NewProductInfoClient() 改为 GetProductInfoClient()，返回唯一对象。



## Lab

- [多态](10_multi-impl/main.go)：多 struct 实现相同 interface

```bash
cd 10_multi-imple
go run main.go interface.go impl1.go impl2.go
```

- [多接口](20_multi-interface/main.go)：一个 struct 实现多个 interface

```bash
cd 20_multi-interface
go run main.go interface1.go interface2.go impl.go
```
