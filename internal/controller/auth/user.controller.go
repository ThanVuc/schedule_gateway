package auth_controller

import (
	"schedule_gateway/global"
	client "schedule_gateway/internal/client/auth"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type UserController struct {
	logger     log.Logger
	userClient client.UserClient
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
	c.JSON(200, gin.H{
		"user_id": "resp.UserId",
		"name":    "resp.Name",
		"email":   "resp.Email",
	})
}

func (uc *UserController) UpdateUserInfo(c *gin.Context) {
	// This method should update user information based on the context
	// For now, we will just return a placeholder response
	c.JSON(200, gin.H{
		"user_id": "",
		"name":    "resp.Name",
		"email":   "resp.Email",
	})
}
