package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type K8sPodsSearch struct {
	business.K8sPods
	request.PageInfo
}
