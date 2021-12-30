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

type K8sNamespacesApi struct {
}

 

// CreateK8sNamespaces 创建K8sNamespaces
// @Tags K8sNamespaces
// @Summary 创建K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sNamespaces true "创建K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNamespaces/createK8sNamespaces [post]
func (k8sNamespacesApi *K8sNamespacesApi) CreateK8sNamespaces(c *gin.Context) {
	var dataObj business.K8sNamespaces
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetK8sNamespacesSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteK8sNamespaces 删除K8sNamespaces
// @Tags K8sNamespaces
// @Summary 删除K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sNamespaces true "删除K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /k8sNamespaces/deleteK8sNamespaces [delete]
func (k8sNamespacesApi *K8sNamespacesApi) DeleteK8sNamespaces(c *gin.Context) {
	var k8sNamespaces business.K8sNamespaces
	_ = c.ShouldBindJSON(&k8sNamespaces)
	if err := bizSev.GetK8sNamespacesSev().Delete(k8sNamespaces); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteK8sNamespacesByIds 批量删除K8sNamespaces
// @Tags K8sNamespaces
// @Summary 批量删除K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /k8sNamespaces/deleteK8sNamespacesByIds [delete]
func (k8sNamespacesApi *K8sNamespacesApi) DeleteK8sNamespacesByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetK8sNamespacesSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateK8sNamespaces 更新K8sNamespaces
// @Tags K8sNamespaces
// @Summary 更新K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.K8sNamespaces true "更新K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /k8sNamespaces/updateK8sNamespaces [put]
func (k8sNamespacesApi *K8sNamespacesApi) UpdateK8sNamespaces(c *gin.Context) {
	var dataObj business.K8sNamespaces
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetK8sNamespacesSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindK8sNamespaces 用id查询K8sNamespaces
// @Tags K8sNamespaces
// @Summary 用id查询K8sNamespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.K8sNamespaces true "用id查询K8sNamespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /k8sNamespaces/findK8sNamespaces [get]
func (k8sNamespacesApi *K8sNamespacesApi) FindK8sNamespaces(c *gin.Context) {
	var k8sNamespaces business.K8sNamespaces
	_ = c.ShouldBindQuery(&k8sNamespaces) 
	 rek8sNamespaces,err:= bizSev.GetK8sNamespacesSev().Get(k8sNamespaces.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"k8sNamespaces": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"k8sNamespaces": rek8sNamespaces}, c)
	}
}

// GetK8sNamespacesList 分页获取K8sNamespaces列表
// @Tags K8sNamespaces
// @Summary 分页获取K8sNamespaces列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.K8sNamespacesSearch true "分页获取K8sNamespaces列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNamespaces/getK8sNamespacesList [get]
func (k8sNamespacesApi *K8sNamespacesApi) GetK8sNamespacesList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.K8sNamespacesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetK8sNamespacesSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.K8sNamespaces true "快速更新K8sNamespaces" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /k8sNamespaces/quickEdit [post] 
func (k8sNamespacesApi *K8sNamespacesApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "k8s_namespaces" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel K8sNamespaces列表
// @Tags K8sNamespaces
// @Summary 分页导出excel K8sNamespaces列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.K8sNamespacesSearch true "分页导出excel K8sNamespaces列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /k8sNamespaces/excelList [get]
func (k8sNamespacesApi *K8sNamespacesApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.K8sNamespacesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetK8sNamespacesSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "命名空间")  
					sheetFields = append(sheetFields, "状态")  
					sheetFields = append(sheetFields, "创建时间") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				arr = append(arr, v.Namespace)
				arr = append(arr, v.Status)
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


 