# Reflect

## 简介

反射（reflection）是在Java出现后迅速流行起来的一种概念。通过反射，你可以获取丰富的类型信息，并可以利用这些类型信息做非常灵活的工作。有时我们需要些一个函数，希望它有能力统一处理各种类型的函数，而这些函数类型可能无法共享同一个接口。但是有些时候，需要搞一些运行时才能确定的东西，例如要从文件或网络中获取一些字典数据，又或者要搞一些不同类型的数据。在这种情况下，`reflection`就有用啦。reflection 能够让你拥有以下能力：

- 在运行时检查 type
- 在运行时检查/修改/创建 值/函数/结构

总的来说，Go 的`reflection`围绕者三个概念`Types`, `Kinds`, `Values`，所有关于反射的操作都在`reflect`包里。



## 操作

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
- 

## Lab

- [TypeOf](10_typeof.go)



## Ref

1. [Golang 反射](https://www.jianshu.com/p/1cf328cfe82b)
2. 