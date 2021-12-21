package system

import (
	"go-cms/global"
	"go-cms/model/common/response"
	"go-cms/model/system/request"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type DBApi struct {
}

// @Tags InitDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Param data body request.InitDB true "初始化数据库参数"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"自动创建数据库成功"}"
// @Router /init/initdb [post]
func (i *DBApi) InitDB(c *gin.Context) {
	if global.DB != nil {
		global.LOG.Error("已存在数据库配置!")
		response.FailWithMessage("已存在数据库配置", c)
		return
	}
	var dbInfo request.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.LOG.Error("参数校验不通过!", zap.Any("err", err))
		response.FailWithMessage("参数校验不通过", c)
		return
	}
	if err := initDBService.InitDB(dbInfo); err != nil {
		global.LOG.Error("自动创建数据库失败!", zap.Any("err", err))
		response.FailWithMessage("自动创建数据库失败，请查看后台日志，检查后在进行初始化", c)
		return
	}
	response.OkWithData("自动创建数据库成功", c)
}

// @Tags CheckDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"探测完成"}"
// @Router /init/checkdb [post]
func (i *DBApi) CheckDB(c *gin.Context) {
	if global.DB != nil {
		global.LOG.Info("数据库无需初始化")
		response.OkWithDetailed(gin.H{"needInit": false}, "数据库无需初始化", c)
		return
	} else {
		global.LOG.Info("前往初始化数据库")
		response.OkWithDetailed(gin.H{"needInit": true}, "前往初始化数据库", c)
		return
	}
}
