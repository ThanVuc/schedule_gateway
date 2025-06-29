package controller

import (
	"schedule_gateway/global"
	"schedule_gateway/internal/client"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/response"

	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	logger           *loggers.LoggerZap
	permissionClient client.PermissionClient
}

func NewPermissionController() *PermissionController {
	return &PermissionController{
		logger:           global.Logger,
		permissionClient: client.NewPermissionClient(),
	}
}

func (pc *PermissionController) GetPermissions(c *gin.Context) {
	response.Ok(c, "GetPermissions called", nil)
}

func (pc *PermissionController) CreatePermission(c *gin.Context) {
	response.Ok(c, "CreatePermission called", nil)
}

func (pc *PermissionController) UpdatePermission(c *gin.Context) {
	response.Ok(c, "UpdatePermission called", nil)
}

func (pc *PermissionController) DeletePermission(c *gin.Context) {
	response.Ok(c, "DeletePermission called", nil)
}

func (pc *PermissionController) AssignPermissionToRole(c *gin.Context) {
	response.Ok(c, "AssignPermissionToRole called", nil)
}
