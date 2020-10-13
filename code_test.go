package gadget

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Base64Encode(t *testing.T) {
	cases := map[string]string{
		"e99p1ant":               "ZTk5cDFhbnQ=",
		"大大大茄子":                  "5aSn5aSn5aSn6IyE5a2Q",
		"":                       "",
		"!@#$%^&*()*&^%$#$%^&*(": "IUAjJCVeJiooKSomXiUkIyQlXiYqKA==",
	}
	Convey("base64 encode string", t, func() {
		for raw, result := range cases {
			So(Base64Encode(raw), ShouldEqual, result)
		}
	})
}

func Test_Base64Decode(t *testing.T) {
	cases := map[string]string{
		"ZTk5cDFhbnQ=":                     "e99p1ant",
		"5aSn5aSn5aSn6IyE5a2Q":             "大大大茄子",
		"":                                 "",
		"IUAjJCVeJiooKSomXiUkIyQlXiYqKA==": "!@#$%^&*()*&^%$#$%^&*(",
	}
	Convey("base64 decode string", t, func() {
		for raw, result := range cases {
			So(Base64Decode(raw), ShouldEqual, result)
		}
	})
}

func Test_URLEncode(t *testing.T) {
	cases := map[string]string{
		"e99p1ant":   "e99p1ant",
		"!@#$%^&*()": "%21%40%23%24%25%5E%26%2A%28%29",
		"":           "",
		"大大大茄子":      "%E5%A4%A7%E5%A4%A7%E5%A4%A7%E8%8C%84%E5%AD%90",
		"\n\r":       "%0A%0D",
	}
	Convey("url encode string", t, func() {
		for raw, result := range cases {
			So(URLEncode(raw), ShouldEqual, result)
		}
	})
}

func Test_URLDecode(t *testing.T) {
	cases := map[string]string{
		"e99p1ant":                       "e99p1ant",
		"%21%40%23%24%25%5E%26%2A%28%29": "!@#$%^&*()",
		"":                               "",
		"%E5%A4%A7%E5%A4%A7%E5%A4%A7%E8%8C%84%E5%AD%90": "大大大茄子",
		"%0A%0D": "\n\r",
	}
	Convey("url decode string", t, func() {
		for raw, result := range cases {
			So(URLDecode(raw), ShouldEqual, result)
		}
	})
}
