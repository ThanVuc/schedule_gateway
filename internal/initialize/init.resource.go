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

		if err != nil || !resp.Success {
			if failAttempt < 5 {
				failAttempt++
				// sleep for a while before retrying
				time.Sleep(time.Second * 5)
				logger.Warn("Failed to save resources, retrying...", zap.Int("attempt", failAttempt))
				continue
			}
			logger.ErrorString("Failed to save resources after multiple attempts", zap.Error(err))
			return
		}
		logger.InfoString("Resources saved successfully")

		break
	}
}
