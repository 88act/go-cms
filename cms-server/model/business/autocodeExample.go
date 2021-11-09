// 自动生成模板SysDictionaryDetail
package business

import (
	"github.com/88act/go-cms/server/global"
)

// 如果含有time.Time 请自行import time包
type BusinessExample struct {
	global.GO_MODEL
	BusinessExampleField string `json:"businessExampleField" form:"businessExampleField" gorm:"column:business_example_field;comment:仅作示例条目无实际作用"` // 展示值
}
