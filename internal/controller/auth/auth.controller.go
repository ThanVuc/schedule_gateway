package auth_controller

import (
	"schedule_gateway/global"
	client "schedule_gateway/internal/client/auth"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/auth"

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
	var body gin.H
	if err := c.ShouldBindJSON(&body); err != nil {
		response.BadRequest(c, "Invalid request body")
		return
	}

	googleAccessToken, ok := body["google_access_token"].(string)
	if !ok || googleAccessToken == "" {
		response.BadRequest(c, "Google access token is required")
		return
	}

	req := &auth.LoginWithGoogleRequest{
		GoogleAccessToken: googleAccessToken,
	}

	ac.authClient.LoginWithGoogle(c, req)

	response.Ok(c, "Login called", gin.H{
		"google_access_token": googleAccessToken,
	})
}
