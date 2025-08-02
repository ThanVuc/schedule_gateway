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

type PermissionController struct {
	logger           log.Logger
	permissionClient client.PermissionClient
}

func NewPermissionController() *PermissionController {
	return &PermissionController{
		logger:           global.Logger,
		permissionClient: client.NewPermissionClient(),
	}
}

func (pc *PermissionController) GetPermissions(c *gin.Context) {
	req := pc.buildGetPermissionRequest(c)
	if req == nil {
		return
	}

	permissions, err := pc.permissionClient.GetPermissions(c, req)

	if err != nil {
		response.InternalServerError(c, "Failed to get permissions")
		return
	}

	if permissions != nil && permissions.Error != nil {
		response.InternalServerError(c, "Failed to get permissions: "+permissions.Error.Message)
		return
	}

	response.Ok(c, "GetPermissions called", dtos.Permissions{
		Items:             permissions.Permissions,
		TotalPermissions:  permissions.TotalPermissions,
		Root:              permissions.Root,
		NonRoot:           permissions.NonRoot,
		RootPercentage:    float32(permissions.RootPercentage),
		NonRootPercentage: float32(permissions.NonRootPercentage),
		TotalItems:        permissions.TotalPermissions,
		TotalPages:        permissions.PageInfo.TotalPages,
		PageSize:          permissions.PageInfo.PageSize,
		Page:              permissions.PageInfo.Page,
		HasPrev:           permissions.PageInfo.HasPrev,
		HasNext:           permissions.PageInfo.HasNext,
	})
}

func (pc *PermissionController) UpsertPermission(c *gin.Context) {
	req := pc.buildUpsertPermissionRequest(c)
	if req == nil {
		return
	}

	resp, err := pc.permissionClient.UpsertPermission(c, req)

	if err != nil {
		response.InternalServerError(c, "Failed to upsert permission")
		return
	}

	if resp != nil && resp.Error != nil {
		response.InternalServerError(c, "Failed to upsert permission")
		return
	}

	response.Ok(c, "Upsert Permission Successfult", gin.H{
		"is_success": resp.IsSuccess,
		"perm_id":    resp.PermissionId,
	})
}

func (pc *PermissionController) GetResources(c *gin.Context) {
	req := &auth.GetResourcesRequest{}
	resources, err := pc.permissionClient.GetResources(c, req)

	if err != nil {
		response.InternalServerError(c, "Failed to get resources")
		return
	}

	response.Ok(c, "GetResources called", resources.Resources)
}

func (pc *PermissionController) GetActions(c *gin.Context) {
	resourceId := c.Query("resource_id")
	if resourceId == "" {
		response.BadRequest(c, "Resource ID is required")
		return
	}

	req := &auth.GetActionsRequest{
		ResourceId: resourceId,
	}
	actions, err := pc.permissionClient.GetActions(c, req)

	if err != nil {
		response.InternalServerError(c, "Internal server error")
		return
	}

	response.Ok(c, "GetActions called", gin.H{
		"actions": actions.Actions,
	})
}

func (pc *PermissionController) GetPermission(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "Permission ID is required")
		return
	}

	req := &auth.GetPermissionRequest{
		PermissionId: id,
	}

	resp, err := pc.permissionClient.GetPermission(c, req)

	if err != nil {
		response.InternalServerError(c, "Failed to get permission: "+err.Error())
		return
	}

	if resp == nil || resp.Error != nil {
		response.InternalServerError(c, "Failed to get permission: "+resp.Error.Message)
		return
	}

	response.Ok(c, "GetPermission successful", resp.Permission)
}

func (pc *PermissionController) DeletePermission(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "Permission ID is required")
		return
	}

	req := &auth.DeletePermissionRequest{
		PermissionId: id,
	}

	resp, err := pc.permissionClient.DeletePermission(c, req)
	if !resp.Success {
		response.BadRequest(c, "Failed to delete permission: "+*resp.Message)
		return
	}

	if err != nil {
		response.InternalServerError(c, "Failed to delete permission: "+err.Error())
		return
	}

	if resp == nil || resp.Error != nil {
		response.InternalServerError(c, "Failed to delete permission: "+resp.Error.Message)
		return
	}

	response.Ok(c, "DeletePermission called", gin.H{
		"is_success": resp.Success,
	})
}

func (pc *PermissionController) buildGetPermissionRequest(c *gin.Context) *auth.GetPermissionsRequest {
	pageQuery := utils.ToPageQuery(c)
	searchString := c.Query("search")
	resourceIdString := c.Query("resource_id")

	req := auth.GetPermissionsRequest{}
	if searchString != "" {
		req.Search = searchString
	}

	if resourceIdString != "" {
		req.ResourceId = resourceIdString
	}

	req.PageQuery = pageQuery

	return &req
}

func (pc *PermissionController) buildUpsertPermissionRequest(c *gin.Context) *auth.UpsertPermissionRequest {
	var req auth.UpsertPermissionRequest
	var dto dtos.UpsertPermissionRequestDTO

	id := c.Param("id")
	if id != "" {
		req.PermissionId = &id
	}

	if err := c.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(c, "Invalid request body: "+err.Error())
		return nil
	}

	if dto.Name == "" {
		response.BadRequest(c, "Name is required")
		return nil
	}

	if dto.ResourceId == "" {
		response.BadRequest(c, "ResourceId is required")
		return nil
	}

	if len(dto.ActionIds) == 0 {
		response.BadRequest(c, "At least one action ID is required")
		return nil
	}

	req.Name = dto.Name
	req.Description = dto.Description
	req.ResourceId = dto.ResourceId
	req.ActionIds = dto.ActionIds

	return &req
}
