package m

import "errors"

var (
	Err_F_SessIsNotExists  = errors.New("front: session is not exists")
	Err_F_MessageIDInvalid = errors.New("front: message_id is invalid")

	Err_B_NodeIDInvalid = errors.New("backend: node_id is invalid")
)
