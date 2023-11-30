// 模板SysRoleMenu
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_SysRoleMenu_mini = "" // "id,created_at,updated_at,role_id,menu_id,status,

// SysRoleMenu 结构体
type SysRoleMenu struct {
	global.BaseModel
	RoleId int64 `json:"roleId" form:"roleId" cn:"角色"  gorm:"column:role_id;comment:角色;type:bigint"`
	MenuId int64 `json:"menuId" form:"menuId" cn:"菜单"  gorm:"column:menu_id;comment:菜单;type:bigint"`
	Status int   `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态;type:smallint"`
}

// TableName SysRoleMenu 表名
func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}

type SysRoleMenuSearch struct {
	request.PageInfo
	global.BaseModel
	RoleId *int64 `json:"roleId"  form:"roleId" `
	MenuId *int64 `json:"menuId"  form:"menuId" `
	Status *int   `json:"status"   form:"status" `
}
