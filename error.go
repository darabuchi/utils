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

func (p *Error) Error() string {
	return p.Message
}

func (p *Error) SetNeedRetry() *Error {
	p.NeedRetry = true
	return p
}

func (p *Error) CheckNeedRetry() bool {
	return p.NeedRetry
}
