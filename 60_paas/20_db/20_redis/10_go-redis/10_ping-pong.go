package main

import (
	"context"
	"fmt"
	"log"

	redis "github.com/go-redis/redis/v8"
)


func main() {
	var ctx = context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB: 0,	//redis默认有0-15共16个数据库，这里设置操作索引为0的数据库
	})

	pong,err := client.Ping(ctx).Result()

	if err != nil {
		log.Fatal(err)
	}

	if pong != "PONG" {
		log.Fatal("fail to connect to the redis server")
	} else {
		fmt.Println("success to connect to the redis server")
	}
}
