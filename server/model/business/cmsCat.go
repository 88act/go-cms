// 模板CmsCat
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_CmsCat_mini = "" // "id,created_at,updated_at,user_id,be_sys,obj_id,type,title,image,be_top,be_nav,sort,status,

// CmsCat 结构体
type CmsCat struct {
	global.BaseModel
	Pid      int    `json:"pid" form:"pid" cn:"父ID"   gorm:"column:pid;comment:父ID;type:int"`
	UserId   int64  `json:"userId" form:"userId" cn:"用户id"  gorm:"column:user_id;comment:用户id;type:bigint"`
	BeSys    bool   `json:"beSys" form:"beSys" cn:"系统分类"   gorm:"column:be_sys;comment:系统分类: 0=否 1=是;type:tinyint"`
	ObjId    int64  `json:"objId" form:"objId" cn:"专题id"  gorm:"column:obj_id;comment:专题id:blog/crm/专题;type:bigint"`
	Type     int    `json:"type" form:"type" cn:"栏目类型"   gorm:"column:type;comment:文章类型;type:smallint"`
	Title    string `json:"title" form:"title" cn:"标题"  gorm:"column:title;comment:标题;type:varchar(200);"`
	Desc     string `json:"desc" form:"desc" cn:"摘要"  gorm:"column:desc;comment:摘要;type:varchar(500);"`
	TagList  string `json:"tagList" form:"tagList" cn:"标签列表"  gorm:"column:tag_list;comment:标签列表:id1,id2,id3,;type:varchar(500);"`
	Image    string `json:"image"  form:"image"  cn:"插图" gorm:"column:image;comment:插图;type:单图片"`
	FileList string `json:"fileList"  form:"fileList"  cn:"媒体列表" gorm:"column:file_list;comment:媒体列表:id列 1,2,3;type:多图片"`
	BeTop    bool   `json:"beTop" form:"beTop" cn:"置顶"   gorm:"column:be_top;comment:置顶:0否 1是;type:tinyint"`
	BeNav    bool   `json:"beNav" form:"beNav" cn:"是否导航"   gorm:"column:be_nav;comment:是否导航;type:tinyint"`
	Sort     int    `json:"sort" form:"sort" cn:"排序"   gorm:"column:sort;comment:排序;type:int"`
	Status   int    `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态:0未审核 1审核通过 2未通过审核 3黑名单;type:smallint"`
}

// TableName CmsCat 表名
func (CmsCat) TableName() string {
	return "cms_cat"
}

type CmsCatSearch struct {
	request.PageInfo
	global.BaseModel
	Pid      *int   `json:"pid"   form:"pid" `
	UserId   *int64 `json:"userId"  form:"userId" `
	BeSys    *bool  `json:"beSys" form:"beSys" `
	ObjId    *int64 `json:"objId"  form:"objId" `
	Type     *int   `json:"type"   form:"type" `
	Title    string `json:"title"   form:"title" `
	Desc     string `json:"desc"   form:"desc" `
	TagList  string `json:"tagList"   form:"tagList" `
	Image    string `json:"image" form:"image" `
	FileList string `json:"fileList" form:"fileList" `
	BeTop    *bool  `json:"beTop" form:"beTop" `
	BeNav    *bool  `json:"beNav" form:"beNav" `
	Sort     *int   `json:"sort"   form:"sort" `
	Status   *int   `json:"status"   form:"status" `
}
