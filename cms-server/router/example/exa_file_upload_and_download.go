package example

import (
	v1 "github.com/88act/go-cms/server/api/v1"
	"github.com/gin-gonic/gin"
)

type FileUploadAndDownloadRouter struct {
}

func (e *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	fileUploadAndDownloadRouter := Router.Group("fileUploadAndDownload")
	var exaFileUploadAndDownloadApi = v1.ApiGroupApp.ExampleApiGroup.FileUploadAndDownloadApi
	{
		fileUploadAndDownloadRouter.POST("upload", exaFileUploadAndDownloadApi.UploadFile)                                 // 上传文件
		fileUploadAndDownloadRouter.POST("getFileList", exaFileUploadAndDownloadApi.GetFileList)                           // 获取上传文件列表
		fileUploadAndDownloadRouter.POST("deleteFile", exaFileUploadAndDownloadApi.DeleteFile)                             // 删除指定文件
		fileUploadAndDownloadRouter.POST("breakpointContinue", exaFileUploadAndDownloadApi.BreakpointContinue)             // 断点续传
		fileUploadAndDownloadRouter.GET("findFile", exaFileUploadAndDownloadApi.FindFile)                                  // 查询当前文件成功的切片
		fileUploadAndDownloadRouter.POST("breakpointContinueFinish", exaFileUploadAndDownloadApi.BreakpointContinueFinish) // 查询当前文件成功的切片
		fileUploadAndDownloadRouter.POST("removeChunk", exaFileUploadAndDownloadApi.RemoveChunk)                           // 查询当前文件成功的切片
	}
}
