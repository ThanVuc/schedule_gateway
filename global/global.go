package global

import (
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/settings"
)

/*
@Author: Sinh
@Date: 2025/6/1
@Description: This package defines global variables that are used throughout the application.
*/
var (
	Config settings.Config
	Logger *loggers.LoggerZap
)
