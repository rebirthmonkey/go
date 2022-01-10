package main

import (
	"fmt"
	"github.com/hyahm/goconfig"
)

type RabbitMQ struct {
	MQUrl string
	Exchange string
	ExchangeType string
	RoutingKey string
}

type BaseConfig struct {
	MsgFrequency int64 // 消息发送频率
	RabbitMQ // MQ信息
}

func ReadBaseConfig(bconfig *BaseConfig, confFile string) {
	//goconfig.InitConf(confFile, goconfig.INI)
	goconfig.InitConf(confFile, goconfig.JSON)
	bconfig.MsgFrequency = goconfig.ReadInt64("Base.messageFrequency", 3)
	bconfig.RabbitMQ.MQUrl = goconfig.ReadString("RabbitMQ.MQUrl", "")
	bconfig.RabbitMQ.Exchange = goconfig.ReadString("RabbitMQ.Exchange", "")
	bconfig.RabbitMQ.ExchangeType = goconfig.ReadString("RabbitMQ.ExchangeType", "")
	bconfig.RabbitMQ.RoutingKey = goconfig.ReadString("RabbitMQ.RoutingKey", "")
}

func main() {
	baseConfig := BaseConfig{}
	// ReadBaseConfig(&baseConfig, "./11_config.ini")
	ReadBaseConfig(&baseConfig, "./11_config.json")
	fmt.Printf("mq.MQUrl = %s \t mq.Exchange = %s \t mq.ExchangeType = %s \t mq.RoutingKey = %s\n", baseConfig.RabbitMQ.MQUrl, baseConfig.RabbitMQ.Exchange, baseConfig.RabbitMQ.ExchangeType, baseConfig.RabbitMQ.RoutingKey)
	fmt.Printf("msgFrequency = %d\n", baseConfig.MsgFrequency)
}
