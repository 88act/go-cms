// 模板MemDepart
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_MemDepart_mini = "" // "id,created_at,updated_at,cu_id,pid,user_id,name,encode,desc,address,phone,email,sort,image,status,

// MemDepart 结构体
type MemDepart struct {
	global.BaseModel
	Pid      int64        `json:"pid" form:"pid" cn:"父id"  gorm:"column:pid;comment:父id;type:bigint"`
	UserId   int64        `json:"userId" form:"userId" cn:"管理员id"  gorm:"column:user_id;comment:管理员id;type:bigint"`
	Name     string       `json:"name" form:"name" cn:"名称"  gorm:"column:name;comment:名称;type:varchar(200);"`
	Encode   string       `json:"encode" form:"encode" cn:"编号"  gorm:"column:encode;comment:编号;type:varchar(100);"`
	Desc     string       `json:"desc" form:"desc" cn:"描述"  gorm:"column:desc;comment:描述;type:varchar(4000);"`
	Address  string       `json:"address" form:"address" cn:"地址"  gorm:"column:address;comment:地址;type:varchar(200);"`
	Phone    string       `json:"phone" form:"phone" cn:"联系电话"  gorm:"column:phone;comment:联系电话;type:varchar(200);"`
	Email    string       `json:"email" form:"email" cn:"邮箱"  gorm:"column:email;comment:邮箱;type:varchar(200);"`
	Sort     int          `json:"sort" form:"sort" cn:"排序"   gorm:"column:sort;comment:排序;type:int"`
	Image    string       `json:"image"  form:"image"  cn:"配图" gorm:"column:image;comment:配图;type:单图片"`
	Status   int          `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态:0未审核 1审核通过 2未通过审核 3草稿;type:smallint"`
	Children []*MemDepart `json:"children,omitempty" gorm:"-" sql:"-"  `
}

// TableName MemDepart 表名
func (MemDepart) TableName() string {
	return "mem_depart"
}

type MemDepartSearch struct {
	request.PageInfo
	global.BaseModel
	Pid     *int64 `json:"pid"  form:"pid" `
	UserId  *int64 `json:"userId"  form:"userId" `
	Name    string `json:"name"   form:"name" `
	Encode  string `json:"encode"   form:"encode" `
	Desc    string `json:"desc"   form:"desc" `
	Address string `json:"address"   form:"address" `
	Phone   string `json:"phone"   form:"phone" `
	Email   string `json:"email"   form:"email" `
	Sort    *int   `json:"sort"   form:"sort" `
	Image   string `json:"image" form:"image" `
	Status  *int   `json:"status"   form:"status" `
}
