package controller

import (
	"schedule_gateway/global"
	"schedule_gateway/internal/client"
	"schedule_gateway/internal/dtos"
	"schedule_gateway/internal/mapper"
	"schedule_gateway/internal/utils"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/auth"

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
	req := pc.buildGetPermissionRequest(c)
	permissions, err := pc.permissionClient.GetPermissions(c, req)

	if err != nil {
		panic(response.InternalServerError("Failed to get permissions"))
	}

	if permissions != nil && permissions.Error != nil {
		panic(response.InternalServerError("Failed to get permissions: " + permissions.Error.Message))
	}

	response.Ok(c, "GetPermissions called", mapper.MapPermissionsToDTO((permissions)))
}

func (pc *PermissionController) UpsertPermission(c *gin.Context) {
	req := pc.buildUpsertPermissionRequest(c)

	resp, err := pc.permissionClient.UpsertPermission(c, req)

	if err != nil {
		panic(response.InternalServerError("Failed to create permission: " + err.Error()))
	}

	if resp != nil && resp.Error != nil {
		panic(response.InternalServerError("Failed to create permission: " + resp.Error.Message))
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
		panic(response.InternalServerError("Failed to get resources"))
	}

	response.Ok(c, "GetResources called", resources.Resources)
}

func (pc *PermissionController) GetActions(c *gin.Context) {
	resourceId := c.Query("resource_id")
	if resourceId == "" {
		panic(response.BadRequest("resource_id is required"))
	}

	req := &auth.GetActionsRequest{
		ResourceId: resourceId,
	}
	actions, err := pc.permissionClient.GetActions(c, req)

	if err != nil {
		panic(response.InternalServerError("Failed to get actions"))
	}

	response.Ok(c, "GetActions called", gin.H{
		"actions": actions.Actions,
	})
}

func (pc *PermissionController) GetPermission(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		panic(response.BadRequest("Permission ID is required"))
	}

	req := &auth.GetPermissionRequest{
		PermissionId: id,
	}

	resp, err := pc.permissionClient.GetPermission(c, req)

	if err != nil {
		panic(response.InternalServerError("Failed to get permission: " + err.Error()))
	}

	if resp == nil || resp.Error != nil {
		panic(response.InternalServerError("Failed to get permission: " + resp.Error.Message))
	}

	response.Ok(c, "GetPermission called", mapper.MapPermissionToDTO(resp))
}

func (pc *PermissionController) DeletePermission(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		panic(response.BadRequest("Permission ID is required"))
	}

	req := &auth.DeletePermissionRequest{
		PermissionId: id,
	}

	resp, err := pc.permissionClient.DeletePermission(c, req)
	if !resp.Success {
		panic(response.BadRequest("Failed to delete permission: " + *resp.Message))
	}

	if err != nil {
		panic(response.InternalServerError("Failed to delete permission: " + err.Error()))
	}

	if resp == nil || resp.Error != nil {
		panic(response.InternalServerError("Failed to delete permission: " + resp.Error.Message))
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
		panic(response.BadRequest("Invalid request body: " + err.Error()))
	}

	if dto.Name == "" {
		panic(response.BadRequest("Name is required"))
	}

	if dto.ResourceId == "" {
		panic(response.BadRequest("ResourceId is required"))
	}

	if len(dto.ActionsIds) == 0 {
		panic(response.BadRequest("At least one action ID is required"))
	}

	req.Name = dto.Name
	req.Description = dto.Description
	req.ResourceId = dto.ResourceId
	req.ActionsIds = dto.ActionsIds

	return &req
}
