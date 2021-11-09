package source

import (
	"time"

	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/system"
	"github.com/gookit/color"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var Admin = new(admin)

type admin struct{}

var admins = []system.SysUser{
	{GO_MODEL: global.GO_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "超级管理员", HeaderImg: "https:///cms.88act.com/res/sys/0/header.jpg", AuthorityId: "888"},
	{GO_MODEL: global.GO_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "a303176530", Password: "3ec063004a6f31642261936a379fde3d", NickName: "QMPlusUser", HeaderImg: "https:///cms.88act.com/res/sys/0/logo.png", AuthorityId: "9528"},
}

//@author: [88act-2](https://github.com/88act)
//@description: sys_users 表数据初始化
func (a *admin) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]system.SysUser{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_users 表初始数据成功!")
		return nil
	})
}
