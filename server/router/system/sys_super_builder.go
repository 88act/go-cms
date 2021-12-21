package system

import (
	v1 "go-cms/api/v1"

	"github.com/gin-gonic/gin"
)

type SuperBuilderRouter struct {
}

func (s *SuperBuilderRouter) InitSuperBuilderRouter(Router *gin.RouterGroup) {
	superBuilderRouter := Router.Group("superBuilder")
	var authorityApi = v1.ApiGroupApp.SystemApiGroup.SuperBuilderApi
	{
		superBuilderRouter.POST("delSysHistory", authorityApi.DelSysHistory) // 删除回滚记录
		superBuilderRouter.POST("getMeta", authorityApi.GetMeta)             // 根据id获取meta信息
		superBuilderRouter.POST("getSysHistory", authorityApi.GetSysHistory) // 获取回滚记录分页
		superBuilderRouter.POST("rollback", authorityApi.RollBack)           // 回滚
		superBuilderRouter.POST("preview", authorityApi.PreviewTemp)         // 获取自动创建代码预览
		superBuilderRouter.POST("createTemp", authorityApi.CreateTemp)       // 创建自动化代码
		superBuilderRouter.GET("getTables", authorityApi.GetTables)          // 获取对应数据库的表
		superBuilderRouter.GET("getDB", authorityApi.GetDB)                  // 获取数据库
		superBuilderRouter.GET("getColumn", authorityApi.GetColumn)          // 获取指定表所有字段信息
	}
}
