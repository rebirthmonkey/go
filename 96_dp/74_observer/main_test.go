package main

import "testing"

func TestObserver(t *testing.T) {
	// 创建张三的信用卡
	creditCard := NewCreditCard("张三")
	// 短信通知订阅信用卡消费及逾期消息
	creditCard.Subscribe(new(shortMessage), ConsumeType, ExpireType)
	// 电子邮件通知订阅信用卡账单及逾期消息
	creditCard.Subscribe(new(email), BillType, ExpireType)
	// 电话通知订阅信用卡逾期消息，同时逾期消息通过三种方式通知
	creditCard.Subscribe(new(telephone), ExpireType)

	creditCard.Consume(500.00) // 信用卡消费
	creditCard.Consume(800.00) // 信用卡消费
	creditCard.SendBill()      // 信用卡发送账单
	creditCard.Expire()        // 信用卡逾期

	// 信用卡逾期消息取消电子邮件及短信通知订阅
	creditCard.Unsubscribe(new(email), ExpireType)
	creditCard.Unsubscribe(new(shortMessage), ExpireType)
	creditCard.Consume(300.00) // 信用卡消费
	creditCard.Expire()        // 信用卡逾期
}
