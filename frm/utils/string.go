package utils

import (
	"unsafe"
)

func Str2Bytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func Bytes2Str(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
