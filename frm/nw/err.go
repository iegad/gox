package nw

import "errors"

var (
	ErrInvalidHeader  = errors.New("header is invalid")
	ErrInvalidBufSize = errors.New("buf size is invalid")
	ErrConfigIsNil    = errors.New("IOServiceConfig is nil")
	ErrNoPoints       = errors.New("must be a tcp or ws point to listen")
	ErrTcpEPIsInvalid = errors.New("tcp_endpoint is invalid")
	ErrWsEPIsInvalid  = errors.New("ws_endpoint is invalid")
)
