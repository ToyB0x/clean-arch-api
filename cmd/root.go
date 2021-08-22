package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/toaru/clean-arch-api/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please specify commands")
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

const configFile = "config"
const configPath = "./config"

func initConfig() {
	GetConfigs(configPath)
}

func GetConfigs(path string) {
	// NOTE: Strength: 3.cmd options > 2.env > 1.config files

	// 2.env (check env on each code execution time)
	viper.AutomaticEnv()

	// read 1.config files
	viper.AddConfigPath(path)
	viper.SetConfigName(configFile)
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
	if err := viper.Unmarshal(&config.Configs); err != nil {
		log.Fatalf("unable to unmarshall config, %v", err)
	}

	// 3.cmd options
	// Write in each command file
}
