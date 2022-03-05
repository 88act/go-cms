package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type CmsArticleSearch struct {
	business.CmsArticle
	request.PageInfo
}
