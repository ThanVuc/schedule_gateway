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

		if failAttempt >= 10 {
			logger.Error("Exceeded maximum retry attempts to save resources", "")
			return
		}

		if err != nil || resp == nil || !resp.Success {
			failAttempt++
			logger.Warn("Failed to save resources", "", zap.Error(err), zap.Int("attempt", failAttempt))
			time.Sleep(sleep)
			sleep = sleep * 2
			continue
		}

		logger.Info("Resources saved successfully", "")
		return
	}
}
