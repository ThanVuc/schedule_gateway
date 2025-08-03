package auth_controller

import (
	"schedule_gateway/global"
	client "schedule_gateway/internal/client/auth"
	"schedule_gateway/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type AuthController struct {
	logger     log.Logger
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
