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

type BbPiStaffApi struct {
}

 

// CreateBbPiStaff 创建BbPiStaff
// @Tags BbPiStaff
// @Summary 创建BbPiStaff
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiStaff true "创建BbPiStaff"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiStaff/createBbPiStaff [post]
func (bbPiStaffApi *BbPiStaffApi) CreateBbPiStaff(c *gin.Context) {
	var dataObj business.BbPiStaff
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetBbPiStaffSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteBbPiStaff 删除BbPiStaff
// @Tags BbPiStaff
// @Summary 删除BbPiStaff
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiStaff true "删除BbPiStaff"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiStaff/deleteBbPiStaff [delete]
func (bbPiStaffApi *BbPiStaffApi) DeleteBbPiStaff(c *gin.Context) {
	var bbPiStaff business.BbPiStaff
	_ = c.ShouldBindJSON(&bbPiStaff)
	if err := bizSev.GetBbPiStaffSev().Delete(bbPiStaff); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBbPiStaffByIds 批量删除BbPiStaff
// @Tags BbPiStaff
// @Summary 批量删除BbPiStaff
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiStaff"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bbPiStaff/deleteBbPiStaffByIds [delete]
func (bbPiStaffApi *BbPiStaffApi) DeleteBbPiStaffByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBbPiStaffSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBbPiStaff 更新BbPiStaff
// @Tags BbPiStaff
// @Summary 更新BbPiStaff
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiStaff true "更新BbPiStaff"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiStaff/updateBbPiStaff [put]
func (bbPiStaffApi *BbPiStaffApi) UpdateBbPiStaff(c *gin.Context) {
	var dataObj business.BbPiStaff
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBbPiStaffSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBbPiStaff 用id查询BbPiStaff
// @Tags BbPiStaff
// @Summary 用id查询BbPiStaff
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BbPiStaff true "用id查询BbPiStaff"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiStaff/findBbPiStaff [get]
func (bbPiStaffApi *BbPiStaffApi) FindBbPiStaff(c *gin.Context) {
	var bbPiStaff business.BbPiStaff
	_ = c.ShouldBindQuery(&bbPiStaff) 
	 rebbPiStaff,err:= bizSev.GetBbPiStaffSev().Get(bbPiStaff.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"bbPiStaff": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"bbPiStaff": rebbPiStaff}, c)
	}
}

// GetBbPiStaffList 分页获取BbPiStaff列表
// @Tags BbPiStaff
// @Summary 分页获取BbPiStaff列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiStaffSearch true "分页获取BbPiStaff列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiStaff/getBbPiStaffList [get]
func (bbPiStaffApi *BbPiStaffApi) GetBbPiStaffList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BbPiStaffSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetBbPiStaffSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.BbPiStaff true "快速更新BbPiStaff" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /bbPiStaff/quickEdit [post] 
func (bbPiStaffApi *BbPiStaffApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "bb_pi_staff" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel BbPiStaff列表
// @Tags BbPiStaff
// @Summary 分页导出excel BbPiStaff列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiStaffSearch true "分页导出excel BbPiStaff列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiStaff/excelList [get]
func (bbPiStaffApi *BbPiStaffApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.BbPiStaffSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetBbPiStaffSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "上传")  
					sheetFields = append(sheetFields, "机构标识")  
					sheetFields = append(sheetFields, "医护人员账号")  
					sheetFields = append(sheetFields, "医护人员姓名")  
					sheetFields = append(sheetFields, "性别")  
					sheetFields = append(sheetFields, "出生日期")  
					sheetFields = append(sheetFields, "身份证号")  
					sheetFields = append(sheetFields, "身份证件类别代码")  
					sheetFields = append(sheetFields, "所属科室代码")  
					sheetFields = append(sheetFields, "专业技术职务代码")  
					sheetFields = append(sheetFields, "专业技术职务类别代码")  
					sheetFields = append(sheetFields, "资质类别名称")  
					sheetFields = append(sheetFields, "资格证书编号")  
					sheetFields = append(sheetFields, "资格获得时间")  
					sheetFields = append(sheetFields, "执业证书编码")  
					sheetFields = append(sheetFields, "发证日期")  
					sheetFields = append(sheetFields, "执业地点")  
					sheetFields = append(sheetFields, "执业范围")  
					sheetFields = append(sheetFields, "主要执业医疗机构代码")  
					sheetFields = append(sheetFields, "主要执业医院名称")  
					sheetFields = append(sheetFields, "全科医生标志")  
					sheetFields = append(sheetFields, "手机号码")  
					sheetFields = append(sheetFields, "专长")  
					sheetFields = append(sheetFields, "民族")  
					sheetFields = append(sheetFields, "参加工作日期")  
					sheetFields = append(sheetFields, "注册日期时间")  
					sheetFields = append(sheetFields, "学历")  
					sheetFields = append(sheetFields, "个人照片存放地址")  
					sheetFields = append(sheetFields, "资格证存放地址")  
					sheetFields = append(sheetFields, "执业证存放地址")  
					sheetFields = append(sheetFields, "数据生成日期时间")  
					sheetFields = append(sheetFields, "填报日期")  
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
				arr = append(arr, v.Yhrygh)
				arr = append(arr, v.Yhryxm)
				arr = append(arr, v.Xb)
				arr = append(arr, v.Csrq)
				arr = append(arr, v.Sfzh)
				arr = append(arr, v.Zjlbdm)
				arr = append(arr, v.Ksdm)
				arr = append(arr, v.Zyjszwdm)
				arr = append(arr, v.Zyjszwlbdm)
				arr = append(arr, v.Zzlbmc)
				arr = append(arr, v.Zgzsbh)
				arr = append(arr, v.Zghdsj)
				arr = append(arr, v.Zyzsbm)
				arr = append(arr, v.Fzrq)
				arr = append(arr, v.Zydd)
				arr = append(arr, v.Zyfw)
				arr = append(arr, v.Zyzyyljgdm)
				arr = append(arr, v.Zyzyyymc)
				arr = append(arr, v.Qkysbz)
				arr = append(arr, v.Sjhm)
				arr = append(arr, v.Zc)
				arr = append(arr, v.Mz)
				arr = append(arr, v.Cjgzrq)
				arr = append(arr, v.Zcsj)
				arr = append(arr, v.Xl)
				arr = append(arr, v.Grzpcfdz)
				arr = append(arr, v.Zgzcfdz)
				arr = append(arr, v.Zyzcfdz)
				arr = append(arr, v.Sjscsj)
				arr = append(arr, v.Tbrqsj)
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


 