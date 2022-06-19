package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"store-admin/internal/apiserver/config"
)

var (
	CONFIG  *config.Server
	Viper   *viper.Viper
	GVA_LOG *zap.Logger
)
