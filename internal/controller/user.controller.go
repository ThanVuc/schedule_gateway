package controller

import (
	"schedule_gateway/global"
	"schedule_gateway/internal/client"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	logger     *loggers.LoggerZap
	userClient client.IUserClient
}

func NewUserController() *UserController {
	return &UserController{
		logger:     global.Logger,
		userClient: client.NewUserClient(),
	}
}

func (uc *UserController) GetUserInfo(c *gin.Context) {
	// This method should retrieve user information based on the context
	// For now, we will just return a placeholder response
	uc.logger.InfoString("GetUserInfo called")
	resp, err := uc.userClient.GetUserInfo("123")
	if err != nil {
		panic(response.InternalServerError("Failed to get user info"))
	}

	c.JSON(200, gin.H{
		"user_id": resp.UserId,
		"name":    resp.Name,
		"email":   resp.Email,
	})
}

func (uc *UserController) UpdateUserInfo(c *gin.Context) {
	// This method should update user information based on the context
	// For now, we will just return a placeholder response
	uc.logger.InfoString("UpdateUserInfo called")
	resp, err := uc.userClient.UpdateUserInfo("123", "New Name", "si@gmail.com")
	if err != nil {
		panic(response.InternalServerError("Failed to update user info"))
	}

	c.JSON(200, gin.H{
		"user_id": resp.UserId,
		"name":    resp.Name,
		"email":   resp.Email,
	})
}
