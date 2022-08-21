package main

import (
	"github.com/rebirthmonkey/pkg/log"
	"golang.org/x/sync/errgroup"
)

func main() {
	options.load("app.yaml")

	// 全局初始化
	config := CreateConfig()

	if err := config.BuildConfigFromOptions(&options); err != nil {
		log.Panic(err.Error())
	}

	log.Init(config.Log)
	defer log.Flush()

	// GIN Server
	ginConfig := CreateGINConfig()

	if err := ginConfig.BuildGINConfigFromOptions(&options); err != nil {
		log.Panic(err.Error())
	}

	ginServer, err := CreateGINServer(ginConfig)
	if err != nil {
		log.Panic(err.Error())
	}

	if err := ginServer.Init(); err != nil {
		log.Panic(err.Error())
	}

	// GRPC Server
	grpcConfig := CreateGRPCConfig()

	if err := grpcConfig.BuildGRPCConfigFromOptions(&options); err != nil {
		log.Panic(err.Error())
	}

	grpcServer, err := CreateGRPCServer(grpcConfig)
	if err != nil {
		log.Panic(err.Error())
	}

	if err := grpcServer.Init(); err != nil {
		log.Panic(err.Error())
	}

	// launch both servers
	var eg errgroup.Group
	eg.Go(func() error {
		if err := grpcServer.Run(); err != nil  {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		if err := ginServer.Run(); err != nil  {
			return err
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
	}

	//if err := ginServer.Run(); err != nil {
	//	log.Panic(err.Error())
	//}
	//
	//if err := grpcServer.Run(); err != nil {
	//	log.Panic(err.Error())
	//}
}
