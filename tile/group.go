package tile

type Group struct {
	tiles []string
}

func NewGroup(tiles []string) *Group {
	return &Group{
		tiles: tiles,
	}
}

func (g Group) Password() int {
	var maxArea int
	for y := range g.tiles {
		for x := range g.tiles[y] {
			area := g.maxRectangleArea(x, y)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func (g Group) maxRectangleArea(x, y int) int {
	var maxArea int
	topLeftTile := g.tiles[y][x]
	xIndicesOfSameTilesInRow := g.findXIndicesOfSameTilesInRow(x, y)
	for _, xIndexOfSameTilesInRow := range xIndicesOfSameTilesInRow {
		yIndicesOfSameTilesInColumn := g.findYIndicesOfSameTilesInColumn(xIndexOfSameTilesInRow, y)
		for _, yIndexOfSameTilesInColumn := range yIndicesOfSameTilesInColumn {
			tile := g.tiles[yIndexOfSameTilesInColumn][x]
			if tile == topLeftTile {
				area := calcRectArea(x, y, xIndexOfSameTilesInRow, yIndexOfSameTilesInColumn)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func (g Group) findXIndicesOfSameTilesInRow(x, y int) []int {
	refTile := g.tiles[y][x]
	var result []int
	for xIndex := x + 1; xIndex < len(g.tiles[y]); xIndex++ {
		tile := g.tiles[y][xIndex]
		if refTile == tile {
			result = append(result, xIndex)
		}
	}
	return result
}

func (g Group) findYIndicesOfSameTilesInColumn(x, y int) []int {
	refTile := g.tiles[y][x]
	var result []int
	for yIndex := y + 1; yIndex < len(g.tiles); yIndex++ {
		tile := g.tiles[yIndex][x]
		if refTile == tile {
			result = append(result, yIndex)
		}
	}
	return result
}

func calcRectArea(topLeftX, topLeftY, bottomRightX, bottomRightY int) int {
	width := bottomRightX - topLeftX + 1
	height := bottomRightY - topLeftY + 1
	return width * height
}
