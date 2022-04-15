package main

import (
	"context"
	"fmt"
	//"time"

	redis "github.com/go-redis/redis/v8"
)


func main() {
	var ctx = context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB: 0,
	})

	msetVal, _ := client.MSet(ctx, map[string]interface{}{"key1":"val1", "key2":"val2", "key3":"val3"}).Result()
	fmt.Println("MSet():", msetVal)

	len, _ := client.Append(ctx, "key4", "val4").Result()

	fmt.Println("the length is:", len)

	getVal, _ := client.Get(ctx, "key1").Result()
	fmt.Println("Get():", getVal)

	mgetVal, _ := client.MGet(ctx, "key1", "key2", "key3").Result()
	for k,v := range mgetVal {
		fmt.Printf("k = %v v = %s\n",k,v)
	}

	getRangeVal, _ := client.GetRange(ctx, "key1", 0, -1).Result()
	fmt.Println("GetRange():", getRangeVal)


	delVal, _ := client.Del(ctx, "key2", "key3").Result()
	fmt.Println("Del():", delVal)

	isExists,_ := client.Exists(ctx, "key4").Result()
	fmt.Println("Exists():", isExists)


	strLen, _ := client.StrLen(ctx, "key1").Result()
	fmt.Println("StrLen():", strLen)
}

