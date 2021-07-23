package gadget

import (
	"encoding/base64"
	"io/ioutil"
	"net/url"
	"strings"
)

// Base64Encode encodes base64 string.
func Base64Encode(raw string) string {
	return base64.StdEncoding.EncodeToString([]byte(raw))
}

// Base64Decode decodes base64 string.
func Base64Decode(raw string) string {
	reader := strings.NewReader(raw)
	decoder := base64.NewDecoder(base64.StdEncoding, reader)
	str, _ := ioutil.ReadAll(decoder)
	return string(str)
}

// URLEncode encodes URL string.
func URLEncode(raw string) string {
	return url.QueryEscape(raw)
}

// URLDecode decode URL string.
func URLDecode(raw string) string {
	str, _ := url.QueryUnescape(raw)
	return str
}
