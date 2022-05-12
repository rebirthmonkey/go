# interface-struct

## 对象

在 Go 中，不同对象间的交互都必须通过方法调用。因此只需要定义了 interface，就等于“类的声明”。而 struct 可以理解为“类的定义”，var struct 则是“类的对象”。

对比与面向对象，GO的面向接口中的interface、struct、var类似于面向对象的类申明、定义、对象。

如下图，可以在函数中对 interface 的对象 var1 进行操作，在具体运行时会 map 到对应的 struct1 上。

![image-20220504163654074](figures/image-20220504163654074.png)



### interface（声明）

interface 定义了对象的对外的行为，是外部操作对象的抓手。



### struct（定义）

#### 结构

struct 分为显性和隐性 2 部分：

- 显性：数据结构
- 隐性：方法，用于实现的 interface

#### 多态

struct 为 interface 提供不同的实现方式，一个 interface 可以由多个不同的 struct 来实现。通过 struct 内不同的数据结构及业务逻辑来实现相同的方法，从而实现 interface 的多态。

#### 继承

在 Go 中，继承通过 struct 的 embed 来实现。当 struct A embed 另一个 struct B 时，A 将继承 B 的所有数据及方法，并作为 A 自身的一部分。

在继承中，如果 A 对 B 的某个方法不满意，可以重新定义，从而覆盖 B 的该方法。

## 最佳实践

见 gRPC productinfo_grpc.go 文件

- 包：对应一个对象的容器
- ProductInfoClient interface：大写，定义对外的“类”。
- productInfoClient struct：小写，该 struct 对 interface 的实现，此处定义 struct 内部的数据结。针对之前提及的通过方法调用交互的原则，所以一般 struct 都是 private 的，不对外公开。而方法（interface 的实现）则是 public 的，被外部调用。
- (c *productInfoClient) AddProduct() 等方法：大写，该 struct 对 interface 的某个方法的具体实现
- NewProductInfoClient()：因为 struct（类的定义）是 private，所以只能通过该函数来创建类的对象。
  - 如果某些类的对象是唯一的，则需要在 pkg 中先创建对象 `var p = &productInfoClient{cc} `，然后 NewProductInfoClient() 改为 GetProductInfoClient()，返回唯一对象。





