package main

import (
	"fmt"
)

type Payment interface {
	order(orderId string) Payment
	getOrderId() string
}

type Wechatpay struct {
	orderId string
}

// 方法接收器
func (o Wechatpay) order(orderId string) Payment {
	fmt.Println("Wechatpay create order： ", o.orderId)
	o.orderId = orderId
	return o
}

func (o Wechatpay) getOrderId() string {
	return o.orderId
}

type Alipay struct {
	orderId string
}

// 指针方法接收器
func (o *Alipay) order(orderId string) Payment {
	fmt.Println("alipay create order： ", o.orderId)
	o.orderId = orderId
	return o
}

func (o *Alipay) getOrderId() string {
	return o.orderId
}

func showPayment(p Payment) {
	fmt.Println("the payment gets order：", p.getOrderId())
}

func main() {
	var wechatpay, alipay Payment

	wechatpay = new(Wechatpay) // 一定要用 new()
	wechatpay.order("20201999293333")
	showPayment(wechatpay)

	alipay = new(Alipay)
	alipay.order("20201999293334")
	showPayment(alipay)
}
