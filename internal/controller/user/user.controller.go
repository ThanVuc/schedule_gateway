package user_controller

import (
	"fmt"
	"schedule_gateway/global"
	client "schedule_gateway/internal/client/user"
	dtos "schedule_gateway/internal/dtos/user"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/user"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
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

func (uc *UserController) GetUserProfile(c *gin.Context) {
	id := c.GetString("user_id")
	fmt.Println("User ID from context:", id)
	if id == "" {
		response.BadRequest(c, "User ID is required")
		return
	}

	req := &user.GetUserProfileRequest{
		Id: id,
	}

	resp, err := uc.userClient.GetUserProfile(c, req)
	if err != nil {
		response.InternalServerError(c, "Service unavailable: "+err.Error())
		return
	}

	if resp == nil {
		response.InternalServerError(c, "Empty response from service")
		return
	}

	if resp.Error != nil {
		switch resp.Error.Code {
		case 1:
			response.NotFound(c, resp.Error.Message)
		default:
			response.BadRequest(c, resp.Error.Message)
		}
		return
	}

	response.Ok(c, "GetUserProfile called", resp.Profiles)
}

func (uc *UserController) UpdateUserInfo(c *gin.Context) {
	req := uc.buildUpsertUserProfileRequest(c)
	if req == nil {
		return
	}

	resp, err := uc.userClient.UpdateUserProfile(c, req)
	if err != nil {
		response.InternalServerError(c, "Failed to update user profile")
		return
	}

	if resp != nil && resp.Error != nil {
		response.InternalServerError(c, "Failed to update user profile: "+resp.Error.Message)
		return
	}
	if resp == nil {
		response.InternalServerError(c, "Failed to update user profile: response is nil")
		return
	}

	response.Ok(c, "UpdateUserProfile Successfult", resp)
}

func (uc *UserController) buildUpsertUserProfileRequest(c *gin.Context) *user.UpdateUserProfileRequest {
	var req user.UpdateUserProfileRequest
	var dto dtos.UpsertUserProfileRequestDTO

	id := c.GetString("user_id")
	if id == "" {
		response.BadRequest(c, "User ID is required")
		return nil
	} else {
		req.Id = id
	}

	if err := c.ShouldBind(&dto); err != nil {
		response.BadRequest(c, "Invalid request body: "+err.Error())
		return nil
	}

	if dto.Fullname == "" {
		response.BadRequest(c, "Fullname is required")
		return nil
	}

	req.Fullname = dto.Fullname
	req.Bio = dto.Bio
	req.DateOfBirth = dto.DateOfBirth
	req.Gender = dto.Gender
	req.Sentence = dto.Sentence
	req.Author = dto.Author

	return &req
}
