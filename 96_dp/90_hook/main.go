package main

import (
	"fmt"
)

type Options struct {
	ConfigPath string
	ServerURL  string
}

func Main(opts Options, preHook func(), postHook func()) {
	if preHook != nil {
		preHook()
	}

	fmt.Println("Start the application!!!")

	if postHook != nil {
		postHook()
	}
}

func main() {
	opts := Options{}
	Main(opts, nil, nil)

	Main(opts,
		func() {
			fmt.Println("preStart Hook!")
		}, func() {
			fmt.Println("postStart Hook!")
		})
}
