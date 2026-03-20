package team_controller

import (
	"fmt"
	"schedule_gateway/global"
	team_client "schedule_gateway/internal/client/team"
	dtos "schedule_gateway/internal/dtos/team_service"
	"schedule_gateway/internal/utils"
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
		fmt.Print("Error code: ", resp.GetError().GetCode())
		gc.logger.Error("Failed to create group: ", "", zap.String("code", resp.Error.Code), zap.String("message", *resp.Error.Details))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
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
			Description: utils.SafeString(group.Description),
			CreatedAt:   group.CreatedAt.AsTime().Format("2006-01-02T15:04:05Z"),
			UpdatedAt:   group.UpdatedAt.AsTime().Format("2006-01-02T15:04:05Z"),
		}
		if group.Owner != nil {
			groupDto.Owner = &dtos.SimpleUserDTO{
				ID:     group.Owner.Id,
				Email:  group.Owner.Email,
				Avatar: utils.SafeString(group.Owner.Avatar),
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

	if resp.GetError() != nil {
		gc.logger.Error("Failed to get group: ", "", zap.String("code", resp.Error.Code), zap.String("message", *resp.Error.Details))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	response.Ok(ctx, "Group retrieved successfully", gin.H{
		"group": groupDto,
	})
}

func (gc *GroupController) UpdateGroup(ctx *gin.Context) {
	id := ctx.Param("group_id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "Group ID is required"})
		return
	}

	fmt.Printf("Group ID: %s\n", id)

	var dto dtos.CreateGroupDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	req := &team_service.UpdateGroupRequest{
		Id:          id,
		Name:        &dto.Name,
		Description: dto.Description,
	}

	resp, err := gc.client.UpdateGroup(ctx, req)
	if err != nil {
		gc.logger.Error("Failed to update group: ", "", zap.Error(err))
		ctx.JSON(500, gin.H{"error": "Failed to update group"})
		return
	}

	if resp.GetError() != nil {
		gc.logger.Error("Failed to update group: ", "", zap.String("code", resp.Error.Code), zap.String("message", *resp.Error.Details))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	dtoResp := gc.builUpdateGroupResponse(resp)
	response.Ok(ctx, "Group updated successfully", dtoResp)

}

func (gc *GroupController) builUpdateGroupResponse(resp *team_service.UpdateGroupResponse) gin.H {
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

func (gc *GroupController) DeleteGroup(ctx *gin.Context) {
	id := ctx.Param("group_id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "Group ID is required"})
		return
	}

	req := &common.IDRequest{Id: id}
	resp, err := gc.client.DeleteGroup(ctx, req)
	if err != nil {
		gc.logger.Error("Failed to delete group: ", "", zap.Error(err))
		ctx.JSON(500, gin.H{"error": "Failed to delete group"})
		return
	}

	if resp.GetError() != nil {
		gc.logger.Error("Failed to delete group: ", "", zap.String("code", resp.Error.Code), zap.String("message", *resp.Error.Details))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	response.Ok(ctx, "Group deleted successfully", nil)
}

func (gc *GroupController) ListMembers(ctx *gin.Context) {
	id := ctx.Param("group_id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "Group ID is required"})
		return
	}

	req := &team_service.ListMembersRequest{GroupId: id}
	resp, err := gc.client.ListMembers(ctx, req)
	if err != nil {
		gc.logger.Error("Failed to list group members: ", "", zap.Error(err))
		ctx.JSON(500, gin.H{"error": "Failed to list group members"})
		return
	}

	var members []dtos.MemberDTO
	for _, member := range resp.Members {
		members = append(members, dtos.MemberDTO{
			ID:       member.Id,
			Email:    member.Email,
			Avatar:   member.Avatar,
			Role:     int32(*member.Role.Enum()),
			JoinedAt: member.JoinedAt.AsTime().Format("2006-01-02T15:04:05Z"),
		})
	}

	if resp.GetError() != nil {
		gc.logger.Error("Failed to list group members: ", "", zap.String("code", resp.Error.Code), zap.String("message", *resp.Error.Details))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	response.Ok(ctx, "Group members retrieved successfully", gin.H{
		"members": members,
		"total":   len(members),
	})

}

func (gc *GroupController) UpdateMemberRole(ctx *gin.Context) {
	groupId := ctx.Param("group_id")
	if groupId == "" {
		ctx.JSON(400, gin.H{"error": "Group ID is required"})
		return
	}

	userId := ctx.Param("user_id")
	if userId == "" {
		ctx.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	var dto dtos.UpdateMemberRoleDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	req := &team_service.UpdateMemberRoleRequest{
		GroupId:  groupId,
		MemberId: userId,
		NewRole:  team_service.GroupRole(dto.Role),
	}

	resp, err := gc.client.UpdateMemberRole(ctx, req)
	if err != nil {
		gc.logger.Error("Failed to update member role: ", "", zap.Error(err))
		ctx.JSON(500, gin.H{"error": "Failed to update member role"})
		return
	}

	if resp.GetError() != nil {
		gc.logger.Error("Failed to update member role: ", "", zap.String("code", resp.Error.Code), zap.String("message", *resp.Error.Details))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	var memberDto *dtos.MemberDTO
	if resp.Member != nil {
		member := resp.Member
		memberDto = &dtos.MemberDTO{
			ID:       member.Id,
			Email:    member.Email,
			Avatar:   member.Avatar,
			Role:     int32(*member.Role.Enum()),
			JoinedAt: member.JoinedAt.AsTime().Format("2006-01-02T15:04:05Z"),
		}
	}

	response.Ok(ctx, "Member role updated successfully", gin.H{
		"member": memberDto,
	})

}

func (gc *GroupController) RemoveMember(ctx *gin.Context) {
	groupId := ctx.Param("group_id")
	if groupId == "" {
		ctx.JSON(400, gin.H{"error": "Group ID is required"})
		return
	}

	userId := ctx.Param("user_id")
	if userId == "" {
		ctx.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	req := &team_service.RemoveMemberRequest{
		GroupId:  groupId,
		MemberId: userId,
	}

	resp, err := gc.client.RemoveMember(ctx, req)
	if err != nil {
		gc.logger.Error("Failed to remove member: ", "", zap.Error(err))
		ctx.JSON(500, gin.H{"error": "Failed to remove member"})
		return
	}

	if resp.GetError() != nil {
		gc.logger.Error("Failed to remove member: ", "", zap.String("code", resp.Error.Code), zap.String("message", *resp.Error.Details))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	response.Ok(ctx, "Member removed successfully", gin.H{
		"success": resp.Success,
	})
}

func (gc *GroupController) CreateInvite(ctx *gin.Context) {
	groupId := ctx.Param("group_id")
	if groupId == "" {
		ctx.JSON(400, gin.H{"error": "Group ID is required"})
		return
	}

	var dto dtos.CreateInviteDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var email *string
	if dto.Email != "" {
		email = &dto.Email
	}

	req := &team_service.CreateInviteRequest{
		GroupId: groupId,
		Email:   email,
		Role:    team_service.GroupRole(dto.Role),
	}
	fmt.Printf("REQ: %+v\n", req)
	resp, err := gc.client.CreateInvite(ctx, req)
	if err != nil {
		gc.logger.Error("Failed to create invite: ", "", zap.Error(err))
		ctx.JSON(500, gin.H{"error": "Failed to create invite"})
		return
	}

	if resp.GetError() != nil {
		gc.logger.Error("Failed to create invite: ", "", zap.String("code", resp.Error.Code), zap.String("message", *resp.Error.Details))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	var inviteDto *dtos.InviteDTO
	if resp.Invite != nil {
		invite := resp.Invite
		inviteDto = &dtos.InviteDTO{
			Code:      invite.Code,
			ExpiresAt: invite.ExpiresAt.AsTime().Format("2006-01-02T15:04:05Z"),
			CreateAt:  invite.CreatedAt.AsTime().Format("2006-01-02T15:04:05Z"),
		}
	}

	response.Ok(ctx, "Invite created successfully", gin.H{
		"invite": inviteDto,
	})
}
