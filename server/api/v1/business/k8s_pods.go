package business

import (
	"errors"
	"fmt"
	"go-cms/global"
	"go-cms/model/business"
	bizReq "go-cms/model/business/request"
	"go-cms/model/common/request"
	"go-cms/model/common/response"
	bizSev "go-cms/service/business"
	commSev "go-cms/service/common"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type K8sPodsApi struct {
}

// CreateK8sPods 创建K8sPods
// @Tags K8sPods
// @Summary 创建K8sPods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sPods true "创建K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sPods/createK8sPods [post]
func (k8sPodsApi *K8sPodsApi) CreateK8sPods(c *gin.Context) {
	var dataObj business.K8sPods
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if id, err := bizSev.GetK8sPodsSev().Create(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteK8sPods 删除K8sPods
// @Tags K8sPods
// @Summary 删除K8sPods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sPods true "删除K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sPods/deleteK8sPods [delete]
func (k8sPodsApi *K8sPodsApi) DeleteK8sPods(c *gin.Context) {
	var k8sPods business.K8sPods
	_ = c.ShouldBindJSON(&k8sPods)
	if err := bizSev.GetK8sPodsSev().Delete(k8sPods); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteK8sPodsByIds 批量删除K8sPods
// @Tags K8sPods
// @Summary 批量删除K8sPods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /k8sPods/deleteK8sPodsByIds [delete]
func (k8sPodsApi *K8sPodsApi) DeleteK8sPodsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetK8sPodsSev().DeleteByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateK8sPods 更新K8sPods
// @Tags K8sPods
// @Summary 更新K8sPods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sPods true "更新K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sPods/updateK8sPods [put]
func (k8sPodsApi *K8sPodsApi) UpdateK8sPods(c *gin.Context) {
	var dataObj business.K8sPods
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetK8sPodsSev().Update(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindK8sPods 用id查询K8sPods
// @Tags K8sPods
// @Summary 用id查询K8sPods
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.K8sPods true "用id查询K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sPods/findK8sPods [get]
func (k8sPodsApi *K8sPodsApi) FindK8sPods(c *gin.Context) {
	var k8sPods business.K8sPods
	_ = c.ShouldBindQuery(&k8sPods)
	rek8sPods, err := bizSev.GetK8sPodsSev().Get(k8sPods.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"k8sPods": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"k8sPods": rek8sPods}, c)
	}
}

// GetK8sPodsList 分页获取K8sPods列表
// @Tags K8sPods
// @Summary 分页获取K8sPods列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.K8sPodsSearch true "分页获取K8sPods列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sPods/getK8sPodsList [get]
func (k8sPodsApi *K8sPodsApi) GetK8sPodsList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.K8sPodsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetK8sPodsSev().GetList(pageInfo, createdAtBetween, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// QuickEdit 快速更新
// @Tags QuickEdit
// @Summary 快速更新
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sPods true "快速更新K8sPods"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /k8sPods/quickEdit [post]
func (k8sPodsApi *K8sPodsApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "k8s_pods"
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel K8sPods列表
// @Tags K8sPods
// @Summary 分页导出excel K8sPods列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.K8sPodsSearch true "分页导出excel K8sPods列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sPods/excelList [get]
func (k8sPodsApi *K8sPodsApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.K8sPodsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetK8sPodsSev().GetListAll(pageInfo, createdAtBetween, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "容器名称")
			sheetFields = append(sheetFields, "容器IP")
			sheetFields = append(sheetFields, "主机IP")
			sheetFields = append(sheetFields, "状态")
			sheetFields = append(sheetFields, "启动时间")
			sheetFields = append(sheetFields, "重启次数")
			sheetFields = append(sheetFields, "命名空间")
			sheetFields = append(sheetFields, "line")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.PodName)
				arr = append(arr, v.PodIp)
				arr = append(arr, v.HostIp)
				arr = append(arr, v.Status)
				arr = append(arr, v.StartTime)
				arr = append(arr, *v.RestartCount)
				arr = append(arr, v.NameSpace)
				arr = append(arr, *v.Line)
				excel.SetSheetRow("Sheet1", axis, &arr)
			}
			filename := fmt.Sprintf("ecl%d.xlsx", time.Now().Unix())
			filePath := global.CONFIG.Local.BasePath + global.CONFIG.Local.Path + "/excel/" + filename
			url := global.CONFIG.Local.BaseUrl + global.CONFIG.Local.Path + "/excel/" + filename
			err := excel.SaveAs(filePath)
			if err != nil {
				global.LOG.Error(err.Error())
				response.FailWithMessage("获取失败", c)
			} else {
				resData := map[string]string{"url": url, "filename": filename}
				response.OkWithData(resData, c)
			}
		}
	}
}
