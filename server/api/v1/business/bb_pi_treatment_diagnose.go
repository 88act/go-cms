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

type BbPiTreatmentDiagnoseApi struct {
}

 

// CreateBbPiTreatmentDiagnose 创建BbPiTreatmentDiagnose
// @Tags BbPiTreatmentDiagnose
// @Summary 创建BbPiTreatmentDiagnose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiTreatmentDiagnose true "创建BbPiTreatmentDiagnose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentDiagnose/createBbPiTreatmentDiagnose [post]
func (m *BbPiTreatmentDiagnoseApi) CreateBbPiTreatmentDiagnose(c *gin.Context) {
	var dataObj business.BbPiTreatmentDiagnose
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetBbPiTreatmentDiagnoseSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteBbPiTreatmentDiagnose 删除BbPiTreatmentDiagnose
// @Tags BbPiTreatmentDiagnose
// @Summary 删除BbPiTreatmentDiagnose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiTreatmentDiagnose true "删除BbPiTreatmentDiagnose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiTreatmentDiagnose/deleteBbPiTreatmentDiagnose [delete]
func (m *BbPiTreatmentDiagnoseApi) DeleteBbPiTreatmentDiagnose(c *gin.Context) {
	var bbPiTreatmentDiagnose business.BbPiTreatmentDiagnose
	_ = c.ShouldBindJSON(&bbPiTreatmentDiagnose)
	if err := bizSev.GetBbPiTreatmentDiagnoseSev().Delete(bbPiTreatmentDiagnose); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBbPiTreatmentDiagnoseByIds 批量删除BbPiTreatmentDiagnose
// @Tags BbPiTreatmentDiagnose
// @Summary 批量删除BbPiTreatmentDiagnose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiTreatmentDiagnose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bbPiTreatmentDiagnose/deleteBbPiTreatmentDiagnoseByIds [delete]
func (m *BbPiTreatmentDiagnoseApi) DeleteBbPiTreatmentDiagnoseByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBbPiTreatmentDiagnoseSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBbPiTreatmentDiagnose 更新BbPiTreatmentDiagnose
// @Tags BbPiTreatmentDiagnose
// @Summary 更新BbPiTreatmentDiagnose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiTreatmentDiagnose true "更新BbPiTreatmentDiagnose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiTreatmentDiagnose/updateBbPiTreatmentDiagnose [put]
func (m *BbPiTreatmentDiagnoseApi) UpdateBbPiTreatmentDiagnose(c *gin.Context) {
	var dataObj business.BbPiTreatmentDiagnose
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBbPiTreatmentDiagnoseSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBbPiTreatmentDiagnose 用id查询BbPiTreatmentDiagnose
// @Tags BbPiTreatmentDiagnose
// @Summary 用id查询BbPiTreatmentDiagnose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BbPiTreatmentDiagnose true "用id查询BbPiTreatmentDiagnose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiTreatmentDiagnose/findBbPiTreatmentDiagnose [get]
func (m *BbPiTreatmentDiagnoseApi) FindBbPiTreatmentDiagnose(c *gin.Context) {
	var bbPiTreatmentDiagnose business.BbPiTreatmentDiagnose
	_ = c.ShouldBindQuery(&bbPiTreatmentDiagnose) 
	 rebbPiTreatmentDiagnose,err:= bizSev.GetBbPiTreatmentDiagnoseSev().Get(bbPiTreatmentDiagnose.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"bbPiTreatmentDiagnose": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"bbPiTreatmentDiagnose": rebbPiTreatmentDiagnose}, c)
	}
}

// GetBbPiTreatmentDiagnoseList 分页获取BbPiTreatmentDiagnose列表
// @Tags BbPiTreatmentDiagnose
// @Summary 分页获取BbPiTreatmentDiagnose列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiTreatmentDiagnoseSearch true "分页获取BbPiTreatmentDiagnose列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentDiagnose/getBbPiTreatmentDiagnoseList [get]
func (m *BbPiTreatmentDiagnoseApi) GetBbPiTreatmentDiagnoseList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BbPiTreatmentDiagnoseSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetBbPiTreatmentDiagnoseSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.BbPiTreatmentDiagnose true "快速更新BbPiTreatmentDiagnose" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /bbPiTreatmentDiagnose/quickEdit [post] 
func (m *BbPiTreatmentDiagnoseApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "bb_pi_treatment_diagnose" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel BbPiTreatmentDiagnose列表
// @Tags BbPiTreatmentDiagnose
// @Summary 分页导出excel BbPiTreatmentDiagnose列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiTreatmentDiagnoseSearch true "分页导出excel BbPiTreatmentDiagnose列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentDiagnose/excelList [get]
func (m *BbPiTreatmentDiagnoseApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.BbPiTreatmentDiagnoseSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetBbPiTreatmentDiagnoseSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "上传")  
					sheetFields = append(sheetFields, "机构标识")  
					sheetFields = append(sheetFields, "服务网点代码")  
					sheetFields = append(sheetFields, "处方编号")  
					sheetFields = append(sheetFields, "处方明细ID")  
					sheetFields = append(sheetFields, "诊断信息 ID")  
					sheetFields = append(sheetFields, "卡号")  
					sheetFields = append(sheetFields, "卡类型")  
					sheetFields = append(sheetFields, "门诊号")  
					sheetFields = append(sheetFields, "姓名")  
					sheetFields = append(sheetFields, "性别")  
					sheetFields = append(sheetFields, "出生日期")  
					sheetFields = append(sheetFields, "年龄岁")  
					sheetFields = append(sheetFields, "年龄月")  
					sheetFields = append(sheetFields, "就诊时间")  
					sheetFields = append(sheetFields, "诊断类型编码")  
					sheetFields = append(sheetFields, "西医诊断编码")  
					sheetFields = append(sheetFields, "西医诊断名称")  
					sheetFields = append(sheetFields, "西医诊断编码类型")  
					sheetFields = append(sheetFields, "中医诊断编码")  
					sheetFields = append(sheetFields, "中医诊断名称")  
					sheetFields = append(sheetFields, "中医诊断编码类型")  
					sheetFields = append(sheetFields, "诊断说明")  
					sheetFields = append(sheetFields, "诊断标志")  
					sheetFields = append(sheetFields, "诊断医生工号")  
					sheetFields = append(sheetFields, "诊断医生姓名")  
					sheetFields = append(sheetFields, "诊断时间")  
					sheetFields = append(sheetFields, "预留一")  
					sheetFields = append(sheetFields, "数据生成时间")  
					sheetFields = append(sheetFields, "填报日期")  
					sheetFields = append(sheetFields, "撤销标志")  
					sheetFields = append(sheetFields, "密级") 

			excel := excelize.NewFile()
			excel.SetSheetRow("Sheet1", "A1", &sheetFields)
			for i, v := range list {
				axis := fmt.Sprintf("A%d", i+2)
				var arr = []interface{}{}
				if v.Status == nil {
					arr = append(arr, "")
				} else {
					arr = append(arr, *v.Status)
				}
				arr = append(arr, v.Jgdm)
				arr = append(arr, v.Fwwddm)
				arr = append(arr, v.Cfbh)
				arr = append(arr, v.Cfmxid)
				arr = append(arr, v.Zdxxid)
				arr = append(arr, v.Kh)
				arr = append(arr, v.Klx)
				arr = append(arr, v.Mzh)
				arr = append(arr, v.Xm)
				arr = append(arr, v.Xbdm)
				arr = append(arr, v.Csrq)
				arr = append(arr, v.Nls)
				arr = append(arr, v.Nly)
				arr = append(arr, v.Jzrqsj)
				arr = append(arr, v.Zdlxbm)
				arr = append(arr, v.Xyzdbm)
				arr = append(arr, v.Xyzdmc)
				arr = append(arr, v.Xyzdbmlx)
				arr = append(arr, v.Zyzdbm)
				arr = append(arr, v.Zyzdmc)
				arr = append(arr, v.Zyzdbmlx)
				arr = append(arr, v.Zdsm)
				arr = append(arr, v.Zdbz)
				arr = append(arr, v.Zdysgh)
				arr = append(arr, v.Zdysxm)
				arr = append(arr, v.Zdsj)
				arr = append(arr, v.Yl1)
				arr = append(arr, v.Sjscsj)
				arr = append(arr, v.Tbrqsj)
				arr = append(arr, v.Cxbz)
				arr = append(arr, v.Mj)   
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


 