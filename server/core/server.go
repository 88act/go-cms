package core

import (
	"fmt"
	"go-cms/global"
	"go-cms/initialize"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.CONFIG.Redis.Open == 1 {
		global.LOG.Debug("开启redis")
		initialize.Redis()
	} else {
		global.LOG.Debug("redis未开启")
	}

	// 从db加载jwt数据
	// if global.DB != nil {
	// 	system.LoadAll()
	// }

	Router := initialize.Routers()
	//Router.Static("/form-generator", "./resource/page")

	initPluginServer()

	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)

	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(20 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(` 
	欢迎使用 go-cms 
	服务地址:http://127.0.0.1%s/
	`, address)
	global.LOG.Error(s.ListenAndServe().Error())

}

// 启动插件的各种服务
func initPluginServer() {
	//fmt.Println("启动插件的各种服务...")
	//collect.GetCollectManager().Init()
	//txim.GetTximManager().Init()
}
