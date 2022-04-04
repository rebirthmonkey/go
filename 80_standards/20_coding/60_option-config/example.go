package main

import "fmt"

func main() {
	opts := NewOptions()

	cfg, err := CreateConfigFromOptions(opts)

	fmt.Println("the config is:", cfg)
	fmt.Println("the error is:", err)
}
