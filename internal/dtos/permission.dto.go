package dtos

import "schedule_gateway/proto/auth"

type Permissions struct {
	Items             []*auth.PermissionItem `json:"items"`
	TotalPermissions  int32                  `json:"total_permissions"`
	Root              int32                  `json:"root"`
	NonRoot           int32                  `json:"non_root"`
	RootPercentage    float32                `json:"root_percentage"`
	NonRootPercentage float32                `json:"non_root_percentage"`
	TotalItems        int32                  `json:"total_items"`
	TotalPages        int32                  `json:"total_pages"`
	PageSize          int32                  `json:"page_size"`
	Page              int32                  `json:"page"`
	HasPrev           bool                   `json:"has_prev"`
	HasNext           bool                   `json:"has_next"`
}
