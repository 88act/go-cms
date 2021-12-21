package source

import (
	"time"

	"go-cms/global"
	"go-cms/model/system"

	"github.com/gookit/color"

	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

var menus = []system.SysBaseMenu{
	{GO_MODEL: global.GO_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1, Meta: system.Meta{Title: "仪表盘", Icon: "setting"}},
	{GO_MODEL: global.GO_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7, Meta: system.Meta{Title: "关于我们", Icon: "info"}},
	{GO_MODEL: global.GO_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: system.Meta{Title: "超级管理员", Icon: "user-solid"}},
	{GO_MODEL: global.GO_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: system.Meta{Title: "角色管理", Icon: "s-custom"}},
	{GO_MODEL: global.GO_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: system.Meta{Title: "菜单管理", Icon: "s-order", KeepAlive: true}},
	{GO_MODEL: global.GO_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: system.Meta{Title: "API管理", Icon: "s-platform", KeepAlive: true}},
	{GO_MODEL: global.GO_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: system.Meta{Title: "用户管理", Icon: "coordinate"}},
	{GO_MODEL: global.GO_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: system.Meta{Title: "个人信息", Icon: "message-solid"}},
	{GO_MODEL: global.GO_MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 6, Meta: system.Meta{Title: "示例文件", Icon: "s-management"}},
	{GO_MODEL: global.GO_MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "excel", Name: "excel", Component: "view/example/excel/excel.vue", Sort: 4, Meta: system.Meta{Title: "excel导入导出", Icon: "s-marketing"}},
	{GO_MODEL: global.GO_MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: system.Meta{Title: "媒体库（上传下载）", Icon: "upload"}},
	{GO_MODEL: global.GO_MODEL{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: system.Meta{Title: "断点续传", Icon: "upload"}},
	{GO_MODEL: global.GO_MODEL{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: system.Meta{Title: "客户列表（资源示例）", Icon: "s-custom"}},
	{GO_MODEL: global.GO_MODEL{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: system.Meta{Title: "系统工具", Icon: "s-cooperation"}},
	{GO_MODEL: global.GO_MODEL{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "business", Name: "business", Component: "view/systemTools/business/index.vue", Sort: 1, Meta: system.Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
	{GO_MODEL: global.GO_MODEL{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: system.Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true}},
	{GO_MODEL: global.GO_MODEL{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: system.Meta{Title: "系统配置", Icon: "s-operation"}},
	{GO_MODEL: global.GO_MODEL{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: system.Meta{Title: "字典管理", Icon: "notebook-2"}},
	{GO_MODEL: global.GO_MODEL{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: system.Meta{Title: "字典详情", Icon: "s-order"}},
	{GO_MODEL: global.GO_MODEL{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: system.Meta{Title: "操作历史", Icon: "time"}},
	{GO_MODEL: global.GO_MODEL{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "https://cms.88act.com", Name: "https://cms.88act.com", Hidden: false, Component: "/", Sort: 0, Meta: system.Meta{Title: "官方网站", Icon: "s-home"}},
	{GO_MODEL: global.GO_MODEL{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "view/system/state.vue", Sort: 6, Meta: system.Meta{Title: "服务器状态", Icon: "cloudy"}},
	{GO_MODEL: global.GO_MODEL{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "14", Path: "superBuilderAdmin", Name: "superBuilderAdmin", Hidden: false, Component: "view/systemTools/superBuilderAdmin/index.vue", Sort: 1, Meta: system.Meta{Title: "自动化代码管理", Icon: "s-finance"}},
	{GO_MODEL: global.GO_MODEL{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "14", Path: "superBuilderEdit/:id", Name: "superBuilderEdit", Hidden: true, Component: "view/systemTools/superBuilder/index.vue", Sort: 0, Meta: system.Meta{Title: "自动化代码（复用）", Icon: "s-finance"}},
}

//@author: [88act-2](https://github.com/88act)
//@description: sys_base_menus 表数据初始化
func (m *menu) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 29}).Find(&[]system.SysBaseMenu{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_base_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_base_menus 表初始数据成功!")
		return nil
	})
}
