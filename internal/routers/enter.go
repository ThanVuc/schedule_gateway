package routers

import (
	"schedule_gateway/internal/routers/authentication"
	"schedule_gateway/internal/routers/authorization"
	auth "schedule_gateway/internal/routers/authorization"
)

type RouterGroup struct {
	AuthorizationRouterEnter  *authorization.AuthorizationRouterGroup
	AuthenticationRouterEnter *authentication.AuthenticationRouterGroup
}

var RouterGroupApp *RouterGroup = &RouterGroup{
	AuthorizationRouterEnter:  &auth.AuthorizationRouterGroup{},
	AuthenticationRouterEnter: &authentication.AuthenticationRouterGroup{},
}
