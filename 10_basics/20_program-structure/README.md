# 程序结构

## 源文件

- 文件名
- 包申明：package name
  - package main：可执行文件
- 包引入：
  - import . "fmt"：调用 fmt 包内函数不需要用包名可以直接调用
- 变量&常量
- 函数
  - 内容
    - 语句&表达式
    - 注释
  - 类型
    - main：主函数入口

## 变量&常量

### 变量

#### 声明

- 标准声明

```go
var name string = "XXX"
或
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

- `x.(T)`：x 代表要被判断类型的值，T 代表判断类型
- `T(x)`：类型转换表达式

```go
value, ok := interface{}(container).([]string)
```

- 把container变量的值转换为空接口值的interface{}(container)
- 判断前者的类型是否为切片类型 []string  `.([]string)`

### 常量

### 指针

#### 声明

- 标准声明：var p *int = &b
- 简化声明：p := &b

#### 赋值

- 指针赋值：p = &b
- 值赋值：*p = 2

#### 使用

- a := *p

### type

自定义变量类型

- type xxx int：一般类型申明

- type xxx struct{}：结构型声明

- type xxx interface{}：接口

### nil

空

- if err != nil：如果err不为空，即有错误。

### Lab

- [变量声明](10_variable-constant-declaration.go)
- [变量作用域](12_variable-scope.go)

## 函数

形式：func 函数名 (参数列表) (返回值列表) {函数体}

### 参数列表

- 变量+类型

- ...
  - 函数任意多个该类型的独立参数
  - 不是作为参数列表，而是作为一个个参数直接传！
  
#### 传递参数

- by value：默认是by value
- by pointer

#### Lab

- [函数&传参](30_function.go)

### 返回值列表

- 变量+类型，但变量可以省略
- 如果只有一个输出，可以省略()

### 函数体

#### Loop

##### Condition

###### IF

```go
if <Boolean expression> { 
  <Statements>
}
```

###### IF-ELSE

```go
if <Boolean expression> { 
  <Statements>
} else { 
  <Statements>
}
```

###### Switch

```go
switch {
  case <condition>: 
    <statements>
  case <condition>:
    <statements>
  default:
    <statements>
}
``

##### Loop

The Go for loop has 4forms:

1. for<initialization>;<condition>;<increment/decrement>{} 
2. for<condition>{}-likeawhileloop
3. for{}-aninfinitewhileloop.
4. forwithrange.

###### Range

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

##### Lab

- [loop for](33_loop-for.go)
- [loop switch](34_loop-switch.go)

#### defer

##### 使用场景

- 函数 return 前执行已注册 defer
- 函数执行到最后执行已注册 defer
- panic 前执行已注册 defer

##### 原理

- 先给返回值赋值
- 执行defer语句
- 包裹函数return返回

#### Lab

- [Defer](60_defer.go)
- [Defer](62_defer.go)

#### 异常处理

##### error

程序员可预知的异常

- errors.New()
- fmt.Errorf()
- interface

##### panic

- 程序员无法预知的异常
- recovery
  - 从panic中恢复，并返回panic value
    - 导致panic异常的函数不会继续运行，但能正常返回 Exit Code 0
  - 必须被defer直接调用

#### Lab

- [错误函数](80_error-func.go)
- [Error Interface实现错误](82_error-interface.go)
- [Panic简介](84_panic.go)
- [Panic通过Defer来Recover](85_panic-recover.go)

### 匿名函数

申明：func (参数列表) (返回值列表) {函数体}：没有函数名、只有函数体

#### 赋值

```go
f := func(data int){fmt.Println("hello", data)}
```

- 可以被作为一种类型被赋值给函数类型的变量

```go
f(100) # 调用
```

#### 用途

- Callback 回调函数
