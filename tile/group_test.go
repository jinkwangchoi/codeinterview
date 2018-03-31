package tile

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGroup(t *testing.T) {
	tiles := []string{
		`BCCCCBBCAA`,
		`BACBBABBAA`,
		`BCBCAAABCB`,
		`BBBACBACBA`,
		`AAACACCBAC`,
		`ABBAACBCCC`,
		`CBAACBBCAA`,
	}
	group := NewGroup(tiles)
	Convey("example", t, func() {
		So(group.Password(), ShouldEqual, 32)
	})

	Convey("maxRectangleArea", t, func() {
		So(group.maxRectangleArea(0, 0), ShouldEqual, 24)
	})
}
