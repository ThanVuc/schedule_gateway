package team_controller

import (
	"schedule_gateway/global"
	team_client "schedule_gateway/internal/client/team"
	dtos "schedule_gateway/internal/dtos/team_service"
	"schedule_gateway/internal/utils"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/common"
	"schedule_gateway/proto/team_service"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

type UserController struct {
	logger log.Logger
	client team_client.UserClient
}

func NewUserController() *UserController {
	return &UserController{
		logger: global.Logger,
		client: team_client.NewUserClient(),
	}
}

func (uc *UserController) GetUserInfo(ctx *gin.Context) {
	req := &common.EmptyRequest{}
	resp, err := uc.client.GetUserInfo(ctx, req)
	if err != nil {
		uc.logger.Error("Failed to get user info: ", "", zap.Error(err))
		ctx.JSON(500, gin.H{"error": "Failed to get user info"})
		return
	}

	if resp.GetError() != nil {

		uc.logger.Error("Failed to get user info: ", "", zap.String("code", resp.Error.Code), zap.String("message", *resp.Error.Details))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	userInfo := uc.buildGetUserInfoRespon(resp)

	response.Ok(ctx, "User info retrieved successfully", gin.H{
		"user": userInfo,
	})

}

func (uc *UserController) buildGetUserInfoRespon(resp *team_service.GetUserInfoResponse) *dtos.TeamUserDTO {
	if resp == nil {
		return nil
	}

	return &dtos.TeamUserDTO{
		Email:                resp.GetEmail(),
		UseEmailNotification: resp.GetUseEmailNotification(),
		UseAppNotification:   resp.GetUseAppNotification(),
		CreateAt:             utils.TimestampToISO8601(resp.GetCreatedAt()),
	}

}

func (uc *UserController) NotificationConfiguration(ctx *gin.Context) {
	var dto dtos.NotificationConfigurationRequestDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		uc.logger.Error("Invalid request body: ", "", zap.Error(err))
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	req := &team_service.NotificationConfigurationRequest{
		UseEmailNotification: dto.UseEmailNotification,
		UseAppNotification:   dto.UseAppNotification,
	}

	resp, err := uc.client.NotificationConfiguration(ctx, req)
	if err != nil {
		uc.logger.Error("Failed to update notification configuration: ", "", zap.Error(err))
		ctx.JSON(500, gin.H{"error": "Failed to update notification configuration"})
		return
	}

	if resp.GetError() != nil {
		uc.logger.Error("Failed to update notification configuration: ", "", zap.String("code", resp.Error.Code), zap.String("message", *resp.Error.Details))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	response.Ok(ctx, "Notification configuration updated successfully", gin.H{
		"success": resp.GetSuccess(),
	})
}
