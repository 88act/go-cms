package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type K8sNodesSearch struct{
    business.K8sNodes
    request.PageInfo
}