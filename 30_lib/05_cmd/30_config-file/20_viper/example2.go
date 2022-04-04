package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main(){
	viper.SetConfigName("config2")  // 配置文件名
	viper.SetConfigType("json") // 配置文件类型，可以是yaml、json、xml
	viper.AddConfigPath(".")  // 配置文件路径
	err := viper.ReadInConfig()  // 读取配置文件信息
	if err != nil{
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}


	type config struct {
		Port int
		Name string
		PathMap string `mapstructure:"path_map"`
	}

	var Conf config

	err2 := viper.Unmarshal(&Conf)
	if err2 != nil {
		panic(fmt.Errorf("unable to decode into struct: %s \n", err))
	}

	fmt.Println("port is:", Conf.Port)
	fmt.Println("name is:", Conf.Name)
	fmt.Println("PathMap is:", Conf.PathMap)
}

