package controller

import (
	"schedule_gateway/pkg/response"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

func NewRoleController() *RoleController {
	return &RoleController{}
}

func (rc *RoleController) GetRoles(c *gin.Context) {
	response.Ok(c, "GetRoles called", nil)
}

func (rc *RoleController) CreateRole(c *gin.Context) {
	response.Ok(c, "CreateRole called", nil)
}

func (rc *RoleController) UpdateRole(c *gin.Context) {
	response.Ok(c, "UpdateRole called", nil)
}

func (rc *RoleController) DeleteRole(c *gin.Context) {
	response.Ok(c, "DeleteRole called", nil)
}

func (rc *RoleController) DisableOrEnableRole(c *gin.Context) {
	response.Ok(c, "DisableOrEnableRole called", nil)
}

func (rc *RoleController) AssignRoleToUser(c *gin.Context) {
	response.Ok(c, "AssignRoleToUser called", nil)
}
