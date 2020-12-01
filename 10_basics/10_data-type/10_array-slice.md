# Array & Slice
## 简介
数组类型的值（简称数组）的长度是固定的，而切片类型的值（切片）是可变长的。

## Array
<<<<<<< HEAD
数组是由相同类型元素的集合组成的数据结构，系统会为数组分配一块连续的内存来保存其中的元素，可以利用数组中元素的索引快速访问元素对应的存储地址。常见的数组大多都是一维的线性数组，而多维数组在数值和图形计算领域却有比较常见的应用。

Go 语言中数组在初始化之后大小就无法改变，存储元素类型相同、但是大小不同的数组类型在 Go 语言看来也是完全不同的两个数组。

An array is a collection of variables of the same data type

### 申明/创建

```go
var a [10]int
var b [100]interface{}
```

### 初始化

```go
arr1 := [3]int{1,2,3} // 可以跳过声明
arr2 := [...]int{1,2,3} // 会被转化为上一个
```

### 访问/赋值

```
arr1[3]
arr2[2] = 3
```
=======
数组的长度在声明的时候必须给定，并且之后不会再改变。可以说，数组的长度是其类型的一部分。比如，[1]string和[2]string就是两个不同的数组类型。其实可以把切片看做是对数组的一层简单的封装，因为在每个切片的底层数据结构中，一定会包含一个数组。数组可以被叫做切片的底层数组，而切片也可以被看作是对数组的某个连续片段的引用。

An array is a collection of variables of the same data type

### func

- len()：长度
- cap()：容量
>>>>>>> de31d1d040ee13ebba7944684f3bebb3f7196274



## Slice

<<<<<<< HEAD
更常用的数据结构其实是切片，切片就是动态数组，它的长度并不固定，我们可以随意向切片中追加元素，而切片会在容量不足时自动扩容。

Slice is an abstraction over Array, it actually uses arrays as an underlying structure. To define a slice, you can declare it as an array without specifying its size.
=======
切片的类型字面量中只有元素的类型，而没有长度。切片的长度可以自动地随着其中元素数量的增长而增长，但不会随着元素数量的减少而减小。

在每个切片的底层数据结构中，一定会包含一个数组。数组可以被叫做切片的底层数组，而切片也可以被看作是对数组的某个连续片段的引用。切片可以想象成有一个窗口，可以通过这个窗口看到一个数组，但是不一定能看到该数组中的所有元素，有时候只能看到连续的一部分元素。

Slice is an abstraction over Array, it actually uses arrays as an underlying structure. 
To define a slice, you can declare it as an array without specifying its size.
>>>>>>> de31d1d040ee13ebba7944684f3bebb3f7196274

The various operations over slice are:
- append(): add the elements to a slice. If the size of underlying array is not enough then automatically a new array is created and content of the old array is copied to it.
- len(): returns the number of elements presents in the slice.
- cap(): returns the capacity of the underlying array of the slice. 
- copy(): the contents of a source slice are copied to a destination slice.
- <SliceName>[start:end]: returns a slice object containing the elements of base slice from index start to end- 1.

### 申明

```go
var s []int
a := make([]int, 10)
b := make([]int, 0, 10)
c := s[2:4]
```

<<<<<<< HEAD
### 初始化

```go
s := []int{1,2,3,4,5,6,7,8,9,10} // 跳过声明
```
=======
上例中，由于c是通过在b上施加切片操作得来的，所以b的底层数组就是c的底层数组。

### func

- len()
- cap()
- append()：在slice尾部添加跟多的元素或切片。在无需扩容时，append函数返回的是指向原底层数组的新切片，而在需要扩容时，append函数返回的是指向新底层数组的新切片。


>>>>>>> de31d1d040ee13ebba7944684f3bebb3f7196274

