package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type K8sClustersSearch struct {
	business.K8sClusters
	request.PageInfo
}
