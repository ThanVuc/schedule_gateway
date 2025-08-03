package user_controller

import (
	"schedule_gateway/global"
	client "schedule_gateway/internal/client/user"
	"schedule_gateway/proto/user"

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

func (uc *UserController) GetUserProfile(c *gin.Context) {
	resp, err := uc.userClient.GetUserProfile(c, &user.GetUserProfileRequest{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get user profile"})
		return
	}

	println("User Profile:", resp)

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
