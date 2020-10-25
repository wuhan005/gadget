package gadget

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Abs(t *testing.T) {
	cases := map[int]int{
		0:    0,
		199:  199,
		-233: 233,
	}
	Convey("Abs", t, func() {
		for raw, result := range cases {
			So(Abs(raw), ShouldEqual, result)
		}
	})
}
