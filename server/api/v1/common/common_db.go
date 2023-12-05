package common

import (
	"fmt"
	"go-cms/global"
	"go-cms/model/common/request"
	commSev "go-cms/service/common"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommonDbApi struct {
	commSev.BaseApi
}

// QuickEdit 更新QuickEdit
// @Tags common
// @Summary 更新QuickEdit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.CmsAd true "更新 QuickEdit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"

func (m *CommonDbApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	//var_dump.Dump(quickEdit)

	if err := commSev.GetCommonDbSev().QuickEdit(c, quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		m.FailWithMessage("更新失败", c)
	} else {
		m.OkWithMessage("更新成功", c)
	}
}

// PidData 获取父子关系的数据，通常用在下拉列表
//
//	js:  let taskReq = {
//		table: "ims_task",
//		pidField: "id",
//		nameField: "title"
//		pidValue: "title"
//	}
//
// await this.getPidData('task', taskReq)
func (m *CommonDbApi) PidData(c *gin.Context) {
	var req request.PidDataReq
	_ = c.ShouldBindJSON(&req)
	// cuId := utils.GetCuId(c)
	// req.CuId = &cuId
	if list, err := commSev.GetCommonDbSev().GetPidData(c, req); err == nil {
		m.OkWithDetailed(request.PidDataResp{
			List: list,
		}, "获取成功", c)
	} else {
		m.OkWithMessage("获取pidData失败", c)
	}
}

// PidTreeData 树状数据 ，通常用在多级选择

func (m *CommonDbApi) PidTreeData(c *gin.Context) {
	var req request.PidDataReq
	_ = c.ShouldBindJSON(&req)

	var list []*request.PidTreeData
	if req.Table == "ims_job_task" {
		list = commSev.GetCommonDbSev().GetTreeData_jobTask(c, req)
	} else if req.Table == "mem_depart" {
		list = commSev.GetCommonDbSev().GetTreeData_memDepart(c, req)
	} else if req.Table == "devCustomer" {
		list = commSev.GetCommonDbSev().GetTreeData_devCustomer(c, req)
	} else if req.Table == "basic_file_cat" {
		list = commSev.GetCommonDbSev().GetTreeData_basicFileCat(c, req)
	} else {
		list = commSev.GetCommonDbSev().GetTreeData(c, req)
	}

	m.OkWithDetailed(request.PidDataTreeResp{
		List: list,
	}, "获取成功", c)

}

// GetDict  通常单级选择
//
//	js:  let taskReq = {
//		table: "ims_task",
//		pidField: "id",
//		nameField: "title"
//		pidValue: "title"
//	}
//
// await this.getDictData('task', taskReq)
func (m *CommonDbApi) GetDict(c *gin.Context) {
	var req request.GetDictReq
	_ = c.ShouldBindJSON(&req)
	if list := commSev.GetCommonDbSev().GetDict(c, req); list != nil {
		m.OkWithDetailed(request.PidDataResp{
			List: list,
		}, "获取成功", c)
	} else {
		m.OkWithMessage("获取pidData失败", c)
	}
}

func (m *CommonDbApi) GetDict2(c *gin.Context) {
	var req request.GetDictReq
	_ = c.ShouldBindJSON(&req)

	if list := commSev.GetCommonDbSev().GetDict2(c, req); list != nil {
		m.OkWithDetailed(request.PidDataResp2{
			List: list,
		}, "获取成功", c)
	} else {
		m.OkWithMessage("获取GetDict2失败", c)
	}
}

// GetDictTree 树状数据 ，通常用在多级选择
// PidData 获取父子关系的数据，通常用在下拉列表
//
//	js:  let taskReq = {
//		table: "ims_task",
//		pidField: "id",
//		nameField: "title"
//		pidValue: "title"
//	}
//
// await this.getDictData('task', taskReq)
func (m *CommonDbApi) GetDictTree(c *gin.Context) {
	var req request.GetDictReq
	_ = c.ShouldBindJSON(&req)
	fmt.Println("GetDictTree 111")
	fmt.Println(req)
	if list := commSev.GetCommonDbSev().GetDictTree(c, req); list != nil {
		m.OkWithDetailed(request.PidDataTreeResp{
			List: list,
		}, "获取成功", c)
	} else {
		m.OkWithMessage("获取GetDictTree失败", c)
	}
}

// 更新树状结构的 pid和sort
func (m *CommonDbApi) UpdatePidSort(c *gin.Context) {
	var req request.UpdatePidSortReq
	_ = c.ShouldBindJSON(&req)

	fmt.Print(req)
	if err := commSev.GetCommonDbSev().UpdatePidSort(c, req); err != nil {
		m.FailWithMessage("修改失败:"+err.Error(), c)
	} else {
		m.OkWithMessage("修改成功", c)
	}

}
