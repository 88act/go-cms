package plugin

import (
	"fmt"
	"strconv"
	"sync"

	"go-cms/global"
	"go-cms/model/business"
	"go-cms/myError"

	"github.com/gogf/gf/util/gconv"
	"github.com/robfig/cron/v3"
)

var once_TximManager sync.Once = sync.Once{}
var obj_TximManager *TximManager
var cornTxim *cron.Cron
var cornTxim_id cron.EntryID

/**
* 数据采集管理器，单例
 */
type TximManager struct {
	Lock     sync.Mutex
	List     []*Txim
	runChan  chan business.ImTxim
	stopChan chan business.ImTxim
	DownChan chan int
	ExitChan chan int
}

/**
*获取单例
 */
func GetTximManager() *TximManager {
	once_TximManager.Do(func() {
		obj_TximManager = new(TximManager)
		//instanse.init()
	})
	return obj_TximManager
}

func (m *TximManager) Init() {
	global.LOG.Info("腾讯IM消息下载插件 TximManager init ....")
	m.List = []*Txim{}
	m.runChan = make(chan business.ImTxim, 10)
	m.stopChan = make(chan business.ImTxim, 10)
	m.DownChan = make(chan int, 10)
	m.ExitChan = make(chan int, 1)
	cornTxim = cron.New(cron.WithSeconds())
	go m.manager_run()
}

func (m *TximManager) Start(data business.ImTxim, opt int) (err error) {
	txim_old := new(Txim)
	var idx int = -1
	for i, v := range m.List {
		if v.Data.ID == data.ID {
			txim_old = v
			idx = i
			break
		}
	}
	global.LOG.Debug("List存在旧的collect，idx=" + gconv.String(idx))
	if opt == 1 {
		//启动
		if idx >= 0 {
			if *txim_old.Data.StatusRun == 1 {
				global.LOG.Debug(data.Name + "启动，旧的已经在运行了")
				return myError.New(myError.ErrOK, "启动成功,旧的已经在运行了")
			} else {
				err = txim_old.Start()
				return err
			}
		} else {
			txim := new(Txim)
			txim.Data = data
			err = txim.Start()
			m.Lock.Lock()
			m.List = append(m.List, txim)
			m.Lock.Unlock()
			global.LOG.Debug(data.Name + " 正常启动 ")
			return err
		}
	} else if opt == 0 {
		//停止
		if idx >= 0 {
			txim_old.Stop()
			return myError.New(myError.ErrOK, "停止成功")
		} else {
			global.LOG.Debug(data.Name + " 停止,原来没启动")
			return myError.New(myError.ErrOK, "停止成功,原来没启动")
		}
	}
	return myError.New(myError.ErrFail, "失败，未知原因")
}

func (m *TximManager) manager_run() {
	global.LOG.Info("TximManager 开始运行 manager_run.....")
	for {
		select {
		// 启动
		case runCol := <-m.runChan:
			global.LOG.Info("TximManager 运行。。 " + runCol.Name)
		// 停止
		case stopCol := <-m.stopChan:
			global.LOG.Info("TximManager 停止。。 " + stopCol.Name)
		// 通知下载文件
		case down := <-m.DownChan:
			global.LOG.Info("TximManager 启动下载。。 " + strconv.Itoa(down))
			go GetTxFileDown().StartDownload()
		//退出Manager
		case ex := <-m.ExitChan:
			//关闭所有的collect
			fmt.Println(ex)
			global.LOG.Info("TximManager 退出")
			return
		}
	}
}
