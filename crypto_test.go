package gadget

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Sha1(t *testing.T) {
	cases := map[string]string{
		"e99p1ant":               "48076e66a051950bd5cd7fc489924a5d67865dac",
		"大大大茄子":                  "26e73b13780881121765168ff86a4b2190ca39b1",
		"":                       "da39a3ee5e6b4b0d3255bfef95601890afd80709",
		"!@#$%^&*()*&^%$#$%^&*(": "ff579ddb1ca1a4ab90b85a9846340f54579b0e08",
	}
	Convey("Sha1 string", t, func() {
		for raw, result := range cases {
			So(Sha1(raw), ShouldEqual, result)
		}
	})
}

func Test_Md5(t *testing.T) {
	cases := map[string]string{
		"e99p1ant":               "0b6032c6ec55da4e03587e48d7e05c69",
		"大大大茄子":                  "8f9bfffe4a2a839e63a5f4f8f430a43c",
		"":                       "d41d8cd98f00b204e9800998ecf8427e",
		"!@#$%^&*()*&^%$#$%^&*(": "6abc01d9e72d9623e4768c433c37ad1b",
	}
	Convey("md5 string", t, func() {
		for raw, result := range cases {
			So(Md5(raw), ShouldEqual, result)
		}
	})
}

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
