package team_controller

import (
	"net/http"
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
		gc.logger.Error("Failed to create group: ", "", zap.String("code", resp.Error.Code), zap.String("message", *resp.Error.Details))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	// reuse shared BuildGroupResponse from work.controller.go
	dto := BuildGroupResponse(resp.GetGroup())

	response.Ok(ctx, "Group created successfully", gin.H{"item": dto})
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

func (gc *GroupController) GetGroup(ctx *gin.Context) {
	id := ctx.Param("group_id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "Group ID is required"})
		return
	}

	req := &common.IDRequest{Id: id}
	resp, err := gc.client.GetGroup(ctx, req)
	if err != nil {
		gc.logger.Error("Failed to get group: ", "", zap.Error(err))
		ctx.JSON(500, gin.H{"error": "Failed to get group"})
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if resp.GetError() != nil {
		gc.logger.Error("Failed to get group: ", "", zap.String("code", resp.Error.Code), zap.String("message", utils.SafeString(resp.Error.Details)))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	// reuse shared BuildGroupResponse
	dto := BuildGroupResponse(resp.GetGroup())

	response.Ok(ctx, "Group retrieved successfully", gin.H{
		"item": dto,
	})
}

func (gc *GroupController) ListGroups(ctx *gin.Context) {
	resp, err := gc.client.ListGroups(ctx, &common.IDRequest{})
	if err != nil {
		gc.logger.Error("Failed to list groups: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to list groups")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if resp.GetError() != nil {
		gc.logger.Error("Failed to list groups: ", "", zap.String("code", resp.Error.Code), zap.String("message", utils.SafeString(resp.Error.Details)))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	items := gc.buildListGroupsResponse(resp.GetGroups())
	response.Ok(ctx, "List groups successful", gin.H{
		"items": items,
		"total": resp.GetTotal(),
	})
}

func (gc *GroupController) buildListGroupsResponse(groups []*team_service.GroupMessage) []dtos.ListGroupItemDTO {
	items := make([]dtos.ListGroupItemDTO, 0, len(groups))
	for _, group := range groups {
		if group == nil {
			continue
		}

		createdAt := ""
		if group.GetCreatedAt() != nil {
			createdAt = group.GetCreatedAt().AsTime().Format("2006-01-02T15:04:05Z")
		}

		updatedAt := ""
		if group.GetUpdatedAt() != nil {
			updatedAt = group.GetUpdatedAt().AsTime().Format("2006-01-02T15:04:05Z")
		}

		items = append(items, dtos.ListGroupItemDTO{
			ID:          group.GetId(),
			Name:        group.GetName(),
			MyRole:      int32(group.GetMyRole()),
			MemberTotal: group.GetMemberCount(),
			AvatarURL:   group.GetAvatar(),
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		})
	}

	return items
}

func (gc *GroupController) ListSimpleUsers(ctx *gin.Context) {
	groupID := ctx.Param("group_id")
	if groupID == "" {
		response.BadRequest(ctx, "Group ID is required")
		return
	}

	resp, err := gc.client.GetSimpleUserByGroupID(ctx, &common.IDRequest{Id: groupID})
	if err != nil {
		gc.logger.Error("Failed to list simple users: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to list simple users")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if resp.GetError() != nil {
		gc.logger.Error("Failed to list simple users: ", "", zap.String("code", resp.Error.Code), zap.String("message", utils.SafeString(resp.Error.Details)))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	items := gc.buildSimpleUsersResponse(resp.GetUsers())
	response.Ok(ctx, "List simple users successful", gin.H{
		"items": items,
		"total": len(items),
	})
}

func (gc *GroupController) buildSimpleUsersResponse(users []*team_service.SimpleUserMessage) []dtos.SimpleUserDTO {
	items := make([]dtos.SimpleUserDTO, 0, len(users))
	for _, user := range users {
		if user == nil {
			continue
		}

		items = append(items, dtos.SimpleUserDTO{
			ID:     user.GetId(),
			Email:  user.GetEmail(),
			Avatar: user.GetAvatar(),
		})
	}

	return items
}

func (gc *GroupController) UpdateGroup(ctx *gin.Context) {
	id := ctx.Param("group_id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "Group ID is required"})
		return
	}

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

	// reuse shared BuildGroupResponse
	dtoResp := BuildGroupResponse(resp.GetGroup())
	response.Ok(ctx, "Group updated successfully", gin.H{"item": dtoResp})

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

	response.Ok(ctx, "Group deleted successfully", gin.H{"item": nil})
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

	if resp.GetError() != nil {
		gc.logger.Error("Failed to list group members: ", "", zap.String("code", resp.Error.Code), zap.String("message", *resp.Error.Details))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	members := gc.buiListMembersResponse(resp.Members)

	response.Ok(ctx, "Group members retrieved successfully", gin.H{
		"items": members,
		"total": len(members),
	})
}

func (gc *GroupController) buiListMembersResponse(members []*team_service.MemberMessage) []dtos.MemberDTO {
	var memberDtos []dtos.MemberDTO
	for _, member := range members {
		memberDtos = append(memberDtos, dtos.MemberDTO{
			ID:       member.GetId(),
			Email:    member.GetEmail(),
			Avatar:   member.GetAvatar(),
			Role:     int32(member.GetRole()),
			JoinedAt: utils.TimestampToISO8601(member.GetJoinedAt()),
		})
	}
	return memberDtos
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
		NewRole:  team_service.GroupRole(dto.NewRole),
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

	memberDto := gc.buildMemberDTO(resp.Member)

	response.Ok(ctx, "Member role updated successfully", gin.H{
		"item": memberDto,
	})
}

func (gc *GroupController) buildMemberDTO(member *team_service.MemberMessage) dtos.MemberDTO {
	return dtos.MemberDTO{
		ID:       member.GetId(),
		Email:    member.GetEmail(),
		Avatar:   member.GetAvatar(),
		Role:     int32(member.GetRole()),
		JoinedAt: utils.TimestampToISO8601(member.GetJoinedAt()),
	}
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
		"item": gin.H{"is_success": resp.Success},
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

	inviteDto := gc.buildCreateInviteResponse(resp.GetInvite())

	response.Ok(ctx, "Invite created successfully", gin.H{
		"item": inviteDto,
	})
}

func (gc GroupController) buildCreateInviteResponse(invite *team_service.InviteMessage) *dtos.InviteDTO {
	if invite == nil {
		return nil
	}

	return &dtos.InviteDTO{
		Code:      invite.GetCode(),
		ExpiresAt: utils.TimestampToISO8601(invite.GetExpiresAt()),
		CreateAt:  utils.TimestampToISO8601(invite.GetCreatedAt()),
	}
}

func (gc *GroupController) AcceptInvite(ctx *gin.Context) {
	var dto dtos.CodeDataDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	req := &team_service.AcceptInviteRequest{
		Code: dto.Data.Code,
	}

	_, exists := ctx.Get("user_id")
	if !exists {
		ctx.Redirect(http.StatusFound, "https://www.schedulr.site/login")
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	resp, err := gc.client.AcceptInvite(ctx, req)
	if err != nil {
		ctx.Redirect(http.StatusFound, "https://www.schedulr.site/login1231111111")
		gc.logger.Error("Failed to accept invite: ", "", zap.Error(err))
		ctx.JSON(500, gin.H{"error": "Failed to accept invite"})
		return
	}

	if resp.GetError() != nil {
		ctx.Redirect(http.StatusFound, "https://www.schedulr.site/login123333")
		gc.logger.Error("Failed to accep2222222222t invite: ", "", zap.String("code", resp.Error.Code), zap.String("message", *resp.Error.Details))
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	ctx.Redirect(http.StatusFound, resp.GetLocation())

}

func BuildGroupResponse(group *team_service.GroupMessage) gin.H {
	if group == nil {
		return gin.H{}
	}

	var owner gin.H
	if group.Owner != nil {
		owner = gin.H{
			"id":     group.Owner.GetId(),
			"email":  group.Owner.GetEmail(),
			"avatar": utils.SafeString(group.Owner.Avatar),
		}
	} else {
		owner = nil
	}

	return gin.H{
		"id":            group.GetId(),
		"name":          group.GetName(),
		"description":   utils.SafeString(group.Description),
		"my_role":       int32(group.GetMyRole()),
		"active_sprint": group.GetActiveSprint(),
		"avatar":        group.GetAvatar(),
		"members_total": group.GetMemberCount(),
		"created_at":    utils.TimestampToISO8601(group.GetCreatedAt()),
		"updated_at":    utils.TimestampToISO8601(group.GetUpdatedAt()),
		"owner":         owner,
	}
}
