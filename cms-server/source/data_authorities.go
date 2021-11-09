package source

import (
	"github.com/88act/go-cms/server/global"

	"github.com/gookit/color"
	"gorm.io/gorm"
)

var DataAuthorities = new(dataAuthorities)

type dataAuthorities struct{}

type DataAuthority struct {
	AuthorityId   string `gorm:"column:sys_authority_authority_id"`
	DataAuthority string `gorm:"column:data_authority_id_authority_id"`
}

var infos = []DataAuthority{
	{"888", "888"},
	{"888", "8881"},
	{"888", "9528"},
	{"9528", "8881"},
	{"9528", "9528"},
}

//@author: [88act-2](https://github.com/88act)
//@description: sys_data_authority_id 表数据初始化
func (d *dataAuthorities) Init() error {
	return global.DB.Table("sys_data_authority_id").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_authority_authority_id IN ('888', '9528') ").Find(&[]DataAuthority{}).RowsAffected == 5 {
			color.Danger.Println("\n[Mysql] --> sys_data_authority_id 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&infos).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_data_authority_id 表初始数据成功!")
		return nil
	})
}
