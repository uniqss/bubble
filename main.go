package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	"bubble/setting"
	Util "bubble/util"
	"fmt"
	"go.uber.org/zap"
	"os"
)

func main() {
	log := Util.ZLog

	if len(os.Args) < 2 {
		log.Error("Usage：./bubble conf/config.ini")
		return
	}
	// 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		log.Error("load config from file failed", zap.Error(err))
		return
	}
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		log.Error("init mysql failed", zap.Error(err))
		return
	}
	defer dao.Close() // 程序退出关闭数据库连接
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		log.Error("server startup failed", zap.Error(err))
	}
}
