package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type K8sNamespacesSearch struct{
    business.K8sNamespaces
    request.PageInfo
}