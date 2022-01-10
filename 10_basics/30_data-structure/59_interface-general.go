package main

import (
	"fmt"
)

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	names := []string{"stanley", "david", "oscar"}

	// 手动将 []string 转为 []interface{}
	values := make([]interface{}, len(names))
	for i, v := range names {
		values[i] = v
	}
	PrintAll(values)
}
