package main

import (
	"github.com/88act/go-cms/server/core"
	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.VP = core.Viper()      // 初始化Viper
	global.LOG = core.Zap()       // 初始化zap日志库
	global.DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	if global.DB != nil {
		// 不需要初始化表 by ljd 20211105
		//initialize.MysqlTables(global.DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()

}
