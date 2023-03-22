package main

import (
	"fmt"
	"os/exec"
)

func main() {
	if output, err := exec.Command("/tmp/test.sh", "1", "2").Output(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}
