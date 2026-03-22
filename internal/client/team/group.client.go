package team_client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/team_service"

	"schedule_gateway/proto/common"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type groupClient struct {
	logger      log.Logger
	groupClient team_service.GroupServiceClient
}

func (wc *groupClient) Ping(c *gin.Context, req *common.EmptyRequest) (*common.EmptyResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.groupClient.Ping(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *groupClient) CreateGroup(c *gin.Context, req *team_service.CreateGroupRequest) (*team_service.CreateGroupResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.groupClient.CreateGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *groupClient) GetGroup(c *gin.Context, req *common.IDRequest) (*team_service.GetGroupResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.groupClient.GetGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *groupClient) GetSimpleUserByGroupID(c *gin.Context, req *common.IDRequest) (*team_service.GetSimpleUserByGroupIDResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.groupClient.GetSimpleUserByGroupID(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *groupClient) UpdateGroup(c *gin.Context, req *team_service.UpdateGroupRequest) (*team_service.UpdateGroupResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.groupClient.UpdateGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *groupClient) DeleteGroup(c *gin.Context, req *common.IDRequest) (*team_service.DeleteGroupResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.groupClient.DeleteGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *groupClient) ListMembers(c *gin.Context, req *team_service.ListMembersRequest) (*team_service.ListMembersResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.groupClient.ListMembers(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *groupClient) UpdateMemberRole(c *gin.Context, req *team_service.UpdateMemberRoleRequest) (*team_service.UpdateMemberRoleResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.groupClient.UpdateMemberRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *groupClient) RemoveMember(c *gin.Context, req *team_service.RemoveMemberRequest) (*team_service.RemoveMemberResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.groupClient.RemoveMember(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *groupClient) CreateInvite(c *gin.Context, req *team_service.CreateInviteRequest) (*team_service.CreateInviteResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.groupClient.CreateInvite(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
