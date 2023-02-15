package main

import (
	"fmt"

	"github.com/rebirthmonkey/go/30_lib/12_version/config"
)

func main() {
	fmt.Println("Config Version is:\t", config.Version)
	fmt.Println("Config BuildTime is:\t", config.BuildTime)
}
