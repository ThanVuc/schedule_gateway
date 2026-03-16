package team_controller

import (
	"fmt"
	"schedule_gateway/global"
	team_client "schedule_gateway/internal/client/team"
	dtos "schedule_gateway/internal/dtos/team_service"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/common"
	"schedule_gateway/proto/team_service"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

type GroupController struct {
	logger log.Logger
	client team_client.GroupClient
}

func NewGroupController() *GroupController {
	return &GroupController{
		logger: global.Logger,
		client: team_client.NewGroupClient(),
	}
}

func (gc *GroupController) Ping(ctx *gin.Context) {
	resp, err := gc.client.Ping(ctx, &common.EmptyRequest{})
	if err != nil {
		gc.logger.Error("Failed to ping GroupService: ", "")
	}
	ctx.JSON(200, resp)
}

func (gc *GroupController) CreateGroup(ctx *gin.Context) {
	req := gc.buildCreateGroupRequest(ctx)
	if req == nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := gc.client.CreateGroup(ctx, req)
	if err != nil {
		gc.logger.Error("Failed to create group: ", "", zap.Error(err))
		ctx.JSON(500, gin.H{"error": "Failed to create group"})
		return
	}

	if resp.GetError() != nil {
		gc.logger.Error("Failed to create group: ", "", zap.String("code", resp.Error.Code), zap.String("message", *resp.Error.Details))
		response.UnprocessableEntity(ctx, resp.Error.Code, resp.Error.Message, *resp.Error.Details)
		return
	}

	dto := gc.buildGetGroupResponse(resp)
	response.Ok(ctx, "Group created successfully", dto)
}

func (gc *GroupController) buildCreateGroupRequest(c *gin.Context) *team_service.CreateGroupRequest {
	var req team_service.CreateGroupRequest
	var dto dtos.CreateGroupDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		return nil
	}

	req.Name = dto.Name
	req.Description = dto.Description

	return &req

}

func (gc *GroupController) buildGetGroupResponse(resp *team_service.CreateGroupResponse) gin.H {
	var groupDto *dtos.GroupDTO
	if resp.Group != nil {
		group := resp.Group
		groupDto = &dtos.GroupDTO{
			ID:          group.Id,
			Name:        group.Name,
			Description: *group.Description,
			CreatedAt:   group.CreatedAt.AsTime().Format("2006-01-02T15:04:05Z"),
			UpdatedAt:   group.UpdatedAt.AsTime().Format("2006-01-02T15:04:05Z"),
		}
		if group.Owner != nil {
			groupDto.Owner = &dtos.SimpleUserDTO{
				ID:     group.Owner.Id,
				Email:  group.Owner.Email,
				Avatar: *group.Owner.Avatar,
			}
		}
	}

	return gin.H{
		"group": groupDto,
	}
}

func (gc *GroupController) GetGroup(ctx *gin.Context) {
	id := ctx.Param("group_id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "Group ID is required"})
		return
	}
	fmt.Printf("Group ID: %s\n", id)

	req := &common.IDRequest{Id: id}
	resp, err := gc.client.GetGroup(ctx, req)
	if err != nil {
		gc.logger.Error("Failed to get group: ", "", zap.Error(err))
		ctx.JSON(500, gin.H{"error": "Failed to get group"})
		return
	}

	var groupDto *dtos.GroupDetailDTO
	if resp.Group != nil {
		group := resp.Group
		groupDto = &dtos.GroupDetailDTO{
			ID:           group.Id,
			Name:         group.Name,
			Description:  *group.Description,
			MyRole:       int32(*group.MyRole.Enum()),
			ActiveSprint: *group.ActiveSprint,
			Avatar:       group.Avatar,
			MembersTotal: group.MemberCount,
			CreatedAt:    group.CreatedAt.AsTime().Format("2006-01-02T15:04:05Z"),
			UpdatedAt:    group.UpdatedAt.AsTime().Format("2006-01-02T15:04:05Z"),
		}
		if group.Owner != nil {
			groupDto.Owner = &dtos.SimpleUserDTO{
				ID:     group.Owner.Id,
				Email:  group.Owner.Email,
				Avatar: *group.Owner.Avatar,
			}
		}
	}

	response.Ok(ctx, "Group retrieved successfully", gin.H{
		"group": groupDto,
	})
}
