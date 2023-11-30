package system

import (
	"go-cms/config"
	"go-cms/global"
	"go-cms/model/system"
	"go-cms/utils"

	"go.uber.org/zap"
)

//@author:  [linjd] 10512203@qq.com
//@function: GetSystemConfig
//@description: 读取配置文件
//@return: err error, conf config.Server

type SystemConfigService struct {
}

func (systemConfigService *SystemConfigService) GetSystemConfig() (err error, conf config.Server) {
	return nil, global.CONFIG
}

// @description   set system config,
//@author:  [linjd] 10512203@qq.com
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func (systemConfigService *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.VP.Set(k, v)
	}
	err = global.VP.WriteConfig()
	return err
}

//@author: [88act-2](https://github.com/88act)
//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error

func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil {
		global.LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.LOG.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
