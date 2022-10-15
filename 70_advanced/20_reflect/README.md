# Reflect

## 简介

反射（reflection）是在 Java 出现后迅速流行起来的一种概念。通过反射，可以获取丰富的类型信息，并可以利用这些类型信息做非常灵活的工作。例如，我们需要一个能统一处理各种类型的函数，但有些信息只有在运行时才能确定（要从文件或网络中获取一些字典数据）。在这种情况下就需要`reflection`，它能帮你：

- 在运行时检查 type
- 在运行时检查/修改/创建 值/函数/结构

总的来说，Go 的`reflection`围绕者三个概念`Types`, `Kinds`, `Values`，所有关于反射的操作都在`reflect`包里。

### (value, type) 对

Go 中的变量包括 type 和 value 两部分：type 包括 static type 和 concrete type，static type 是 Go 默认自带的类型（如 int、string），而 concrete type 是用户定义的 struct 或 interface， runtime 看见的类型。类型断言能否成功，取决于变量的 concrete type，而不是 static type。因此，一个 reader 变量如果它的 concrete type 也实现了 write 方法，它也可以被类型断言为 writer。反射就是建立在 type 之上的，Golang 的指定 type 的变量类型是 static type（也就是指定 int、string 这些的变量），在创建变量的时候就已经确定。而 Go 中反射主要与 interface 类型相关，只有 interface 类型才有反射一说。

value 是实际变量值，type 是实际变量的类型。一个 interface{} 类型的变量包含了 2 个指针，一个指针指向值的concrete type，另外一个指针指向实际的 value。例如，创建类型为 ``*os.File` 的变量，然后将其赋给一个接口变量 r：

```go
tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

var r io.Reader
r = tty
```

接口变量 r 的 (value, type) 中将记录信息 (tty, *os.File)，这个 (value, type) 在接口变量的连续赋值过程中是不变的，将接口变量 r 赋给另一个接口变量 w：

```go
var w io.Writer
w = r.(io.Writer)
```

接口变量 w 的 (value, type) 与 r 的  (value, type)  相同，都是:(tty, *os.File)，即使 w 是空接口类型， (value, type)  也是不变的。

interface 及其  (value, type)  的存在，是 Go 中实现反射的前提，理解了  (value, type) ，就更容易理解反射。反射就是用来检测存储在接口变量内部 (value, concrete type) 对的一种机制。

## 操作

既然反射就是用来检测存储在接口变量内部 (value, concrete type) 对的一种机制，Go 提供了两种方法，可以很容易的访问接口变量内容，分别是 `reflect.ValueOf()` 和 `reflect.TypeOf()`。

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var num float64 = 1.2345

    fmt.Println("type: ", reflect.TypeOf(num))
    fmt.Println("value: ", reflect.ValueOf(num))
}

运行结果:
type:  float64
value:  1.2345
```

- reflect.TypeOf：直接给了 type，如 float64、int、各种pointer、struct 等真实的类型
- reflect.ValueOf：直接给了具体的值，如 1.2345 这个具体数值，或者类似 `&{1 "Allen.Wu" 25}` 这种 struct 的值

### TypeOf

反射来获取值得类型

```go
varType := reflect.TypeOf(var)
```

### ValueOf

除了检查变量的类型，还可以通过`reflection`来读/写/新建一个值。

```go
refVal := reflect.ValueOf(var) 
refPtrVal := reflect.ValueOf(&var)
```

改变该变量的值

```go
refPtrVal.Elem().Set(newRefValue)
```

创建一个新的值

```go
newPtrVal := reflect.New(varType)
```

- 反射开发中最常用的可能就是reflect.DeepEqual



## interface

当执行 `reflect.ValueOf(interface)` 之后，就得到了 interface 的 value。可以通过它本身的 `interface()` 方法获得接口变量的真实内容，然后可以通过类型判断进行转换，转换为原有真实类型。不过，我们可能是已知原有类型，也有可能是未知原有类型，因此，下面分两种情况进行说明。

### 已知原有类型

已知类型后转换为其对应的类型的做法如下，直接通过 `interface` 方法然后强制转换，如 `realValue := value.Interface().(已知的类型)`。

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var num float64 = 1.2345

    pointer := reflect.ValueOf(&num)
    value := reflect.ValueOf(num)

    // 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
    // Golang 对类型要求非常严格，类型一定要完全符合
    // 如下两个，一个是*float64，一个是float64，如果弄混，则会panic
    convertPointer := pointer.Interface().(*float64)
    convertValue := value.Interface().(float64)

    fmt.Println(convertPointer)
    fmt.Println(convertValue)
}

运行结果：
0xc42000e238
1.2345
```

- 转换时，如果转换的类型不完全符合，则直接panic，类型要求非常严格！
- 转换时，要区分是指针还是值。

### 未知原有类型

很多情况下，可能并不知道其具体类型，这时需要进行遍历探测其 Filed 来得知，示例如下:

```go
package main

import (
    "fmt"
    "reflect"
)

type User struct {
    Id   int
    Name string
    Age  int
}

func (u User) ReflectCallFunc() {
    fmt.Println("Wukong.SUN ReflectCallFunc")
}

func main() {

    user := User{1, "Wukong.SUN", 30}

    DoFiledAndMethod(user)

}

// 通过接口来获取任意参数，然后一一揭晓
func DoFiledAndMethod(input interface{}) {

    getType := reflect.TypeOf(input)
    fmt.Println("get Type is :", getType.Name())

    getValue := reflect.ValueOf(input)
    fmt.Println("get all Fields is:", getValue)

    // 获取方法字段
    // 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
    // 2. 再通过reflect.Type的Field获取其Field
    // 3. 最后通过Field的Interface()得到对应的value
    for i := 0; i < getType.NumField(); i++ {
        field := getType.Field(i)
        value := getValue.Field(i).Interface()
        fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
    }

    // 获取方法
    // 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
    for i := 0; i < getType.NumMethod(); i++ {
        m := getType.Method(i)
        fmt.Printf("%s: %v\n", m.Name, m.Type)
    }
}

运行结果：
get Type is : User
get all Fields is: {1 Wukong.SUN 30}
Id: int = 1
Name: string = Wukong.SUN
Age: int = 25
ReflectCallFunc: func(main.User)
```

通过运行结果可以得知获取未知类型的 interface 的具体变量及其类型的步骤为：

1. 先获取 interface 的 reflect.Type，然后通过 NumField 进行遍历
2. 再通过 reflect.Type 的 Field 获取其 Field
3. 最后通过 Field 的 Interface() 得到对应的 value

通过运行结果可以得知获取未知类型的 interface 的所属方法的步骤为：

1. 先获取 interface 的 reflect.Type，然后通过 NumMethod 进行遍历
2. 再分别通过 reflect.Type 的 Method 获取对应的真实的方法
3. 最后对结果取其 Name 和 Type 得知具体的方法名
4. struct 或 struct 嵌套都是一样的判断处理方式

### 设置实际变量值

通过 reflect.Value 设置实际变量的值：reflect.Value 是通过 reflect.ValueOf(X) 获得的，只有当 X 是指针的时候，才可以通过 reflec.Value 修改实际变量 X 的值，即：要修改反射类型的对象就一定要保证其值是 “addressable” 的。示例如下：

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {

    var num float64 = 1.2345
    fmt.Println("old value of pointer:", num)

    // 通过reflect.ValueOf获取num中的reflect.Value，注意，参数必须是指针才能修改其值
    pointer := reflect.ValueOf(&num)
    newValue := pointer.Elem()

    fmt.Println("type of pointer:", newValue.Type())
    fmt.Println("settability of pointer:", newValue.CanSet())

    // 重新赋值
    newValue.SetFloat(77)
    fmt.Println("new value of pointer:", num)

    ////////////////////
    // 如果reflect.ValueOf的参数不是指针，会如何？
    pointer = reflect.ValueOf(num)
    //newValue = pointer.Elem() // 如果非指针，这里直接panic，“panic: reflect: call of reflect.Value.Elem on float64 Value”
}

运行结果：
old value of pointer: 1.2345
type of pointer: float64
settability of pointer: true
new value of pointer: 77
```

- 需要传入的参数是 `* float64` 这个指针，然后可以通过`pointer.Elem()` 去获取所指向的 `Value`，**注意一定要是指针**。
- 如果传入的参数不是指针，而是变量，那么
  - 通过 Elem 获取原始值对应的对象则直接 panic
  - 通过 CanSet 方法查询是否可以设置返回 false
- newValue.CantSet() 表示是否可以重新设置其值，如果输出的是 true 则可修改，否则不能修改，修改完之后再进行打印发现真的已经修改了。
- reflect.Value.Elem() 表示获取原始值对应的反射对象，只有原始对象才能修改，当前反射对象是不能修改的
- 也就是说如果要修改反射类型对象，其值必须是“addressable”（对应的要传入的是指针，同时要通过Elem方法获取原始值对应的反射对象）
- struct 或 struct 的嵌套都是一样的判断处理方式

### 方法调用

通过 `reflect.ValueOf` 来进行方法的调用算是一个高级用法了，前面我们只说到对类型、变量的几种反射的用法，包括如何获取其值、其类型、如果重新设置新值。但是在工程应用中，另外一个常用并且属于高级的用法，就是通过 reflect 来进行方法的调用。比如，要做框架工程的时候，需要可以随意扩展方法，或者说用户可以自定义方法，那么就需要通过 reflect 来搞定。示例如下：

```go
package main

import (
    "fmt"
    "reflect"
)

type User struct {
    Id   int
    Name string
    Age  int
}

func (u User) ReflectCallFuncHasArgs(name string, age int) {
    fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u User) ReflectCallFuncNoArgs() {
    fmt.Println("ReflectCallFuncNoArgs")
}

// 如何通过反射来进行方法的调用？
// 本来可以用u.ReflectCallFuncXXX直接调用的，但是如果要通过反射，那么首先要将方法注册，也就是MethodByName，然后通过反射调动mv.Call

func main() {
    user := User{1, "Allen.Wu", 25}
    
    // 1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value，得到“反射类型对象”后才能做下一步处理
    getValue := reflect.ValueOf(user)

    // 一定要指定参数为正确的方法名
    // 2. 先看看带有参数的调用方法
    methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
    args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)}
    methodValue.Call(args)

    // 一定要指定参数为正确的方法名
    // 3. 再看看无参数的调用方法
    methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
    args = make([]reflect.Value, 0)
    methodValue.Call(args)
}


运行结果：
ReflectCallFuncHasArgs name:  wudebao , age: 30 and origal User.Name: Allen.Wu
ReflectCallFuncNoArgs
```

- 要通过反射来调用起对应的方法，必须要先通过`reflect.ValueOf(interface)` 来获取到 `reflect.Value`
- `reflect.Value.MethodByName`这 MethodByName，需要指定准确真实的方法名字，如果错误将直接 panic，MethodByName 返回一个函数值对应的 reflect.Value 方法的名字。
- []reflect.Value 这个是最终需要调用的方法的参数，可以没有或者一个或者多个，根据实际参数来定。
- reflect.Value 的 Call 这个方法将最终调用真实的方法，参数务必保持一致，如果 `reflect.Value` 不是一个方法，那么将直接panic。
- 本来可以用 `u.ReflectCallFuncXXX` 直接调用的，但是如果要通过反射，那么首先要将方法注册，也就是MethodByName，然后通过反射调用 methodValue.Call。

## 总结

上述详细说明了Go 的反射 reflect 的各种功能和用法：

- 反射可以大大提高程序的灵活性，使得 interface{} 有更大的发挥余地
  - 反射必须结合 interface 才玩得转
  - 变量的 type 要是 concrete type 的（也就是interface变量）才有反射一说
- 反射可以将“接口类型变量”转换为“反射类型对象”
  - 反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息
- 反射可以将“反射类型对象”转换为“接口类型变量”
  - reflect.value.Interface().(已知的类型)
  - 遍历 reflect.Type 的 Field 获取其 Field
- 反射可以修改反射类型对象，但是其值必须是“addressable”
  - 想要利用反射修改对象状态，前提是 interface.data 是 settable，即 pointer-interface
- 通过反射可以“动态”调用方法
- 因为 Go 本身不支持模板，因此在以往需要使用模板的场景下往往就需要使用反射 reflect 来实现

## Lab

- [TypeOf](10_typeof.go)

```bash
go run 10_typeof.go
```


- [ValueOf](20_valueof.go)

```bash
go run 20_valueof.go
```

## Ref

1. [Golang 反射](https://www.jianshu.com/p/1cf328cfe82b)
1. [Golang的反射reflect深入理解和示例](https://www.jianshu.com/p/b46b1ccd2757)
