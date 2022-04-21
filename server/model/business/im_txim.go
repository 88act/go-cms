// 自动生成模板ImTxim
package business

import (
	"go-cms/global"
	"time"
)

// ImTxim 结构体
// 如果含有time.Time 请自行import time包
type ImTxim struct {
	ImTximMini
	UserSig string `json:"userSig" cn:"签名" form:"userSig" gorm:"column:user_sig;comment:签名;type:varchar(500);"`
}

// ImTximMini 结构体
type ImTximMini struct {
	global.GO_MODEL
	Name       string     `json:"name" cn:"名称" form:"name" gorm:"column:name;comment:名称;type:varchar(100);"`
	AppId      string     `json:"appId" cn:"appid" form:"appId" gorm:"column:app_id;comment:appid;type:varchar(100);"`
	Identifier string     `json:"identifier" cn:"管理员帐号" form:"identifier" gorm:"column:identifier;comment:管理员帐号;type:varchar(100);"`
	RunTimes   *int       `json:"runTimes" cn:"运行次数" form:"runTimes" gorm:"column:run_times;comment:运行次数;type:int"`
	BeginTime  *time.Time `json:"beginTime" cn:"开始时间" form:"beginTime" gorm:"column:begin_time;comment:开始时间;type:datetime"`
	NowTime    *time.Time `json:"nowTime" cn:"当前时间" form:"nowTime" gorm:"column:now_time;comment:当前时间;type:datetime"`
	Status     *int       `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0未上传,1正常;type:smallint"`
	StatusRun  *int       `json:"statusRun" cn:"运行状态" form:"statusRun" gorm:"column:status_run;comment:运行状态;type:smallint"`
}

// TableName ImTxim 表名
func (ImTxim) TableName() string {
	return "im_txim"
}
