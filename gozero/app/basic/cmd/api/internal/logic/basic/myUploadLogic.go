package basic

import (
	"context"
	"fmt"
	"go-cms/app/basic/cmd/api/internal/svc"
	"go-cms/app/basic/cmd/api/internal/types"
	. "go-cms/common/baseModel"
	"go-cms/common/ctxdata"
	"go-cms/common/utils"
	"io"
	"math"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/fishtailstudio/imgo"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

const DEFAULT_MAX_WIDTH float64 = 200
const DEFAULT_MAX_HEIGHT float64 = 200

const MEDIA_TYPE_img int = 1
const MEDIA_TYPE_video = 2
const MEDIA_TYPE_audio int = 3
const MEDIA_TYPE_doc int = 4

var allow_ext map[string]string = map[string]string{
	".jpg":  "jpg",
	".png":  "png",
	".jpeg": "jpeg",
	".mp3":  "mp3",
	".wav":  "wav",
	".mp4":  "mp4",
	".pdf":  "pdf",
	".zip":  "zip",
	".flv":  "flv",
	".docx": "docx",
}

var img_ext map[string]string = map[string]string{
	".jpg":  "jpg",
	".png":  "png",
	".jpeg": "jpeg",
}
var video_ext map[string]string = map[string]string{
	".mp4": "mp4",
	".flv": "flv",
}
var audio_ext map[string]string = map[string]string{
	".mp3": "mp3",
	".wav": "wav",
}
var doc_ext map[string]string = map[string]string{
	".pdf":  "pdf",
	".zip":  "zip",
	".docx": "docx",
}

type MyUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewMyUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *MyUploadLogic {
	return &MyUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *MyUploadLogic) MyUpload(req *types.UploadReq) (resp *types.FileDetailResp, err error) {

	basicFile := BasicFile{}
	_ = copier.Copy(&basicFile, &req)
	if req.UserId > 0 {
		basicFile.UserId = req.UserId
	} else {
		basicFile.UserId = ctxdata.GetUidFromCtx(l.ctx)
	}
	//basicFile.UserId = ctxdata.GetUidFromCtx(l.ctx)
	//basicFile.UserId = req.UserId
	//logx.Errorf("===== MyUpload req BeCut ==%v", req.BeCut)
	_, file, err := l.r.FormFile("file")
	if err != nil {
		return nil, err
	}
	// 读取文件后缀
	basicFile.Ext = path.Ext(file.Filename)
	_, ok := allow_ext[basicFile.Ext]
	if !ok {
		return nil, errors.New("文件类型错误," + basicFile.Ext) // errors.Wrapf(xerr.NewErrCode(xerr.Fail), "非法文件类型 ,Ext = %v", basicFile.Ext)
		//return nil, errors.New("文件类型错误")
	}
	basicFile.MediaType = 0
	if _, ok = img_ext[basicFile.Ext]; ok {
		basicFile.MediaType = MEDIA_TYPE_img
	} else if _, ok = video_ext[basicFile.Ext]; ok {
		basicFile.MediaType = MEDIA_TYPE_video
	} else if _, ok = audio_ext[basicFile.Ext]; ok {
		basicFile.MediaType = MEDIA_TYPE_audio
	} else if _, ok = doc_ext[basicFile.Ext]; ok {
		basicFile.MediaType = MEDIA_TYPE_doc
	}

	// 读取文件名并加密
	//name := strings.TrimSuffix(file.Filename, ext)
	//name = utils.MD5V([]byte(name))
	// 20位 的 Nanoid uuid.Nanoid(20)
	basicFile.Name = utils.GUID()
	// 拼接新文件名
	// filename := name + "_" + time.Now().Format("20060102150405") + ext
	//filename := name + "_" + time.Now().Format("20060102150405") + ext
	path := ""
	// 如果客户端指定了 路径
	if req.FilePath != "" {
		path = l.svcCtx.Config.LocalRes.PathUser + "/" + req.FilePath
	} else {
		path = l.svcCtx.Config.LocalRes.PathUser + "/" + time.Now().Format("20060102")
	}
	logx.Infof("UploadFile path = %s", path)
	//fmt.Println("UploadFile path = ", path)
	// 尝试创建此路径
	//mkdirErr := os.MkdirAll(global.REDIS.Local.Path, os.ModePerm)
	mkdirErr := os.MkdirAll(l.svcCtx.Config.LocalRes.BasePath+path, os.ModePerm)
	if mkdirErr != nil {
		logx.Errorf("尝试创建此路径 error%v", mkdirErr)
		return nil, mkdirErr
		//global.LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		//return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	//path := global.REDIS.Local.Path + "/" + filename
	path_src := ""
	path2 := ""
	if !utils.IsEmptyStr(req.FileName) {
		// 客户端指定文件名
		path_src = path + "/" + req.FileName
		path_src = strings.Replace(path_src, ".jpg", "_src.jpg", 1)
		path_src = strings.Replace(path_src, ".png", "_src.png", 1)
		path2 = path + "/" + req.FileName
	} else {
		path_src = path + "/" + basicFile.Name + "_src" + basicFile.Ext //原图
		path2 = path + "/" + basicFile.Name + basicFile.Ext             // 缩略图
	}

	pathAll := l.svcCtx.Config.LocalRes.BasePath + path2
	pathAll_src := l.svcCtx.Config.LocalRes.BasePath + path_src
	// fmt.Println("UploadFile path   = ", path2)
	// fmt.Println("UploadFile path_src   = ", path_src)
	fmt.Println("UploadFile pathAll   = ", pathAll)
	// fmt.Println("UploadFile pathAll_src = ", pathAll_src)

	f, openError := file.Open() // 读取文件
	if openError != nil {
		logx.Errorv("读取文件 出错," + openError.Error())
		//global.LOG.Error("AAA file.Open() Filed", zap.Any("err", openError.Error()))
		//return "", "", errors.New("AAA file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	okPath := pathAll
	if basicFile.MediaType == MEDIA_TYPE_img {
		//_src 为原文件
		okPath = pathAll_src
	}
	out, createErr := os.Create(okPath)
	if createErr != nil {
		logx.Errorv("创建文件 出错," + createErr.Error())
		//global.LOG.Error("BBB os.Create() Filed", zap.Any("err", createErr.Error()))
		//return "", "", errors.New("BBB os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close()             // 创建文件 defer 关闭
	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		logx.Errorv("传输文件 出错," + copyErr.Error())
		//	global.LOG.Error("CCC io.Copy() Filed", zap.Any("err", copyErr.Error()))
		//return "", "", errors.New("CCC io.Copy() Filed, err:" + copyErr.Error())
	}
	//生成缩略图
	if basicFile.MediaType == MEDIA_TYPE_img {
		go makeThumbnail(pathAll_src, pathAll, basicFile.Ext)
		//go exeCmd2(pathAll_src, pathAll, basicFile.Ext, l.svcCtx.Config.LocalRes.BasePath+path)
	}
	//需要分割的全景图
	// if req.BeCut == 1 {
	// 	go cutTile(pathAll_src, pathAll, basicFile.Ext, l.svcCtx.Config.LocalRes.BasePath+path, basicFile.Name)
	// }

	//保存数据库
	basicFile.Driver = 0 //local
	basicFile.Path = path2
	basicFile.Guid = basicFile.Name
	basicFile.Name = file.Filename
	basicFile.Status = 1
	_, err = l.svcCtx.BasicFileSev.Create(l.ctx, basicFile)
	//info := types.FileInfo{Path: path}
	//resp := new(types.UploadResp)
	//*resp.Info = info
	return &types.FileDetailResp{Info: types.FileInfo{Path: l.svcCtx.Config.LocalRes.BaseUrl + basicFile.Path, Guid: basicFile.Guid, MediaType: basicFile.MediaType}}, nil
}

// 计算图片缩放后的尺寸
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(DEFAULT_MAX_WIDTH/float64(srcWidth), DEFAULT_MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// 生成缩略图
func makeThumbnail(imagePath string, savePath string, ext string) error {
	myImgo := imgo.Load(imagePath)
	w, h := calculateRatioFit(myImgo.Width(), myImgo.Height())
	//myImgo.Resize(w, h).Save(savePath, 50)
	myImgo.Resize(w, h).Save(savePath, 30)
	return nil
}

// 分割tile
func cutTile(imagePath string, savePath string, ext string, path string, guid string) {
	// 创建目录
	dir := path + "/" + guid + "/"

	logx.Errorf("创建目录 dir= %v", dir)
	mkdirErr := os.MkdirAll(dir, os.ModePerm)
	if mkdirErr != nil {
		logx.Errorf("cutTile 创建目录error= %v", mkdirErr.Error())
		return
	}

	// 调整大小
	toPath := dir + "vr" + ext
	toPathMini := dir + "mini" + ext
	toPathIcon := dir + "icon" + ext

	ctx2, cancel := context.WithTimeout(context.Background(), 720*time.Second)
	defer cancel()
	command := "convert -resize 16000x8000 -quality 70 " + imagePath + " " + toPath
	command1 := "convert -resize 800x400 " + toPath + " " + toPathMini
	command2 := "convert -resize 200x100 " + toPathMini + " " + toPathIcon
	output, err := exec.CommandContext(ctx2, "bash", "-c", command, command1, command2).Output()
	// cmd.Output（）返回的错误将特定于操作系统，具体取决于进程被杀死时发生的情况。
	if ctx2.Err() == context.DeadlineExceeded {
		logx.Errorf("cutTile 调整大小 error= %v", err.Error())
		return
	}
	logx.Errorf("cutTile 调整 output= %v", string(output))

}

// 分割tile
func cutTile3(imagePath string, savePath string, ext string, path string, guid string) {
	// 创建目录
	dir := path + "/" + guid + "/"

	logx.Errorf("创建目录 dir= %v", dir)
	mkdirErr := os.MkdirAll(dir, os.ModePerm)
	if mkdirErr != nil {
		logx.Errorf("cutTile 创建目录error= %v", mkdirErr.Error())
		return
	}

	// 调整大小
	toPath := dir + "vr" + ext
	toPathMini := dir + "mini" + ext
	toPathIcon := dir + "icon" + ext

	myImgo := imgo.Load(imagePath)
	//myImgo.Resize(w, h).Save(savePath, 50)
	myImgo.Resize(16000, 8000).Save(toPath, 70)
	myImgo.Resize(800, 400).Save(toPathMini, 70)
	myImgo.Resize(200, 100).Save(toPathIcon, 50)

	myImgo2 := imgo.Load(imagePath)

	for x := 0; x <= 32; x++ { // 常见的 for 循环，支持初始化语句。
		for y := 0; y <= 16; y++ { // 常见的 for 循环，支持初始化语句。
			pathTile := fmt.Sprintf("%svr%d_%d.jpg", dir, x, y)
			xx := x * 500
			yy := y * 500
			newImg := myImgo2.Crop(xx, yy, 500, 500)
			newImg.Save(pathTile)
			logx.Errorf(" xx =%d yy= %d", xx, yy)
			logx.Errorf(pathTile)
			//img.Save(pathTile, 70)
		}
	}

	// 创建一个上下文并且设置超时时间
	// ctx, cancel := context.WithTimeout(context.Background(), 720*time.Second)
	// defer cancel()
	// command := "convert " + toPath + " -crop 500x500  -set filename:tile '%[fx:floor(page.x/500)]_%[fx:floor(page.y/500)]' -set filename:orig " + dir + "%t %[filename:orig]_%[filename:tile].jpg"
	// output, err := exec.CommandContext(ctx, "bash", "-c", command).Output()
	// //output, err := exec.CommandContext(ctx, "cmd", "-c", command).Output()
	// // cmd.Output（）返回的错误将特定于操作系统，具体取决于进程被杀死时发生的情况。
	// if ctx.Err() == context.DeadlineExceeded {
	// 	logx.Errorf("cutTile 分割图片  超时", ctx.Err().Error())
	// 	return
	// }
	// logx.Errorf("cutTile 分割图片output= %v", string(output))
	// // 如果没有上下文错误，我们知道命令已经执行完成（或者执行出错）
	// //fmt.Println("已经执行完成 Output=", string(out))
	// if err != nil {
	// 	logx.Errorf("cutTile 分割图片  error= %s", err.Error())
	// 	logx.Errorf("cutTile 分割图片 command= %s", command)
	// }

}

// 分割tile
func cutTile2(imagePath string, savePath string, ext string, path string, guid string) {
	// 创建目录
	dir := path + "/" + guid + "/"

	//logx.Errorf("创建目录 dir= %v", dir)
	mkdirErr := os.MkdirAll(dir, os.ModePerm)
	if mkdirErr != nil {
		logx.Errorf("cutTile 创建目录error= %v", mkdirErr.Error())
		return
	}
	// 调整大小
	toPath := dir + "vr" + ext
	toPathMini := dir + "mini" + ext
	toPathIcon := dir + "icon" + ext
	// 创建一个上下文并且设置超时时间
	ctx2, cancel := context.WithTimeout(context.Background(), 720*time.Second)
	defer cancel()
	command := "convert -resize 16000x8000 -quality 70 " + imagePath + " " + toPath
	command1 := "convert -resize 800x400 " + toPath + " " + toPathMini
	command2 := "convert -resize 200x100 " + toPathMini + " " + toPathIcon
	output, err := exec.CommandContext(ctx2, "bash", "-c", command, command1, command2).Output()
	// cmd.Output（）返回的错误将特定于操作系统，具体取决于进程被杀死时发生的情况。
	if ctx2.Err() == context.DeadlineExceeded {
		logx.Errorf("cutTile 调整大小 error= %v", err.Error())
		return
	}
	logx.Errorf("cutTile 调整 output= %v", string(output))

	// 创建一个上下文并且设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 720*time.Second)
	defer cancel()
	command = "convert " + toPath + " -crop 500x500  -set filename:tile '%[fx:floor(page.x/500)]_%[fx:floor(page.y/500)]' -set filename:orig " + dir + "%t %[filename:orig]_%[filename:tile].jpg"
	_, err = exec.CommandContext(ctx, "bash", "-c", command).Output()
	// cmd.Output（）返回的错误将特定于操作系统，具体取决于进程被杀死时发生的情况。
	if ctx.Err() == context.DeadlineExceeded {
		logx.Errorf("cutTile 分割图片  超时")
		return
	}
	// 如果没有上下文错误，我们知道命令已经执行完成（或者执行出错）
	//fmt.Println("已经执行完成 Output=", string(out))
	if err != nil {
		logx.Errorf("cutTile 分割图片  error= %s", err.Error())
		logx.Errorf("cutTile 分割图片 command= %s", command)
	}

	// 使用我们的上下文创建命令
	//cmd := exec.CommandContext(ctx, "ping", "-c 4", "-i 1", "8.8.8.8")
	// 这次我们可以简单的使用Output()获取输出结果
	//out, err := cmd.Output()
	// 我们要通过检查上下文错误来确定超时是否执行
	// command1 := "cd " + path
	// command2 := "convert " + imagePath + "  -crop 500x500  aaa.jpg"
	// logx.Errorv(command1)
	// logx.Errorv(command2)
	// output, err := exec.Command("/bin/sh", "-c", command1, command2).Output()
	// if err != nil {
	// 	logx.Errorv(err.Error())
	// }
	// logx.Errorv(string(output))
}

/*
// 生成缩略图
func exeCmd(imagePath string, savePath string, ext string) {
	f, _ := os.Create(savePath)
	f.Close()
	//str := "C:\\Program Files\\ImageMagick-7.1.0-Q16-HDRI\\magick.exe convert -resize 16000x8000  -quality 70  " + imagePath + " " + savePath
	//str := "convert -resize 16000x8000  -quality 70  " + imagePath + " " + savePath
	output, err := exec.Command("magick", "convert", "-resize", "16000x8000", "-quality", "70", imagePath, savePath).Output()
	if err != nil {
		logx.Errorv(err.Error())
	}
	// 因为结果是字节数组，需要转换成string
	logx.Errorv(string(output))
	//output, err = exec.Command("magick", "-crop", "500x500", savePath, "-set", "filename:tile '%[fx:floor(page.x/500)]_%[fx:floor(page.y/500)]'", "-set", "filename:orig %t %[filename:orig]_%[filename:tile].jpg").Output()
	output, err = exec.Command("magick", savePath, "-crop", "500x500", "aaa.jpg").Output()
	if err != nil {
		logx.Errorv(err.Error())
	}
	logx.Errorv(string(output))
	//cmd := exec.Command(str)
	// 执行命令，返回命令是否执行成功
	// err := cmd.Run()
	// if err != nil {
	// 	// 命令执行失败
	// 	panic(err)
	// }

	// output, err := cmd.Output()

	// 通过exec.Command函数执行命令或者shell
	// 第一个参数是命令路径，当然如果PATH路径可以搜索到命令，可以不用输入完整的路径
	// 第二到第N个参数是命令的参数
	// 下面语句等价于执行命令: ls -l /var/
	// cmd := exec.Command("/bin/ls", "-l", "/var/")
	// // 执行命令，并返回结果
	// output, err := cmd.Output()
	// if err != nil {
	// 	panic(err)
	// }
	// // 因为结果是字节数组，需要转换成string
	// fmt.Println(string(output))

}
*/
