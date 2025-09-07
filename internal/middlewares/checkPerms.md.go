package middlewares

import (
	auth_client "schedule_gateway/internal/client/auth"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/auth"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	clientInstance auth_client.AuthClient
	once           sync.Once
)

func getAuthClient() auth_client.AuthClient {
	once.Do(func() {
		clientInstance = auth_client.NewAuthClient()
	})
	return clientInstance
}

func CheckPerm(resource string, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if err != nil || accessToken == "" {
			response.Unauthorized(c, "Unauthorized: missing access token")
			return
		}

		authClient := getAuthClient()
		req := &auth.CheckPermissionRequest{
			AccessToken:  accessToken,
			ResourceName: resource,
			ActionName:   action,
		}

		resp, err := authClient.CheckPermission(c, req)
		if err != nil {
			response.InternalServerError(c, "Failed to check permission: "+err.Error())
			return
		}

		switch resp.Status {
		case auth.PERMISSION_STATUS_UNAUTHORIZED:
			response.Unauthorized(c, "Unauthorized: invalid access token")
			return
		case auth.PERMISSION_STATUS_FORBIDDEN:
			response.Forbidden(c, "Forbidden: insufficient permissions")
			return
		case auth.PERMISSION_STATUS_PERMISSION_UNSPECIFIED:
			response.Forbidden(c, "Forbidden: unspecified permission")
			return
		case auth.PERMISSION_STATUS_ALLOWED:
			c.Set("user_id", resp.UserId)
			c.Next()
		default:
			response.InternalServerError(c, "Internal server error: unknown permission status")
			return
		}
	}
}
