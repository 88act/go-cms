package upload

import (
	"mime/multipart"

	"github.com/88act/go-cms/server/global"
)

//@author: [88act-4](https://github.com/88act)
//@author: [88act-2](https://github.com/88act)
//@interface_name: OSS
//@description: OSS接口

type OSS interface {
	// module 模块, userType ,1管理用户 2普通用户
	UploadFile(file *multipart.FileHeader, module int, userType int) (string, string, error)
	DeleteFile(key string) error
}

//@author: [88act-4](https://github.com/88act)
//@author: [88act-2](https://github.com/88act)
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
