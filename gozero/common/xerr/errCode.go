package xerr

//成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

//全局错误码
const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002
const TOKEN_EXPIRE_ERROR uint32 = 100003
const TOKEN_GENERATE_ERROR uint32 = 100004
const DB_ERROR uint32 = 200000
const DB_RecordNotFound uint32 = 200001
const DB_InsertErr uint32 = 200002
const DB_UpdateErr uint32 = 200003
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 200004

//用户模块
