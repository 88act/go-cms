package result

import (
	"fmt"
	"net/http"

	"go-cms/common/ctxdata"
	"go-cms/common/xerr"

	"github.com/gogf/gf/util/gconv"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

// http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error, req interface{}) {

	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := "服务器开小差啦，稍后再来试一试"
		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			} else {
				//普通的错误
				errcode = xerr.Fail
				errmsg = err.Error()
			}
		}
		reqJson := gconv.String(req)
		logMsg := fmt.Sprint("API错误 Path=", r.URL.Path, ",errcode=", errcode, ",errmsg=", errmsg, ",req=", reqJson)
		useLogMsg := ctxdata.GetUserLogInfo(r.Context())
		logMsg = logMsg + "," + useLogMsg
		logc.Error(r.Context(), logMsg)
		httpx.WriteJson(w, http.StatusOK, Error(errcode, errmsg))

	}
}

// 授权的http方法
func AuthHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := "服务器开小差啦，稍后再来试一试"
		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			} else {
				//普通的错误
				errcode = xerr.Fail
				errmsg = err.Error()
			}
		}
		logMsg := fmt.Sprint("GATEWAY错误 Path=", r.URL.Path, ",errcode=", errcode, ",msg=", errmsg) // ", Error=", err.Error())
		logc.Error(r.Context(), logMsg)
		//fmt.Println(logMsg)
		//logx.WithContext(r.Context()).Errorf("【GATEWAY-ERR】 : %+v ", err)
		httpx.WriteJson(w, http.StatusUnauthorized, Error(errcode, errmsg))
	}
}

// http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", xerr.MapErrMsg(xerr.REUQEST_PARAM_ERROR), err.Error())
	logMsg := fmt.Sprintf("参数错误 Path=%s , err=%s", r.URL.Path, errMsg)
	useLogMsg := ctxdata.GetUserLogInfo(r.Context())
	logMsg = logMsg + "," + useLogMsg
	logc.Error(r.Context(), logMsg)
	httpx.WriteJson(w, http.StatusBadRequest, Error(xerr.REUQEST_PARAM_ERROR, errMsg))
}
