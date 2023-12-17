package basic

import "errors"

var (
	Err_CFG_NodeCodeInvalid = errors.New("config: node_code is invalid")
	Err_CFG_FrontInvalid    = errors.New("config: front is invalid")
	Err_CFG_FrontMaxConn    = errors.New("config: front.max_conn is invalid")
	Err_CFG_FrontEndpoint   = errors.New("config: front.tcp_endpoint or ws_endpoint is invalid")
	Err_CFG_FrontTimeout    = errors.New("config: front.timeout is invalid")
	Err_CFG_BackendInvalid  = errors.New("config: backend is invalid")
	Err_CFG_BackendEndpoint = errors.New("config: front.tcp_endpoint is invalid")
)
