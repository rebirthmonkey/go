package main

import (
	"errors"
	"fmt"
)

func div(n, m int) (int, error) {
	if m == 0 {
		return 0, errors.New("零不能作为分母")
	}
	return m / n, nil
}

func main() {
	res1, err1 := div(1, 1)
	fmt.Println(res1, err1)

	res2, err2 := div(1, 0)
	fmt.Println(res2, err2)

	e := fmt.Errorf("自定义error")
	fmt.Println(e)
}
