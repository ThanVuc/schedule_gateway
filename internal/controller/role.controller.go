package controller

import (
	"schedule_gateway/global"
	"schedule_gateway/internal/client"
	"schedule_gateway/internal/mapper"
	"schedule_gateway/internal/utils"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/auth"

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
	req := rc.buildGetRolesRequest(c)
	roles, err := rc.roleClient.GetRoles(c, req)
	if err != nil {
		panic(response.InternalServerError("Failed to get roles: " + err.Error()))
	}

	if roles != nil && roles.Error != nil {
		panic(response.InternalServerError("Failed to get roles: " + roles.Error.Message))
	}

	response.Ok(c, "GetRoles called", mapper.MapRolesToDTO(roles))
}

func (rc *RoleController) GetRole(c *gin.Context) {
	roleId := c.Param("id")
	if roleId == "" {
		panic(response.BadRequest("Role ID is required"))
	}

	req := &auth.GetRoleRequest{RoleId: roleId}
	role, err := rc.roleClient.GetRole(c, req)
	if err != nil {
		panic(response.InternalServerError("Failed to get role: " + err.Error()))
	}

	if role != nil && role.Error != nil {
		panic(response.InternalServerError("Failed to get role: " + role.Error.Message))
	}

	response.Ok(c, "GetRole called", mapper.MapRoleToDTO(role))
}

func (rc *RoleController) DeleteRole(c *gin.Context) {
	roleId := c.Param("id")
	if roleId == "" {
		panic(response.BadRequest("Role ID is required"))
	}
	println("Deleting role with ID:", roleId)

	req := &auth.DeleteRoleRequest{RoleId: roleId}
	resp, err := rc.roleClient.DeleteRole(c, req)

	if !resp.Success {
		panic(response.BadRequest("Failed to delete role: " + *resp.Message))
	}

	if err != nil {
		panic(response.InternalServerError("Failed to delete role: " + err.Error()))
	}

	if resp != nil && resp.Error != nil {
		panic(response.InternalServerError("Failed to delete role: " + resp.Error.Message))
	}

	response.Ok(c, "DeleteRole called", gin.H{"is_success": resp.Success})
}

func (rc *RoleController) DisableOrEnableRole(c *gin.Context) {
	roleId := c.Param("id")
	if roleId == "" {
		panic(response.BadRequest("Role ID is required"))
	}

	req := &auth.DisableOrEnableRoleRequest{RoleId: roleId}
	resp, err := rc.roleClient.DisableOrEnableRole(c, req)

	if !resp.Success {
		panic(response.BadRequest("Failed to disable or enable role: " + *resp.Message))
	}

	if err != nil {
		panic(response.InternalServerError("Failed to disable or enable role: " + err.Error()))
	}

	if resp != nil && resp.Error != nil {
		panic(response.InternalServerError("Failed to disable or enable role: " + resp.Error.Message))
	}

	response.Ok(c, "DisableOrEnableRole called", gin.H{"is_success": resp.Success})
}

func (rc *RoleController) buildGetRolesRequest(c *gin.Context) *auth.GetRolesRequest {
	pageQuery := utils.ToPageQuery(c)
	searchString := c.Query("search")

	req := auth.GetRolesRequest{}
	if searchString != "" {
		req.Search = searchString
	}

	req.PageQuery = pageQuery
	req.Search = searchString

	return &req
}
