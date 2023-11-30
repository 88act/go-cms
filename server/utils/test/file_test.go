package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-cms/global"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_unzip(t *testing.T) {

	// file, _ := os.Open("C:/work_go/cuckooserver/res/sys/txim/2021122317/111.zip")
	// destPath, err := unpackit.Unpack(file, "C:/work_go/cuckooserver/res/sys/txim/2021122317")
	// if err != nil {
	// 	fmt.Println("txim解压失败:" + err.Error())
	// 	return
	// }
	// fmt.Println(destPath)
	// // err := utils.UnPack("C:/work_go/cuckooserver/res/sys/txim/2021122317/2021122317.gz", "C:/work_go/cuckooserver/res/sys/txim/2021122317")
	// // if err != nil {
	// // 	fmt.Println("txim解压失败:" + err.Error())
	// // 	return
	// // }

	err := json2struct("C:/work_go/cuckooserver/res/sys/txim/202112/2317/unknown-pack")
	if err != nil {
		fmt.Println(err.Error())
	}

	// inputString := strings.Trim(",sdsd{\"sdsd sdsdsdsd{zS dfsf}sadasdda\"} ,", ",")
	// fmt.Println(inputString)
}

func json2struct(filepath string) (err error) {
	//打开文件
	inputFile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	//按行读取文件
	msgList := []ImTxMsg2{}
	fileReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := fileReader.ReadString('\n')
		if readerError == io.EOF {
			//fmt.Printf("读错误")
			break
		}
		inputString = strings.Trim(inputString, ",\n") // inputString[:len(inputString)-1] //
		msg := ImTxMsg2{}
		if err := json.Unmarshal([]byte(inputString), &msg); err == nil {
			//fmt.Println(" MsgType ==  " + msg.FromAccount)
			msgList = append(msgList, msg)
		} else {
			global.LOG.Error("转换失败 : " + err.Error() + ", String =" + inputString)
			//fmt.Println(inputString)
		}
	}
	// imTxMsgList := []business.ImTxMsg{}

	// for _, v := range msgList {
	// 	//fmt.Println(v)
	// 	obj := &business.ImTxMsg{}
	// 	if err := gconv.Struct(v, obj); err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	imTxMsgList = append(imTxMsgList, *obj)
	// }

	// for _, v := range imTxMsgList {
	// 	fmt.Println(v)
	// }
	return nil
}

// ImTxMsg 结构体
// 如果含有time.Time 请自行import time包
type ImTxMsg2 struct {
	FromAccount     string        `json:"from_Account" cn:"发送人" form:"fromAccount" gorm:"column:from_account;comment:发送人;type:varchar(100);"`
	ToAccount       string        `json:"to_Account" cn:"接收人" form:"toAccount" gorm:"column:to_account;comment:接收人;type:varchar(100);"`
	MsgTimestamp    int           `json:"msgTimestamp" cn:"时间撮" form:"msgTimestamp" gorm:"column:msg_timestamp;comment:时间撮;type:int"`
	MsgSeq          int           `json:"msgSeq" cn:"seq" form:"msgSeq" gorm:"column:msg_seq;comment:seq;type:int"`
	MsgRandom       int           `json:"msgRandom" cn:"random" form:"msgRandom" gorm:"column:msg_random;comment:random;type:varchar(100);"`
	MsgBody         []MsgBodyStru `json:"MsgBody" cn:"random" form:"msgRandom" gorm:"column:msg_random;comment:random;type:varchar(100);"`
	MsgFromPlatform string        `json:"msgFromPlatform" cn:"平台" form:"msgFromPlatform" gorm:"column:msg_from_platform;comment:平台;type:varchar(255);"`
	ClientIp        string        `json:"clientIp" cn:"IP地址" form:"clientIp" gorm:"column:client_ip;comment:IP地址;type:varchar(255);"`
}

type MsgBodyStru struct {
	MsgType    string      `json:"msgType" cn:"消息类型" form:"msgType" gorm:"column:msg_type;comment:消息类型;type:varchar(100);"`
	MsgContent ContentStru `json:"msgContent" cn:"内容" form:"msgContent" gorm:"column:msg_content;comment:内容;type:text;"`
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
