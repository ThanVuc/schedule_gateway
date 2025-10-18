package settings

/*
	@Author: Sinh
	@Date: 2025/6/1
	@Description: This file create a configuration struct for the application.
	@Note: The configuration struct is used to store the configuration of the application.
	All the configuration is loaded from a YAML file using Viper.
	And it is stored in the global variable `Config` in the `global` package.
*/

type Config struct {
	Server                  Server                  `mapstructure:"server" json:"server" yaml:"server"`
	Log                     Log                     `mapstructure:"log" json:"log" yaml:"log"`
	AuthService             AuthService             `mapstructure:"auth_service" json:"auth_service" yaml:"auth_service"`
	UserService             UserService             `mapstructure:"user_service" json:"user_service" yaml:"user_service"`
	CsrfSecret              string                  `mapstructure:"csrf_secret" json:"csrf_secret" yaml:"csrf_secret"`
	SessionSecret           string                  `mapstructure:"session_secret" json:"session_secret" yaml:"session_secret"`
	PersonalScheduleService PersonalScheduleService `mapstructure:"personal_schedule_service" json:"personal_schedule_service" yaml:"personal_schedule_service"`
}

type Server struct {
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Mode string `mapstructure:"mode" json:"mode" yaml:"mode"`
}

type Log struct {
	Level string `mapstructure:"level" json:"level" yaml:"level"`
}

type AuthService struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

type UserService struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

type PersonalScheduleService struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}
