package dtos

import "schedule_gateway/proto/auth"

type UserAuthInfo struct {
	UserId      string `json:"user_id"`
	Email       string `json:"email"`
	Permissions []*auth.PermissionAuthItem
}
