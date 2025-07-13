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
	sleep := time.Second
	for {
		resp, err := client.NewAuthClient().SaveRouteResource(context.Background(), &auth.SaveRouteResourceRequest{
			Items: helper.GetResources(),
		})

		if err != nil || resp == nil || !resp.Success {
			failAttempt++
			logger.ErrorString("Failed to save resources", zap.Error(err), zap.Int("attempt", failAttempt))
			time.Sleep(sleep)
			sleep *= 2
			if sleep > 30*time.Second {
				logger.ErrorString("Max retry attempts reached, giving up on saving resources", zap.Int("attempt", failAttempt))
				return
			}
			continue
		}

		logger.InfoString("Resources saved successfully")
		return
	}
}
