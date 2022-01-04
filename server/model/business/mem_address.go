// 自动生成模板MemAddress
package business

import (
	"go-cms/global"
)

// MemAddress 结构体
// 如果含有time.Time 请自行import time包
type MemAddress struct {
     MemAddressMini
      Address  string `json:"address" cn:"详细地址" form:"address" gorm:"column:address;comment:详细地址;type:varchar(120);"`
}

// MemAddressMini 结构体
type MemAddressMini struct {
      global.GO_MODEL
      UserId  *int `json:"userId" cn:"用户id" form:"userId" gorm:"column:user_id;comment:用户id;type:bigint"`
      Realname  string `json:"realname" cn:"收货人" form:"realname" gorm:"column:realname;comment:收货人;type:varchar(60);"`
      Phone  string `json:"phone" cn:"手机" form:"phone" gorm:"column:phone;comment:手机;type:varchar(20);"`
      Email  string `json:"email" cn:"邮箱地址" form:"email" gorm:"column:email;comment:邮箱地址;type:varchar(60);"`
      AreaCode  string `json:"areaCode" cn:"地区编码" form:"areaCode" gorm:"column:area_code;comment:地区编码;type:varchar(20);"`
      ZipCode  string `json:"zipCode" cn:"邮政编码" form:"zipCode" gorm:"column:zip_code;comment:邮政编码;type:varchar(60);"`
      IsDefault  *bool `json:"isDefault" cn:"默认地址" form:"isDefault" gorm:"column:is_default;comment:默认地址;type:tinyint"`
      Status  *int `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态;type:smallint"`

}


// TableName MemAddress 表名
func (MemAddress) TableName() string {
  return "mem_address"
}

