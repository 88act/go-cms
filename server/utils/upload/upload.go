package upload

import (
	"mime/multipart"

	"go-cms/global"
)

//@author: 10512203@qq.com
//@author: 10512203@qq.com
//@interface_name: OSS
//@description: OSS接口

type OSS interface {
	// module 模块, userType ,1管理用户 2普通用户
	UploadFile(file *multipart.FileHeader, module int, userType int) (string, string, error)
	DeleteFile(key string) error
}

//@author: 10512203@qq.com
//@author: 10512203@qq.com
//@function: NewOss
//@description: OSS接口
//@description: OSS的实例化方法
//@return: OSS

func NewOss() OSS {
	switch global.CONFIG.System.OssType {
	case "local":
		return &Local{}
	case "qiniu":
		return &Qiniu{}
	case "tencent-cos":
		return &TencentCOS{}
	case "aliyun-oss":
		return &AliyunOSS{}
	default:
		return &Local{}
	}
}
