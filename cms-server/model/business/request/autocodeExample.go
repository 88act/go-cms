// 自动生成模板SysDictionaryDetail
package request

import (
	"github.com/88act/go-cms/server/model/business"
	"github.com/88act/go-cms/server/model/common/request"
)

// 如果含有time.Time 请自行import time包
type BusinessExampleSearch struct {
	business.BusinessExample
	request.PageInfo
}
