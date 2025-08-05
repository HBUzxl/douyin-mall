package errorx

var message map[int]string

func init() {
	message = make(map[int]string)
	message[ErrInvalidParams] = "参数错误"
	message[ErrServerCommon] = "服务器开小差啦,稍后再来试一试"
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
