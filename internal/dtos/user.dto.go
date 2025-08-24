package dtos

import "schedule_gateway/proto/auth"

type Users struct {
	Items             []*auth.UserItem `json:"items"`
	TotalUsers        int32            `json:"total_users"`
	PageSize          int32            `json:"page_size"`
	Page              int32            `json:"page"`
	HasPrev           bool             `json:"has_prev"`
	HasNext           bool             `json:"has_next"`
}
