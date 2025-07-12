package initialize

import (
	"context"
	"schedule_gateway/global"
	"schedule_gateway/internal/client"
	"schedule_gateway/internal/helper"
	"schedule_gateway/proto/auth"
	"time"

	"go.uber.org/zap"
)

func InitResource() {
	logger := global.Logger
	failAttempt := 0
	for {
		resp, err := client.NewAuthClient().SaveRouteResource(context.Background(), &auth.SaveRouteResourceRequest{
			Items: helper.GetResources(),
		})

		if err != nil || resp == nil || !resp.Success {
			failAttempt++
			logger.ErrorString("Failed to save resources", zap.Error(err), zap.Int("attempt", failAttempt))
			time.Sleep(5 * time.Second)
			continue
		}

		logger.InfoString("Resources saved successfully")
		return
	}
}
