package tile

import (
	"testing"

	"bytes"
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
		So(group.maxRectangleArea(-1, -1), ShouldEqual, 0)
		So(group.maxRectangleArea(-1, -1), ShouldEqual, 0)
	})

	Convey("findXIndicesOfSameTilesInRow", t, func() {
		So(group.findXIndicesOfSameTilesInRow(0, 0), ShouldResemble, []int{5, 6})
		So(group.findXIndicesOfSameTilesInRow(1, 0), ShouldResemble, []int{2, 3, 4, 7})
		So(group.findXIndicesOfSameTilesInRow(-1, -1), ShouldBeNil)
		So(group.findXIndicesOfSameTilesInRow(-1, 0), ShouldBeNil)
	})

	Convey("findYIndicesOfSameTilesInColumn", t, func() {
		So(group.findYIndicesOfSameTilesInColumn(0, 0), ShouldResemble, []int{1, 2, 3})
		So(group.findYIndicesOfSameTilesInColumn(1, 0), ShouldResemble, []int{2})
		So(group.findYIndicesOfSameTilesInColumn(-1, -1), ShouldBeNil)
		So(group.findYIndicesOfSameTilesInColumn(-1, 0), ShouldBeNil)
	})

	Convey("calcRectArea", t, func() {
		So(calcRectArea(0, 0, 0, 0), ShouldEqual, 1)
		So(calcRectArea(0, 0, 1, 0), ShouldEqual, 2)
		So(calcRectArea(0, 0, 0, 1), ShouldEqual, 2)
		So(calcRectArea(0, 0, 1, 1), ShouldEqual, 4)

		// range check
		So(calcRectArea(1, 0, 0, 0), ShouldEqual, 0)
		So(calcRectArea(0, 1, 0, 0), ShouldEqual, 0)
		So(calcRectArea(1, 1, 0, 0), ShouldEqual, 0)
		So(calcRectArea(2, 0, 0, 0), ShouldEqual, 0)
		So(calcRectArea(0, 2, 0, 0), ShouldEqual, 0)
		So(calcRectArea(2, 2, 0, 0), ShouldEqual, 0)
	})
}

func benchmarkGroup(width, height int, b *testing.B) {
	var buffer bytes.Buffer
	for i := 0; i < width; i++ {
		buffer.WriteString("A")
	}
	atiles := buffer.String()
	tiles := make([]string, height)
	for i := range tiles {
		tiles[i] = atiles
	}
	group := NewGroup(tiles)
	for i := 0; i < b.N; i++ {
		group.Password()
	}
}

func BenchmarkGroup_Password10x10(b *testing.B) {
	benchmarkGroup(10, 10, b)
}

func BenchmarkGroup_Password20x20(b *testing.B) {
	benchmarkGroup(20, 20, b)
}

func BenchmarkGroup_Password30x30(b *testing.B) {
	benchmarkGroup(30, 30, b)
}

func BenchmarkGroup_Password40x40(b *testing.B) {
	benchmarkGroup(40, 40, b)
}

func BenchmarkGroup_Password50x50(b *testing.B) {
	benchmarkGroup(50, 50, b)
}

func BenchmarkGroup_Password60x60(b *testing.B) {
	benchmarkGroup(60, 60, b)
}
