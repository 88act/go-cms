package request

import (
	"go-cms/model/business"
	"go-cms/model/common/request"
)

type K8sDeploymentsSearch struct{
    business.K8sDeployments
    request.PageInfo
}