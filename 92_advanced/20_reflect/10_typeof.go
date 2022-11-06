package main

import (
	"fmt"
	"reflect"
)

type FooIF interface {
	DoSomething()
	DoSomethingWithArg(a string)
	DoSomethingWithUnCertenArg(a ...string)
}

type Foo struct {
	A int
	B string
	C struct {
		C1 int
	}
}

func (f *Foo) DoSomething() {
	fmt.Println(f.A, f.B)
}

func (f *Foo) DoSomethingWithArg(a string) {
	fmt.Println(f.A, f.B, a)
}

func (f *Foo) DoSomethingWithUnCertenArg(a ...string) {
	fmt.Println(f.A, f.B, a[0])
}

func (f *Foo) returnOneResult() int {
	return 2
}

func main() {
	var simpleObj Foo
	var pointer2obj = &simpleObj
	var simpleIntArray = [3]int{1, 2, 3}
	var simpleMap = map[string]string{
		"a": "b",
	}
	var simpleChan = make(chan int, 1)
	var x uint64
	var y uint32

	varType := reflect.TypeOf(simpleObj)
	varPointerType := reflect.TypeOf(pointer2obj)

	// 对齐之后要多少容量
	fmt.Println("Align: ", varType.Align())
	// 作为结构体的`field`要对其之后要多少容量
	fmt.Println("FieldAlign: ", varType.FieldAlign())
	// 叫啥
	fmt.Println("Name: ", varType.Name())
	// 绝对引入路径
	fmt.Println("PkgPath: ", varType.PkgPath())
	// 实际上用了多少内存
	fmt.Println("Size: ", varType.Size())
	// 到底啥类型的
	fmt.Println("Kind: ", varType.Kind())

	// 有多少函数
	fmt.Println("NumMethod: ", varPointerType.NumMethod())

	// 通过名字获取一个函数
	m, success := varPointerType.MethodByName("DoSomethingWithArg")
	if success {
		m.Func.Call([]reflect.Value{
			reflect.ValueOf(pointer2obj),
			reflect.ValueOf("sad"),
		})
	}

	// 通过索引获取函数
	m = varPointerType.Method(1)
	m.Func.Call([]reflect.Value{
		reflect.ValueOf(pointer2obj),
		reflect.ValueOf("sad2"),
	})

	// 是否实现了某个接口
	fmt.Println("Implements:", varPointerType.Implements(reflect.TypeOf((*FooIF)(nil)).Elem()))

	//  看看指针多少bit
	fmt.Println("Bits: ", reflect.TypeOf(x).Bits())

	// 查看array, chan, map, ptr, slice的元素类型
	fmt.Println("Elem: ", reflect.TypeOf(simpleIntArray).Elem().Kind())

	// 查看Array长度
	fmt.Println("Len: ", reflect.TypeOf(simpleIntArray).Len())

	// 查看结构体field
	fmt.Println("Field", varType.Field(1))

	// 查看结构体field
	fmt.Println("FieldByIndex", varType.FieldByIndex([]int{2, 0}))

	// 查看结构提field
	fi, success2 := varType.FieldByName("A")
	if success2 {
		fmt.Println("FieldByName", fi)
	}

	// 查看结构体field
	fi, success2 = varType.FieldByNameFunc(func(fieldName string) bool {
		return fieldName == "A"
	})
	if success2 {
		fmt.Println("FieldByName", fi)
	}

	//  查看结构体数量
	fmt.Println("NumField", varType.NumField())

	// 查看map的key类型
	fmt.Println("Key: ", reflect.TypeOf(simpleMap).Key().Name())

	// 查看函数有多少个参数
	fmt.Println("NumIn: ", reflect.TypeOf(pointer2obj.DoSomethingWithUnCertenArg).NumIn())

	// 查看函数参数的类型
	fmt.Println("In: ", reflect.TypeOf(pointer2obj.DoSomethingWithUnCertenArg).In(0))

	// 查看最后一个参数，是否解构了
	fmt.Println("IsVariadic: ", reflect.TypeOf(pointer2obj.DoSomethingWithUnCertenArg).IsVariadic())

	// 查看函数有多少输出
	fmt.Println("NumOut: ", reflect.TypeOf(pointer2obj.DoSomethingWithUnCertenArg).NumOut())

	// 查看函数输出的类型
	fmt.Println("Out: ", reflect.TypeOf(pointer2obj.returnOneResult).Out(0))

	// 查看通道的方向, 3双向。
	fmt.Println("ChanDir: ", int(reflect.TypeOf(simpleChan).ChanDir()))

	// 查看该类型是否可以比较。不能比较的slice, map, func
	fmt.Println("Comparable: ", varPointerType.Comparable())

	// 查看类型是否可以转化成另外一种类型
	fmt.Println("ConvertibleTo: ", varPointerType.ConvertibleTo(reflect.TypeOf("a")))

	// 该类型的值是否可以另外一个类型
	fmt.Println("AssignableTo: ", reflect.TypeOf(x).AssignableTo(reflect.TypeOf(y)))
}
