package initialize

import (
	"schedule_gateway/global"
	"schedule_gateway/pkg/loggers"
)

func InitLogger() {
	global.Logger = loggers.NewLogger(
		global.Config.Log,
	)
}
