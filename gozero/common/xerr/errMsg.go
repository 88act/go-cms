package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "成功"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"
	message[DB_RecordNotFound] = "数据不存在"
	message[DB_InsertErr] = "新增数据失败"
	message[DB_UpdateErr] = "更新数据失败"
	message[VerifyCode_ERROR] = "验证码错误"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	//eturn false
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
