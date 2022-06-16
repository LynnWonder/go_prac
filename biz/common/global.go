package common

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Viper         *viper.Viper
	Logger        *zap.Logger
	SugaredLogger *zap.SugaredLogger
)
