package conf

import (
	"time"
)

var Conf Config

type Config struct {
	RunMode        string
	ServerConfig   ServerConfig
	DatabaseConfig DatabaseConfig
}

//服务端配置
type ServerConfig struct {
	HTTPPort          int
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	ServerContextPath string
}

//数据库配置
type DatabaseConfig struct {
	Host        string
	Type        string
	User        string
	Password    string
	Name        string
	TablePrefix string
}
