package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type SysSuperBuilderHistoriesSearch struct {
	business.SysSuperBuilderHistories
	request.PageInfo
}
