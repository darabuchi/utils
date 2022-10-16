package utils

import (
	"fmt"
)

type Error struct {
	ErrCode   int32  `json:"err_code,omitempty"`
	Message   string `json:"message,omitempty"`
	NeedRetry bool   `json:"need_retry,omitempty"`
}

func NewError(format string, a ...interface{}) *Error {
	return &Error{
		Message:   fmt.Sprintf(format, a...),
		NeedRetry: false,
	}
}

func NewErrorWithCode(code int32, format string, a ...interface{}) *Error {
	return &Error{
		ErrCode:   code,
		Message:   fmt.Sprintf(format, a...),
		NeedRetry: false,
	}
}

func (p *Error) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", p.ErrCode, p.Message)
}

func (p *Error) SetNeedRetry() *Error {
	p.NeedRetry = true
	return p
}

func (p *Error) CheckNeedRetry() bool {
	return p.NeedRetry
}

func (p *Error) SetErrCode(errcode int32) *Error {
	p.ErrCode = errcode
	return p
}

func GetErrCode(err error) int32 {
	if err == nil {
		return 0
	}
	if x, ok := err.(*Error); ok {
		return x.ErrCode
	} else {
		return -1
	}
}
