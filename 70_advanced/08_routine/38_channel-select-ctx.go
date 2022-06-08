package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"sync"
)

func Rpc(ctx context.Context, url string) error {
	result := make(chan int)
	err := make(chan error)

	go func() {
		// 进行RPC调用，并且返回是否成功，成功通过result传递成功信息，错误通过error传递错误信息
		isSuccess := true
		if isSuccess {
			result <- 1
		} else {
			err <- errors.New("some error happen")
		}
	}()

	select {
	case <-ctx.Done():
		// 其他RPC调用调用失败
		fmt.Println("context.Done")
		return ctx.Err()
	case e := <-err:
		// 本RPC调用失败，返回错误信息
		fmt.Println("RPC:", url, "err", e.Error())
		return e
	case <-result:
		// 本RPC调用成功，不返回错误信息
		fmt.Println("RPC:", url, "success")
		return nil
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.TODO())

	// RPC1调用
	err := Rpc(ctx, "http://rpc_1_url")
	if err != nil {
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(4)

	// RPC2调用
	go func() {
		defer wg.Done()
		err := Rpc(ctx, "http://rpc_2_url")
		if err != nil {
			cancel()
		}
	}()

	// context.Cancel()
	go func() {
		defer wg.Done()
		cancel()
	}()

	// RPC3调用
	go func() {
		defer wg.Done()
		err := Rpc(ctx, "http://rpc_3_url")
		if err != nil {
			cancel()
		}
	}()

	// RPC4调用
	go func() {
		defer wg.Done()
		err := Rpc(ctx, "http://rpc_4_url")
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()
}
