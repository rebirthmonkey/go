# Array & Slice
## Array
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



## Slice

更常用的数据结构其实是切片，切片就是动态数组，它的长度并不固定，我们可以随意向切片中追加元素，而切片会在容量不足时自动扩容。

Slice is an abstraction over Array, it actually uses arrays as an underlying structure. To define a slice, you can declare it as an array without specifying its size.

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
c := s[0:4]
```

### 初始化

```go
s := []int{1,2,3,4,5,6,7,8,9,10} // 跳过声明
```

