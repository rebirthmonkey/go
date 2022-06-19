package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	var version string
	flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flagSet.StringVar(&version, "version", "abc", "Print version information and quit.")
	flagSet.Parse(os.Args[1:])

	fmt.Println("the version is:", version)
}