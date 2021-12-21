package system

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"

	"go-cms/global"
	"go-cms/model/common/response"
	"go-cms/model/system"
	systemReq "go-cms/model/system/request"
	"go-cms/utils"

	var_dump "github.com/favframework/debug"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SuperBuilderApi struct {
}

// @Tags SuperBuilder
// @Summary 删除回滚记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.AutoHistoryByID true "删除回滚记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /superBuilder/delSysHistory [post]
func (autoApi *SuperBuilderApi) DelSysHistory(c *gin.Context) {
	var id systemReq.AutoHistoryByID
	_ = c.ShouldBindJSON(&id)
	err := superBuilderHistoryService.DeletePage(id.ID)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	}
	response.OkWithMessage("删除成功", c)

}

// @Tags SuperBuilder
// @Summary 查询回滚记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SysAutoHistory true "查询回滚记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /superBuilder/getSysHistory [post]
func (autoApi *SuperBuilderApi) GetSysHistory(c *gin.Context) {
	var search systemReq.SysAutoHistory
	_ = c.ShouldBindJSON(&search)
	err, list, total := superBuilderHistoryService.GetSysHistoryPage(search.PageInfo)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     search.Page,
			PageSize: search.PageSize,
		}, "获取成功", c)
	}
}

// @Tags SuperBuilder
// @Summary 回滚
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.AutoHistoryByID true "回滚自动生成代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"回滚成功"}"
// @Router /superBuilder/rollback [post]
func (autoApi *SuperBuilderApi) RollBack(c *gin.Context) {
	var id systemReq.AutoHistoryByID
	_ = c.ShouldBindJSON(&id)
	if err := superBuilderHistoryService.RollBack(id.ID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("回滚成功", c)
}

// @Tags SuperBuilder
// @Summary 回滚
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.AutoHistoryByID true "获取meta信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /superBuilder/getMeta [post]
func (autoApi *SuperBuilderApi) GetMeta(c *gin.Context) {
	var id systemReq.AutoHistoryByID
	_ = c.ShouldBindJSON(&id)
	if v, err := superBuilderHistoryService.GetMeta(id.ID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithDetailed(gin.H{"meta": v}, "获取成功", c)
	}

}

// @Tags SuperBuilder
// @Summary 预览创建后的代码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SuperBuilderStruct true "预览创建代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /superBuilder/preview [post]
func (autoApi *SuperBuilderApi) PreviewTemp(c *gin.Context) {
	var a system.SuperBuilderStruct
	_ = c.ShouldBindJSON(&a)

	if err := utils.Verify(a, utils.SuperBuilderVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	superBuilder, err := superBuilderService.PreviewTemp(a)
	if err != nil {
		global.LOG.Error("预览失败!", zap.Any("err", err))
		response.FailWithMessage("预览失败", c)
	} else {
		response.OkWithDetailed(gin.H{"superBuilder": superBuilder}, "预览成功", c)
	}
}

// @Tags SuperBuilder
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SuperBuilderStruct true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /superBuilder/createTemp [post]
func (autoApi *SuperBuilderApi) CreateTemp(c *gin.Context) {
	global.LOG.Debug(" CreateTemp 网址 = " + c.Request.Host)
	if strings.Contains(strings.ToLower(c.Request.Host), "go-cms.88act.com") || strings.Contains(c.Request.Host, "120.24.150.47") {
		response.FailWithMessage("测试站点禁止自动生成代码", c)
		return
	}

	var a system.SuperBuilderStruct
	_ = c.ShouldBindJSON(&a)

	var_dump.Dump(a)

	// if err := utils.Verify(a, utils.SuperBuilderVerify); err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	var apiIds []uint
	if a.AutoCreateApiToSql {
		if ids, err := superBuilderService.AutoCreateApi(&a); err != nil {
			global.LOG.Error("自动化创建失败!请自行清空垃圾数据!", zap.Any("err", err))
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("msg", url.QueryEscape("自动化创建失败!请自行清空垃圾数据!"))
			return
		} else {
			apiIds = ids
		}
	}
	//var_dump.Dump("11111111111111111")
	err := superBuilderService.CreateTemp(a, apiIds...)
	fmt.Println("代码生成,:err=", err)
	if err != nil {
		//var_dump.Dump("2222222222222222")
		//fmt.Println("bbbb2222")

		if errors.Is(err, system.AutoMoveErr) {

			var_dump.Dump("代码生成,system.AutoMoveErr")
			c.Writer.Header().Add("success", "true")
			c.Writer.Header().Add("msg", url.QueryEscape(err.Error()))
		} else {

			var_dump.Dump("代码生成, 222")
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("msg", url.QueryEscape(err.Error()))
			_ = os.Remove("./ginvueadmin.zip")
		}
	} else {

		var_dump.Dump("代码生成,3333")
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File("./ginvueadmin.zip")
		_ = os.Remove("./ginvueadmin.zip")
	}
}

// @Tags SuperBuilder
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /superBuilder/getTables [get]
func (autoApi *SuperBuilderApi) GetTables(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.CONFIG.Mysql.Dbname)
	err, tables := superBuilderService.GetTables(dbName)
	if err != nil {
		global.LOG.Error("查询table失败!", zap.Any("err", err))
		response.FailWithMessage("查询table失败", c)
	} else {
		response.OkWithDetailed(gin.H{"tables": tables}, "获取成功", c)
	}
}

// @Tags SuperBuilder
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /superBuilder/getDatabase [get]
func (autoApi *SuperBuilderApi) GetDB(c *gin.Context) {
	if err, dbs := superBuilderService.GetDB(); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"dbs": dbs}, "获取成功", c)
	}
}

// @Tags SuperBuilder
// @Summary 获取当前表所有字段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /superBuilder/getColumn [get]
func (autoApi *SuperBuilderApi) GetColumn(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.CONFIG.Mysql.Dbname)
	tableName := c.Query("tableName")
	if err, columns := superBuilderService.GetColumn(tableName, dbName); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"columns": columns}, "获取成功", c)
	}
}
