// 模板CmsVisitor
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_CmsVisitor_mini = "" // "id,created_at,updated_at,user_id,type,obj_id,plat_type,client_ip,status,

// CmsVisitor 结构体
type CmsVisitor struct {
	global.BaseModel
	UserId   int64 `json:"userId" form:"userId" cn:"用户id"  gorm:"column:user_id;comment:用户id;type:bigint"`
	Type     int   `json:"type" form:"type" cn:"类型"   gorm:"column:type;comment:类型:0文章,1blog,2 team;type:smallint"`
	ObjId    int64 `json:"objId" form:"objId" cn:"对象Id"  gorm:"column:obj_id;comment:对象Id;type:bigint"`
	PlatType int   `json:"platType" form:"platType" cn:"平台"   gorm:"column:plat_type;comment:平台:1 wx 2 web,3 andriod,4 ios;type:smallint"`
	ClientIp int64 `json:"clientIp" form:"clientIp" cn:"客户端ip"  gorm:"column:client_ip;comment:客户端ip;type:bigint"`
	Status   int   `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态: 0未审核 1审核 2未通过审核 3 草稿;type:smallint"`
}

// TableName CmsVisitor 表名
func (CmsVisitor) TableName() string {
	return "cms_visitor"
}

type CmsVisitorSearch struct {
	request.PageInfo
	global.BaseModel
	UserId   *int64 `json:"userId"  form:"userId" `
	Type     *int   `json:"type"   form:"type" `
	ObjId    *int64 `json:"objId"  form:"objId" `
	PlatType *int   `json:"platType"   form:"platType" `
	ClientIp *int64 `json:"clientIp"  form:"clientIp" `
	Status   *int   `json:"status"   form:"status" `
}
