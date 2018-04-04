package tile

import (
	"fmt"
	"sync"
)

type Group struct {
	tiles []string
}

func NewGroup(tiles []string) *Group {
	return &Group{
		tiles: tiles,
	}
}

func (g Group) tileAt(x, y int) (uint8, error) {
	err := fmt.Errorf("index out of range (%d, %d)", x, y)
	if y < 0 || y >= len(g.tiles) {
		return 0, err
	}
	if x < 0 || x >= len(g.tiles[y]) {
		return 0, err
	}
	return g.tiles[y][x], nil
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

func (g Group) PasswordWithGoroutines() int {
	var wait sync.WaitGroup
	areaChan := make(chan int, len(g.tiles)*g.maxWidth())
	for y := range g.tiles {
		for x := range g.tiles[y] {
			wait.Add(1)
			go func(targetX, targetY int) {
				areaChan <- g.maxRectangleArea(targetX, targetY)
				wait.Done()
			}(x, y)
		}
	}
	wait.Wait()

	close(areaChan)

	var maxArea int
	for area := range areaChan {
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func (g Group) maxWidth() int {
	var maxWidth int
	for _, tiles := range g.tiles {
		if maxWidth > len(tiles) {
			maxWidth = len(tiles)
		}
	}
	return maxWidth
}

func (g Group) maxRectangleArea(x, y int) int {
	topLeftTile, err := g.tileAt(x, y)
	if err != nil {
		return 0
	}
	var maxArea int
	xIndicesOfSameTilesInRow := g.findXIndicesOfSameTilesInRow(x, y)
	for _, xIndexOfSameTilesInRow := range xIndicesOfSameTilesInRow {
		yIndicesOfSameTilesInColumn := g.findYIndicesOfSameTilesInColumn(xIndexOfSameTilesInRow, y)
		for _, yIndexOfSameTilesInColumn := range yIndicesOfSameTilesInColumn {
			tile, err := g.tileAt(x, yIndexOfSameTilesInColumn)
			if err != nil {
				continue
			}
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
	refTile, err := g.tileAt(x, y)
	if err != nil {
		return nil
	}
	var result []int
	for xIndex := x + 1; xIndex < len(g.tiles[y]); xIndex++ {
		tile, err := g.tileAt(xIndex, y)
		if err != nil {
			continue
		}
		if refTile == tile {
			result = append(result, xIndex)
		}
	}
	return result
}

func (g Group) findYIndicesOfSameTilesInColumn(x, y int) []int {
	refTile, err := g.tileAt(x, y)
	if err != nil {
		return nil
	}
	var result []int
	for yIndex := y + 1; yIndex < len(g.tiles); yIndex++ {
		tile, err := g.tileAt(x, yIndex)
		if err != nil {
			continue
		}
		if refTile == tile {
			result = append(result, yIndex)
		}
	}
	return result
}

func calcRectArea(topLeftX, topLeftY, bottomRightX, bottomRightY int) int {
	if topLeftX > bottomRightX {
		return 0
	}
	if topLeftY > bottomRightY {
		return 0
	}
	width := bottomRightX - topLeftX + 1
	height := bottomRightY - topLeftY + 1
	return width * height
}
