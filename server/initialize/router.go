package initialize

import (
	"context"
	"net/http"
	"time"

	_ "go-cms/docs"
	"go-cms/global"
	"go-cms/middleware"
	"go-cms/router"
	"go-cms/service/common"

	"github.com/gin-gonic/gin"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
)

// timeout middleware wraps the request context with a timeout
func timeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		// wrap the request context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			// check if context timeout was reached
			if ctx.Err() == context.DeadlineExceeded {

				// write response and abort the request
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}

			//cancel to clear resources after finished
			cancel()
		}()

		// replace request with context wrapped request
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func timedHandler(duration time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		// get the underlying request context
		ctx := c.Request.Context()

		// create the response data type to use as a channel type
		type responseData struct {
			status int
			body   map[string]interface{}
		}

		// create a done channel to tell the request it's done
		doneChan := make(chan responseData)
		// here you put the actual work needed for the request
		// and then send the doneChan with the status and body
		// to finish the request by writing the response
		go func() {
			time.Sleep(duration)
			doneChan <- responseData{
				status: 200,
				body:   gin.H{"hello": "world"},
			}
		}()
		// non-blocking select on two channels see if the request
		// times out or finishes
		select {
		// if the context is done it timed out or was cancelled
		// so don't return anything
		case <-ctx.Done():
			return

			// if the request finished then finish the request by
			// writing the response
		case res := <-doneChan:
			c.JSON(res.status, res.body)
		}
	}
}

// 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(timeoutMiddleware(time.Second * 4800)) // 40分钟

	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	//Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	//Router.Static("/favicon.ico", "./dist/favicon.ico")
	//Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	//Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面
	//Router.StaticFS(global.REDIS.Local.Path, http.Dir(global.REDIS.Local.Path)) // 为用户头像和文件提供静态地址

	Router.StaticFS("res", http.Dir("res")) //  提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.LOG.Info("use middleware cors")
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//global.LOG.Info("register swagger handler")

	//add by ljd 20210927
	//为没有配置处理函数的路由添加处理程序，默认情况下它返回404代码
	Router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, common.Response{
			400,
			nil,
			"没有路由配置",
		})
		c.Abort()
		return
	})
	//加载模板
	//Router.LoadHTMLGlob("template/def/*")
	// 方便统一添加路由组前缀 多服务器上线使用
	//获取路由组实例
	systemRouter := router.RouterGroupApp.System
	businessRouter := router.RouterGroupApp.Business
	commonRouter := router.RouterGroupApp.Common
	clientRouter := router.RouterGroupApp.Client
	PublicGroup := Router.Group("")
	{
		// 健康监测 http://localhost:8080/api/health
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok health")
		})
		// 测试   http://localhost:8080/api/test
		// 注意，vue用了代理api前缀    http://127.0.0.1:8888/health
		// add by ljd 20210927
		PublicGroup.GET("/test", func(c *gin.Context) {
			c.JSON(200, "test")
		})
		// PublicGroup.GET("/", func(c *gin.Context) {
		// 	c.JSON(200, "go-cms server")
		// })
		clientRouter.InitCmsRouter(PublicGroup) // 客户端网页 首页
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		//businessRouter.InitReportApiRouter(PublicGroup) //报表
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitBasePriRouter(PrivateGroup) // 注册基础功能路由 有鉴权
		systemRouter.InitMenuRouter(PrivateGroup)    // 注册menu路由
		commonRouter.InitCommonDbRouter(PrivateGroup)
		commonRouter.InitCommonDbRouterPub(PublicGroup)
		commonRouter.InitCommonFileRouter(PublicGroup)
		commonRouter.InitCommonFileRouterPrv(PrivateGroup)
		// golang_coder Begin; DO NOT EDIT.
		businessRouter.InitSysMenuRouter(PrivateGroup)
		businessRouter.InitSysRoleRouter(PrivateGroup)
		businessRouter.InitSysLogsRouter(PrivateGroup)
		businessRouter.InitSysApisRouter(PrivateGroup)
		businessRouter.InitMemUserRouter(PrivateGroup)
		businessRouter.InitMemDepartRouter(PrivateGroup)
		businessRouter.InitMemLogsRouter(PrivateGroup)
		businessRouter.InitMemUserSafeRouter(PrivateGroup)
		businessRouter.InitBasicFileCatRouter(PrivateGroup)
		businessRouter.InitBasicDictRouter(PrivateGroup)
		businessRouter.InitBasicFileRouter(PrivateGroup)
		businessRouter.InitCmsDetailRouter(PrivateGroup)
		businessRouter.InitCmsVisitorRouter(PrivateGroup)
		businessRouter.InitCmsDiscussRouter(PrivateGroup)
		businessRouter.InitCmsCatArtRouter(PrivateGroup)
		businessRouter.InitCmsCatRouter(PrivateGroup)
		businessRouter.InitCmsBlogRouter(PrivateGroup)
		businessRouter.InitCmsGroupRouter(PrivateGroup)
		businessRouter.InitCmsTagRouter(PrivateGroup)
		businessRouter.InitCmsArtRouter(PrivateGroup)
		// golang_coder End; DO NOT EDIT.

	}
	global.LOG.Info("router register success")
	return Router
}
