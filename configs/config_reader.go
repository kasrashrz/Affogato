package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadConfig() Modes {
	var mode Modes
	viper.SetConfigName("devel_config")

	viper.AddConfigPath("./configs")

	viper.AutomaticEnv()

	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&mode)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	return mode
}
