package controller

import (
	"schedule_gateway/global"
	"schedule_gateway/internal/client"
	"schedule_gateway/internal/dtos"
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

	response.Ok(c, "GetPermissions called", gin.H{
		"permissions": permissions.Permissions,
	})
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

func (pc *PermissionController) DeletePermission(c *gin.Context) {
	response.Ok(c, "DeletePermission called", nil)
}

func (pc *PermissionController) AssignPermissionToRole(c *gin.Context) {
	response.Ok(c, "AssignPermissionToRole called", nil)
}

func (pc *PermissionController) GetResources(c *gin.Context) {
	req := &auth.GetResourcesRequest{}
	resources, err := pc.permissionClient.GetResources(c, req)

	if err != nil {
		panic(response.InternalServerError("Failed to get resources"))
	}

	response.Ok(c, "GetResources called", gin.H{
		"resources": resources.Resources,
	})
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

func (pc *PermissionController) buildGetPermissionRequest(c *gin.Context) *auth.GetPermissionsRequest {
	pageQuery := utils.ToPageQuery(c)
	searchString := c.Query("search")
	resourceIdString := c.Query("resource_id")

	req := auth.GetPermissionsRequest{
		PageQuery:  pageQuery,
		Search:     searchString,
		ResourceId: resourceIdString,
	}

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
