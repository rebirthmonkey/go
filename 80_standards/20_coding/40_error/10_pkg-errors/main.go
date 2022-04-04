package main

import (
	"log"

	"github.com/pkg/errors"
)

func main() {
	if err := funcA(); err != nil {
		log.Fatalln("main call func got failed:", err)
		return
	}
	log.Println("call func success")
}

func funcA() error {
	if err := funcB(); err != nil {
		return errors.Wrap(err, "funcA call funcB failed")
	}

	return errors.New("funcA called error")
}

func funcB() error {
	return errors.New("funcB called error")
}
