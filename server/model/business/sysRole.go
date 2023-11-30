// 模板SysRole
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_SysRole_mini = "" // "id,created_at,updated_at,cu_id,pid,role_name,role_code,default_menu,status,

// SysRole 结构体
type SysRole struct {
	global.BaseModel
	UserId      int64  `json:"userId" form:"userId" cn:"录入人"  gorm:"column:user_id;comment:录入人;type:bigint"`
	Pid         int64  `json:"pid" form:"pid" cn:"父id"  gorm:"column:pid;comment:父id;type:bigint"`
	RoleName    string `json:"roleName" form:"roleName" cn:"角色名"  gorm:"column:role_name;comment:角色名;type:varchar(100);"`
	RoleCode    string `json:"roleCode" form:"roleCode" cn:"角色编码"  gorm:"column:role_code;comment:角色编码;type:varchar(10);"`
	DefaultMenu int64  `json:"defaultMenu" form:"defaultMenu" cn:"默认菜单"  gorm:"column:default_menu;comment:默认菜单;type:bigint"`
	Status      int    `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态;type:smallint"`
	MenuList    string `json:"menuList" form:"menuList" cn:"菜单列表"  gorm:"column:menu_list;comment:角色名;type:varchar(3000);"`
	ApiList     string `json:"apiList" form:"apiList" cn:"api列表"  gorm:"column:api_List;comment:角色名;type:varchar(3000);"`
}

// TableName SysRole 表名
func (SysRole) TableName() string {
	return "sys_role"
}

type SysRoleSearch struct {
	request.PageInfo
	global.BaseModel
	Pid         *int64 `json:"pid"  form:"pid" `
	UserId      *int64 `json:"userId" form:"userId" `
	RoleName    string `json:"roleName"   form:"roleName" `
	RoleCode    string `json:"roleCode"   form:"roleCode" `
	DefaultMenu *int64 `json:"defaultMenu"  form:"defaultMenu" `
	Status      *int   `json:"status"   form:"status" `
	MenuList    string `json:"menuList" form:"menuList"`
	ApiList     string `json:"apiList" form:"apiList"`
}
