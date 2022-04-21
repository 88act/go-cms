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
	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
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
	MsgContent interface{} `json:"msgContent" `
}

// TIMCustomElem
type TIMCustomElem struct {
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

type ChatStart struct {
	Status  int    `json:"status"`
	OrderId string `json:"orderId" `
}

//TIMTextElem
type TIMTextElem struct {
	Text string `json:"Text"`
}

//TIMImageElem
type TIMImageElem struct {
	UUID           string           `json:"comment"`
	ImageFormat    int              `json:"ImageFormat" `
	ImageInfoArray []ImageInfoArray `json:"ImageInfoArray" `
}
type ImageInfoArray struct {
	Type   int    `json:"Type"`
	Size   int    `json:"Size" `
	Width  int    `json:"Width" `
	Height int    `json:"Height" `
	URL    string `json:"URL" `
}

func (m *Txim) Stop() (err error) {
	m.statusRun = 2
	mapdata := make(map[string]interface{})
	mapdata["status_run"] = 2
	err = bizSev.GetImTximSev().UpdateByMap(m.Data, mapdata)
	cornTxim.Stop()
	cornTxim.Remove(cornTxim_id)
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
	mapdata := make(map[string]interface{})
	mapdata["status_run"] = 1
	err = bizSev.GetImTximSev().UpdateByMap(m.Data, mapdata)
	if err != nil {
		return myError.New(myError.ErrDbUpdate, "启动失败，更新数据库失败")
	} else {
		global.LOG.Debug("更新数据库 StatusRun成功")
	}
	global.LOG.Debug(m.Data.Name + " 启动，加入队列id=" + gconv.String(m.Data.ID))
	m.statusRun = 1
	m.beginCollect()
	return myError.New(myError.ErrOK, "启动成功")
}

func (m *Txim) beginCollect() {
	spec1 := "*/10 * * * * ?" // 10秒   20分钟执行一次
	id, _ := cornTxim.AddFunc(spec1, m.beginCollect_do)
	cornTxim_id = id
	cornTxim.Start()
}

//读取腾讯接口
func (m *Txim) beginCollect_do() {
	// 加入定时任务
	global.LOG.Info("txim 开始腾讯接口 beginCollect_do...")
	hh, _ := time.ParseDuration("1h")
	h24, _ := time.ParseDuration("-24h")
	//比较时间大小的方法有：Before, After, Equal
	//global.LOG.Debug("NowTime ==1==" + m.Data.NowTime.Add(hh).Format("2006-01-02 15:04:05"))
	//	global.LOG.Debug("NowTime ==2==" + time.Now().Format("2006-01-02 15:04:05"))
	mytime := *m.Data.NowTime
	if mytime.Add(hh).After(time.Now().Add(h24)) {
		//if mytime.Add(hh).After(time.Now()) {
		global.LOG.Info("时间超前,24小时后才下载文件，返回，" + mytime.Format("2006-01-02 15:04:05"))
		//m.Stop()
		return
	}
	//加1 小时
	myNowTime := m.Data.NowTime
	//global.LOG.Debug("NowTime ====" + m.Data.NowTime.Format("2006-01-02 15:04:05"))
	//global.LOG.Debug("NowTime =" + utils.TimeToStr(m.Data.NowTime))
	MsgTime := utils.TimeToStr(myNowTime)
	global.LOG.Debug("MsgTime =" + MsgTime)
	ChatType := "C2C"
	url := "https://console.tim.qq.com/v4/open_msg_svc/get_history?usersig=" +
		m.Data.UserSig + "&identifier=" + m.Data.Identifier + "&sdkappid=" + m.Data.AppId + "&contenttype=json"
	cli := goz.NewClient()

	resp, err := cli.Post(url, goz.Options{
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
		},
		JSON: struct {
			ChatType string `json:"ChatType"`
			MsgTime  string `json:"MsgTime"`
		}{ChatType, MsgTime}, // }{"C2C", "2022031718"}, "2022032211"
	})

	if err != nil {
		global.LOG.Error("txim 请求错误 setp1:" + err.Error())
		return
	}
	body, _ := resp.GetBody()
	fmt.Println("=============body================")
	fmt.Println(body)

	//
	var fileResp FileResp
	err = json.Unmarshal([]byte(body), &fileResp)
	if err != nil {
		global.LOG.Error("txim json解析错误:" + err.Error())
		return
	}

	// 更新数据库 运行次数 当前时间
	*m.Data.NowTime = m.Data.NowTime.Add(hh)
	*m.Data.RunTimes = *m.Data.RunTimes + 1
	mapdata := make(map[string]interface{})
	mapdata["run_times"] = *m.Data.RunTimes
	mapdata["now_time"] = *m.Data.NowTime
	_ = bizSev.GetImTximSev().UpdateByMap(m.Data, mapdata)

	// 文件未准备好 ，退出
	//if fileResp.ErrorInfo == "Err_File_Not_Ready"  {
	if fileResp.ErrorCode == 1005 { //Err_File_Expired
		global.LOG.Info("文件未准备好 ，退出 Err_File_Expired  MsgTime=" + MsgTime)
		return
	}

	//文件入库操作
	//if fileResp.ActionStatus =="OK" {
	var txf = business.ImTxMsgFile{}
	txf.ChatType = ChatType // fileResp.ChatType
	txf.MsgTime = MsgTime   //fileResp.MsgTime
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
	jsons, _ := json.Marshal(txf)
	global.LOG.Info("腾讯返回数据：" + string(jsons))
	//if txf.ErrorInfo == "Err_File_Not_Ready" || txf.ErrorInfo == "Err_File_Expired" {
	if fileResp.ErrorCode == 1005 || fileResp.ErrorCode == 1004 {
		global.LOG.Info("没有文件 Err_File_Not_Ready Err_File_Expired  MsgTime=" + MsgTime)
	} else {
		go downloadFile(txf)
	}

}

//下载文件
func downloadFile(txf business.ImTxMsgFile) {
	if utils.IsEmpty(txf.Url) {
		fmt.Println(txf)
		global.LOG.Error("txf.Url= null")
		return
	}

	yearmonth := string([]rune(txf.MsgTime)[:6])
	dayhour := string([]rune(txf.MsgTime)[6:10]) //string([]byte(data.MsgTime)[4:8])  rune 兼容汉字
	// res/sys/txim/202112/2110/****.gz
	local := global.CONFIG.Local.BasePath + global.CONFIG.Local.Path + "/txim/" + yearmonth + "/" + dayhour + "/"
	//global.LOG.Error("txim  local :" + local)
	zipPath, err := utils.DownloadFile(txf.Url, local, txf.MsgTime+".gz")
	if err != nil {
		global.LOG.Error("txim 下载downloadFile失败:" + err.Error())
		global.LOG.Error("txf.Url=" + txf.Url)
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
	global.LOG.Error(" filepath = " + filepath)
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
			global.LOG.Error("读错误")

			break
		}
		fmt.Println(" inputString ==  " + inputString)
		if utils.IsEmpty(inputString) {
			global.LOG.Error("inputString IsEmpty")

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
		//fmt.Println("v.MsgTimestamp")
		//fmt.Println(obj.MsgTimestamp)
		obj.MsgSeq = &v.MsgSeq
		obj.MsgRandom = &v.MsgRandom
		obj.MsgFromPlatform = v.MsgFromPlatform
		obj.ClientIp = v.ClientIp

		urlList := []string{}
		for _, mbs := range v.MsgBody {
			obj.MsgType = mbs.MsgType
			obj.MsgContent = gconv.String(mbs.MsgContent)

			//obj.MsgText = obj.MsgText +  gconv.String(mbs.MsgContent) + ";"
			//提取图片url
			rxStrict := xurls.Strict()
			mylist := rxStrict.FindAllString(gconv.String(mbs.MsgContent), -1)
			urlList = append(urlList, mylist...)
			bytes, e := json.Marshal(urlList)
			if e == nil {
				txurl := string(bytes)
				if !utils.IsEmpty(txurl) {
					obj.MediaListTx = obj.MediaListTx + txurl + ";"
				}
			}
			switch obj.MsgType {
			case "TIMTextElem":
				global.LOG.Debug("TIMTextElem")
				tIMTextElem := TIMTextElem{}
				if err := json.Unmarshal([]byte(gconv.String(mbs.MsgContent)), &tIMTextElem); err == nil {
					obj.MsgText = tIMTextElem.Text
				} else {
					global.LOG.Error("转换TIMTextElem失败 : " + err.Error() + ", String =" + gconv.String(mbs.MsgContent))
				}
			case "TIMCustomElem":
				global.LOG.Debug("TIMCustomElem")
				tIMCustomElem := TIMCustomElem{}
				if err := json.Unmarshal([]byte(gconv.String(mbs.MsgContent)), &tIMCustomElem); err == nil {
					obj.MsgText = tIMCustomElem.Text
					//订单状态chatstart
					if tIMCustomElem.Data == "chatstart" {
						chatStart := ChatStart{}
						if err := json.Unmarshal([]byte(tIMCustomElem.Ext), &chatStart); err == nil {
							obj.OrderId = chatStart.OrderId
							obj.OrderStatus = &chatStart.Status
						} else {
							global.LOG.Error("转换订单失败 : " + err.Error() + ", String =" + tIMCustomElem.Ext)
						}
					} else {
						global.LOG.Debug("自定义消息 TIMCustomElem.data =" + tIMCustomElem.Data)
					}

				} else {
					global.LOG.Error("转换订单失败 : " + err.Error() + ", String =" + gconv.String(mbs.MsgContent))
				}
			case "TIMVideoFileElem":
				global.LOG.Debug("TIMVideoFileElem")
			case "TIMSoundElem":
				global.LOG.Debug("TIMSoundElem")
			case "TIMImageElem":
				global.LOG.Debug("TIMImageElem")
			default:
				global.LOG.Error("未知类型MsgType = " + obj.MsgType)
			}

		}

		// 判断是否 是开始 结束 订单标记

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
					if url == "https://web.sdk.qcloud.com/im/assets/images/transparent.png" || strings.Contains(url, "?imageView") {
						global.LOG.Info("被忽略的图片 url = " + url)
					} else {
						txFile := business.ImTxFile{}
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
	}
	// 通知下载图片媒体文件
	GetTximManager().DownChan <- 1
	return nil
}

// // 获取sig签名 ， uid 用户名 ， expire 过期时间 秒
func (m *Txim) GetSig(uid string, expire int) (sig string, err error) {
	//sig, err = tencentyun.GenSig("1", "2", "xiaojun", 86400*180)
	sig, err = tencentyun.GenUserSig(global.CONFIG.Txim.Appid, global.CONFIG.Txim.Key, uid, expire)
	if err != nil {
		fmt.Println(err.Error())
		global.LOG.Error(err.Error())
		return "", err
	} else {
		return sig, err
	}
}
