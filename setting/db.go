package setting

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitMySql(cfg *MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = DB.DB().Ping()
	return
}

func DBClose() {
	DB.Close()
}
