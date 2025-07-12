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
	Server      Server      `mapstructure:"server" json:"server" yaml:"server"`
	Log         Log         `mapstructure:"log" json:"log" yaml:"log"`
	AuthService AuthService `mapstructure:"auth_service" json:"auth_service" yaml:"auth_service"`
	UserService UserService `mapstructure:"user_service" json:"user_service" yaml:"user_service"`
}

type Server struct {
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Mode string `mapstructure:"mode" json:"mode" yaml:"mode"`
}

type Log struct {
	Level       string `mapstructure:"level" json:"level" yaml:"level"`
	FileLogPath string `mapstructure:"file_log_path" json:"file_log_path" yaml:"file_log_path"`
	MaxSize     int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxAge      int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	Compress    bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}

type AuthService struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

type UserService struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}
