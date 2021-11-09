package source

import (
	"time"

	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/example"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var File = new(file)

type file struct{}

var files = []example.ExaFileUploadAndDownload{
	{global.GO_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "10.png", "https://logo.png", "png", "158787308910.png"},
	{global.GO_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "logo.png", "https:/ Avatar.png", "png", "1587973709logo.png"},
}

//@author: [88act-2](https://github.com/88act)
//@description: exa_file_upload_and_downloads 表初始化数据
func (f *file) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]example.ExaFileUploadAndDownload{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> exa_file_upload_and_downloads 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&files).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> exa_file_upload_and_downloads 表初始数据成功!")
		return nil
	})
}
