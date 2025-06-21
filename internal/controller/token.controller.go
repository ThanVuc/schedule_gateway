package controller

import (
	"schedule_gateway/pkg/response"

	"github.com/gin-gonic/gin"
)

type TokenController struct{}

func NewTokenController() *TokenController {
	return &TokenController{}
}

func (tc *TokenController) RefreshToken(c *gin.Context) {
	response.Ok(c, "RefreshToken called", nil)
}

func (tc *TokenController) RevokeToken(c *gin.Context) {
	response.Ok(c, "RevokeToken called", nil)
}
