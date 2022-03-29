package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command {
	Use: "version",
	Short: "version subcommand",
	Long: "this is a 'version' subcommand",
	Run: runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(cmd *cobra.Command, args []string)  {
	fmt.Println("version is 1.0.0")
}
