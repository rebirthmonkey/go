package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

func main() {
	if err := funcA2(); err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	log.Println("call main() success")
}

func funcA2() error {
	if err := funcB2(); err != nil {
		return err
	}

	return errors.Errorf("%s error", "funcA")
}

func funcB2() error {
	return errors.Errorf("%s error", "funcB")
}
