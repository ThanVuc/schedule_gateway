package initialize

import (
	"fmt"
	"os"
	"schedule_gateway/global"

	"github.com/spf13/viper"
)

/*
@Author: Sinh
@Date: 2025/6/1
@Description: Load configuration from a YAML file using Viper.
The configuration file file is loaded to the global.Config variable.
*/
func LoadConfig() {
	viper := viper.New()
	// Add both the relative path and current directory for flexibility
	viper.AddConfigPath("./")
	configName := os.Getenv("GO_ENV")
	if configName == "" {
		configName = "dev"
	}

	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")

	// Read in the config file
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// read the config file
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("unable to decode configuration into struct, %v", err))
	}
}
