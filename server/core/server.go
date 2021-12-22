package core

import (
	"fmt"

	"time"

	"go-cms/global"
	"go-cms/initialize"
	plugin "go-cms/plugin/collect"
	"go-cms/service/system"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.CONFIG.Redis.Open == 1 {
		global.LOG.Debug("开启redis")
		initialize.Redis()
	} else {
		global.LOG.Debug("redis未开启")
	}

	// 从db加载jwt数据
	if global.DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))
	var collectManager2 = plugin.GetCollectManager()
	collectManager2.Init()
	fmt.Printf(`
	collectManager2.Init
	欢迎使用 github.com/88act/go-cms 
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080 
`, address)
	global.LOG.Error(s.ListenAndServe().Error())

}