package utils

import (
	"crypto/hmac"
	"crypto/sha1"
)

func HmacSha1Signature(message, key []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}
