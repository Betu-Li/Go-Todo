package main

import (
	"Gin-Todo/models"
	"Gin-Todo/routers"
	"Gin-Todo/setting"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

const ConfFilePath = "./conf/app.ini"

func main() {
	confFile := ConfFilePath
	if len(os.Args) > 2 {
		fmt.Println("use specified conf file: ", os.Args[1])
		confFile = os.Args[1]
	} else {
		fmt.Println("no configuration file was specified, use ./conf/app.ini")
	}
	// 加载配置文件
	if err := setting.Init(confFile); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// Debug: 打印加载的配置和端口值
	fmt.Printf("Loaded configuration: %+v\n", setting.Conf)
	fmt.Printf("Starting server on port: %d\n", setting.Conf.Port)

	// 创建数据库
	// sql:CREATE DATABASE gin_Todo
	// 连接数据库
	err := setting.InitMySql(setting.Conf.MySQLConfig)
	if err != nil {
		panic(err)
	}
	//程序退出时关闭数据库连接
	defer setting.DBClose()
	// 绑定数据库
	setting.DB.AutoMigrate(models.Todo{})

	// 注册路由
	r := routers.InitRouter()

	// Debug: 打印端口值
	fmt.Printf("Starting server on port: %d\n", setting.Conf.Port)

	// 启动服务
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
	}
}
