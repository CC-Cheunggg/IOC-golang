package logger

import (
	"sync"

	"github.com/cc-cheunggg/ioc-golang/logger/common"
	"github.com/cc-cheunggg/ioc-golang/logger/zp"
)

var (
	zapFactory = zp.Factory
	zapOnce    sync.Once
)

func Load(config common.Config) error {
	switch config.Engine {
	case "zap":
		zapOnce.Do(func() {
			zapFactory(config)
			SetZap(zapFactory.NewZap())
		})
	}
	return nil
}
