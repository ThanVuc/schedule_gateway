package auth_controller

import (
	"fmt"
	"schedule_gateway/global"
	client "schedule_gateway/internal/client/auth"
	"schedule_gateway/internal/dtos"
	"schedule_gateway/internal/utils"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

type UserController struct {
	logger     log.Logger
	userClient client.UserClient
}

func NewUserController() *UserController {
	return &UserController{
		logger:     global.Logger,
		userClient: client.NewUserClient(),
	}
}

func (uc *UserController) AssignRoleToUser(c *gin.Context) {
	req, err := uc.buildAssignRoleToUserRequest(c)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	resp, err := uc.userClient.AssignRoleToUser(c, req)
	if err != nil {
		uc.logger.Error("Failed to assign role to user", "", zap.Error(err))
		response.InternalServerError(c, "Failed to assign role to user")
		return
	}

	if resp == nil || !*resp.Success {
		response.InternalServerError(c, "Failed to assign role to user: "+*resp.Message)
		return
	}

	response.Ok(c, "AssignRoleToUser called", resp)
}

func (uc *UserController) buildAssignRoleToUserRequest(c *gin.Context) (*auth.AssignRoleToUserRequest, error) {
	var body gin.H
	if err := c.ShouldBindJSON(&body); err != nil {
		response.BadRequest(c, "Invalid request body")
		return nil, err
	}

	userID, ok := body["user_id"].(string)
	if !ok || userID == "" {
		response.BadRequest(c, "User ID is required")
		return nil, fmt.Errorf("user_id is required")
	}

	roleIDSInterface, ok := body["role_ids"].([]interface{})
	if !ok {
		response.BadRequest(c, "Role ID is required")
		return nil, fmt.Errorf("role_id is required")
	}

	if len(roleIDSInterface) == 0 {
		return &auth.AssignRoleToUserRequest{
			UserId:  userID,
			RoleIds: []string{},
		}, nil
	}

	roleIDs := make([]string, 0)
	for i, v := range roleIDSInterface {
		roleID, ok := v.(string)
		if !ok {
			response.BadRequest(c, fmt.Sprintf("Role ID at index %d is invalid", i))
			return nil, fmt.Errorf("role_id at index %d is invalid", i)
		}
		roleIDs = append(roleIDs, roleID)
	}

	req := &auth.AssignRoleToUserRequest{
		UserId:  userID,
		RoleIds: roleIDs,
	}

	return req, nil
}

func (uc *UserController) GetUsers(c *gin.Context) {
	req := uc.buildGetUserRequest(c)

	if req == nil {
		return
	}

	users, err := uc.userClient.GetUsers(c, req)

	if err != nil {
		response.InternalServerError(c, "Failed to get users")
		return
	}

	if users != nil && users.Error != nil {
		response.InternalServerError(c, "Failed to get users: "+users.Error.Message)
		return
	}

	response.Ok(c, "Get User Successful", dtos.Users{
		Items:      users.Users,
		TotalUsers: users.TotalUsers,
		TotalPages: users.PageInfo.TotalPages,
		PageSize:   users.PageInfo.PageSize,
		Page:       users.PageInfo.Page,
		HasPrev:    users.PageInfo.HasPrev,
		HasNext:    users.PageInfo.HasNext,
	})
}

func (pc *UserController) buildGetUserRequest(c *gin.Context) *auth.GetUsersRequest {
	pageQuery := utils.ToPageQuery(c)
	searchString := c.Query("search")

	req := auth.GetUsersRequest{}
	if searchString != "" {
		req.Search = searchString
	}

	req.PageQuery = pageQuery

	return &req
}

func (pc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "User ID is required")
		return
	}

	req := &auth.GetUserRequest{
		UserId: id,
	}

	resp, err := pc.userClient.GetUser(c, req)

	if err != nil {
		response.InternalServerError(c, "Failed to get user: "+err.Error())
		return
	}

	if resp == nil || resp.Error != nil {
		response.InternalServerError(c, "Failed to get user: "+resp.Error.Message)
		return
	}

	response.Ok(c, "GetUser successful", resp.User)
}

func (pc *UserController) LockOrUnLockUser(c *gin.Context) {
	req, err := pc.buildLockUserRequest(c)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	resp, err := pc.userClient.LockOrUnLockUser(c, req)
	if err != nil {
		pc.logger.Error("Failed to lock user", "", zap.Error(err))
		response.InternalServerError(c, "Failed to lock user")
		return
	}

	if resp == nil || resp.Error != nil {
		msg := "Failed to lock user"
		if resp != nil && resp.Error != nil {
			msg = resp.Error.Message
		}
		response.InternalServerError(c, msg)
		return
	}

	response.Ok(c, "Lock/UnLock User successful", resp)
}

func (pc *UserController) buildLockUserRequest(c *gin.Context) (*auth.LockUserRequest, error) {
	var body struct {
		UserID     string  `json:"user_id"`
		LockReason *string `json:"lock_reason"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		return nil, fmt.Errorf("invalid request body")
	}

	if body.UserID == "" {
		return nil, fmt.Errorf("user_id is required")
	}

	req := &auth.LockUserRequest{
		UserId:     body.UserID,
		LockReason: body.LockReason,
	}

	return req, nil
}
