package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"go-cms/global"
	"go-cms/model/business"
	"go-cms/model/system/request"

	bizSev "go-cms/service/business"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"go.uber.org/zap"
)

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId int64

		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				global.LOG.Error("read body from request error:", zap.Any("err", err))
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}
		if claims, ok := c.Get("claims"); ok {
			waitUse := claims.(*request.CustomClaims)
			userId = waitUse.UserId
		} else {
			userId = gconv.Int64(c.Request.Header.Get("x-user-id"))
		}
		record := business.SysLogs{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   string(body),
			UserId: userId,
		}
		if record.Path == "/base/login" || record.Path == "/memUser/memChangePwd" || record.Path == "/base/changePwdMy" || record.Path == "/base/changePwd" {
			record.Body = "--***--"
		}
		// 存在某些未知错误 TODO
		//values := c.Request.Header.Values("content-type")
		//if len(values) >0 && strings.Contains(values[0], "boundary") {
		//	record.Body = "file"
		//}
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency.Microseconds() // 微秒  0.001.001
		record.Resp = writer.body.String()
		if _, err := bizSev.GetSysLogsSev().Create(c, record); err != nil {
			global.LOG.Error("create SysLogs record error:", zap.Any("err", err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
