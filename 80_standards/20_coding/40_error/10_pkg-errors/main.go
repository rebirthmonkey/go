package main

import (
	"github.com/pkg/errors"
	"log"
)

func main() {
	if err := funcA(); err != nil {
		log.Fatalf("main call func got failed: %v", err)
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
