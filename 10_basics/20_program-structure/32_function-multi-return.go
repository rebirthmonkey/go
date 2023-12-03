package main

import (
	"fmt"
)

// divide 函数尝试将一个数除以另一个数。
// 它返回结果和可能发生的错误。
func divide(dividend, divisor float64) (float64, float64, float64, error) {
	if divisor == 0 {
		return dividend, divisor, 0, fmt.Errorf("cannot divide by zero")
	}
	return dividend, divisor, dividend / divisor, nil
}

func main() {
	// 调用 divide 函数并接收其返回的两个值。
	dividend, divisor, result, err := divide(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(dividend, "divides", divisor, "has the result:", result)
	}

	// 尝试一个会导致错误的除法操作。
	dividend, divisor, result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(dividend, "divides", divisor, "has the result:", result)
	}
}
