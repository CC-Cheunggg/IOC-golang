package logger

import (
	"github.com/alibaba/ioc-golang/logger/common"
	"github.com/alibaba/ioc-golang/logger/zp"
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
