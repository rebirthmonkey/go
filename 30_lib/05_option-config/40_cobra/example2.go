package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "wkctl",
		Short: "wktcl command line tool - short",
		Long:  "wktcl command line tool - long",
		Run:   runRoot,
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(helloCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func runRoot(cmd *cobra.Command, args []string) {
	fmt.Printf("execute %s args:%v \n", cmd.Name(), args)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version subcommand",
	Long:  "this is a 'version' subcommand",
	Run:   runVersion,
}

func runVersion(cmd *cobra.Command, args []string) {
	fmt.Println("version is 1.0.0")
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "hello subcommand",
	Long:  "this is a 'hello' subcommand",
	Args:  cobra.MinimumNArgs(1),
	Run:   runHello,
}

func runHello(cmd *cobra.Command, args []string) {
	fmt.Println("Hello ", args[0])
}
