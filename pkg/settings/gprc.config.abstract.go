package settings

type GrpcBase interface {
	GetHost() string
	GetPort() int
}

func (ac *AuthService) GetHost() string {
	return ac.Host
}

func (ac *AuthService) GetPort() int {
	return ac.Port
}

func (us *UserService) GetHost() string {
	return us.Host
}

func (us *UserService) GetPort() int {
	return us.Port
}

func (ps *PersonalScheduleService) GetHost() string {
	return ps.Host
}

func (ps *PersonalScheduleService) GetPort() int {
	return ps.Port
}

func (ns *NotificationService) GetHost() string {
	return ns.Host
}

func (ns *NotificationService) GetPort() int {
	return ns.Port
}
