package connection

import (
	"store-hexagon-architecture/internal/utils/logger"

	"github.com/spf13/viper"
)

type Connection struct {
	Config *viper.Viper
	Logger *logger.Logger
}
