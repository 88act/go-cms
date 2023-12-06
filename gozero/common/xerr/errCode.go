package xerr

//成功返回
const OK uint32 = 200
const Fail uint32 = 400

/**(前3位代表业务,后三位代表具体功能)**/

//全局错误码
//const SERVER_COMMON_ERROR uint32 = 400101
const SERVER_COMMON_ERROR uint32 = 400
const REUQEST_PARAM_ERROR uint32 = 400102
const TOKEN_EXPIRE_ERROR uint32 = 400103
const TOKEN_GENERATE_ERROR uint32 = 400104
const DB_ERROR uint32 = 400200
const DB_RecordNotFound uint32 = 400201
const DB_InsertErr uint32 = 400202
const DB_UpdateErr uint32 = 400203
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 400204

const VerifyCode_ERROR uint32 = 400300

//用户模块
