package auth_controller

import (
	"net/http"
	"schedule_gateway/global"
	client "schedule_gateway/internal/client/auth"
	"schedule_gateway/internal/dtos"
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

func (ac *AuthController) RefreshToken(c *gin.Context) {

	req := ac.buildRefreshTokenRequest(c)
	if req == nil {
		return
	}

	resp, err := ac.authClient.RefreshToken(c, req)
	if err != nil {
		response.InternalServerError(c, "Refresh token failed")
		return
	}

	if resp == nil || resp.Error != nil {
		response.BadRequest(c, "Refresh token fail: "+resp.Error.Message)
		return
	}

	accessTokenCookie := utils.GetHttpOnlyCookie("access_token", resp.AccessToken)
	refreshTokenCookie := utils.GetHttpOnlyCookie("refresh_token", resp.RefreshToken)

	http.SetCookie(c.Writer, accessTokenCookie)
	http.SetCookie(c.Writer, refreshTokenCookie)
	response.Ok(c, "Refresh token", gin.H{
		"status": "Refresh successful",
	})
}

func (ac *AuthController) buildRefreshTokenRequest(c *gin.Context) *auth.RefreshTokenRequest {
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		response.BadRequest(c, "Access token cookie not found")
		return nil
	}

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		response.BadRequest(c, "Refresh token cookie not found")
		return nil
	}

	req := &auth.RefreshTokenRequest{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return req
}

func (ac *AuthController) Logout(c *gin.Context) {
	req := ac.buildLogoutRequest(c)
	if req == nil {
		return
	}

	resp, err := ac.authClient.Logout(c, req)
	if err != nil {
		response.InternalServerError(c, "Logout failed")
		return
	}

	if resp == nil || resp.Error != nil {
		response.BadRequest(c, "Logout failed: "+resp.Error.Message)
		return
	}

	// Clear cookies
	rmAccessTokenCookie := utils.ClearCookie("access_token")
	rmRefreshTokenCookie := utils.ClearCookie("refresh_token")

	http.SetCookie(c.Writer, rmAccessTokenCookie)
	http.SetCookie(c.Writer, rmRefreshTokenCookie)

	response.Ok(c, "Logout", gin.H{
		"status": "logout successful",
	})
}

func (ac *AuthController) buildLogoutRequest(c *gin.Context) *auth.LogoutRequest {
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		response.BadRequest(c, "Access token cookie not found")
		return nil
	}

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		response.BadRequest(c, "Refresh token cookie not found")
		return nil
	}

	req := &auth.LogoutRequest{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return req
}

func (ac *AuthController) GetUserActionsAndResources(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")
	if err != nil || accessToken == "" {
		response.BadRequest(c, "Access token cookie not found")
		return
	}

	req := &auth.GetUserActionsAndResourcesRequest{
		AccessToken: accessToken,
	}

	resp, err := ac.authClient.GetUserActionsAndResources(c, req)
	if err != nil {
		response.InternalServerError(c, "Get user auth info failed")
		return
	}

	if resp == nil || resp.Error != nil {
		response.InternalServerError(c, "Get user auth info failed: "+resp.Error.Message)
		return
	}

	response.Ok(c, "Get user auth info", dtos.UserAuthInfo{
		UserId:      resp.UserId,
		Email:       resp.Email,
		Permissions: resp.Permissions,
	})
}
