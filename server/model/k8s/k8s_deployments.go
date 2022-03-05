package k8s

type K8sDeployment struct {
	ID         int    `json:"id"`
	Namespace  string `json:"namespace"`
	Deployment string `json:"deployment"`
	Replicas   int32  `json:"replicas"`
	CreateTime string `json:"createTime"`
}
