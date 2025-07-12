package routers

import (
	"schedule_gateway/internal/routers/authentication"
	"schedule_gateway/internal/routers/authorization"
	user_route "schedule_gateway/internal/routers/user"
)

type RouterGroup struct {
	AuthorizationRouterEnter  *authorization.AuthorizationRouterGroup
	AuthenticationRouterEnter *authentication.AuthenticationRouterGroup
	UserRouterEnter           *user_route.UserRouterGroup
}

var RouterGroupApp *RouterGroup = &RouterGroup{
	AuthorizationRouterEnter:  &authorization.AuthorizationRouterGroup{},
	AuthenticationRouterEnter: &authentication.AuthenticationRouterGroup{},
	UserRouterEnter:           &user_route.UserRouterGroup{},
}
