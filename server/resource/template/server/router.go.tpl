package business

import (
	"go-cms/api/v1"
	"go-cms/middleware"
	"github.com/gin-gonic/gin"
)

type {{.StructName}}Router struct {
}

// Init{{.StructName}}Router 初始化 {{.StructName}} 路由信息
func (s *{{.StructName}}Router) Init{{.StructName}}Router(Router *gin.RouterGroup) {
	router := Router.Group("{{.Abbreviation}}").Use(middleware.OperationRecord())
	routerNoRecord := Router.Group("{{.Abbreviation}}")
	var apiV1 = v1.ApiGroupApp.BusinessApiGroup.{{.StructName}}Api
	{
		router.POST("create{{.StructName}}", apiV1.Create{{.StructName}})   // 新建{{.StructName}}
		router.POST("delete{{.StructName}}", apiV1.Delete{{.StructName}}) // 删除{{.StructName}}
		router.POST("delete{{.StructName}}ByIds", apiV1.Delete{{.StructName}}ByIds) // 批量删除{{.StructName}}
		router.POST("update{{.StructName}}", apiV1.Update{{.StructName}})    // 更新{{.StructName}}
	    router.POST("quickEdit", apiV1.QuickEdit)  // 快速编辑	
	}
	{
		routerNoRecord.GET("find{{.StructName}}", apiV1.Find{{.StructName}})        // 根据ID获取{{.StructName}}
		routerNoRecord.GET("get{{.StructName}}List", apiV1.Get{{.StructName}}List)  // 获取{{.StructName}}列表
        routerNoRecord.GET("excelList", apiV1.ExcelList)  // 分页导出excel {{.StructName}}列表
	}
}
