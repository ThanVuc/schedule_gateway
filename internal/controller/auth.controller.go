package controller

import (
	"schedule_gateway/global"
	"schedule_gateway/internal/client"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	logger     *loggers.LoggerZap
	authClient client.AuthClient
}

func NewAuthController() *AuthController {
	return &AuthController{
		logger:     global.Logger,
		authClient: client.NewAuthClient(),
	}
}

func (ac *AuthController) LoginWithGoogle(c *gin.Context) {

	ac.authClient.LoginWithGoogle(c, nil)

	response.Ok(c, "Login called", gin.H{
		"access_token":  "",
		"refresh_token": "",
	})
}
