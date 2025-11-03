package helper

import (
	constant "schedule_gateway/internal/routers/constant"
	"schedule_gateway/proto/auth"
)

type ResourcePredefine struct {
	AuthResource       *auth.Resource
	RoleResource       *auth.Resource
	PermissionResource *auth.Resource
	UserResource       *auth.Resource
	AdminUserResource  *auth.Resource
	LabelResource      *auth.Resource
	GoalResource       *auth.Resource
}

func InitResources() *ResourcePredefine {
	r := &ResourcePredefine{
		AuthResource: &auth.Resource{
			Id:   "1",
			Name: constant.AUTH_RESOURCE,
		},
		RoleResource: &auth.Resource{
			Id:   "2",
			Name: constant.ROLE_RESOURCE,
		},
		PermissionResource: &auth.Resource{
			Id:   "3",
			Name: constant.PERMISSION_RESOURCE,
		},
		UserResource: &auth.Resource{
			Id:   "4",
			Name: constant.USER_RESOURCE,
		},
		AdminUserResource: &auth.Resource{
			Id:   "5",
			Name: constant.ADMIN_USER_RESOURCE,
		},
		LabelResource: &auth.Resource{
			Id:   "6",
			Name: constant.LABEL_RESOURCE,
		},
		GoalResource: &auth.Resource{
			Id:   "7",
			Name: constant.GOAL_RESOURCE,
		},
	}
	return r
}
