package main

import (
	"fmt"
)

type Animal interface {
	GetName() string
}

type Cat struct {
	name string
}

func (c *Cat)GetName() string {
	return c.name
}

type Dog struct {
	name string
}

func (d *Dog)GetName() string {
	return d.name
}

type Monkey struct {
	name string
}


func (m Monkey)GetName() string {
	return m.name
}

type Pig struct {
	name string
}

func (p Pig)GetName() string {
	return p.name
}


func main() {
	monkey := Monkey{name: "wukong"}
	cat := Cat{name: "kitty"}

	// 子类->基类：struct->interface
	// 值接收器
	animal1 := Animal(monkey)
	fmt.Println("I'm animal1", animal1.GetName())

	// 指针接收器
	animal2 := Animal(&cat) // 括号里需要传递一个*Cat类型而不能是Cat类型，因为是*Cat类型实现了GetName()方法，而不是Cat类型
	fmt.Println("I'm animal2", animal2.GetName())

	// 子类->子类：struct->struct
	// 值接收器
	pig := Pig(monkey)
	fmt.Println("I'm pig", pig.GetName())

	// 指针接收器
	// 结构体
	dog1 := Dog(cat) // 括号里需要传递一个Cat类型，因为Cat=Dog, *Cat=*Dog
	fmt.Println("I'm dog1", dog1.GetName())

	// 结构体指针
	dog2 := (*Dog)(&cat)
	fmt.Println("I'm dog2", dog2.GetName())  // 如上所述，*Cat = *Dog

	// 基类->子类：interface->struct
	// 值接收器
	monkey2, ok := animal1.(Monkey)
	if ok {
		fmt.Println("convert animal1 to monkey", monkey2.GetName())
	} else {
		fmt.Println("can not convert animal1 to monkey")
	}

	// 指针接收器：结构体
	cat2, ok := animal2.(*Cat) // 类型断言，左边必须是接口类型的对象，当接口对象的实际类型和要转换的目标类型匹配时，转换成功，否则失败
	if ok {
		fmt.Println("convert animal2 to cat", cat2.GetName())
	} else {
		fmt.Println("can not convert animal2 to cat")
	}

	// 指针接收器：结构体
	dog3 , ok := animal2.(*Dog) // 类型断言，接口对象的实际类型和要转换的目标类型不匹配
	if ok {
		fmt.Println("convert animal2 to dog - " + dog3.GetName())
	} else {
		fmt.Println("can not convert animal2 to dog")
	}
}

