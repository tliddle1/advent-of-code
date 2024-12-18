package day18

import (
	"math"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

func part1(filePath string, n, numBytes int) int {
	input := parse.GetInput(filePath)
	_ = input
	grid := newGrid[string](n, ".")
	for i := range numBytes {
		line := input[i]
		x, y := parseLine(line)
		grid[y][x] = "#"
	}
	distanceGrid := newDistanceGrid(n, math.MaxInt32)
	distanceGrid.updateDistance(n-1, n-1, 0, grid)
	return distanceGrid[0][0]
}

func newDistanceGrid(n int, value int) DistanceGrid {
	return newGrid[int](n, value)
}

func part2(filePath string, n, numBytes int) string {
	input := parse.GetInput(filePath)
	_ = input
	grid := newGrid[string](n, ".")
	for i := range numBytes {
		line := input[i]
		x, y := parseLine(line)
		grid[y][x] = "#"
	}
	for i := numBytes; i < len(input); i++ {
		line := input[i]
		x, y := parseLine(line)
		grid[y][x] = "#"
		distanceGrid := newDistanceGrid(n, math.MaxInt32)
		distanceGrid.updateDistance(n-1, n-1, 0, grid)
		if distanceGrid[0][0] == math.MaxInt32 {
			return line
		}
	}
	return "none blocked"
}

func parseLine(line string) (int, int) {
	xy := strings.Split(line, ",")
	return parse.StringToInt(xy[0]), parse.StringToInt(xy[1])
}

type DistanceGrid [][]int

func (g DistanceGrid) updateDistance(x, y, distance int, grid [][]string) {
	if y < 0 || x < 0 || y >= len(g) || x >= len(g[y]) {
		return
	}
	if grid[y][x] == "#" {
		return
	}
	if distance < g[y][x] {
		g[y][x] = distance
		g.updateDistance(x-1, y, distance+1, grid)
		g.updateDistance(x+1, y, distance+1, grid)
		g.updateDistance(x, y-1, distance+1, grid)
		g.updateDistance(x, y+1, distance+1, grid)
	}

}

func newGrid[T any](n int, value T) [][]T {
	var grid [][]T
	for _ = range n {
		var row []T
		for _ = range n {
			row = append(row, value)
		}
		grid = append(grid, row)
	}
	return grid
}
