// 自动生成模板SysDictionaryDetail
package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

// 如果含有time.Time 请自行import time包
type BusinessExampleSearch struct {
	business.BusinessExample
	request.PageInfo
}
