package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(tempStr))
}

func MakePassword(plainpwd, salt string) string {
	return MD5Encode(plainpwd + salt)
}
