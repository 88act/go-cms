package plugin

import (
	"sync"

	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/business"
	"github.com/88act/go-cms/server/myError"
	"github.com/gogf/gf/util/gconv"
)

var once sync.Once = sync.Once{}
var instanse *CollectManager

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
	once.Do(func() {
		instanse = new(CollectManager)
		//instanse.init()
		//instanse.Name = "这是一个单例"
	})
	return instanse
}

func (m *CollectManager) Init() {
	global.LOG.Debug("CollectManager init ....")
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
			if 1 == *collect_old.Data.StatusRun {
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
	global.LOG.Debug("CollectManager 开始运行 manager_run.....")
	for {
		select {
		// 启动
		case runCol := <-m.runChan:
			global.LOG.Debug("CollectManager 运行。。 " + runCol.Name)
		// 停止
		case stopCol := <-m.stopChan:
			global.LOG.Debug("CollectManager 停止。。 " + stopCol.Name)
		//退出Manager
		case _ = <-m.exitChan:
			//关闭所有的collect
			global.LOG.Debug("CollectManager 退出")
			return
		}
	}
}
