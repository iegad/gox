package nw

import "unsafe"

const (
	DEFAULT_MAX_CONN  int32 = 1000
	DEFAULT_TIMEOUT   int32 = 300
	MAX_TIMEOUT       int32 = 600
	UINT32_SIZE             = int(unsafe.Sizeof(uint32(0)))
	MAX_BUF_SIZE            = 1024 * 1024
	DEFAULT_RBUF_SIZE       = 2000
)
