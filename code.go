package gadget

import (
	"encoding/base64"
	"io/ioutil"
	"net/url"
	"strings"
)

// Base64Encode: base64 encode string.
func Base64Encode(raw string) string {
	return base64.StdEncoding.EncodeToString([]byte(raw))
}

// Base64Decode: base64 decode string.
func Base64Decode(raw string) string {
	reader := strings.NewReader(raw)
	decoder := base64.NewDecoder(base64.StdEncoding, reader)
	str, _ := ioutil.ReadAll(decoder)
	return string(str)
}

// URLEncode: url encode string.
func URLEncode(raw string) string {
	return url.QueryEscape(raw)
}

// URLDecode: url decode string.
func URLDecode(raw string) string {
	str, _ := url.QueryUnescape(raw)
	return str
}
