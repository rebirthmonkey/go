package main

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
)


func main() {
	var ctx = context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB: 0,
	})


	lpushVal, _ := client.LPush(ctx, "mylist", "val1", "val2", "val3").Result()
	fmt.Println("LPush():", lpushVal)

	rpushVal, _ := client.RPush(ctx, "mylist", "val4", "val5", "val6").Result()
	fmt.Println("RPush():", rpushVal)

	lpopVal, _ := client.LPop(ctx, "mylist").Result()
	fmt.Println("LPop():", lpopVal)

	rpopVal, _ := client.RPop(ctx, "mylist").Result()
	fmt.Println("RPop():", rpopVal)

	lremVal, _ := client.LRem(ctx, "mylist", 1, "val1").Result()
	fmt.Println("LRem():", lremVal)

}

