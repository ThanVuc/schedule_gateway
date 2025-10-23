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
	R2                      R2                      `mapstructure:"r2" json:"r2" yaml:"r2"`
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

type R2 struct {
	AccountID       string `mapstructure:"account_id" json:"account_id" yaml:"account_id"`
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKeyID     string `mapstructure:"access_key_id" json:"access_key_id" yaml:"access_key_id"`
	SecrecAccessKey string `mapstructure:"secret_access_key" json:"secret_access_key" yaml:"secret_access_key"`
	BucketName      string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	UseSSL          bool   `mapstructure:"use_ssl" json:"use_ssl" yaml:"use_ssl"`
	PublicURL       string `mapstructure:"public_url" json:"public_url" yaml:"public_url"`
}
