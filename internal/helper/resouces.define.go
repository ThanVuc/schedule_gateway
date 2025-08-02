package helper

import (
	constant "schedule_gateway/internal/routers/constant"
	"schedule_gateway/proto/auth"
)

type ResourcePredefine struct {
	AuthResource       *auth.Resource
	TokenResource      *auth.Resource
	RoleResource       *auth.Resource
	PermissionResource *auth.Resource
	UserResource       *auth.Resource
}

func InitResources() *ResourcePredefine {
	r := &ResourcePredefine{
		AuthResource: &auth.Resource{
			Id:   "1",
			Name: constant.AUTH_RESOURCE,
		},
		TokenResource: &auth.Resource{
			Id:   "2",
			Name: constant.TOKEN_RESOURCE,
		},
		RoleResource: &auth.Resource{
			Id:   "3",
			Name: constant.ROLE_RESOURCE,
		},
		PermissionResource: &auth.Resource{
			Id:   "4",
			Name: constant.PERMISSION_RESOURCE,
		},
		UserResource: &auth.Resource{
			Id:   "5",
			Name: constant.USER_RESOURCE,
		},
	}
	return r
}
