package main

import (
	"fmt"
	"github.com/hyahm/goconfig"
)

type RabbitMQ struct {
	MQUrl        string
	Exchange     string
	ExchangeType string
	RoutingKey   string
}

type BaseConfig struct {
	MsgFrequency int64 // 消息发送频率
	RabbitMQ           // MQ信息
}

func ReadBaseConfig(bconfig *BaseConfig, mode string, confFile string) {
	goconfig.InitConf(confFile)

	bconfig.MsgFrequency = goconfig.ReadInt64("Base.messageFrequency", 3)
	bconfig.MQUrl = goconfig.ReadString("RabbitMQ.MQUrl", "")
	bconfig.RabbitMQ.MQUrl = goconfig.ReadString("RabbitMQ.MQUrl", "")
	bconfig.RabbitMQ.Exchange = goconfig.ReadString("RabbitMQ.Exchange", "")
	bconfig.RabbitMQ.ExchangeType = goconfig.ReadString("RabbitMQ.ExchangeType", "")
	bconfig.RabbitMQ.RoutingKey = goconfig.ReadString("RabbitMQ.RoutingKey", "")
}

func main() {
	baseConfigIni := BaseConfig{}
	ReadBaseConfig(&baseConfigIni, "ini", "./config.ini")
	fmt.Println("msgFrequency = ", baseConfigIni.MsgFrequency)
	fmt.Println("mq.MQUrl = ", baseConfigIni.MQUrl)
	fmt.Println("mq.MQUrl = ", baseConfigIni.RabbitMQ.MQUrl) // embedded struct
	fmt.Println("mq.Exchange = ", baseConfigIni.RabbitMQ.Exchange)
	fmt.Println("mq.ExchangeType = ", baseConfigIni.RabbitMQ.ExchangeType)
	fmt.Println("mq.RoutingKey = ", baseConfigIni.RabbitMQ.RoutingKey)

}
