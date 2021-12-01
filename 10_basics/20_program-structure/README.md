# 程序结构

## 变量&常量

### 变量

#### 声明

- 标准声明

```go
var name string 
name = "XXX"
```

- 合并声明：采用了Go的类型判断

```go
var name = "XXX"
```

- 短变量声明：采用了Go的类型判断，但只能在函数体内使用短变量声明

```go
name := "XXX"
```

<img src="figures/image-20201007094955159.png" alt="image-20201007094955159" style="zoom:33%;" />

#### 作用域

作用域最大的用处就是对程序的访问权限的控制，一个程序的作用域总是会被限制在某个代码块中。

- 代码引用变量总会最优先查找当前代码块中的那个变量
- 如果当前代码块中没有，则会沿着代码块的嵌套关系一层层地向上查找，一直查到当前包
- 如果仍然找不到，那么就会报错了

#### 类型断言 表达式

- `x.(T)`：x代表要被判断类型的值，T带包判断类型
- `T(x)`：类型转换表达式

```go
value, ok := interface{}(container).([]string)
```

- 把container变量的值转换为空接口值的interface{}(container)
- 判断前者的类型是否为切片类型 []string 的 .([]string)

### 常量

### 代码

- [变量声明](10_variable-declaration.go)
- [变量作用域](12_variable-scope.go)

## Loop

### Condition

#### IF

```go
if <Boolean expression> { 
  <Statements>
}
```

#### IF-ELSE

```go
if <Boolean expression> { 
  <Statements>
} else { 
  <Statements>
}
```

#### Switch

```go
switch {
  case <condition>: 
    <statements>
  case <condition>:
    <statements>
  default:
    <statements>
}
```


### Loop

The Go for loop has 4forms:

1. for<initialization>;<condition>;<increment/decrement>{} 
2. for<condition>{}-likeawhileloop
3. for{}-aninfinitewhileloop.
4. forwithrange.

#### Range

The range keyword is used in for loop to iterate data in data structures (arrays, slices, string, maps etc.). 

```go
func main() {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr)
}
```

### 代码

- [loop for]()
- [loop switch]()

## 函数

### 传递参数
#### by value
默认是by value

#### by pointer



