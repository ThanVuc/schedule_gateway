package controller

import (
	"schedule_gateway/global"
	"schedule_gateway/internal/client"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/response"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	logger     *loggers.LoggerZap
	roleClient client.RoleClient
}

func NewRoleController() *RoleController {
	return &RoleController{
		logger:     global.Logger,
		roleClient: client.NewRoleClient(),
	}
}

func (rc *RoleController) GetRoles(c *gin.Context) {
	response.Ok(c, "GetRoles called", nil)
}

func (rc *RoleController) DeleteRole(c *gin.Context) {
	response.Ok(c, "DeleteRole called", nil)
}

func (rc *RoleController) DisableOrEnableRole(c *gin.Context) {
	response.Ok(c, "DisableOrEnableRole called", nil)
}
