package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile2 string
	name     string
	age      int
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "MyApp is a sample app",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				name = viper.GetString("name")
			}

			if age == 0 {
				age = viper.GetInt("age")
			}

			fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
		},
	}

	rootCmd.Flags().StringVar(&name, "name", "", "your name")
	rootCmd.Flags().IntVar(&age, "age", 0, "your age")

	cobra.OnInitialize(initConfig2)
	rootCmd.PersistentFlags().StringVar(&cfgFile2, "config", "", "config file (default is $HOME/.myapp.yaml)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig2() {
	viper.SetConfigType("yaml")

	if cfgFile2 != "" {
		viper.SetConfigFile(cfgFile2)
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigName("myapp")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to read config file:", err)
	}
}
