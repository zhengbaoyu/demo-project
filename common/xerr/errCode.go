package xerr

//成功返回
const OK uint32 = 200

//请求参数错误
const ReuqestParamError uint32 = 1002

//生成token失败
const TokenGenerateError uint32 = 1004

//数据不存在
const DataNoExistError uint32 = 1007

//服务器请求失败
const ServerCommonError uint32 = 5001

//用户未登录
const UserIsNotLoginError uint32 = 1001

//token过期
const TokenExpireError uint32 = 1003
