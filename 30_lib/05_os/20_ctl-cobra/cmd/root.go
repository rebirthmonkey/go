package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use: "rmctl",
	Short: "rmtcl command line tool - short",
	Long: "rmtcl command line tool - long",
	Run: runRoot,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func runRoot(cmd *cobra.Command, args []string)  {
	fmt.Printf("execute %s args:%v \n", cmd.Name(), args)
	// TODO 这里处理无参数启动时程序处理
}
