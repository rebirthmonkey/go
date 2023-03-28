package main

import (
	"fmt"
	"github.com/rebirthmonkey/pkg/log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config3") // 配置文件名
	viper.SetConfigType("yaml")    // 配置文件类型，可以是yaml、json、xml
	viper.AddConfigPath(".")       // 配置文件路径
	err := viper.ReadInConfig()    // 读取配置文件信息
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	type Options struct {
		Insecure struct {
			Address string `mapstructure:"bind-address"`
			Port    int    `mapstructure:"bind-port"`
		}
		Server struct {
			Mode    string `mapstructure:"mode"`
			Healthz bool   `mapstructure:"healthz"`
		}
		Log log.Options
	}

	var opts Options

	err2 := viper.Unmarshal(&opts)
	if err2 != nil {
		panic(fmt.Errorf("unable to decode into struct: %s \n", err))
	}

	fmt.Println("address is:", opts.Insecure.Address)
	fmt.Println("port is:", opts.Insecure.Port)
	fmt.Println("server mode is:", opts.Server.Mode)
	fmt.Printf("server mode is: %t \n", opts.Server.Healthz)
	fmt.Printf("log DisableCaller is: %t\n", opts.Log.DisableCaller)
}
