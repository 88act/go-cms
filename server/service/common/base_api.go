package common

import (
	"go-cms/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseApi struct {
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// // 返回id
// type IdResp struct {
// 	Id int64 `json:"id" form:"id"`
// }

const (
	//ERROR   = 7
	//SUCCESS = 0
	ERROR   = 400
	SUCCESS = 200
)

func (m *BaseApi) Result(code int, data interface{}, msg string, c *gin.Context) {
	if code == ERROR {
		global.LOG.Error("api返回错误BaseApi ,", zap.Any("msg", msg))
	}
	c.AbortWithStatusJSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})

	// 这里直接结束 函数 add by ljd 20211105
	//https://zhujianghan.github.io/zjh-blog/posts/backend/gin-middleware-%E6%B3%A8%E6%84%8F%E4%BA%8B%E9%A1%B9/
	c.Abort()
	return

}

func (m *BaseApi) Ok(c *gin.Context) {
	m.Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func (m *BaseApi) OkWithMessage(message string, c *gin.Context) {
	m.Result(SUCCESS, map[string]interface{}{}, message, c)
}

func (m *BaseApi) OkWithData(data interface{}, c *gin.Context) {
	m.Result(SUCCESS, data, "操作成功", c)
}

func (m *BaseApi) OkWithDetailed(data interface{}, message string, c *gin.Context) {
	m.Result(SUCCESS, data, message, c)
}

func (m *BaseApi) Fail(c *gin.Context) {
	m.Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func (m *BaseApi) FailWithMessage(message string, c *gin.Context) {
	m.Result(ERROR, map[string]interface{}{}, message, c)
}

func (m *BaseApi) FailWithDetailed(data interface{}, message string, c *gin.Context) {
	m.Result(ERROR, data, message, c)
}
