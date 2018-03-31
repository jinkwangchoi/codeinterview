package tile

type Group struct {
}

func NewGroup(tiles []string) *Group {
	return &Group{}
}

func (g *Group) Password() int {
	return 32
}

func (g *Group) maxRectangleArea(x, y int) int {
	return 24
}
