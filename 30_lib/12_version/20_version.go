package main

import (
	"fmt"

	"github.com/rebirthmonkey/go/pkg/version"
)

func main() {
	fmt.Println("Git Version is:\t", version.GitVersion)
	fmt.Println("Build Date is:\t", version.BuildDate)
}
