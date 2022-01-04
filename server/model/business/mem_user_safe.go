// 自动生成模板MemUserSafe
package business

import (
	"go-cms/global"
)

// MemUserSafe 结构体
// 如果含有time.Time 请自行import time包
type MemUserSafe struct {
     MemUserSafeMini
}

// MemUserSafeMini 结构体
type MemUserSafeMini struct {
      global.GO_MODEL
      UserId  *int `json:"userId" cn:"用户id" form:"userId" gorm:"column:user_id;comment:用户id;type:bigint"`
      Type  *int `json:"type" cn:"类型" form:"type" gorm:"column:type;comment:类型:1改密码 2改手机号 3改用户名 4实名认证;type:smallint"`
      ValOld  *int `json:"valOld" cn:"旧值" form:"valOld" gorm:"column:val_old;comment:旧值;type:smallint"`
      ValNew  *int `json:"valNew" cn:"新值" form:"valNew" gorm:"column:val_new;comment:新值;type:smallint"`
      Status  *int `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0=审核中,1=审核通过,2=审核未通过;type:smallint"`

}


// TableName MemUserSafe 表名
func (MemUserSafe) TableName() string {
  return "mem_user_safe"
}

