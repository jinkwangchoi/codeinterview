package tile

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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
		So(group.maxRectangleArea(1, 0), ShouldEqual, 9)
	})

	Convey("findXIndicesOfSameTilesInRow", t, func() {
		So(group.findXIndicesOfSameTilesInRow(0, 0), ShouldResemble, []int{5, 6})
		So(group.findXIndicesOfSameTilesInRow(1, 0), ShouldResemble, []int{2, 3, 4, 7})
	})

	Convey("findYIndicesOfSameTilesInColumn", t, func() {
		So(group.findYIndicesOfSameTilesInColumn(0, 0), ShouldResemble, []int{1, 2, 3})
		So(group.findYIndicesOfSameTilesInColumn(1, 0), ShouldResemble, []int{2})
	})

	Convey("calcRectArea", t, func() {
		So(calcRectArea(0, 0, 0, 0), ShouldEqual, 1)
		So(calcRectArea(0, 0, 1, 0), ShouldEqual, 2)
		So(calcRectArea(0, 0, 0, 1), ShouldEqual, 2)
		So(calcRectArea(0, 0, 1, 1), ShouldEqual, 4)
	})
}
