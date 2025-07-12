package helper

import "schedule_gateway/proto/auth"

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
			Name: "auth",
		},
		TokenResource: &auth.Resource{
			Id:   "2",
			Name: "token",
		},
		RoleResource: &auth.Resource{
			Id:   "3",
			Name: "roles",
		},
		PermissionResource: &auth.Resource{
			Id:   "4",
			Name: "permissions",
		},
		UserResource: &auth.Resource{
			Id:   "5",
			Name: "users",
		},
	}
	return r
}
