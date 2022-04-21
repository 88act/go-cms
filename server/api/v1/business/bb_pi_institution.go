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

type BbPiInstitutionApi struct {
}

 

// CreateBbPiInstitution 创建BbPiInstitution
// @Tags BbPiInstitution
// @Summary 创建BbPiInstitution
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiInstitution true "创建BbPiInstitution"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiInstitution/createBbPiInstitution [post]
func (bbPiInstitutionApi *BbPiInstitutionApi) CreateBbPiInstitution(c *gin.Context) {
	var dataObj business.BbPiInstitution
	_ = c.ShouldBindJSON(&dataObj)
	
	if err := gvalid.CheckStruct(c,dataObj, nil); err != nil {
		response.FailWithMessage("创建失败,"+err.FirstString(), c)
		return
	}

 
	if id,err := bizSev.GetBbPiInstitutionSev().Create(dataObj); err != nil {
        global.LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
	    idResp := &response.IdResp{Id: id}
		response.OkWithData(idResp, c)
	}
}

// DeleteBbPiInstitution 删除BbPiInstitution
// @Tags BbPiInstitution
// @Summary 删除BbPiInstitution
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiInstitution true "删除BbPiInstitution"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bbPiInstitution/deleteBbPiInstitution [delete]
func (bbPiInstitutionApi *BbPiInstitutionApi) DeleteBbPiInstitution(c *gin.Context) {
	var bbPiInstitution business.BbPiInstitution
	_ = c.ShouldBindJSON(&bbPiInstitution)
	if err := bizSev.GetBbPiInstitutionSev().Delete(bbPiInstitution); err != nil {
        global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBbPiInstitutionByIds 批量删除BbPiInstitution
// @Tags BbPiInstitution
// @Summary 批量删除BbPiInstitution
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BbPiInstitution"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bbPiInstitution/deleteBbPiInstitutionByIds [delete]
func (bbPiInstitutionApi *BbPiInstitutionApi) DeleteBbPiInstitutionByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := bizSev.GetBbPiInstitutionSev().DeleteByIds(IDS); err != nil {
        global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBbPiInstitution 更新BbPiInstitution
// @Tags BbPiInstitution
// @Summary 更新BbPiInstitution
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.BbPiInstitution true "更新BbPiInstitution"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bbPiInstitution/updateBbPiInstitution [put]
func (bbPiInstitutionApi *BbPiInstitutionApi) UpdateBbPiInstitution(c *gin.Context) {
	var dataObj business.BbPiInstitution
	_ = c.ShouldBindJSON(&dataObj)

	if err := gvalid.CheckStruct(c, dataObj, nil); err != nil {
		response.FailWithMessage("更新失败,"+err.FirstString(), c)
		return
	}

	if err := bizSev.GetBbPiInstitutionSev().Update(dataObj); err != nil {
        global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBbPiInstitution 用id查询BbPiInstitution
// @Tags BbPiInstitution
// @Summary 用id查询BbPiInstitution
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.BbPiInstitution true "用id查询BbPiInstitution"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bbPiInstitution/findBbPiInstitution [get]
func (bbPiInstitutionApi *BbPiInstitutionApi) FindBbPiInstitution(c *gin.Context) {
	var bbPiInstitution business.BbPiInstitution
	_ = c.ShouldBindQuery(&bbPiInstitution) 
	 rebbPiInstitution,err:= bizSev.GetBbPiInstitutionSev().Get(bbPiInstitution.ID,""); 
	 if errors.Is(err, gorm.ErrRecordNotFound) { 
		response.OkWithData(gin.H{"bbPiInstitution": nil}, c)
	} else if err != nil { 
        global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else { 
		response.OkWithData(gin.H{"bbPiInstitution": rebbPiInstitution}, c)
	}
}

// GetBbPiInstitutionList 分页获取BbPiInstitution列表
// @Tags BbPiInstitution
// @Summary 分页获取BbPiInstitution列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiInstitutionSearch true "分页获取BbPiInstitution列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiInstitution/getBbPiInstitutionList [get]
func (bbPiInstitutionApi *BbPiInstitutionApi) GetBbPiInstitutionList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")

	var pageInfo bizReq.BbPiInstitutionSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if  list, total, err := bizSev.GetBbPiInstitutionSev().GetList(pageInfo,createdAtBetween,""); err != nil {
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
// @Param data body business.BbPiInstitution true "快速更新BbPiInstitution" 
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router  /bbPiInstitution/quickEdit [post] 
func (bbPiInstitutionApi *BbPiInstitutionApi) QuickEdit(c *gin.Context) {
	var quickEdit request.QuickEdit
	_ = c.ShouldBindJSON(&quickEdit)
	quickEdit.Table = "bb_pi_institution" 
	if err := commSev.GetCommonDbSev().QuickEdit(quickEdit); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// excelList 分页导出excel BbPiInstitution列表
// @Tags BbPiInstitution
// @Summary 分页导出excel BbPiInstitution列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.BbPiInstitutionSearch true "分页导出excel BbPiInstitution列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bbPiInstitution/excelList [get]
func (bbPiInstitutionApi *BbPiInstitutionApi) ExcelList(c *gin.Context) {
	createdAtBetween, _ := c.GetQueryArray("createdAtBetween[]")
	var pageInfo bizReq.BbPiInstitutionSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list,_,err:= bizSev.GetBbPiInstitutionSev().GetListAll(pageInfo,createdAtBetween,""); err != nil {
	    global.LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        if len(list) == 0 {
			response.FailWithMessage("没有数据", c)
		} else { 
			sheetFields := []string{}  
					sheetFields = append(sheetFields, "上传")  
					sheetFields = append(sheetFields, "机构标识")  
					sheetFields = append(sheetFields, "统一信用代码")  
					sheetFields = append(sheetFields, "机构名称")  
					sheetFields = append(sheetFields, "行政区划代码")  
					sheetFields = append(sheetFields, "机构类型")  
					sheetFields = append(sheetFields, "机构成立日期")  
					sheetFields = append(sheetFields, "单位隶属关系代码")  
					sheetFields = append(sheetFields, "机构分类代码")  
					sheetFields = append(sheetFields, "机构分类代码")  
					sheetFields = append(sheetFields, "经济类型代码")  
					sheetFields = append(sheetFields, "地址")  
					sheetFields = append(sheetFields, "组织机构代码")  
					sheetFields = append(sheetFields, "医院名称")  
					sheetFields = append(sheetFields, "医疗机构级别")  
					sheetFields = append(sheetFields, "医院机构等级")  
					sheetFields = append(sheetFields, "医院网址")  
					sheetFields = append(sheetFields, "许可证号码")  
					sheetFields = append(sheetFields, "许可项目名称")  
					sheetFields = append(sheetFields, "许可证有效期")  
					sheetFields = append(sheetFields, "信息安等级保护")  
					sheetFields = append(sheetFields, "信息安等保编号")  
					sheetFields = append(sheetFields, "开办资金金额数")  
					sheetFields = append(sheetFields, "法人代表")  
					sheetFields = append(sheetFields, "机构民族自治标志")  
					sheetFields = append(sheetFields, "是否分支机构")  
					sheetFields = append(sheetFields, "机构第二名称")  
					sheetFields = append(sheetFields, "机构描述")  
					sheetFields = append(sheetFields, "邮政编码")  
					sheetFields = append(sheetFields, "电话号码")  
					sheetFields = append(sheetFields, "电子信箱")  
					sheetFields = append(sheetFields, "第三方信用代码")  
					sheetFields = append(sheetFields, "第三方机构名称")  
					sheetFields = append(sheetFields, "协议有效期")  
					sheetFields = append(sheetFields, "部署平台")  
					sheetFields = append(sheetFields, "架构描述")  
					sheetFields = append(sheetFields, "机房面积")  
					sheetFields = append(sheetFields, "是否双路发电")  
					sheetFields = append(sheetFields, "备注")  
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
				arr = append(arr, v.Zzjgdm)
				arr = append(arr, v.Jgmc)
				arr = append(arr, v.Xzqhdm)
				arr = append(arr, v.Jglx)
				arr = append(arr, v.Jgclrq)
				arr = append(arr, v.Dwlsgxdm)
				arr = append(arr, v.Jgflgllbdm)
				arr = append(arr, v.Jgfldm)
				arr = append(arr, v.Jjlxdm)
				arr = append(arr, v.Dz)
				arr = append(arr, v.Styyzzjgdm)
				arr = append(arr, v.Styymc)
				arr = append(arr, v.Styljgjb)
				arr = append(arr, v.Styljgdj)
				arr = append(arr, v.Hlwyywz)
				arr = append(arr, v.Xkzhm)
				arr = append(arr, v.Xkxmmc)
				arr = append(arr, v.Xkzyxq)
				arr = append(arr, v.Xxaqdjbh)
				arr = append(arr, v.Xxaqdjbhbh)
				arr = append(arr, v.Kbzjjes)
				arr = append(arr, v.Frdb)
				arr = append(arr, v.Jgszdmzzzdfbz)
				arr = append(arr, v.Sffzjg)
				arr = append(arr, v.Jgdemc)
				arr = append(arr, v.Jgms)
				arr = append(arr, v.Yzbm)
				arr = append(arr, v.Dhhm)
				arr = append(arr, v.Dwdzyx)
				arr = append(arr, v.Dsfjgdm)
				arr = append(arr, v.Dsfjgmc)
				arr = append(arr, v.Xyyxq)
				arr = append(arr, v.Bspt)
				arr = append(arr, v.Jgmsxx)
				arr = append(arr, v.Jfmj)
				arr = append(arr, v.Sfslgd)
				arr = append(arr, v.Bz)
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


 