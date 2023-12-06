package basic

import (
	"context"
	"net/http"

	"go-cms/app/basic/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) UploadLogic {
	return UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

// func (l *UploadLogic) Upload(basicFile model.BasicFile) (resp *types.UploadResp, err error) {
// 	//logx.Errorv(" Upload==================")
// 	//userId := ctxdata.GetUidFromCtx(l.ctx)
// 	//logx.Errorf("userId ====%v ", userId)
// 	basicFile.UserId = ctxdata.GetUidFromCtx(l.ctx)
// 	_, file, err := l.r.FormFile("file")
// 	if err != nil {
// 		return nil, err
// 	}
// 	// 读取文件后缀

// 	basicFile.Ext = path.Ext(file.Filename)

// 	_, ok := allow_ext[basicFile.Ext]
// 	if !ok {

// 		return nil, errors.Wrapf(xerr.NewErrCode(xerr.Fail), "非法文件类型 ,Ext = %v", basicFile.Ext)
// 		//return nil, errors.New("文件类型错误")
// 	}
// 	basicFile.MediaType = 0
// 	if _, ok = img_ext[basicFile.Ext]; ok {
// 		basicFile.MediaType = MEDIA_TYPE_img
// 	} else if _, ok = video_ext[basicFile.Ext]; ok {
// 		basicFile.MediaType = MEDIA_TYPE_video
// 	} else if _, ok = audio_ext[basicFile.Ext]; ok {
// 		basicFile.MediaType = MEDIA_TYPE_audio
// 	} else if _, ok = doc_ext[basicFile.Ext]; ok {
// 		basicFile.MediaType = MEDIA_TYPE_doc
// 	}

// 	// 读取文件名并加密
// 	//name := strings.TrimSuffix(file.Filename, ext)
// 	//name = utils.MD5V([]byte(name))
// 	basicFile.Name = utils.GUID()
// 	// 拼接新文件名
// 	// filename := name + "_" + time.Now().Format("20060102150405") + ext
// 	//filename := name + "_" + time.Now().Format("20060102150405") + ext
// 	var path string
// 	path = l.svcCtx.Config.LocalRes.PathUser + "/" + strconv.Itoa(1) + "/" + time.Now().Format("20060102")
// 	fmt.Println("UploadFile path = ", path)
// 	// 尝试创建此路径
// 	//mkdirErr := os.MkdirAll(global.REDIS.Local.Path, os.ModePerm)
// 	mkdirErr := os.MkdirAll(l.svcCtx.Config.LocalRes.BasePath+path, os.ModePerm)
// 	if mkdirErr != nil {
// 		logx.Errorf("尝试创建此路径 error%v", mkdirErr)
// 		return nil, mkdirErr
// 		//global.LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
// 		//return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
// 	}
// 	// 拼接路径和文件名
// 	//path := global.REDIS.Local.Path + "/" + filename
// 	path_src := path + "/" + basicFile.Name + "_src" + basicFile.Ext //原图
// 	path = path + "/" + basicFile.Name + basicFile.Ext               // 缩略图
// 	pathAll := l.svcCtx.Config.LocalRes.BasePath + path
// 	pathAll_src := l.svcCtx.Config.LocalRes.BasePath + path_src
// 	//fmt.Println("UploadFile path   = ", path)
// 	//fmt.Println("UploadFile path_src   = ", path_src)
// 	//fmt.Println("UploadFile pathAll   = ", pathAll)
// 	//fmt.Println("UploadFile pathAll_src = ", pathAll_src)

// 	f, openError := file.Open() // 读取文件
// 	if openError != nil {
// 		//global.LOG.Error("AAA file.Open() Filed", zap.Any("err", openError.Error()))
// 		//return "", "", errors.New("AAA file.Open() Filed, err:" + openError.Error())
// 	}

// 	defer f.Close() // 创建文件 defer 关闭
// 	okPath := pathAll
// 	if basicFile.MediaType == MEDIA_TYPE_img {
// 		okPath = pathAll_src
// 	}
// 	out, createErr := os.Create(okPath)
// 	if createErr != nil {
// 		//global.LOG.Error("BBB os.Create() Filed", zap.Any("err", createErr.Error()))
// 		//return "", "", errors.New("BBB os.Create() Filed, err:" + createErr.Error())
// 	}
// 	defer out.Close()             // 创建文件 defer 关闭
// 	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
// 	if copyErr != nil {
// 		//	global.LOG.Error("CCC io.Copy() Filed", zap.Any("err", copyErr.Error()))
// 		//return "", "", errors.New("CCC io.Copy() Filed, err:" + copyErr.Error())
// 	}
// 	//生成缩略图
// 	if basicFile.MediaType == MEDIA_TYPE_img {
// 		go makeThumbnail(pathAll_src, pathAll, basicFile.Ext)
// 	}

// 	//保存数据库
// 	basicFile.Driver = 0 //local
// 	basicFile.Path = path
// 	basicFile.Guid = basicFile.Name
// 	basicFile.Name = file.Filename
// 	_, err = l.svcCtx.BasicFileSev.Create(l.ctx, basicFile)
// 	//info := types.FileInfo{Path: path}
// 	//resp := new(types.UploadResp)
// 	//*resp.Info = info
// 	return &types.UploadResp{Info: types.FileInfo{Path: l.svcCtx.Config.LocalRes.BaseUrl + basicFile.Path, Guid: basicFile.Guid, MediaType: basicFile.MediaType}}, nil
// }

// // 计算图片缩放后的尺寸
// func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
// 	ratio := math.Min(DEFAULT_MAX_WIDTH/float64(srcWidth), DEFAULT_MAX_HEIGHT/float64(srcHeight))
// 	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
// }

// // 生成缩略图
// func makeThumbnail(imagePath string, savePath string, ext string) error {
// 	// logx.Errorv(imagePath)
// 	// logx.Errorv(savePath)
// 	// logx.Errorv(ext)

// 	imgo.Load(imagePath).
// 		Resize(200, 100).
// 		Save(savePath)
// 		//	Insert("gopher.png", 50, 50).

// 		// file, _ := os.Open(imagePath)
// 		// defer file.Close()

// 		// img, _, err := image.Decode(file)
// 		// if err != nil {
// 		// 	return err
// 		// }

// 		// b := img.Bounds()
// 		// width := b.Max.X
// 		// height := b.Max.Y

// 		// w, h := calculateRatioFit(width, height)

// 		// fmt.Println("width = ", width, " height = ", height)
// 		// fmt.Println("w = ", w, " h = ", h)

// 		// // 调用resize库进行图片缩放
// 		// m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

// 		// // 需要保存的文件
// 		// imgfile, _ := os.Create(savePath)
// 		// defer imgfile.Close()

// 		// // 以PNG格式保存文件
// 		// if strings.ToLower(ext) == ".png" {
// 		// 	err = png.Encode(imgfile, m)
// 		// } else {
// 		// 	err = jpeg.Encode(imgfile, m, nil)
// 		// }
// 		// if err != nil {
// 		// 	return err
// 		// }
// 	return nil
// }

// // func (l *UploadLogic) Upload(basicFile model.BasicFile) (resp *types.UploadResp, err error) {

// // 	_, file, err := l.r.FormFile("file")
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	// 读取文件后缀

// // 	basicFile.Ext = path.Ext(file.Filename)

// // 	_, ok := allow_ext[basicFile.Ext]
// // 	if !ok {

// // 		return nil, errors.Wrapf(xerr.NewErrCode(xerr.Fail), "非法文件类型 ,Ext = %v", basicFile.Ext)
// // 		//return nil, errors.New("文件类型错误")
// // 	}
// // 	if _, ok = img_ext[basicFile.Ext]; ok {
// // 		basicFile.MediaType = MEDIA_TYPE_img
// // 	}
// // 	if _, ok = video_ext[basicFile.Ext]; ok {
// // 		basicFile.MediaType = 2
// // 	}
// // 	if _, ok = audio_ext[basicFile.Ext]; ok {
// // 		basicFile.MediaType = 3
// // 	}
// // 	if _, ok = doc_ext[basicFile.Ext]; ok {
// // 		basicFile.MediaType = 4
// // 	}

// // 	// 读取文件名并加密
// // 	//name := strings.TrimSuffix(file.Filename, ext)
// // 	//name = utils.MD5V([]byte(name))
// // 	basicFile.Name = utils.GUID()
// // 	// 拼接新文件名
// // 	// filename := name + "_" + time.Now().Format("20060102150405") + ext
// // 	//filename := name + "_" + time.Now().Format("20060102150405") + ext
// // 	var path string
// // 	path = l.svcCtx.Config.LocalRes.PathUser + "/" + strconv.Itoa(1) + "/" + time.Now().Format("20060102")
// // 	fmt.Println("UploadFile path = ", path)
// // 	// 尝试创建此路径
// // 	//mkdirErr := os.MkdirAll(global.REDIS.Local.Path, os.ModePerm)
// // 	mkdirErr := os.MkdirAll(l.svcCtx.Config.LocalRes.BasePath+path, os.ModePerm)
// // 	if mkdirErr != nil {
// // 		logx.Errorf("尝试创建此路径 error%v", mkdirErr)
// // 		return nil, mkdirErr
// // 		//global.LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
// // 		//return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
// // 	}
// // 	// 拼接路径和文件名
// // 	//path := global.REDIS.Local.Path + "/" + filename
// // 	path_src := path + "/" + basicFile.Name + "_src" + basicFile.Ext //原图
// // 	path = path + "/" + basicFile.Name + basicFile.Ext               // 缩略图
// // 	pathAll := l.svcCtx.Config.LocalRes.BasePath + path
// // 	pathAll_src := l.svcCtx.Config.LocalRes.BasePath + path_src
// // 	fmt.Println("UploadFile path   = ", path)
// // 	fmt.Println("UploadFile path_src   = ", path_src)
// // 	fmt.Println("UploadFile pathAll   = ", pathAll)
// // 	fmt.Println("UploadFile pathAll_src = ", pathAll_src)

// // 	f, openError := file.Open() // 读取文件
// // 	if openError != nil {
// // 		//global.LOG.Error("AAA file.Open() Filed", zap.Any("err", openError.Error()))
// // 		//return "", "", errors.New("AAA file.Open() Filed, err:" + openError.Error())
// // 	}

// // 	defer f.Close() // 创建文件 defer 关闭
// // 	okPath := pathAll
// // 	if basicFile.MediaType == MEDIA_TYPE_img {
// // 		okPath = pathAll_src
// // 	}
// // 	out, createErr := os.Create(okPath)
// // 	if createErr != nil {
// // 		//global.LOG.Error("BBB os.Create() Filed", zap.Any("err", createErr.Error()))
// // 		//return "", "", errors.New("BBB os.Create() Filed, err:" + createErr.Error())
// // 	}
// // 	defer out.Close()             // 创建文件 defer 关闭
// // 	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
// // 	if copyErr != nil {
// // 		//	global.LOG.Error("CCC io.Copy() Filed", zap.Any("err", copyErr.Error()))
// // 		//return "", "", errors.New("CCC io.Copy() Filed, err:" + copyErr.Error())
// // 	}
// // 	//生成缩略图
// // 	if basicFile.MediaType == MEDIA_TYPE_img {
// // 		makeThumbnail(pathAll_src, pathAll, basicFile.Ext)
// // 	}

// // 	//保存数据库
// // 	basicFile.Driver = 0 //local
// // 	basicFile.Path = path
// // 	basicFile.Guid = basicFile.Name
// // 	basicFile.Name = file.Filename
// // 	_, err = l.svcCtx.BasicFileSev.Create(l.ctx, basicFile)
// // 	//info := types.FileInfo{Path: path}
// // 	//resp := new(types.UploadResp)
// // 	//*resp.Info = info
// // 	return &types.UploadResp{Info: types.FileInfo{Path: l.svcCtx.Config.LocalRes.BaseUrl + basicFile.Path, Guid: basicFile.Guid, MediaType: basicFile.MediaType}}, nil
