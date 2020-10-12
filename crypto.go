package gadget

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)

func Sha1(raw string) string {
	h := sha1.New()
	h.Write([]byte(raw))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Md5(raw string) string {
	h := md5.New()
	h.Write([]byte(raw))
	return hex.EncodeToString(h.Sum(nil))
}

func Base64Encode(raw string) string {
	return base64.StdEncoding.EncodeToString([]byte(raw))
}

func Base64Decode(raw string) string {
	reader := strings.NewReader(raw)
	decoder := base64.NewDecoder(base64.StdEncoding, reader)
	str, _ := ioutil.ReadAll(decoder)
	return string(str)
}
