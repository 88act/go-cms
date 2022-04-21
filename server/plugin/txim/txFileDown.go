package plugin

import (
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	bizSev "go-cms/service/business"
	"go-cms/utils"
	"path"
	"strconv"
	"strings"
	"sync"
)

//文件下载器
type TxFileDown struct {
	statusRun int
}

var once_TxFileDown sync.Once = sync.Once{}
var obj_TxFileDown *TxFileDown
var wg sync.WaitGroup

/**
*获取单例
 */
func GetTxFileDown() *TxFileDown {
	once_TxFileDown.Do(func() {
		obj_TxFileDown = new(TxFileDown)
	})
	return obj_TxFileDown
}

func (m *TxFileDown) StartDownload() {
	if m.statusRun == 1 {
		global.LOG.Info("文件下载器txFileDown已经在下载中....statusRun =1")
		return
	}
	m.statusRun = 1
	//获取未下载的文件
	go m.doDownload()

}

func (m *TxFileDown) doDownload() {
	global.LOG.Info("下载  。。。。。 doDownload")
	var pageInfo bizReq.ImTxFileSearch
	pageInfo.Page = 1
	pageInfo.PageSize = 20
	pageInfo.OrderKey = "id"
	pageInfo.OrderDesc = false
	pageInfo.Status = utils.IntPtr(0)
	createdAtBetween := []string{}
	wg = sync.WaitGroup{}
	if list, _, err := bizSev.GetImTxFileSev().GetListAll(pageInfo, createdAtBetween, ""); err != nil {
		global.LOG.Error("下载文件ImTxFile获取失败" + err.Error())
		return
	} else {
		global.LOG.Info("等待下载文件 len = " + strconv.Itoa(len(list)))
		if len(list) == 0 {
			m.statusRun = 0
			global.LOG.Info("下载完成.正常退出doDownload")
			return
		}
		for _, v := range list {
			wg.Add(1)
			go downOne(v)
		}
	}
	wg.Wait()
	global.LOG.Info("下载完成 ...继续 doDownload ")
	m.doDownload()
}

func downOne(v business.ImTxFile) {
	filename := path.Base(v.Url)
	if fullPath, err := utils.DownloadFile(v.Url, v.Local, filename); err == nil {
		global.LOG.Info("downOne fullPath = " + fullPath)
		//更新数据库
		mp := make(map[string]interface{})
		///字符串替换, -1表示全部替换, 0表示不替换, 1表示替换第一个, 2表示替换第二个...
		local := strings.Replace(fullPath, global.CONFIG.Local.BasePath, "", -1)
		mp["local"] = local
		mp["status"] = 1
		ext := path.Ext(local)
		media_type := 2
		if ext == ".mp4" {
			media_type = 3
		} else if ext == ".m4a" {
			media_type = 4
		}
		mp["media_type"] = media_type

		global.LOG.Info(" 成功下载fullPath = " + fullPath)
		bizSev.GetImTxFileSev().UpdateByMap(v, mp)
	} else {
		global.LOG.Error("下载失败 v.Url =" + v.Url + err.Error())
	}
	wg.Done()
}
