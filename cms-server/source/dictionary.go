package source

import (
	"time"

	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/system"
	"github.com/gookit/color"

	"gorm.io/gorm"
)

var Dictionary = new(dictionary)

type dictionary struct{}

var status = new(bool)

//@author: [88act-2](https://github.com/88act)
//@description: sys_dictionaries 表数据初始化
func (d *dictionary) Init() error {
	*status = true
	var dictionaries = []system.SysDictionary{
		{GO_MODEL: global.GO_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "性别", Type: "sex", Status: status, Desc: "性别字典"},
		{GO_MODEL: global.GO_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库int类型", Type: "int", Status: status, Desc: "int类型对应的数据库类型"},
		{GO_MODEL: global.GO_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库时间日期类型", Type: "time.Time", Status: status, Desc: "数据库时间日期类型"},
		{GO_MODEL: global.GO_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库浮点型", Type: "float64", Status: status, Desc: "数据库浮点型"},
		{GO_MODEL: global.GO_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库字符串", Type: "string", Status: status, Desc: "数据库字符串"},
		{GO_MODEL: global.GO_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库bool类型", Type: "bool", Status: status, Desc: "数据库bool类型"},
	}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 6}).Find(&[]system.SysDictionary{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_dictionaries 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&dictionaries).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_dictionaries 表初始数据成功!")
		return nil
	})
}
