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
