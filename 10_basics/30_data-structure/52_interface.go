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

// 指针方法接收器
func (o *Alipay) order(orderId string) Payment{
	fmt.Println("alipay create order： ", o.orderId)
	o.orderId = orderId
	return o
}

func (o *Alipay) getOrderId() string{
	return o.orderId
}

func main(){
	var alipay Payment
	alipay = new(Alipay) // 一定要用 new
	alipay.order( "20201999293334")
	fmt.Println("alipay get order：", alipay.getOrderId())
}
