package service

import (
	"fmt"
	"github.com/spf13/viper"
	"jirou/models"
	"os"
)

func GetConfig() (config models.Config) {
	loadConfig()

	return readConfig()
}

func loadConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.jirou")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Printf("Could not load config file: %s\n", fmt.Errorf("%w", err))
		os.Exit(1)
	}
}

func readConfig() (config models.Config) {
	err := viper.Unmarshal(&config)
	if err != nil { // Handle errors reading the config file
		fmt.Printf("Could not read configuration file: %s\n", fmt.Errorf("%w", err))
		os.Exit(1)
	}

	return
}
