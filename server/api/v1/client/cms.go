package client

import (
	commSev "go-cms/service/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CmsApi struct {
	commSev.BaseApi
}

// 首页
func (m *CmsApi) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试"})
}

// 首页
func (m *CmsApi) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "我是Login测试"})
}
