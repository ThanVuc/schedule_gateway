package controller

import (
	"schedule_gateway/global"
	"schedule_gateway/internal/client"
	"schedule_gateway/internal/dtos"
	"schedule_gateway/internal/utils"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type RoleController struct {
	logger     log.Logger
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
		response.InternalServerError(c, "Failed to get roles: "+err.Error())
		return
	}

	if roles != nil && roles.Error != nil {
		response.InternalServerError(c, "Failed to get roles: "+roles.Error.Message)
		return
	}

	response.Ok(c, "GetRoles called", dtos.Roles{
		Items:             roles.Roles,
		TotalRoles:        roles.TotalRoles,
		Root:              roles.Root,
		NonRoot:           roles.NonRoot,
		RootPercentage:    float32(roles.RootPercentage),
		NonRootPercentage: float32(roles.NonRootPercentage),
		TotalItems:        roles.PageInfo.TotalItems,
		TotalPages:        roles.PageInfo.TotalPages,
		PageSize:          roles.PageInfo.PageSize,
		Page:              roles.PageInfo.Page,
		HasPrev:           roles.PageInfo.HasPrev,
		HasNext:           roles.PageInfo.HasNext,
	})
}

func (rc *RoleController) GetRole(c *gin.Context) {
	roleId := c.Param("id")
	if roleId == "" {
		response.BadRequest(c, "Role ID is required")
		return
	}

	req := &auth.GetRoleRequest{RoleId: roleId}
	role, err := rc.roleClient.GetRole(c, req)
	if err != nil {
		response.InternalServerError(c, "Failed to get role: "+err.Error())
		return
	}

	if role != nil && role.Error != nil {
		response.InternalServerError(c, "Failed to get role: "+role.Error.Message)
		return
	}

	response.Ok(c, "GetRole called", role)
}

func (rc *RoleController) DeleteRole(c *gin.Context) {
	roleId := c.Param("id")
	if roleId == "" {
		response.BadRequest(c, "Role ID is required")
		return
	}
	println("Deleting role with ID:", roleId)

	req := &auth.DeleteRoleRequest{RoleId: roleId}
	resp, err := rc.roleClient.DeleteRole(c, req)

	if !resp.Success {
		response.BadRequest(c, "Failed to delete role: "+*resp.Message)
		return
	}

	if err != nil {
		response.InternalServerError(c, "Failed to delete role: "+err.Error())
		return
	}

	if resp != nil && resp.Error != nil {
		response.InternalServerError(c, "Failed to delete role: "+resp.Error.Message)
		return
	}

	response.Ok(c, "DeleteRole called", gin.H{"is_success": resp.Success})
}

func (rc *RoleController) DisableOrEnableRole(c *gin.Context) {
	roleId := c.Param("id")
	if roleId == "" {
		response.BadRequest(c, "Role ID is required")
		return
	}

	req := &auth.DisableOrEnableRoleRequest{RoleId: roleId}
	resp, err := rc.roleClient.DisableOrEnableRole(c, req)

	if !resp.Success {
		response.BadRequest(c, "Failed to disable or enable role: "+*resp.Message)
		return
	}

	if err != nil {
		response.InternalServerError(c, "Failed to disable or enable role: "+err.Error())
		return
	}

	if resp != nil && resp.Error != nil {
		response.InternalServerError(c, "Failed to disable or enable role: "+resp.Error.Message)
		return
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

func (rc *RoleController) UpsertRole(c *gin.Context) {
	req := rc.buildUpsertRoleRequest(c)

	resp, err := rc.roleClient.UpsertRole(c, req)
	if err != nil {
		response.InternalServerError(c, "Failed to upsert role: "+err.Error())
		return
	}

	if resp == nil || resp.Error != nil {
		response.InternalServerError(c, "Failed to upsert role: "+resp.Error.Message)
		return
	}

	if !resp.IsSuccess {
		response.BadRequest(c, "Failed to upsert role: "+resp.Message)
		return
	}

	response.Ok(c, "UpsertRole called", gin.H{
		"is_success": resp.IsSuccess,
	})

}

func (rc *RoleController) buildUpsertRoleRequest(c *gin.Context) *auth.UpsertRoleRequest {
	var req auth.UpsertRoleRequest
	var dto dtos.UpsertRoleRequestDTO
	roleId := c.Param("id")
	if roleId == "" {
		req.RoleId = nil
	}

	c.ShouldBindJSON(&dto)

	if dto.Name == "" {
		response.BadRequest(c, "Role name is required")
		return nil
	}

	if dto.Description == "" {
		response.BadRequest(c, "Role description is required")
		return nil
	}

	req.RoleId = dto.RoleId
	req.Name = dto.Name
	req.Description = dto.Description
	req.PermissionIds = dto.PermissionIds

	return &req
}
