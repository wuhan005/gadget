package gadget

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

// Sha1 sha1 string.
func Sha1(raw string) string {
	h := sha1.New()
	_, _ = h.Write([]byte(raw))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// HmacSha1 HMAC SHA1 string.
func HmacSha1(input string, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	_, _ = io.WriteString(h, input)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Md5 md5 string.
func Md5(raw string) string {
	h := md5.New()
	_, _ = h.Write([]byte(raw))
	return hex.EncodeToString(h.Sum(nil))
}
