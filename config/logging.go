package config

import (
	"go.uber.org/zap"
)

var Sugar *zap.SugaredLogger = nil

func init(){
	Sugar = zap.NewExample().Sugar()
	defer Sugar.Sync()
}
