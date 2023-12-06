package basic

import (
	"errors"
	"net/http"

	"go-cms/common/result"

	"go-cms/app/basic/cmd/api/internal/logic/basic"
	"go-cms/app/basic/cmd/api/internal/svc"
	"go-cms/app/basic/cmd/api/internal/types"

	"github.com/gogf/gf/util/gconv"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//文件上传一般会采用 POST multipart/form-data 的形式，处理这类请求要调用 r.ParseMultipartForm，无论是显式调用，
		//还是在 r.FormFile 里面的隐式调用 那 32 Mb 是对文件上传大小的限制吗？不是，上传的文件们按顺序存入内存中，累加大小不得超出 32Mb

		// basicFile := model.BasicFile{}
		// r.ParseMultipartForm(32 << 20) //32 Mb
		// r.ParseForm()
		// if r.MultipartForm != nil {
		// 	//userId, _ := strconv.Atoi(r.MultipartForm.Value["userId"][0])

		// 	basicFile.Md5 = r.MultipartForm.Value["md5"][0]
		// 	basicFile.Sha1 = r.MultipartForm.Value["sha1"][0]
		// 	basicFile.Size, _ = strconv.Atoi(r.MultipartForm.Value["size"][0])
		// 	basicFile.FileType = r.MultipartForm.Value["fileType"][0]
		// 	catid, _ := strconv.Atoi(r.MultipartForm.Value["catId"][0])
		// 	basicFile.CatId = int64(catid)
		// }
		// l := basic.NewUploadLogic(r.Context(), svcCtx, r)
		// resp, err := l.Upload(basicFile)
		// result.HttpResult(r, w, resp, err)

		//logx.Errorf("UploadHandler开始文件上传....")
		//fmt.Println("UploadHandler开始文件上传....")

		var req types.UploadReq
		//basicFile := model.BasicFile{}
		r.ParseMultipartForm(32 << 20) //32 Mb
		r.ParseForm()

		if r.MultipartForm != nil {
			if len(r.MultipartForm.Value["userId"]) > 0 {
				req.UserId = gconv.Int64(r.MultipartForm.Value["userId"][0])
			}
			if len(r.MultipartForm.Value["filePath"]) > 0 {
				req.FilePath = r.MultipartForm.Value["filePath"][0]
			} else {
				result.HttpResult(r, w, nil, errors.New("缺少参数filePath"), req)
			}
			if len(r.MultipartForm.Value["fileName"]) > 0 {
				req.FileName = r.MultipartForm.Value["fileName"][0]
			} else {
				result.HttpResult(r, w, nil, errors.New("fileName"), req)
			}
			// if len(r.MultipartForm.Value["md5"]) > 0 {
			// 	req.Md5 = r.MultipartForm.Value["md5"][0]
			// } else {
			// 	result.HttpResult(r, w, nil, errors.New("缺少参数md5"), req)
			// }
			// if len(r.MultipartForm.Value["size"]) > 0 {
			// 	req.Size = gconv.Int(r.MultipartForm.Value["size"][0])
			// } else {

			// 	result.HttpResult(r, w, nil, errors.New("缺少参数size"), req)
			// }
			// if len(r.MultipartForm.Value["fileType"]) > 0 {
			// 	req.FileType = r.MultipartForm.Value["fileType"][0]
			// } else {
			// 	result.HttpResult(r, w, nil, errors.New("缺少参数fileType"), req)
			// }
		}
		//logx.Errorv(req)
		l := basic.NewMyUploadLogic(r.Context(), svcCtx, r)
		resp, err := l.MyUpload(&req)

		result.HttpResult(r, w, resp, err, req)
	}
}
