package response

import (
	"net/http"

	"douyin-mall/common/errorx"

	"github.com/zeromicro/go-zero/core/trace"
)

type Body struct {
	Code    int         `json:"code"`
	Msg     string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	TraceId string      `json:"traceId,omitempty"`
}

const (
	ErrUnknown    = "未知错误"
	ErrParamsSSet = "参数设置错误，请核对参数是否输入完整"
	Success       = "Ok"
)

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	body.TraceId = w.Header().Get(trace.TraceIdKey)
	errcode := errorx.ErrServerCommon
}
