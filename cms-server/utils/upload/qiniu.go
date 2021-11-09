package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/88act/go-cms/server/global"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"go.uber.org/zap"
)

type Qiniu struct{}

//@author: [88act](https://github.com/88act)
//@author: [88act-4](https://github.com/88act)
//@author: [88act-2](https://github.com/88act)
//@object: *Qiniu
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*Qiniu) UploadFile(file *multipart.FileHeader, module int, userType int) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: global.CONFIG.Qiniu.Bucket}
	mac := qbox.NewMac(global.CONFIG.Qiniu.AccessKey, global.CONFIG.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiniuConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	f, openError := file.Open()
	if openError != nil {
		global.LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))

		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close()                                                  // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra)
	if putErr != nil {
		global.LOG.Error("function formUploader.Put() Filed", zap.Any("err", putErr.Error()))
		return "", "", errors.New("function formUploader.Put() Filed, err:" + putErr.Error())
	}
	return global.CONFIG.Qiniu.ImgPath + "/" + ret.Key, ret.Key, nil
}

//@author: [88act](https://github.com/88act)
//@author: [88act-4](https://github.com/88act)
//@author: [88act-2](https://github.com/88act)
//@object: *Qiniu
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (*Qiniu) DeleteFile(key string) error {
	mac := qbox.NewMac(global.CONFIG.Qiniu.AccessKey, global.CONFIG.Qiniu.SecretKey)
	cfg := qiniuConfig()
	bucketManager := storage.NewBucketManager(mac, cfg)
	if err := bucketManager.Delete(global.CONFIG.Qiniu.Bucket, key); err != nil {
		global.LOG.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

//@author: [88act-2](https://github.com/88act)
//@object: *Qiniu
//@function: qiniuConfig
//@description: 根据配置文件进行返回七牛云的配置
//@return: *storage.Config

func qiniuConfig() *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      global.CONFIG.Qiniu.UseHTTPS,
		UseCdnDomains: global.CONFIG.Qiniu.UseCdnDomains,
	}
	switch global.CONFIG.Qiniu.Zone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}
