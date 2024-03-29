package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[ServerCommonError] = "服务请求失败"
	message[ReuqestParamError] = "参数错误"

	message[TokenGenerateError] = "生成token失败"
	message[DataNoExistError] = "数据不存在"

	message[UserIsNotLoginError] = "用户未登录"
	message[TokenExpireError] = "token失效，请重新登陆"

}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
