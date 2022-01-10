package main

import (
	"fmt"
)

type Payment interface {
	order(orderId string) Payment
	getOrderId() string
}

type Alipay struct{
	orderId string
}

// Alipay 方法，*Alipay 指针传值
func (o *Alipay) order(orderId string) Payment{
	fmt.Println("alipay create order", o.orderId)
	o.orderId = orderId
	return o
}

// Alipay 方法，不能直接访问属性，添加方法
func (o *Alipay) getOrderId() string{
	return o.orderId
}

func main(){
	var alipay Payment
	alipay = new(Alipay) // 一定要用 new
	alipay.order( "20201999293334")
	fmt.Println( alipay.getOrderId())
}
