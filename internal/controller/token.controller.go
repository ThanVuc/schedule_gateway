package controller

import (
	"schedule_gateway/global"
	"schedule_gateway/internal/client"
	"schedule_gateway/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type TokenController struct {
	logger      log.Logger
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
