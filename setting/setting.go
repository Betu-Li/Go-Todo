package setting

import (
	"gopkg.in/ini.v1"
)

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	DBName   string `ini:"dbname"`
}

var (
	Conf = new(AppConfig)
)

// AppConfig 应用配置
type AppConfig struct {
	Release      bool `ini:"release"`
	Port         int  `ini:"port"`
	*MySQLConfig `ini:"database"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
