package upload

import (
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"

	"github.com/nfnt/resize"

	"errors"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/utils"
	"go.uber.org/zap"
)

const DEFAULT_MAX_WIDTH float64 = 200
const DEFAULT_MAX_HEIGHT float64 = 200

type Local struct{}

//@author: [88act](https://github.com/88act)
//@author: [88act-4](https://github.com/88act)
//@author: [88act-2](https://github.com/88act)
//@object: *Local
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*Local) UploadFile(file *multipart.FileHeader, module int, userType int) (string, string, error) {
	// 读取文件后缀
	fmt.Println("UploadFile userType  = ", userType)
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	//name := strings.TrimSuffix(file.Filename, ext)
	//name = utils.MD5V([]byte(name))

	name := utils.GUID()
	// 拼接新文件名
	// filename := name + "_" + time.Now().Format("20060102150405") + ext

	//filename := name + "_" + time.Now().Format("20060102150405") + ext
	var path string
	if userType == 1 { //管理用户
		path = global.CONFIG.Local.Path + "/" + strconv.Itoa(module) +
			"/" + time.Now().Format("20060102")
	} else if userType == 2 { //普通用户
		path = global.CONFIG.Local.PathUser + "/" + strconv.Itoa(module) +
			"/" + time.Now().Format("20060102")
	}
	fmt.Println("UploadFile path = ", path)

	// 尝试创建此路径
	//mkdirErr := os.MkdirAll(global.REDIS.Local.Path, os.ModePerm)
	mkdirErr := os.MkdirAll(global.CONFIG.Local.BasePath+path, os.ModePerm)
	if mkdirErr != nil {
		global.LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	//path := global.REDIS.Local.Path + "/" + filename
	path_src := path + "/" + name + "_src" + ext //原图
	path = path + "/" + name + ext               // 缩略图
	pathAll := global.CONFIG.Local.BasePath + path
	pathAll_src := global.CONFIG.Local.BasePath + path_src
	fmt.Println("UploadFile path   = ", path)
	fmt.Println("UploadFile path_src   = ", path_src)
	fmt.Println("UploadFile pathAll   = ", pathAll)
	fmt.Println("UploadFile pathAll_src = ", pathAll_src)

	f, openError := file.Open() // 读取文件
	if openError != nil {
		global.LOG.Error("AAA file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("AAA file.Open() Filed, err:" + openError.Error())
	}

	defer f.Close() // 创建文件 defer 关闭
	out, createErr := os.Create(pathAll_src)
	if createErr != nil {
		global.LOG.Error("BBB os.Create() Filed", zap.Any("err", createErr.Error()))
		return "", "", errors.New("BBB os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close()             // 创建文件 defer 关闭
	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.LOG.Error("CCC io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("CCC io.Copy() Filed, err:" + copyErr.Error())
	}
	//生成缩略图
	makeThumbnail(pathAll_src, pathAll, ext)
	return path, name, nil
}

// 计算图片缩放后的尺寸
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(DEFAULT_MAX_WIDTH/float64(srcWidth), DEFAULT_MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// 生成缩略图
func makeThumbnail(imagePath string, savePath string, ext string) error {
	file, _ := os.Open(imagePath)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	w, h := calculateRatioFit(width, height)

	fmt.Println("width = ", width, " height = ", height)
	fmt.Println("w = ", w, " h = ", h)

	// 调用resize库进行图片缩放
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	// 需要保存的文件
	imgfile, _ := os.Create(savePath)
	defer imgfile.Close()

	// 以PNG格式保存文件
	if strings.ToLower(ext) == ".png" {
		err = png.Encode(imgfile, m)
	} else {
		err = jpeg.Encode(imgfile, m, nil)
	}
	if err != nil {
		return err
	}
	return nil
}

//@author: [88act](https://github.com/88act)
//@author: [88act-4](https://github.com/88act)
//@author: [88act-2](https://github.com/88act)
//@object: *Local
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (*Local) DeleteFile(key string) error {
	p := global.CONFIG.Local.Path + "/" + key
	if strings.Contains(p, global.CONFIG.Local.Path) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
