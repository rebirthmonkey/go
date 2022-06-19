package main

import (
	"os"

	"github.com/rebirthmonkey/pkg/errors"
	"github.com/rebirthmonkey/pkg/log"
)

func main() {
	if err := os.Chdir("/root"); err != nil {
		log.Errorf("change dir failed err: %+v", err)

		err2 := errors.Wrap(err, "wrap the error")
		log.Errorf("change dir failed err2: %+v", err2)
	}
}
