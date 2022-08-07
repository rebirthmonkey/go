# 数据结构

## 基础

### boolean

```go
var B bool = true
```

### integer

int8、int16、int32、int64：8 表示变量在内存中的大小（以位为单位）

```go
var i int = -1000
```

#### unsigned integer

uint8、uint16、uint32、uint64

```go
var j uint = 1000
```

#### byte

byte = int8

#### rune

rune = int32, used to store Unicode characters

### float

float32、float64

```go
var f float32 = 1.2345
```

### complex

复数：complex64、complex128


### Lab
- [基础数据结构](01_basic.go)


## 复合

### 数组（array）

数组类型的值（简称数组）的长度是固定的，而切片类型的值（切片）是可变长的。

数组是由相同类型元素的集合组成的数据结构，系统会为数组分配一块连续的内存来保存其中的元素，可以利用数组中元素的索引快速访问元素对应的存储地址。常见的数组大多都是一维的线性数组，而多维数组在数值和图形计算领域却有比较常见的应用。

Go 中数组在初始化之后大小就无法改变，存储元素类型相同、但是大小不同的数组类型在 Go 看来也是完全不同的两个数组。

#### 申明/创建

```go
var a [10]int
var b [100]interface{}
```

#### 初始化

```go
arr1 := [3]int{1,2,3} // 可以跳过声明
arr2 := [...]int{1,2,3} // 会被转化为上一个
```

#### 访问/赋值

```
arr1[3]
arr2[2] = 3
```

#### func

- len()：长度
- cap()：容量

### 切片（slice）

更常用的数据结构其实是切片，切片就是动态数组，它的长度并不固定，可以随意向切片中追加元素，而切片会在容量不足时自动扩容。

Slice is an abstraction over Array, it actually uses arrays as an underlying structure. To define a slice, you can declare it as an array without specifying its size.

切片的类型字面量中只有元素的类型，而没有长度。切片的长度可以自动地随着其中元素数量的增长而增长，但不会随着元素数量的减少而减小。

在每个切片的底层数据结构中，一定会包含一个数组。数组可以被叫做切片的底层数组，而切片也可以被看作是对数组的某个连续片段的引用。切片可以想象成有一个窗口，可以通过这个窗口看到一个数组，但是不一定能看到该数组中的所有元素，有时候只能看到连续的一部分元素。

Slice is an abstraction over Array, it actually uses arrays as an underlying structure. 
To define a slice, you can declare it as an array without specifying its size.

The various operations over slice are:

- append(): add the elements to a slice. If the size of underlying array is not enough then automatically a new array is created and content of the old array is copied to it.
- len(): returns the number of elements presents in the slice.
- cap(): returns the capacity of the underlying array of the slice. 
- copy(): the contents of a source slice are copied to a destination slice.
- <SliceName>[start:end]: returns a slice object containing the elements of base slice from index start to end- 1.

#### 申明

```go
var s []int
a := make([]int, 10)
b := make([]int, 0, 10)
c := s[2:4]
```

#### 初始化

```go
s := []int{1,2,3,4,5,6,7,8,9,10} // 跳过声明
```

上例中，由于 c 是通过在 s 上施加切片操作得来的，所以 s 的底层数组就是c的底层数组。

#### func

- len()
- cap()
- append()：在slice尾部添加跟多的元素或切片。在无需扩容时，append函数返回的是指向原底层数组的新切片，而在需要扩容时，append函数返回的是指向新底层数组的新切片。

### 字典（map）

哈希表是除了数组之外，最常见的数据结构，几乎所有的语言都会有数组和哈希表这两种集合元素，有的语言将数组实现成列表，有的语言将哈希表称作结构体或者字典，但是它们是两种设计集合元素的思路，数组用于表示元素的序列，而哈希表示的是键值对之间映射关系，只是不同语言的叫法和实现稍微有些不同。

A map is a collection of Key-Value pairs. Hash-Table is used to store elements in a Map so it is unordered.

字典（map）存储的不是单一值的集合，而是键值对的集合。在 Go 中，为了避免歧义，将键值对换了一种称呼，叫做：“键 - 元素对”。A map is a collection of Key-Value pairs. Hash-Table is used to store elements in a Map so it is unordered.

Go 的字典类型其实是一个哈希表（hash table）的实现。在这个实现中，键和元素的最大不同在于，键的类型是受限的，而元素却可以是任意类型的。可以把键理解为元素的一个索引，可以在哈希表中通过键查找与它成对的那个元素。键和元素的这种对应关系，在数学里就被称为“映射”，这也是“map”这个词的本意，哈希表的映射过程就存在于对键 - 元素对的增、删、改、查的操作之中。

#### 声明/创建

```go
hash := make(map[string]int, 3)
```

#### 初始化

```go
hash := map[string]int{
	"1": 2,
	"3": 4,
	"5": 6,
}
```

#### 访问/赋值

如要在哈希表中查找与某个键值对应的那个元素值，那么需要先把键值作为参数传给这个哈希表。哈希表会先用哈希函数（hash function）把键值转换为哈希值。哈希值通常是一个无符号的整数。一个哈希表会持有一定数量的哈希桶（bucket），这些哈希桶会均匀地储存其所属哈希表收纳的“键-元素”对。因此，哈希表会先用这个键哈希值的低几位去定位到一个哈希桶，然后再去这个哈希桶中查找这个键。由于“键 - 元素”对总是被捆绑在一起存储的，所以一旦找到了键，就一定能找到对应的元素值。随后，哈希表就会把相应的元素值作为结果返回。只要这个“键-元素”对存在哈希表中就一定会被查找到，因为哈希表增、改、删“键-元素”对时的映射过程，与前文所述如出一辙。

```go
hash["1"] = 2
hash["3"] = 4
hash["5"] = 6
```

### 列表（list）

Go 的链表实现在标准库的 container/list 包中，实现了一个双向链表，而 Element 则代表了链表中元素的结构。

List 和 Element 都是结构体类型。结构体类型的特点就是它们的零值都会是拥有特定结构，但是没有任何定制化内容的值，相当于一个空壳。值中的字段也都会被分别赋予各自类型的零值。广义来讲，所谓的零值就是只做了声明，但还未做初始化的变量被给予的缺省值。每个类型的零值都会依据该类型的特性而被设定。

### Ring

container/ring 包中的 Ring 类型实现的是一个循环链表，也就是俗称的环。循环链表一旦被创建，其长度是不可变的。

### 字符串（string）

- A String is a sequence of Unicode character. 
- String is an immutable type variable. 
- Double quotes are used to declare strings.

#### 初始化

```go
var s string = "Hello, World!" // 标准
```

```go
s := "Hello, World!" // 简短
```

#### Operation

- len(mystring) --> 12: Used to find the number of characters in mystring
- “hello”+“world” --> “helloworld”: 2 strings are concatenated into a single string
- “world” == “hello” --> False: Equality can be tested using “==” sign
- “a” < “b” --> True: Unicode value can also be used to “b” < “a” False compare
- mystring[0] --> “h” Indexing: String are indexed same as array.
- mystring[1:4] --> "ell": Slicing

### Lab
- [数组 Array](10_array.go)
- [切片 Slice](11_slice.go)
- [切片 Slice](12_slice2.go)
- [切片 Slice Sum](14_slice-sum.go)：the sum of all the elements of the integer list，given list as an input argument.
- [切片 Slice Sequential Search](15_slice-sequential-search.go)：for unsorted values
- [切片 Slice Binary Search](16_slice-binary-search.go)：for sorted values
- [字典 Map](22_map.go)
- [双向链表 List](25_list.go)
- [字符串 String](29_string.go)


## 结构体（struct）

structures are a collection of multiple data types as a single entity.


### 声明

```go
type Point struct { 
    x int
    y int 
}
```

#### 初始化

```go
point1 := Point{50, 50}
point2 := Point{x: 100, y: 100}
```

#### 访问/赋值

```go
point1.x = 3
a := point2.y
```

### 成员

#### 变量

- point1.X
- 结构体不能包含自身，但可包含自身的结构体指针

#### 方法

##### 方法接收器

在 Go 中，将函数绑定到结构体上，则称该函数是该结构体的方法，其定义的方式是在 func 与函数名间加上具体结构体变量，这个结构体变量称为"方法接收器"。

```go
type Member struct {
    Id     int
    Name   string
    Email  string
    Gender int
    Age    int
}

//普通函数
func setName(m Member,name string){
    m.Name = name
}

//绑定到Member结构体的方法，只有指针才能传递
func (m *Member)setName(name string){
    m.Name = name
}

m := Member{}
m.setName("小明")
fmt.Println(m.Name)//输出小明，如果非指针则为空
```

###### 值接收器

值传递，修改时不会修改原变量

```go
func (p Point) Distance(q Point) float64 {
  ...
}
```

###### 指针接收器

指针传递，修改时会修改值中原变量

```go
func (p *Point) Distance(q Point) float64 {
  ...
}
```

### 匿名组合

通过结构体嵌套实现继承，Point 的成员直接变为 Circle 的成员，而不需要加一层嵌套。

```go
type Circle struct { 
  Point 
}
c := Circle{x:8, y:10}
```

### 结构体指针

```go
var ppoint3 *Point = &point1
```

- (*ppoint3).x 与ppoint3.y 等价：可以通过指针名直接调用值的变量或方法。指针可用来代替值，但值无法代替指针，因为可能有多个指针指向同一个值，Go会自动把指针转换为值。
- 函数的输入/输出一般采用结构体指针

### Lab

- [结构体 struct](50_struct.go)
- [结构体方法 method](51_struct-method.go)


## 接口（interface）

### 简介

Go 中的接口定义了一组**方法**的集合，但这些方法不会在接口上直接实现，而是需要用户自定义的方法来实现。

在接口类型中的方法都是没有实际结构体的，仅仅只是在接口中存放一些方法的签名（签名 = 函数名+参数(类型)+返回值(类型)）。

#### 优点

- 代码扩展性更强了
- 可以解耦上下游的实现：为不同层级的模块提供一个定义好的中间层。这样，上游不再需要依赖下游的具体实现，充分地对上下游进行了解耦。
- 提高了代码的可测性
- 代码更健壮、更稳定了

#### Go语言是面向接口编程

面向接口编程是根据结构体可以执行的操作而不是其所包含的数据来设计抽象。接口可以看做结构体的“基类”，它定义了结构体的行为。结构体则是接口的实现，通过实现所有接口声明的方法来实现该接口。当结构体中包含了该接口，则表示结构体实现该“基类”

### 声明

在接口声明中只能定义方法签名，不能包含变量。

```go
package main

import (
	"fmt"
)

// Shaper 接口类型
type Shaper interface {
	Area() float64
}

// Circle struct类型
type Circle struct {
	radius float64
}

// Circle类型实现Shaper中的方法Area()
func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// Square struct类型
type Square struct {
	length float64
}

// Square类型实现Shaper中的方法Area()
func (s *Square) Area() float64 {
	return s.length * s.length
}

func main() {
	// Circle类型的指针类型实例
	c := new(Circle)
	c.radius = 2.5

	// Square类型的值类型实例
	s := Square{3.2}

	// Sharpe接口实例ins1，它自身是指针类型的
	var ins1 Shaper

  // 将Circle实例c赋值给接口实例ins1，那么ins1中就保存了实例c
	ins1 = c
	fmt.Println(ins1)

	// 使用类型推断将Square实例s赋值给接口实例
	ins2 := s
	fmt.Println(ins2)
  
	fmt.Println(ins1.Area())   // 输出19.625
	fmt.Println(ins2.Area())   // 输出10.24
}
```

当用户自定义的结构体实现了接口中定义的方法时，那么自定义结构体的实例可以赋值给接口类型的实例，这个赋值过程使得接口实例中保存了用户自定义结构体实例。如：结构体实例 c 与接口实例 ins1 包含了两个地址：【1】

- 第一部分是实例的类型信息
- 第二个部分是实例自身信息

<img src="figures/image-20211126113650944.png" alt="image-20211126113650944" style="zoom: 25%;" />

### 实现

当一个结构体为一个接口中所有的方法提供定义时，它被称为实现该接口。而判断一个结构体是否实现了一个接口是完全是自动地。

### 多态

通过接口定义“基类”，多个结构体实现接口中定义的所有方法，从而实现这个“基类”。当通过结构体调用该接口的方法时，所有符合该接口的结构体都可被调用，从而实现多态调用。

#### 基类（接口）定义

重点是实现方法的主体，无论是结构体还是结构体指针

```go
animals := []Animal{Dog{}, Cat{}} # Animal是个接口
animals := []Animal{Dog{}, &Cat{}} # 既有结构体，也有接头体指针
func TotalPerimeter(shapes ...Shape) float64 {...}
TotalPerimeter(a, b, c, d) # 实现Shape的结构体或结构体指针
```



### interface{}

interface{} 作为所有类的“基类”被使用

```go
func PrintAll(vals []interface{}) {...}
```

可将 []string 转为 []interface{} 类型



## 类型转换

### 基本

b=type(a) 如：`b = int32(a)`

### 结构体

<img src="figures/image-20220410132213918.png" alt="image-20220410132213918" style="zoom:50%;" />

#### 子类->基类

struct->interface

- 值接收器：animal1 := Animal(monkey)

- 指针接收器：animal2 := Animal(&cat)

#### 子类->子类

struct->struct，因为两个结构体都实现了该接口（方法）

- 值接收器：pig := Pig(monkey)

- 指针接收器：

  - 结构体：dog1 := Dog(cat)

  - 结构体指针：dog2 := (*Dog)(&cat)

#### 基类->子类

interface->struct

- 值接收器：monkey2, ok := animal1.(Monkey)

- 指针接收器

  - 结构体

    - 原路返回（animal2 从 cat 转换过来）：cat2, ok := animal2.(*Cat) 

    - 非原路返回（不可行）：dog3 , ok := animal2.(*Dog) 

  - 结构体指针：太复杂，不考虑

### Lab
- [基础数据结构转换](80_basic-type-trans.go)
- [字符串String转换](81_basic-string-trans.go)
- [Interface-Struct转换](85_interface-trans.go)

## 非结构化

无法预知数据结构的数据类型属于非结构化范畴，在Go中，无法通过构建预定的struct数据结构来序列化或反序列化，在 Go 中，一般通过如下数据结构来解决：

```go
var result map[string]interface{}
```

其中 `interface{}` 是 Go 中的通用类型，可以转换为任何类型。在 Go 中，可以通过断言来进行类型转换：

```go
if description, ok := result["description"].(string); ok {
  fmt.Println(description)
}
```





## Ref

1. [Golang 之 interface接口全面理解](https://blog.csdn.net/Webben/article/details/110448404?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-0.queryctr&spm=1001.2101.3001.4242.1&utm_relevant_index=2)

