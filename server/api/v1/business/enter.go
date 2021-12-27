package business

import (
	plugin "go-cms/plugin/collect"
)

type ApiGroup struct {
	// Code generated by github.com/88act/go-cms/server Begin; DO NOT EDIT.
	ColHsjApi
	ColCollectApi2
	K8sPodsApi
	K8sNodesApi
	K8sNamespacesApi
	K8sDeploymentsApi
	K8sClustersApi
	CmsAdApi
	CmsArticleApi
	CmsCatApi
	BasicFileApi
	CmsAdSeatApi
	MemUserApi
	MemUserSafeApi
	ImTximApi2
	ColKeyFieldApi
	ColCollectApi
	ImTxFileApi
	ImTximApi
	ImTxMsgApi
	ImTxMsgFileApi
	// Code generated by github.com/88act/go-cms/server End; DO NOT EDIT.
}

var collectManager = plugin.GetCollectManager()
