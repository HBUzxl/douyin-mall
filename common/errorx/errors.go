package errorx

import (
	"fmt"
	"net/http"
)

const (
	defaultCode = 10001
)

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func ErrorHandler(err error) (int, interface{}) {
	switch e := err.(type) {
	case *CodeError:
		return http.StatusOK, e.Data()
	default:
		return http.StatusOK, NewCodeError(-1, err.Error())
	}
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.Code, e.Msg)
}

func (e *CodeError) Data() CodeError {
	return CodeError{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

// 返回给前端的错误码
func (e *CodeError) GetErrCode() int {
	return e.Code
}

// 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.Msg
}

func NewErrCode(errCode int) *CodeError {
	return &CodeError{Code: errCode, Msg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{Code: ErrServerCommon, Msg: errMsg}
}

func NewErrMsgf(errMsg string, args ...any) *CodeError {
	return &CodeError{Code: ErrServerCommon, Msg: fmt.Sprintf(errMsg, args...)}
}
