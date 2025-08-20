package auth_controller

import (
	"fmt"
	"schedule_gateway/global"
	client "schedule_gateway/internal/client/auth"
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
	if !ok || len(roleIDSInterface) == 0 {
		response.BadRequest(c, "Role ID is required")
		return nil, fmt.Errorf("role_id is required")
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
