package main

import (
	"log"

	"github.com/pkg/errors"
)

func main() {
	if err := funcA(); err != nil {
		log.Fatalln("main() call funcA failed:", err)
		return
	}
	log.Println("call main() success")
}

func funcA() error {
	if err := funcB(); err != nil {
		return errors.Wrap(err, "funcA call funcB failed")
	}

	return errors.Errorf("%s error", "funcA")
	//return errors.New("funcA called error")
}

func funcB() error {
	//return errors.New("funcB error")
	return errors.Errorf("%s error", "funcB")
}
