package plugin

import (
	"sync"

	"go-cms/global"
	"go-cms/model/business"
	"go-cms/myError"

	"github.com/gogf/gf/util/gconv"
)

var once_CollectManager sync.Once = sync.Once{}
var obj_CollectManager *CollectManager

/**
* 数据采集管理器，单例
 */
type CollectManager struct {
	Lock     sync.Mutex
	List     []*Collect
	runChan  chan business.ColCollect
	stopChan chan business.ColCollect
	exitChan chan int
}

/**
*获取单例
 */
func GetCollectManager() *CollectManager {
	once_CollectManager.Do(func() {
		obj_CollectManager = new(CollectManager)
		//instanse.init()
	})
	return obj_CollectManager
}

func (m *CollectManager) Init() {
	global.LOG.Info("爬虫采集器插件服务器 CollectManager init ....")
	m.List = []*Collect{}
	m.runChan = make(chan business.ColCollect, 1)
	m.stopChan = make(chan business.ColCollect, 1)
	m.exitChan = make(chan int, 1)
	go m.manager_run()
}

func (m *CollectManager) Start(data business.ColCollect, opt int) (err error) {
	collect_old := new(Collect)
	var idx int = -1
	for i, v := range m.List {
		if v.Data.ID == data.ID {
			collect_old = v
			idx = i
			break
		}
	}
	global.LOG.Debug("List存在旧的collect，idx=" + gconv.String(idx))
	if opt == 1 {
		//启动
		if idx >= 0 {
			if *collect_old.Data.StatusRun == 1 {
				global.LOG.Debug(data.Name + "启动，旧的已经在运行了")
				return myError.New(myError.ErrOK, "启动成功,旧的已经在运行了")
			} else {
				err = collect_old.Start()
				return err
			}
		} else {
			collect := new(Collect)
			collect.Data = data
			err = collect.Start()
			m.Lock.Lock()
			m.List = append(m.List, collect)
			m.Lock.Unlock()
			global.LOG.Debug(data.Name + " 正常启动 ")
			return err
		}
	} else if opt == 0 {
		//停止
		if idx >= 0 {
			collect_old.Stop()
			return myError.New(myError.ErrOK, "停止成功")
		} else {
			global.LOG.Debug(data.Name + " 停止,原来没启动")
			return myError.New(myError.ErrOK, "停止成功,原来没启动")
		}
	}
	return myError.New(myError.ErrFail, "失败，未知原因")
}

func (m *CollectManager) manager_run() {
	global.LOG.Info("CollectManager 开始运行 manager_run.....")
	for {
		select {
		// 启动
		case runCol := <-m.runChan:
			global.LOG.Info("CollectManager 运行。。 " + runCol.Name)
		// 停止
		case stopCol := <-m.stopChan:
			global.LOG.Info("CollectManager 停止。。 " + stopCol.Name)
		//退出Manager
		case _ = <-m.exitChan:
			//关闭所有的collect
			global.LOG.Info("CollectManager 退出")
			return
		}
	}
}
