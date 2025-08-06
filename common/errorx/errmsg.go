package errorx

var message map[int]string

func init() {
	message = make(map[int]string)
	message[ErrInvalidParams] = "参数错误"
	message[ErrServerCommon] = "服务器开小差啦,稍后再来试一试"

	message[UNKNOWN_CLIENT_ERROR] = "未知客户端错误"
	message[INVALID_PARAMS_ERROR] = "参数无效"
	message[INVALID_TOKEN_ERROR] = "令牌无效"
	message[INVALID_ACCOUNT_PASSWORD_ERROR] = "账号或密码错误"
	message[PASSWORD_NOT_MATCH_ERROR] = "密码不匹配"
	message[CHARGE_FAILED_ERROR] = "支付失败，请稍后重试"
	message[STOCK_NOT_ENOUGH_ERROR] = "库存不足，请稍后再试"
	message[PRODUCT_NOT_FOUND_ERROR] = "商品未找到"
	message[TRANSACTION_NOT_FOUND_ERROR] = "交易记录未找到"

	message[UNKNOWN_SERVER_ERROR] = "服务器开小差啦，稍后再来试一试"
	message[AUTH_DELIBER_TOKEN_ERROR] = "认证令牌发放失败"
	message[USER_REGISTER_ERROR] = "用户注册失败"
	message[CHECKOUT_ERROR] = "结算服务异常"
}

func MapErrMsg(errCode int) string {
	if msg, ok := message[errCode]; ok {
		return msg
	} else {
		return "服务器开小差啦，请稍后再试"
	}
}

func IsCodeErr(errcode int) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
