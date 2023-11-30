package main

import (
	"fmt"
	"os"

	"go-cms/core"
	"go-cms/global"
	"go-cms/initialize"
	"go-cms/utils"

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
	physicalID, err := utils.GetDevInfo()
	if err == nil {
		fmt.Println("服务器机器码 =>", physicalID)
		//utils.GetCpuInfo()
		beLicense, _ := utils.CheckLicense(physicalID)
		if !beLicense {
			fmt.Println("服务器未授权 ...  ")
			//return
		} else {
			fmt.Println("服务器授权成功")
		}
	} else {
		fmt.Println("读取硬件信息失败...  ", err.Error())
		//return
	}
	// 打印 数据库信息
	//mssql := &global.CONFIG.Mysql
	//fmt.Println("CONFIG.Mysql 变量 ... ")
	//fmt.Println(global.CONFIG.Mysql)
	//从 docker 环境变量获取
	// mssql.Path = getEnv("db_path", mssql.Path)
	// mssql.Dbname = getEnv("db_name", mssql.Dbname)
	// mssql.Username = getEnv("db_user", mssql.Username)
	// mssql.Password = getEnv("db_pwd", mssql.Password)
	// fmt.Println("从docker环境变量Mysql更新.... ")
	//fmt.Println(mssql)
	//system := &global.CONFIG.System
	//system.SuperAdmin = getEnvBool("super_admin", system.SuperAdmin)
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
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}
	core.RunServer()
}
func getEnv(key string, defaultValue string) string {
	myValue := defaultValue
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println("从docker环境1 key =" + key + " value = " + value)
		if !utils.IsEmpty(value) {
			myValue = value
		}
	}
	fmt.Println("从docker环境2 key =" + key + " myValue =" + myValue)
	return myValue
}
func getEnvBool(key string, defaultValue bool) bool {
	myValue := defaultValue
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println("从docker环境1 key =" + key + " value = " + value)
		if value == "true" {
			myValue = true
		}
		if value == "false" {
			myValue = false
		}
	}
	fmt.Println("从docker环境2  myValue =", myValue)
	return myValue
}
