package logger

import (
	"github.com/cc-cheunggg/ioc-golang/logger/common"
	"github.com/cc-cheunggg/ioc-golang/logger/zp"
)

var zapFactory = zp.Factory

func Load(config common.Config) error {
	switch config.Engine {
	case "zap":
		zapFactory(config)
		SetZap(zapFactory.NewZap())
	}
	return nil
}
