package global

import (
	"github.com/kbsonlong/gin-demo/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	CONFIG config.Server
	LOG    *zap.Logger
	CFG    *viper.Viper
)
