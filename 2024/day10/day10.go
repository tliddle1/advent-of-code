package day10

import (
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

type Grid [][]int

func (g Grid) at(coordinate Coordinate) int {
	return g[coordinate[0]][coordinate[1]]
}

func (g Grid) contains(coordinate Coordinate) bool {
	x, y := coordinate[0], coordinate[1]
	return !(x < 0 || y < 0 || x >= len(g) || y >= len(g[0]))
}

type Coordinate [2]int

func newGrid(input []string) Grid {
	var grid Grid
	for _, line := range input {
		var row []int
		for _, num := range strings.Split(line, "") {
			row = append(row, parse.StringToInt(num))
		}
		grid = append(grid, row)
	}
	return grid
}

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	return solve1(newGrid(input))
}

func solve1(grid Grid) int {
	var result int
	for i := range grid {
		for j := range grid[i] {
			if grid.at(Coordinate{i, j}) == 0 {
				result += countTrailPeaks(Coordinate{i, j}, grid)
			}
		}
	}
	return result
}
func countTrailPeaks(coordinate Coordinate, grid Grid) int {
	return len(parse.RemoveDuplicates[Coordinate](getTrailPeaks(coordinate, grid)))
}

func getTrailPeaks(coordinate Coordinate, grid Grid) []Coordinate {
	cur := grid.at(coordinate)
	if cur == 9 {
		return []Coordinate{coordinate}
	}
	var trailPeaks []Coordinate
	x, y := coordinate[0], coordinate[1]
	// North
	nextCoordinate := Coordinate{x, y + -1}
	if grid.contains(nextCoordinate) && grid.at(nextCoordinate) == cur+1 {
		trailPeaks = append(trailPeaks, getTrailPeaks(nextCoordinate, grid)...)
	}
	// East
	nextCoordinate = Coordinate{x + 1, y}
	if grid.contains(nextCoordinate) && grid.at(nextCoordinate) == cur+1 {
		trailPeaks = append(trailPeaks, getTrailPeaks(nextCoordinate, grid)...)
	}
	// South
	nextCoordinate = Coordinate{x, y + 1}
	if grid.contains(nextCoordinate) && grid.at(nextCoordinate) == cur+1 {
		trailPeaks = append(trailPeaks, getTrailPeaks(nextCoordinate, grid)...)
	}
	// West
	nextCoordinate = Coordinate{x + -1, y}
	if grid.contains(nextCoordinate) && grid.at(nextCoordinate) == cur+1 {
		trailPeaks = append(trailPeaks, getTrailPeaks(nextCoordinate, grid)...)
	}
	return trailPeaks
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	return solve2(newGrid(input))
}

func solve2(grid Grid) int {
	var result int
	for i := range grid {
		for j := range grid[i] {
			if grid.at(Coordinate{i, j}) == 0 {
				result += countTrails(Coordinate{i, j}, grid)
			}
		}
	}
	return result
}

func countTrails(coordinate Coordinate, grid Grid) int {
	return len(getTrailPeaks(coordinate, grid))
}
