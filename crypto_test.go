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

func Test_HmacSha1(t *testing.T) {
	// empty key
	key := ""
	cases := map[string]string{
		"e99p1ant":               "58f04bb1bab728a57996f38d34f3e00584a8c5f8",
		"大大大茄子":                  "3a996c5b3d1e5aad530416170d28c59a0f5bb932",
		"":                       "fbdb1d1b18aa6c08324b7d64b71fb76370690e1d",
		"!@#$%^&*()*&^%$#$%^&*(": "e8cd4d613044a333a8cf3de2e80b2477be873c87",
	}
	Convey("HmacSha1 string, empty key", t, func() {
		for raw, result := range cases {
			So(HmacSha1(raw, key), ShouldEqual, result)
		}
	})

	key = "e99"
	cases = map[string]string{
		"e99p1ant":               "819dbada1c4ae382980251906612cce1ab481d87",
		"大大大茄子":                  "71a4122e455fa78cd49027babc9675923819d81c",
		"":                       "f8744be95d7f51864d77b003e642931ff0f62e25",
		"!@#$%^&*()*&^%$#$%^&*(": "4563ff913bc7d5fd65f833ad0cd1a35c7e0d7ddd",
	}
	Convey("HmacSha1 string, key: e99", t, func() {
		for raw, result := range cases {
			So(HmacSha1(raw, key), ShouldEqual, result)
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
