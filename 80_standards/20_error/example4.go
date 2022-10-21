package main

import (
	"os"

	"github.com/rebirthmonkey/go/pkg/errors"
	"github.com/rebirthmonkey/go/pkg/log"
)

func main() {
	if err := os.Chdir("/rooot"); err != nil {
		log.Errorf("change dir failed err: %+v", err)

		err2 := errors.Wrap(err, "wrap the error")
		log.Errorf("change dir failed err2: %+v", err2)
	}
}
