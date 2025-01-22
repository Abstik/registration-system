package main

import (
	"fmt"

	"registration_system/dao/mysql"
	"registration_system/logger"
	"registration_system/routers"
	"registration_system/settings"
	"registration_system/utils"
)

func main() {
	// 初始化加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	//初始化zap日志库
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	//初始化mysql
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()  // 程序退出关闭数据库连接
	utils.InitSqlTable() // 建表

	// 注册路由
	r := routers.SetupRouter(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
