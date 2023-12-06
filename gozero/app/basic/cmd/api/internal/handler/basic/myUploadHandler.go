package basic

import (
	"go-cms/app/basic/cmd/api/internal/logic/basic"
	"go-cms/app/basic/cmd/api/internal/svc"
	"go-cms/app/basic/cmd/api/internal/types"
	"go-cms/common/result"
	"net/http"
	"strconv"
)

func MyUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// if err := httpx.Parse(r, &req); err != nil {
		// 	result.ParamErrorResult(r, w, err)
		// 	return
		// }
		//fmt.Println("MyUploadHandler开始文件上传....")
		var req types.UploadReq
		//basicFile := model.BasicFile{}
		r.ParseMultipartForm(32 << 20) //32 Mb
		r.ParseForm()
		if r.MultipartForm != nil {
			// userId, _ := strconv.Atoi(r.MultipartForm.Value["userId"][0])
			//basicFile.UserId = int64(userId)
			req.Md5 = r.MultipartForm.Value["md5"][0]
			req.Sha1 = r.MultipartForm.Value["sha1"][0]
			req.Size, _ = strconv.Atoi(r.MultipartForm.Value["size"][0])
			req.FileType = r.MultipartForm.Value["fileType"][0]
			catid, _ := strconv.Atoi(r.MultipartForm.Value["catId"][0])
			req.CatId = int64(catid)
			beCut, _ := strconv.Atoi(r.MultipartForm.Value["beCut"][0])
			req.BeCut = beCut
		}
		//logx.Errorf("beCut = ", req.BeCut)
		l := basic.NewMyUploadLogic(r.Context(), svcCtx, r)
		resp, err := l.MyUpload(&req)
		result.HttpResult(r, w, resp, err, req)
	}
}
