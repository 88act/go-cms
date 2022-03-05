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

type K8sClustersApi struct {
}

// CreateK8sClusters 创建K8sClusters
// @Tags K8sClusters
// @Summary 创建K8sClusters
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sClusters true "创建K8sClusters"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sClusters/createK8sClusters [post]
func (k8sClustersApi *K8sClustersApi) CreateK8sClusters(c *gin.Context) {
	var dataObj business.K8sClusters
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

	if id, err := bizSev.GetK8sClustersSev().Create(dataObj); err != nil {
		global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteK8sClusters 删除K8sClusters
// @Tags K8sClusters
// @Summary 删除K8sClusters
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sClusters true "删除K8sClusters"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sClusters/deleteK8sClusters [delete]
func (k8sClustersApi *K8sClustersApi) DeleteK8sClusters(c *gin.Context) {
	var k8sClusters business.K8sClusters
	_ = c.ShouldBindJSON(&k8sClusters)
	if err := bizSev.GetK8sClustersSev().Delete(k8sClusters); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteK8sClustersByIds 批量删除K8sClusters
// @Tags K8sClusters
// @Summary 批量删除K8sClusters
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sClusters"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /k8sClusters/deleteK8sClustersByIds [delete]
func (k8sClustersApi *K8sClustersApi) DeleteK8sClustersByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetK8sClustersSev().DeleteByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateK8sClusters 更新K8sClusters
// @Tags K8sClusters
// @Summary 更新K8sClusters
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sClusters true "更新K8sClusters"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sClusters/updateK8sClusters [put]
func (k8sClustersApi *K8sClustersApi) UpdateK8sClusters(c *gin.Context) {
	var dataObj business.K8sClusters
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetK8sClustersSev().Update(dataObj); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindK8sClusters 用id查询K8sClusters
// @Tags K8sClusters
// @Summary 用id查询K8sClusters
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.K8sClusters true "用id查询K8sClusters"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sClusters/findK8sClusters [get]
func (k8sClustersApi *K8sClustersApi) FindK8sClusters(c *gin.Context) {
	var k8sClusters business.K8sClusters
	_ = c.ShouldBindQuery(&k8sClusters)
	rek8sClusters, err := bizSev.GetK8sClustersSev().Get(k8sClusters.ID, "")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithData(gin.H{"k8sClusters": nil}, c)
	} else if err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"k8sClusters": rek8sClusters}, c)
	}
}

// GetK8sClustersList 分页获取K8sClusters列表
// @Tags K8sClusters
// @Summary 分页获取K8sClusters列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.K8sClustersSearch true "分页获取K8sClusters列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sClusters/getK8sClustersList [get]
func (k8sClustersApi *K8sClustersApi) GetK8sClustersList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.K8sClustersSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := bizSev.GetK8sClustersSev().GetList(pageInfo, createdAtBetween, ""); err != nil {
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
// @Param data body business.K8sClusters true "快速更新K8sClusters"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /k8sClusters/quickEdit [post]
func (k8sClustersApi *K8sClustersApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "k8s_clusters"
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// excelList 分页导出excel K8sClusters列表
// @Tags K8sClusters
// @Summary 分页导出excel K8sClusters列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.K8sClustersSearch true "分页导出excel K8sClusters列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sClusters/excelList [get]
func (k8sClustersApi *K8sClustersApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.K8sClustersSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, _, err := bizSev.GetK8sClustersSev().GetListAll(pageInfo, createdAtBetween, ""); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else {
			sheetFields := []string{}
			sheetFields = append(sheetFields, "集群名称")
			sheetFields = append(sheetFields, "Kubeconfig内容")
			sheetFields = append(sheetFields, "集群版本")

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.ClusterName)
				arr = append(arr, v.KubeConfig)
				arr = append(arr, v.ClusterVersion)
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
