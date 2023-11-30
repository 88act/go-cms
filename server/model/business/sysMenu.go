// 模板SysMenu
// Author 10512203@qq.com
package business

import (
	"go-cms/global"
	"go-cms/model/common/request"
)

// 表格字段
const Field_SysMenu_mini = "" // "id,created_at,updated_at,cu_id,pid,name,path,title,level,auths,show_link,active_path,frame_src,component,param,sort,keep_alive,icon,status,

// SysMenu 结构体
type SysMenu struct {
	global.BaseModel
	Pid        int64      `json:"pid" form:"pid" cn:"父菜单ID"  gorm:"column:pid;comment:父菜单ID;type:bigint"`
	Name       string     `json:"name" form:"name" cn:"路由name"  gorm:"column:name;comment:路由name;type:varchar(255);"`
	Path       string     `json:"path" form:"path" cn:"路由path"  gorm:"column:path;comment:路由path;type:varchar(255);"`
	Title      string     `json:"title" form:"title" cn:"标题"  gorm:"column:title;comment:标题;type:varchar(255);"`
	Level      int        `json:"level" form:"level" cn:"等级"   gorm:"column:level;comment:等级;type:int"`
	Auths      string     `json:"auths" form:"auths" cn:"按钮权限"  gorm:"column:auths;comment:按钮权限;type:varchar(255);"`
	ShowLink   bool       `json:"showLink" form:"showLink" cn:"是否隐藏"   gorm:"column:show_link;comment:是否隐藏;type:tinyint"`
	ActivePath string     `json:"activePath" form:"activePath" cn:"活动链接"  gorm:"column:active_path;comment:活动链接;type:varchar(255);"`
	FrameSrc   string     `json:"frameSrc" form:"frameSrc" cn:"外链"  gorm:"column:frame_src;comment:外链;type:varchar(255);"`
	Component  string     `json:"component" form:"component" cn:"前端文件路径"  gorm:"column:component;comment:前端文件路径;type:varchar(255);"`
	Param      string     `json:"param" form:"param" cn:"参数"  gorm:"column:param;comment:参数;type:varchar(255);"`
	Sort       int64      `json:"sort" form:"sort" cn:"排序"  gorm:"column:sort;comment:排序;type:bigint"`
	KeepAlive  bool       `json:"keepAlive" form:"keepAlive" cn:"保持连接"   gorm:"column:keep_alive;comment:保持连接;type:tinyint"`
	Icon       string     `json:"icon" form:"icon" cn:"图标"  gorm:"column:icon;comment:图标;type:varchar(255);"`
	Status     int        `json:"status" form:"status" cn:"状态"   gorm:"column:status;comment:状态;type:smallint"`
	Children   []*SysMenu `json:"children,omitempty" gorm:"-" sql:"-"  `
}

// TableName SysMenu 表名
func (SysMenu) TableName() string {
	return "sys_menu"
}

type SysMenuSearch struct {
	request.PageInfo
	global.BaseModel
	Pid        *int64 `json:"pid"  form:"pid" `
	Name       string `json:"name"   form:"name" `
	Path       string `json:"path"   form:"path" `
	Title      string `json:"title"   form:"title" `
	Level      *int   `json:"level"   form:"level" `
	Auths      string `json:"auths"   form:"auths" `
	ShowLink   *bool  `json:"showLink" form:"showLink" `
	ActivePath string `json:"activePath"   form:"activePath" `
	FrameSrc   string `json:"frameSrc"   form:"frameSrc" `
	Component  string `json:"component"   form:"component" `
	Param      string `json:"param"   form:"param" `
	Sort       *int64 `json:"sort"  form:"sort" `
	KeepAlive  *bool  `json:"keepAlive" form:"keepAlive" `
	Icon       string `json:"icon"   form:"icon" `
	Status     *int   `json:"status"   form:"status" `
}

type MenuMeta struct {
	Title string `json:"title" form:"title" cn:"标题"   `
	Rank  int64  `json:"rank" form:"rank" cn:"排序"  `
	//Sort      int64  `json:"sort" form:"sort" cn:"排序"  `
	KeepAlive  bool   `json:"keepAlive" form:"keepAlive"   `
	Icon       string `json:"icon" form:"icon" cn:"图标"`
	Roles      string `json:"roles" form:"role" cn:"角色"  `
	Auths      string `json:"auths" form:"auths" cn:"按钮权限"  `
	FrameSrc   string `json:"frameSrc" form:"frameSrc" cn:"外链"  `
	ShowLink   bool   `json:"showLink" form:"showLink" cn:"是否显示" `
	ActivePath string `json:"activePath" form:"activePath" cn:"激活链接"  `
	//Hidden    bool           `json:"hidden" form:"hidden" cn:"是否隐藏"   gorm:"column:hidden;comment:是否隐藏;type:tinyint"`
	Param string `json:"param" form:"param" cn:"参数"  gorm:"column:param;comment:参数;type:varchar(255);"`
}

type SysMenuMini struct {
	Id int64 `form:"id" json:"id"` // 主键ID
	//CuId      int64          `json:"cuId" form:"cuId" cn:"商户"  gorm:"column:cu_id;comment:商户:非0表示特定的商户菜单;type:bigint"`
	Pid       int64          `json:"pid" form:"pid" cn:"父菜单ID"  gorm:"column:pid;comment:父菜单ID;type:bigint"`
	Name      string         `json:"name" form:"name" cn:"路由name"  gorm:"column:name;comment:路由name;type:varchar(255);"`
	Path      string         `json:"path" form:"path" cn:"路由path"  gorm:"column:path;comment:路由path;type:varchar(255);"`
	Level     int            `json:"level" form:"level" cn:"等级"   gorm:"column:level;comment:等级;type:int"`
	Component string         `json:"component" form:"component" cn:"前端文件路径"  gorm:"column:component;comment:前端文件路径;type:varchar(255);"`
	Meta      MenuMeta       `json:"meta" form:"meta"  `
	Children  []*SysMenuMini `json:"children,omitempty" gorm:"-" sql:"-"  `
}

type SysMenuMiniResp struct {
	List        []*SysMenuMini `json:"list" gorm:"-" sql:"-"   `
	DefaultMenu string         `json:"defaultMenu"  form:"defaultMenu" `
}
