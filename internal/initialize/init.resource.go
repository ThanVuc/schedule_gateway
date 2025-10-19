package initialize

import (
	"context"
	"schedule_gateway/global"
	client "schedule_gateway/internal/client/auth"
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
			logger.Warn("Failed to save resources", "", zap.Error(err), zap.Int("attempt", failAttempt))
			time.Sleep(sleep)
			sleep = 5 * time.Second
			continue
		}

		logger.Info("Resources saved successfully", "")
		return
	}
}
