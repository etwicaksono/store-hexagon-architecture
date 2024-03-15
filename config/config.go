package config

import (
	"fmt"
	"store-hexagon-architecture/internal/utils/projectpath"

	"github.com/spf13/viper"
)

func GetConfig() *viper.Viper {
	// Config
	config := viper.New()
	config.AddConfigPath(fmt.Sprint(projectpath.Root, "/config"))
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	// set up default value
	config.SetDefault("environment", "development")
	config.SetDefault("fiber.prefork", false)

	err := config.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	return config
}
