package tile

type Group struct {
}

func NewGroup(tiles []string) *Group {
	return &Group{}
}

func (g Group) Password() int {
	return 32
}
