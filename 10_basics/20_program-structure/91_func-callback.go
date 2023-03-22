package main

import "fmt"

type callback func(int) int

func testCallBack(x int, cb callback) {
	cb(x)
}

func callBack(x int) int {
	fmt.Println("I'm a callback func with parameter = ", x)
	return x
}

func main() {
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Println("I'm another callback func with parameter = ", x)
		return x
	})
}
