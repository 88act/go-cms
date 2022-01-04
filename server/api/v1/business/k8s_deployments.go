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

type K8sDeploymentsApi struct {
}

 

// CreateK8sDeployments 创建K8sDeployments
// @Tags K8sDeployments
// @Summary 创建K8sDeployments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sDeployments true "创建K8sDeployments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sDeployments/createK8sDeployments [post]
func (k8sDeploymentsApi *K8sDeploymentsApi) CreateK8sDeployments(c *gin.Context) {
	var dataObj business.K8sDeployments
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetK8sDeploymentsSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteK8sDeployments 删除K8sDeployments
// @Tags K8sDeployments
// @Summary 删除K8sDeployments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sDeployments true "删除K8sDeployments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sDeployments/deleteK8sDeployments [delete]
func (k8sDeploymentsApi *K8sDeploymentsApi) DeleteK8sDeployments(c *gin.Context) {
	var k8sDeployments business.K8sDeployments
	_ = c.ShouldBindJSON(&k8sDeployments)
	if err := bizSev.GetK8sDeploymentsSev().Delete(k8sDeployments); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteK8sDeploymentsByIds 批量删除K8sDeployments
// @Tags K8sDeployments
// @Summary 批量删除K8sDeployments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sDeployments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /k8sDeployments/deleteK8sDeploymentsByIds [delete]
func (k8sDeploymentsApi *K8sDeploymentsApi) DeleteK8sDeploymentsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetK8sDeploymentsSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateK8sDeployments 更新K8sDeployments
// @Tags K8sDeployments
// @Summary 更新K8sDeployments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sDeployments true "更新K8sDeployments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sDeployments/updateK8sDeployments [put]
func (k8sDeploymentsApi *K8sDeploymentsApi) UpdateK8sDeployments(c *gin.Context) {
	var dataObj business.K8sDeployments
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetK8sDeploymentsSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindK8sDeployments 用id查询K8sDeployments
// @Tags K8sDeployments
// @Summary 用id查询K8sDeployments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.K8sDeployments true "用id查询K8sDeployments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sDeployments/findK8sDeployments [get]
func (k8sDeploymentsApi *K8sDeploymentsApi) FindK8sDeployments(c *gin.Context) {
	var k8sDeployments business.K8sDeployments
	_ = c.ShouldBindQuery(&k8sDeployments) 
	 rek8sDeployments,err:= bizSev.GetK8sDeploymentsSev().Get(k8sDeployments.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"k8sDeployments": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"k8sDeployments": rek8sDeployments}, c)
	}
}

// GetK8sDeploymentsList 分页获取K8sDeployments列表
// @Tags K8sDeployments
// @Summary 分页获取K8sDeployments列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.K8sDeploymentsSearch true "分页获取K8sDeployments列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sDeployments/getK8sDeploymentsList [get]
func (k8sDeploymentsApi *K8sDeploymentsApi) GetK8sDeploymentsList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.K8sDeploymentsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetK8sDeploymentsSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.K8sDeployments true "快速更新K8sDeployments" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /k8sDeployments/quickEdit [post] 
func (k8sDeploymentsApi *K8sDeploymentsApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "k8s_deployments" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel K8sDeployments列表
// @Tags K8sDeployments
// @Summary 分页导出excel K8sDeployments列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.K8sDeploymentsSearch true "分页导出excel K8sDeployments列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sDeployments/excelList [get]
func (k8sDeploymentsApi *K8sDeploymentsApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.K8sDeploymentsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetK8sDeploymentsSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "命名空间")  
					sheetFields = append(sheetFields, "应用")  
					sheetFields = append(sheetFields, "实例数")  
					sheetFields = append(sheetFields, "创建时间") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.Namespace)
				arr = append(arr, v.Deployment)
				arr = append(arr, *v.Replicas)
				arr = append(arr, v.CreateTime)   
			    excel.SetSheetRow("Sheet1", axis,&arr)  
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


 