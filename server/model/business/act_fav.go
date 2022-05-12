// 自动生成模板ActFav
package business

import (
	"go-cms/global"
)

// ActFav 结构体 
type ActFav struct {
     ActFavMini
}

// ActFavMini 结构体
type ActFavMini struct {
      global.GO_MODEL
      UserId  *int `json:"userId" cn:"用户id" form:"userId" gorm:"column:user_id;comment:用户id;type:bigint"`
      Type  *int `json:"type" cn:"收藏类型" form:"type" gorm:"column:type;comment:收藏类型:1活动,2文章;type:smallint"`
      ObjId  *int `json:"objId" cn:"对象id" form:"objId" gorm:"column:obj_id;comment:对象id;type:bigint"`
      Status  *int `json:"status" cn:"状态" form:"status" gorm:"column:status;comment:状态:0未审核 1审核通过 2未通过审核 3黑名单;type:smallint"`

}


// TableName ActFav 表名
func (ActFav) TableName() string {
  return "act_fav"
}

