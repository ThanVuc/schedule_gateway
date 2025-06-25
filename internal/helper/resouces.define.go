package helper

import v1 "schedule_gateway/internal/grpc/auth.v1"

type ResourcePredefine struct {
	AuthResource       *v1.Resource
	TokenResource      *v1.Resource
	RoleResource       *v1.Resource
	PermissionResource *v1.Resource
	UserResource       *v1.Resource
}

func InitResources() *ResourcePredefine {
	r := &ResourcePredefine{
		AuthResource: &v1.Resource{
			ResourceId: "1",
			Resource:   "auth",
		},
		TokenResource: &v1.Resource{
			ResourceId: "2",
			Resource:   "token",
		},
		RoleResource: &v1.Resource{
			ResourceId: "3",
			Resource:   "roles",
		},
		PermissionResource: &v1.Resource{
			ResourceId: "4",
			Resource:   "permissions",
		},
		UserResource: &v1.Resource{
			ResourceId: "5",
			Resource:   "uesrs",
		},
	}
	return r
}
