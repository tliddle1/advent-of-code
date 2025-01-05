package day4

import (
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
	"github.com/tliddle1/advent-of-code/2024/spatial"
)

func makeGrid(input []string) spatial.Grid[string] {
	var grid spatial.Grid[string]
	for _, row := range input {
		grid = append(grid, strings.Split(row, ""))
	}
	return grid
}

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	g := makeGrid(input)

	var result int
	for y := range g.N() {
		for x := range g.M() {
			pos := spatial.Pos{x, y}
			if g.Get(pos) == "X" {
				result += countXMAS(pos, g)
			}
		}
	}
	return result
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	grid := makeGrid(input)

	var result int
	for y := range grid.N() {
		for x := range grid.M() {
			pos := spatial.Pos{x, y}
			if grid.Get(pos) == "A" {
				if xmas(pos, grid) {
					result++
				}
			}
		}
	}
	return result
}

func countXMAS(pos spatial.Pos, g spatial.Grid[string]) int {
	var result int
	for d := range 8 {
		dir := spatial.Dir(d)
		if getWord(g, pos, dir, 4) == "XMAS" {
			result++
		}
	}
	return result
}

func getWord(g spatial.Grid[string], pos spatial.Pos, dir spatial.Dir, length int) string {
	if !g.Contains(pos.Move(dir, 3)) {
		return ""
	}
	var word string
	for i := range length {
		word += g.Get(pos.Move(dir, i))
	}
	return word
}

func xmas(pos spatial.Pos, grid spatial.Grid[string]) bool {
	if grid.OnEdge(pos) {
		return false
	}
	SE := grid.Get(pos.Move(spatial.SE, 1))
	NW := grid.Get(pos.Move(spatial.NW, 1))
	SW := grid.Get(pos.Move(spatial.SW, 1))
	NE := grid.Get(pos.Move(spatial.NE, 1))
	cross1 := (SE == "M" && NW == "S") || (SE == "S" && NW == "M")
	cross2 := (SW == "M" && NE == "S") || (SW == "S" && NE == "M")
	return cross1 && cross2
}
