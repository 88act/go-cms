package main

import (
	"fmt"
	"os"

	"go-cms/core"
	"go-cms/global"
	"go-cms/initialize"
	"go-cms/model/business"

	"github.com/gin-gonic/gin"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {

	global.VP = core.Viper() // 初始化Viper
	global.LOG = core.Zap()  // 初始化zap日志库
	// 打印 数据库信息
	mssql := &global.CONFIG.Mysql
	fmt.Println("CONFIG.Mysql 变量 ... ")
	fmt.Println(global.CONFIG.Mysql)
	//从 docker 环境变量获取
	mssql.Path = getEnv("db_path", mssql.Path)
	mssql.Dbname = getEnv("db_name", mssql.Dbname)
	mssql.Username = getEnv("db_user", mssql.Username)
	mssql.Password = getEnv("db_pwd", mssql.Password)
	fmt.Println("从docker环境变量Mysql更新.... ")
	fmt.Println(mssql)
	if global.CONFIG.System.GoMode == "debug" {
		fmt.Println("GoMode = DebugMode. ")
		gin.SetMode(gin.DebugMode)
	} else {
		fmt.Println("GoMode = ReleaseMode. ")
		gin.SetMode(gin.ReleaseMode)
	}
	global.DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	if global.DB != nil {
		// 不需要初始化表 by ljd 20211105
		initialize.MysqlTables(global.DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}
	var obj = business.CmsCat{}
	err := global.DB.First(&obj).Error
	if err != nil {
		fmt.Println("数据库error = " + err.Error())
	}
	fmt.Println(obj)

	core.RunWindowsServer()

}

func getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println("从docker环境 key =" + key + " value = " + value)
		return value
	}
	fmt.Println("从docker环境 key =" + key + " value = null")
	return defaultValue
}
