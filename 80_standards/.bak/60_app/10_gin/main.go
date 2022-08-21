package main

import (
	"fmt"
	"github.com/rebirthmonkey/pkg/log"
)

func main() {
	options.load("app.yaml")

	ginConfig := CreateGINConfig()

	if err := ginConfig.BuildGINConfigFromOptions(&options); err != nil {
		log.Panic(err.Error())
	}

	// 初始化全局logger
	log.Init(ginConfig.Log)
	defer log.Flush()

	fmt.Println(ginConfig.Log)

	ginServer, err := CreateGINServer(ginConfig)
	if err != nil {
		log.Panic(err.Error())
	}

	if err := ginServer.Init(); err != nil {
		log.Panic(err.Error())
	}

	if err := ginServer.Run(); err != nil {
		log.Panic(err.Error())
	}
}
