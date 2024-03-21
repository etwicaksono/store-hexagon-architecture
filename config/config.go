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
	config.SetDefault("app.name", "Store-Hexagon-Architecture")
	config.SetDefault("app.version", "0.0.0")
	config.SetDefault("app.host", "localhost")
	config.SetDefault("app.port", 3000)
	config.SetDefault("app.url", "http://localhost:3000")
	config.SetDefault("app.logLevel", "trace")
	config.SetDefault("fiber.idleTimeout", 5)
	config.SetDefault("fiber.writeTimeout", 5)
	config.SetDefault("fiber.readTimeout", 5)
	config.SetDefault("fiber.prefork", false)
	config.SetDefault("fiber.enableStackTrace", true)
	config.SetDefault("otel.grpcHost", "18.141.146.167:4317")
	config.SetDefault("db.address", "localhost:27017")
	config.SetDefault("db.name", "db_store")
	config.SetDefault("db.maxConnectionOpen", 10)
	config.SetDefault("db.maxConnectionIdle", 10)
	config.SetDefault("db.maxConnectionLifetime", "60s")

	err := config.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	return config
}
