package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

type Config struct {
	Name string
	Age  int
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "MyApp is a sample app",
		Run: func(cmd *cobra.Command, args []string) {
			config := &Config{}
			if err := viper.Unmarshal(config); err != nil {
				fmt.Println("Failed to unmarshal config file:", err)
				return
			}
			fmt.Printf("Hello, %s! You are %d years old.\n", config.Name, config.Age)
		},
	}

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config.yaml)")
	//viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(cfgFile)
	//viper.AddConfigPath(".")
	//viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to read config file:", err)
	}
}
