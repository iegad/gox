package pb

import "errors"

var (
	Err_NodeCodeInvalid   = errors.New("node code is invalid")
	Err_MessageIDInvalid  = errors.New("message id is invalid")
	Err_IdempotentInvalid = errors.New("idempotent is invalid")
)
