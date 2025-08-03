package user_client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/user"

	"github.com/gin-gonic/gin"
)

type userClient struct {
	userClient user.UserServiceClient
}

func (uc *userClient) GetUserProfile(c *gin.Context, req *user.GetUserProfileRequest) (*user.GetUserProfileResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := uc.userClient.GetUserProfile(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
