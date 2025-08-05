package response

import (
	"net/http"

	"douyin-mall/common/errorx"

	"github.com/pingcap/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/trace"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
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
	Success       = "OK"
)

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	body.TraceId = w.Header().Get(trace.TraceIdKey)
	errcode := errorx.ErrServerCommon
	errmsg := ErrUnknown
	if err != nil {
		if resp == nil {
			// 处理参数校验错误
			errmsg = ErrParamsSSet
		}
		causeErr := errors.Cause(err)
		// 自定义错误类型
		if e, ok := causeErr.(*errorx.CodeError); ok {
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok {
				// grpc err 错误
				grpcCode := int(gstatus.Code())
				if errorx.IsCodeErr(grpcCode) {
					// 区分自定义错误跟系统底层，db等错误；底层，db不返回给前端
					errcode = grpcCode
					errmsg = gstatus.Message()
				} else {
					errmsg = gstatus.Message()
				}
			}
		}
		body.Code = errcode
		body.Msg = errmsg
		logx.Errorf("[GATEWAY-ERR]: %+v", err) // %+v 打印错误堆栈
	} else {
		body.Msg = Success
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
