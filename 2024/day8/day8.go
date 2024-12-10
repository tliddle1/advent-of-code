package day8

import (
	"fmt"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

type Coordinate [2]int

type Board struct {
	grid [][]string
}

func newBoard(input []string) *Board {
	var grid [][]string
	for _, line := range input {
		grid = append(grid, strings.Split(line, ""))
	}
	return &Board{grid}
}

func (this *Board) inGrid(coordinate Coordinate) bool {
	n := len(this.grid)
	m := len(this.grid[0])
	x, y := coordinate[0], coordinate[1]
	return !(x < 0 || y < 0 || x >= m || y >= n)
}

func (this *Board) Print(antinodes map[Coordinate]struct{}) {
	for antinode := range antinodes {
		x, y := antinode[0], antinode[1]
		this.grid[x][y] = "#"
	}
	for row := range this.grid {
		fmt.Println(this.grid[row])
	}
}

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	board := newBoard(input)
	return solve1(board)
}

func solve1(board *Board) int {
	allAntinodes := make(map[Coordinate]struct{})
	nodeLocations := make(map[string][]Coordinate)
	for i := range board.grid {
		for j, char := range board.grid[i] {
			if char != "." {
				nodeLocations[char] = append(nodeLocations[char], Coordinate{i, j})
			}
		}
	}
	for char, coordinates := range nodeLocations {
		_ = char
		for i := range len(coordinates) - 1 {
			for j := i + 1; j < len(coordinates); j++ {
				antinodes := getAntinodes(coordinates[i], coordinates[j])
				for _, antinode := range antinodes {
					if board.inGrid(antinode) {
						allAntinodes[antinode] = struct{}{}
					}
				}
			}
		}
	}
	//board.Print(allAntinodes)
	return len(allAntinodes)
}

func getAntinodes(coordinate1, coordinate2 Coordinate) []Coordinate {
	x1, y1 := coordinate1[0], coordinate1[1]
	x2, y2 := coordinate2[0], coordinate2[1]
	result := []Coordinate{
		{x2 + x2 - x1, y2 + y2 - y1},
		{x1 + x1 - x2, y1 + y1 - y2},
	}
	return result
}

func getAntinodes2(coordinate1, coordinate2 Coordinate, board *Board) []Coordinate {
	x1, y1 := coordinate1[0], coordinate1[1]
	x2, y2 := coordinate2[0], coordinate2[1]
	var result []Coordinate
	xDiff := x2 - x1
	yDiff := y2 - y1
	nextAntinode := Coordinate{x2 + xDiff, y2 + yDiff}
	for board.inGrid(nextAntinode) {
		result = append(result, nextAntinode)
		nextAntinode[0] += xDiff
		nextAntinode[1] += yDiff
	}
	xDiff = x1 - x2
	yDiff = y1 - y2
	nextAntinode = Coordinate{x1 + xDiff, y1 + yDiff}
	for board.inGrid(nextAntinode) {
		result = append(result, nextAntinode)
		nextAntinode[0] += xDiff
		nextAntinode[1] += yDiff
	}
	return result
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	board := newBoard(input)
	return solve2(board)
}

func solve2(board *Board) int {
	allAntinodes := make(map[Coordinate]struct{})
	nodeLocations := make(map[string][]Coordinate)
	for i := range board.grid {
		for j, char := range board.grid[i] {
			if char != "." {
				nodeLocations[char] = append(nodeLocations[char], Coordinate{i, j})
			}
		}
	}
	for char, coordinates := range nodeLocations {
		_ = char
		for i := range len(coordinates) - 1 {
			for j := i + 1; j < len(coordinates); j++ {
				antinodes := getAntinodes2(coordinates[i], coordinates[j], board)
				allAntinodes[coordinates[i]] = struct{}{}
				allAntinodes[coordinates[j]] = struct{}{}
				for _, antinode := range antinodes {
					if board.inGrid(antinode) {
						allAntinodes[antinode] = struct{}{}
					}
				}
			}
		}
	}
	//board.Print(allAntinodes)
	return len(allAntinodes)
}
