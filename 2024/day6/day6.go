package day6

import (
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

type Grid [][]string
type Pos [2]int

func PosLess(a, b Pos) bool {
	return a[0] < b[0]
}

const (
	North = iota
	East
	South
	West
)

func moveDirection(direction int, pos Pos) Pos {
	switch direction {
	case North:
		return [2]int{pos[0] - 1, pos[1]}
	case East:
		return [2]int{pos[0], pos[1] + 1}
	case South:
		return [2]int{pos[0] + 1, pos[1]}
	case West:
		return [2]int{pos[0], pos[1] - 1}
	}
	panic("invalid direction")
}

func (g Grid) StartingLocation() Pos {
	for i := range g {
		for j := range g[i] {
			if g[i][j] == "^" {
				return Pos{i, j}
			}
		}
	}
	return Pos{-1, -1}
}

func (g Grid) At(pos Pos) string {
	if pos[0] < 0 || pos[1] < 0 || pos[0] >= len(g) || pos[1] >= len(g[0]) {
		return "OUT_OF_RANGE"
	}
	return g[pos[0]][pos[1]]
}

func Part1(filePath string) int {
	input := parse.GetInput(filePath)
	grid := inputToGrid(input)
	return solve(grid)
}

func inputToGrid(input []string) Grid {
	var grid Grid
	for _, line := range input {
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}

func solve(grid Grid) int {
	return len(getVisitedPositions(grid))
}

func getVisitedPositions(grid Grid) map[Pos]struct{} {
	visitedPositions := make(map[Pos]struct{})
	currentPos := grid.StartingLocation()
	currentDir := North
	for grid.At(currentPos) != "OUT_OF_RANGE" {
		visitedPositions[currentPos] = struct{}{}
		nextPos := moveDirection(currentDir, currentPos)
		if grid.At(nextPos) == "#" {
			currentDir = rotate90(currentDir)
		} else {
			currentPos = nextPos
		}
	}
	return visitedPositions
}

func rotate90(dir int) int {
	switch dir {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	}
	panic("invalid direction")
}

func Part2(filePath string) int {
	input := parse.GetInput(filePath)
	grid := inputToGrid(input)
	return solve2(grid)
}

func solve2(grid Grid) int {
	var result int
	for position := range getVisitedPositions(grid) {
		if hasInfiniteLoop(grid, position) {
			result++
		}
	}
	return result
}

func hasInfiniteLoop(grid Grid, obstacle Pos) bool {
	currentPos := grid.StartingLocation()
	currentDir := North
	timesVisited := make(map[Pos]int)
	for grid.At(currentPos) != "OUT_OF_RANGE" {
		timesVisited[currentPos]++
		if timesVisited[currentPos] > 4 {
			return true
		}
		nextPos := moveDirection(currentDir, currentPos)
		if grid.At(nextPos) == "#" || nextPos == obstacle {
			currentDir = rotate90(currentDir)
		} else {
			currentPos = nextPos
		}
	}
	return false
}
