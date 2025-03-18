package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 MD5加密
func EncodeMD5(value string) string {
	if value == "" {
		return ""
	}
	
	h := md5.New()
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}
