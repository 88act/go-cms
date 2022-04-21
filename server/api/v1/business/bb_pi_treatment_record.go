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

type BbPiTreatmentRecordApi struct {
}

 

// CreateBbPiTreatmentRecord 创建BbPiTreatmentRecord
// @Tags BbPiTreatmentRecord
// @Summary 创建BbPiTreatmentRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiTreatmentRecord true "创建BbPiTreatmentRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentRecord/createBbPiTreatmentRecord [post]
func (m *BbPiTreatmentRecordApi) CreateBbPiTreatmentRecord(c *gin.Context) {
	var dataObj business.BbPiTreatmentRecord
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetBbPiTreatmentRecordSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteBbPiTreatmentRecord 删除BbPiTreatmentRecord
// @Tags BbPiTreatmentRecord
// @Summary 删除BbPiTreatmentRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiTreatmentRecord true "删除BbPiTreatmentRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiTreatmentRecord/deleteBbPiTreatmentRecord [delete]
func (m *BbPiTreatmentRecordApi) DeleteBbPiTreatmentRecord(c *gin.Context) {
	var bbPiTreatmentRecord business.BbPiTreatmentRecord
	_ = c.ShouldBindJSON(&bbPiTreatmentRecord)
	if err := bizSev.GetBbPiTreatmentRecordSev().Delete(bbPiTreatmentRecord); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBbPiTreatmentRecordByIds 批量删除BbPiTreatmentRecord
// @Tags BbPiTreatmentRecord
// @Summary 批量删除BbPiTreatmentRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiTreatmentRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bbPiTreatmentRecord/deleteBbPiTreatmentRecordByIds [delete]
func (m *BbPiTreatmentRecordApi) DeleteBbPiTreatmentRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBbPiTreatmentRecordSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBbPiTreatmentRecord 更新BbPiTreatmentRecord
// @Tags BbPiTreatmentRecord
// @Summary 更新BbPiTreatmentRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiTreatmentRecord true "更新BbPiTreatmentRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiTreatmentRecord/updateBbPiTreatmentRecord [put]
func (m *BbPiTreatmentRecordApi) UpdateBbPiTreatmentRecord(c *gin.Context) {
	var dataObj business.BbPiTreatmentRecord
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBbPiTreatmentRecordSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBbPiTreatmentRecord 用id查询BbPiTreatmentRecord
// @Tags BbPiTreatmentRecord
// @Summary 用id查询BbPiTreatmentRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BbPiTreatmentRecord true "用id查询BbPiTreatmentRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiTreatmentRecord/findBbPiTreatmentRecord [get]
func (m *BbPiTreatmentRecordApi) FindBbPiTreatmentRecord(c *gin.Context) {
	var bbPiTreatmentRecord business.BbPiTreatmentRecord
	_ = c.ShouldBindQuery(&bbPiTreatmentRecord) 
	 rebbPiTreatmentRecord,err:= bizSev.GetBbPiTreatmentRecordSev().Get(bbPiTreatmentRecord.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"bbPiTreatmentRecord": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"bbPiTreatmentRecord": rebbPiTreatmentRecord}, c)
	}
}

// GetBbPiTreatmentRecordList 分页获取BbPiTreatmentRecord列表
// @Tags BbPiTreatmentRecord
// @Summary 分页获取BbPiTreatmentRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiTreatmentRecordSearch true "分页获取BbPiTreatmentRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentRecord/getBbPiTreatmentRecordList [get]
func (m *BbPiTreatmentRecordApi) GetBbPiTreatmentRecordList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BbPiTreatmentRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetBbPiTreatmentRecordSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.BbPiTreatmentRecord true "快速更新BbPiTreatmentRecord" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /bbPiTreatmentRecord/quickEdit [post] 
func (m *BbPiTreatmentRecordApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "bb_pi_treatment_record" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel BbPiTreatmentRecord列表
// @Tags BbPiTreatmentRecord
// @Summary 分页导出excel BbPiTreatmentRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiTreatmentRecordSearch true "分页导出excel BbPiTreatmentRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiTreatmentRecord/excelList [get]
func (m *BbPiTreatmentRecordApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.BbPiTreatmentRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetBbPiTreatmentRecordSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
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
					sheetFields = append(sheetFields, "就诊类型")  
					sheetFields = append(sheetFields, "卡号")  
					sheetFields = append(sheetFields, "卡类型")  
					sheetFields = append(sheetFields, "门诊号")  
					sheetFields = append(sheetFields, "科室编码")  
					sheetFields = append(sheetFields, "科室名称")  
					sheetFields = append(sheetFields, "姓名")  
					sheetFields = append(sheetFields, "性别代码")  
					sheetFields = append(sheetFields, "出生日期")  
					sheetFields = append(sheetFields, "年龄岁")  
					sheetFields = append(sheetFields, "年龄月")  
					sheetFields = append(sheetFields, "过敏史标识")  
					sheetFields = append(sheetFields, "过敏史")  
					sheetFields = append(sheetFields, "参保类别")  
					sheetFields = append(sheetFields, "就诊日期时间")  
					sheetFields = append(sheetFields, "就诊诊断说明")  
					sheetFields = append(sheetFields, "初诊标志代码")  
					sheetFields = append(sheetFields, "主诉")  
					sheetFields = append(sheetFields, "现病史")  
					sheetFields = append(sheetFields, "既往史")  
					sheetFields = append(sheetFields, "辅助检查项目")  
					sheetFields = append(sheetFields, "辅助检查结果")  
					sheetFields = append(sheetFields, "门诊症状-诊断代码")  
					sheetFields = append(sheetFields, "门诊症状-诊断名称")  
					sheetFields = append(sheetFields, "门诊症状诊断编码类型")  
					sheetFields = append(sheetFields, "症状描述")  
					sheetFields = append(sheetFields, "辨证依据")  
					sheetFields = append(sheetFields, "治则治法")  
					sheetFields = append(sheetFields, "健康指导")  
					sheetFields = append(sheetFields, "处置计划")  
					sheetFields = append(sheetFields, "应诊医生工号")  
					sheetFields = append(sheetFields, "应诊医师签名")  
					sheetFields = append(sheetFields, "初诊医疗卫生机构代码")  
					sheetFields = append(sheetFields, "初诊机构名称")  
					sheetFields = append(sheetFields, "就诊结束时间")  
					sheetFields = append(sheetFields, "转诊标志")  
					sheetFields = append(sheetFields, "患者满意度")  
					sheetFields = append(sheetFields, "预留一")  
					sheetFields = append(sheetFields, "预留二")  
					sheetFields = append(sheetFields, "电子处方地址")  
					sheetFields = append(sheetFields, "医生端视频地址")  
					sheetFields = append(sheetFields, "患者端视频地址")  
					sheetFields = append(sheetFields, "交流音频地址")  
					sheetFields = append(sheetFields, "数据生成日期时间")  
					sheetFields = append(sheetFields, "填报日期")  
					sheetFields = append(sheetFields, "密级")  
					sheetFields = append(sheetFields, "撤销标志") 

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
				arr = append(arr, v.Jzlx)
				arr = append(arr, v.Kh)
				arr = append(arr, v.Klx)
				arr = append(arr, v.Mzh)
				arr = append(arr, v.Ksbm)
				arr = append(arr, v.Ksmc)
				arr = append(arr, v.Xm)
				arr = append(arr, v.Xbdm)
				arr = append(arr, v.Csrq)
				arr = append(arr, v.Nls)
				arr = append(arr, v.Nly)
				arr = append(arr, v.Gmsbs)
				arr = append(arr, v.Gms)
				arr = append(arr, v.Cblb)
				arr = append(arr, v.Jzrqsj)
				arr = append(arr, v.Jzzdsm)
				arr = append(arr, v.Czbzdm)
				arr = append(arr, v.Zs)
				arr = append(arr, v.Xbs)
				arr = append(arr, v.Jws)
				arr = append(arr, v.Fzjcxm)
				arr = append(arr, v.Fzjcjg)
				arr = append(arr, v.Mzzzzddm)
				arr = append(arr, v.Mzzzzdmc)
				arr = append(arr, v.Mzzzzdbmlx)
				arr = append(arr, v.Zzms)
				arr = append(arr, v.Bzyj)
				arr = append(arr, v.Zzzf)
				arr = append(arr, v.Jkzd)
				arr = append(arr, v.Czjh)
				arr = append(arr, v.Yzysgh)
				arr = append(arr, v.Yzysqm)
				arr = append(arr, v.Czylwsjgdm)
				arr = append(arr, v.Czylwsjgmc)
				arr = append(arr, v.Jzjssj)
				arr = append(arr, v.Zzbz)
				arr = append(arr, v.Hzmyd)
				arr = append(arr, v.Yl1)
				arr = append(arr, v.Yl2)
				arr = append(arr, v.Dzcfwjcfdz)
				arr = append(arr, v.Ysdspwjcfdz)
				arr = append(arr, v.Hzdspwjcfdz)
				arr = append(arr, v.Jlypcfdz)
				arr = append(arr, v.Sjscsj)
				arr = append(arr, v.Tbrqsj)
				arr = append(arr, v.Mj)
				arr = append(arr, v.Cxbz)   
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


 