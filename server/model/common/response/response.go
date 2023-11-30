package response

import (
	"go-cms/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	//ERROR   = 7
	//SUCCESS = 0
	ERROR   = 400
	SUCCESS = 200
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	if code == ERROR {
		global.LOG.Error("api返回错误response ,", zap.Any("msg", msg))
	}
	// 开始时间
	c.AbortWithStatusJSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
	// 这里直接结束 函数 add by ljd 20211105  c.Abort();  c.Next();
	c.Abort()

}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
	c.Abort()
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
	c.Abort()
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
	c.Abort()
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
	c.Abort()
}

func FailWithMessage(message string, c *gin.Context) {

	Result(ERROR, map[string]interface{}{}, message, c)
	c.Abort()
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {

	Result(ERROR, data, message, c)
	c.Abort()
}
