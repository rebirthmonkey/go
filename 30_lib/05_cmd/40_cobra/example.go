package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var version bool
	var rootCmd = &cobra.Command{
		Use:   "wkctl",
		Short: "wktcl command line tool - short",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("inside rootCmd Run with args: %v\n", args)
			if version {
				fmt.Printf("version:1.0\n")
			}
		},
	}

	flags := rootCmd.Flags()
	flags.BoolVarP(&version, "version", "v", false, "Print version information and quit")
	_ = rootCmd.Execute()
}
