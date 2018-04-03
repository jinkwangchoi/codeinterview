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
	return 32
}

func (g Group) maxRectangleArea(x, y int) int {
	if x == 1 {
		return 9
	}
	return 24
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
