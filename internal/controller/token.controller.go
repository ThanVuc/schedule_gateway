package controller

import (
	"schedule_gateway/global"
	"schedule_gateway/internal/client"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/response"

	"github.com/gin-gonic/gin"
)

type TokenController struct {
	logger      *loggers.LoggerZap
	tokenClient client.TokenClient
}

func NewTokenController() *TokenController {
	return &TokenController{
		logger:      global.Logger,
		tokenClient: client.NewTokenClient(),
	}
}

func (tc *TokenController) RefreshToken(c *gin.Context) {
	response.Ok(c, "RefreshToken called", nil)
}

func (tc *TokenController) RevokeToken(c *gin.Context) {
	response.Ok(c, "RevokeToken called", nil)
}
