package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)

	var version string
	flagSet.StringVar(&version, "version", "abc", "Print version information and quit.")
	flagSet.Parse(os.Args[1:])

	fmt.Println("the version is:", version)
}
