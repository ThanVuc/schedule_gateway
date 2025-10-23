package global

import (
	"schedule_gateway/pkg/settings"

	"github.com/thanvuc/go-core-lib/log"
	"github.com/thanvuc/go-core-lib/storage"
)

/*
@Author: Sinh
@Date: 2025/6/1
@Description: This package defines global variables that are used throughout the application.
*/
var (
	Config settings.Config
	Logger log.Logger
	R2Client  *storage.R2Client
)
