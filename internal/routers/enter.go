package routers

import (
	"schedule_gateway/internal/routers/authentication"
	"schedule_gateway/internal/routers/authorization"
	personalschedule_router "schedule_gateway/internal/routers/personalschedule"
	user_route "schedule_gateway/internal/routers/user"
)

type RouterGroup struct {
	AuthorizationRouterEnter    *authorization.AuthorizationRouterGroup
	AuthenticationRouterEnter   *authentication.AuthenticationRouterGroup
	UserRouterEnter             *user_route.UserRouterGroup
	PersonalScheduleRouterEnter *personalschedule_router.PersonalscheduleRouterGroup
}

var RouterGroupApp *RouterGroup = &RouterGroup{
	AuthorizationRouterEnter:    &authorization.AuthorizationRouterGroup{},
	AuthenticationRouterEnter:   &authentication.AuthenticationRouterGroup{},
	UserRouterEnter:             &user_route.UserRouterGroup{},
	PersonalScheduleRouterEnter: &personalschedule_router.PersonalscheduleRouterGroup{},
}
