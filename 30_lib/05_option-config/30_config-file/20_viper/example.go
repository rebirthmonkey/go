package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config") // 配置文件名
	viper.SetConfigType("json")   // 配置文件类型，可以是yaml、json、xml
	viper.AddConfigPath(".")      // 配置文件路径
	err := viper.ReadInConfig()   // 读取配置文件信息
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println("port is:", viper.GetInt("port"))
	fmt.Println("RabbitMQ.MQUrl is:", viper.GetString("RabbitMQ.MQUrl"))

	healthz := viper.GetBool("healthz")
	if healthz {
		fmt.Println("healthz is activated")
	}
}
