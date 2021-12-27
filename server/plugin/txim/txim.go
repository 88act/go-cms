package plugin

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-cms/global"
	"go-cms/model/business"
	"go-cms/myError"
	bizSev "go-cms/service/business"
	"go-cms/utils"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gogf/gf/util/gconv"
	"github.com/idoubi/goz"
	"mvdan.cc/xurls/v2"
)

//单个下载器
type Txim struct {
	Data      business.ImTxim
	statusRun int
}

// ImTxMsgFileMini 结构体
type FileResp struct {
	ChatType     string `json:"chatType"`
	MsgTime      string `json:"msgTime" `
	File         []TxFile
	ActionStatus string `json:"ActionStatus"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorCode    int32  `json:"ErrorCode"`
}
type TxFile struct {
	URL        string `json:"url"`
	ExpireTime string `json:"ExpireTime" `
	FileSize   int    `json:"fileSize"  `
	FileMd5    string `json:"fileMd5" `
	GzipSize   int    `json:"gzipSize" `
	GzipMd5    string `json:"gzipMd5"  `
}

// JsonStruct 结构体， json文件转换为 struct
type JsonStruct struct {
	FromAccount     string        `json:"from_Account" `
	ToAccount       string        `json:"to_Account"  `
	MsgTimestamp    int           `json:"msgTimestamp" `
	MsgSeq          int           `json:"msgSeq" `
	MsgRandom       int           `json:"msgRandom" `
	MsgBody         []MsgBodyStru `json:"MsgBody" `
	MsgFromPlatform string        `json:"msgFromPlatform" `
	ClientIp        string        `json:"clientIp" `
}

type MsgBodyStru struct {
	MsgType    string      `json:"msgType" `
	MsgContent ContentStru `json:"msgContent" `
}

type ContentStru struct {
	Text string `json:"Text" `
	Desc string `json:"Desc" `
	Data string `json:"Data" `
	Ext  string `json:"Ext" `
}

type ExtStru struct {
	Comment  string `json:"comment"`
	Score    int32  `json:"score" `
	Title    string `json:"title" `
	Sound    string `json:"sound"`
	ImageUrl string `json:"imageUrl"`
	Price    string `json:"price"`
}

func (m *Txim) Stop() (err error) {
	m.statusRun = 0
	err = global.DB.Model(&m.Data).Update("status_run", 0).Error
	if err != nil {
		return myError.New(myError.ErrDbUpdate, "停止失败，更新数据库失败")
	} else {
		global.LOG.Debug("更新数据库 StatusRun 成功")
	}
	return myError.New(myError.ErrOK, "停止成功")
}

func (m *Txim) Start() (err error) {
	//更新数据库

	// err = global.DB.Model(&m.Data).Update("status_run", 1).Error
	var statusRun int = 1
	m.Data.StatusRun = &statusRun
	err = bizSev.GetImTximSev().Update(m.Data)
	if err != nil {
		return myError.New(myError.ErrDbUpdate, "启动失败，更新数据库失败")
	} else {
		global.LOG.Debug("更新数据库 StatusRun成功")
	}
	global.LOG.Debug(m.Data.Name + " 启动，加入队列id=" + gconv.String(m.Data.ID))
	m.statusRun = 1
	fmt.Println("beginCollect----1-----")
	m.beginCollect()
	return myError.New(myError.ErrOK, "启动成功")
}

//读取腾讯接口
func (m *Txim) beginCollect() {
	global.LOG.Info("txim 请求beginCollect...")
	cli := goz.NewClient()
	resp, err := cli.Post("https://console.tim.qq.com/v4/open_msg_svc/get_history?usersig=eJwtjMEKgkAURf9ltoU97b0hhBaRENgQRbUI2hgz2aMcxEaTon9vUpf3nMv5iIPaB42pRCyiAMS426yNdXzlDme6YDuIp75nZclaxCECSAgRp70xbcmV8ZyIIgDoqePiz6Q-E6GUQ4Vz38Xa3lbN6PQ*T1KdbB*ZoRzbXO3qxeWo1s68ZsuNUzZNCOfi*wNdMjHF&identifier=admin&sdkappid=1400601443&contenttype=json", goz.Options{
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
		},
		JSON: struct {
			ChatType string `json:"ChatType"`
			MsgTime  string `json:"MsgTime"`
		}{"C2C", "2021122317"},
	})

	if err != nil {
		global.LOG.Error("txim 请求错误:" + err.Error())
		return
	}

	body, _ := resp.GetBody()

	fmt.Println(body)

	var fileResp FileResp

	err = json.Unmarshal([]byte(body), &fileResp)
	if err != nil {
		global.LOG.Error("txim json解析错误:" + err.Error())
		return
	}
	//入库操作
	//if fileResp.ActionStatus =="OK" {

	var txf = business.ImTxMsgFile{}
	txf.ChatType = fileResp.ChatType
	txf.MsgTime = fileResp.MsgTime
	txf.ActionStatus = fileResp.ActionStatus
	txf.ErrorCode = string(fileResp.ErrorCode)
	txf.ErrorInfo = fileResp.ErrorInfo
	status := 1
	txf.Status = &status
	//TODO: 这里还需要处理多个文件的情况
	if len(fileResp.File) > 0 {
		txf.Url = fileResp.File[0].URL
		txf.GzipSize = &fileResp.File[0].GzipSize
		txf.GzipMd5 = fileResp.File[0].GzipMd5
		txf.FileMd5 = fileResp.File[0].FileMd5
		txf.FileSize = &fileResp.File[0].FileSize
		txf.ExpireTime = utils.Str2Time(fileResp.File[0].ExpireTime)
	}
	id, err2 := bizSev.GetImTxMsgFileSev().Create(txf)
	if err2 != nil {
		global.LOG.Error("txim 数据库操作失败:" + err.Error())
		return
	}
	txf.ID = id
	go downloadFile(txf)
}

//下载文件
func downloadFile(txf business.ImTxMsgFile) {
	yearmonth := string([]rune(txf.MsgTime)[:6])
	dayhour := string([]rune(txf.MsgTime)[6:10]) //string([]byte(data.MsgTime)[4:8])  rune 兼容汉字
	// res/sys/txim/202112/2110/****.gz
	local := global.CONFIG.Local.BasePath + global.CONFIG.Local.Path + "/txim/" + yearmonth + "/" + dayhour + "/"
	global.LOG.Error("txim  local :" + local)
	zipPath, err := utils.DownloadFile(txf.Url, local, txf.MsgTime+".gz")
	if err != nil {
		global.LOG.Info("txim 下载downloadFile失败:" + err.Error())
		return
	}

	//	zipPath := local + txf.MsgTime + ".gz"
	unzipfile, err := utils.Unzip_tx(zipPath, local)
	if err != nil {
		global.LOG.Info("txim 解压失败:" + err.Error() + " path = " + zipPath)
		return
	}
	global.LOG.Info("txim  解压后  unzipfile :" + unzipfile)

	err = json2struct(unzipfile, txf, local)
	if err != nil {
		global.LOG.Info("txim json入库失败 :" + err.Error())
		return
	}

}

// 转换 json数据到数据库
func json2struct(filepath string, txf business.ImTxMsgFile, local string) (err error) {
	//打开文件
	inputFile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	//按行读取文件
	msgList := []JsonStruct{}
	fileReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := fileReader.ReadString('\n')
		if readerError == io.EOF {
			//fmt.Printf("读错误")
			break
		}
		//去掉首尾回车和，号
		inputString = strings.Trim(inputString, ",\n") // inputString[:len(inputString)-1] //
		msg := JsonStruct{}
		if err := json.Unmarshal([]byte(inputString), &msg); err == nil {
			//fmt.Println(" MsgType ==  " + msg.FromAccount)
			msgList = append(msgList, msg)
		} else {
			global.LOG.Error("转换失败 : " + err.Error() + ", String =" + inputString)
			//fmt.Println(inputString)
		}
	}
	//imTxMsgList := []business.ImTxMsg{}
	status := 0
	for _, v := range msgList {
		//	fmt.Println(v)
		obj := business.ImTxMsg{}
		obj.FromAccount = v.FromAccount
		obj.ToAccount = v.ToAccount
		// fmt.Println(v.MsgTimestamp)
		// ti := utils.intto // gconv.Time(v.MsgTimestamp, "2006-01-02 15:04:05")
		// fmt.Println("v.MsgTimestamp")
		// fmt.Println(ti)

		obj.MsgTimestamp = utils.Int2Time(int64(v.MsgTimestamp))
		fmt.Println("v.MsgTimestamp")
		fmt.Println(obj.MsgTimestamp)
		obj.MsgSeq = &v.MsgSeq
		obj.MsgRandom = &v.MsgRandom
		obj.MsgFromPlatform = v.MsgFromPlatform
		obj.ClientIp = v.ClientIp

		urlList := []string{}
		if len(v.MsgBody) > 0 {
			obj.MsgType = v.MsgBody[0].MsgType
			b, err := json.Marshal(v.MsgBody[0].MsgContent)
			if err == nil {
				obj.MsgContent = string(b)
			}
			obj.MsgText = v.MsgBody[0].MsgContent.Text
			//提取图片url
			rxStrict := xurls.Strict()
			urlList = rxStrict.FindAllString(v.MsgBody[0].MsgContent.Ext, -1)
			bytes, e := json.Marshal(urlList)
			if e == nil {
				obj.MediaListTx = string(bytes)
			}
		}

		//imTxMsgList = append(imTxMsgList, obj)

		//写数据库
		obj.CreatedAt = time.Now()
		obj.ChatType = txf.ChatType
		obj.MsgTime = txf.MsgTime
		obj.Status = &status
		obj.StatusMedia = &status
		//保存 ImTxMsg
		id, err := bizSev.GetImTxMsgSev().Create(obj)
		if err != nil {
			global.LOG.Error("创建 ImTxMsg 出错 " + err.Error())
		} else {
			// 保存  ImTxFile
			fileID := []uint{}
			for _, url := range urlList {
				if !utils.IsEmpty(url) {
					var txFile business.ImTxFile
					var idint int = int(id)
					txFile.MsgId = &idint
					txFile.Url = url
					txFile.Local = local // 图片等 下载到相同的目录内
					txFile.Status = &status
					//txFile.MediaType = 1
					fid, err2 := bizSev.GetImTxFileSev().Create(txFile)
					if err2 != nil {
						global.LOG.Error("创建 ImTxFile 出错 " + err2.Error())
					} else {
						fileID = append(fileID, fid)
					}
				}
			}

		}
	}
	// 通知下载文件
	GetTximManager().DownChan <- 1
	return nil
}
