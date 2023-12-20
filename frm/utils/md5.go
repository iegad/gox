package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(raw []byte) []byte {
	h := md5.New()
	h.Write(raw)
	return h.Sum(nil)
}

func MD5Hex(raw string) string {
	h := md5.New()
	h.Write([]byte(raw))
	return hex.EncodeToString(h.Sum(nil))
}
