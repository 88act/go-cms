package k8s

type K8sPods struct {
	ID           int    `json:"id" gorm:"primarykey;AUTO_INCREMENT"`
	PodName      string `form:"podName" json:"podName"`
	PodIP        string `json:"podIP"`
	HostIP       string `json:"hostIP"`
	Status       string `json:"status"`
	StartTime    string `json:"startTime"`
	RestartCount int32  `json:"restartCount"`
	NameSpace    string `form:"namespace" json:"namespace"`
	Line         int64  `form:"line" json:"line"`
}
