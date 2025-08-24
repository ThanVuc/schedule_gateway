package auth_controller

import (
	"net/http"
	"schedule_gateway/global"
	client "schedule_gateway/internal/client/auth"
	"schedule_gateway/internal/utils"
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

	resp, err := ac.authClient.LoginWithGoogle(c, req)
	if err != nil {
		response.InternalServerError(c, "Login failed"+err.Error())
		return
	}

	if resp == nil || resp.Error != nil {
		response.BadRequest(c, "Login failed: "+resp.Error.Message)
		return
	}

	accessTokenCookie := utils.GetHttpOnlyCookie("access_token", resp.AccessToken)
	refreshTokenCookie := utils.GetHttpOnlyCookie("refresh_token", resp.RefreshToken)

	http.SetCookie(c.Writer, accessTokenCookie)
	http.SetCookie(c.Writer, refreshTokenCookie)

	response.Ok(c, "Login", gin.H{
		"status": "login with google successful",
	})
}

func (ac *AuthController) Logout(c *gin.Context) {
	req := &auth.LogoutRequest{
		AccessToken:  "",
		RefreshToken: "",
	}

	resp, err := ac.authClient.Logout(c, req)
	if err != nil {
		response.InternalServerError(c, "Logout failed")
		return
	}

	if !*resp.Success {
		response.BadRequest(c, "Logout failed")
		return
	}

	response.Ok(c, "Login with GitHub called", gin.H{
		"message": "This endpoint is not implemented yet",
	})
}
