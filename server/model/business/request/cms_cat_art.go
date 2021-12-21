package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type CmsCatArtSearch struct {
	business.CmsCatArt
	request.PageInfo
}
