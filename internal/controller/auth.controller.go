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

func (ac *AuthController) Login(c *gin.Context) {

	response.Ok(c, "Login called", gin.H{
		"access_token":  "",
		"refresh_token": "",
	})
}

func (ac *AuthController) Register(c *gin.Context) {
	response.Ok(c, "Register called", nil)
}

func (ac *AuthController) ConfirmEmail(c *gin.Context) {
	response.Ok(c, "ConfirmEmail called", nil)
}

func (ac *AuthController) Logout(c *gin.Context) {
	response.Ok(c, "Logout called", nil)
}

func (ac *AuthController) ResetPassword(c *gin.Context) {
	response.Ok(c, "ResetPassword called", nil)
}

func (ac *AuthController) ForgotPassword(c *gin.Context) {
	response.Ok(c, "ForgotPassword called", nil)
}

func (ac *AuthController) ConfirmForgotPassword(c *gin.Context) {
	response.Ok(c, "ConfirmForgotPassword called", nil)
}
