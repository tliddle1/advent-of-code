package day4

import (
	"bufio"
	"os"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

func part1(filePath string) int {
	file, err := os.Open(filePath)
	parse.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	///////////////////////////////////
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	result := solve(input)
	///////////////////////////////////
	parse.CheckError(scanner.Err())
	return result
}

func solve(input []string) int {
	var result int
	var grid [][]string
	for _, line := range input {
		grid = append(grid, strings.Split(line, ""))
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "X" {
				result += count(i, j, grid)
			}
		}
	}
	return result
}

func count(i int, j int, grid [][]string) int {
	var result int
	for d := range 8 {
		if getWord(i, j, grid, d) == "XMAS" {
			result++
		}
	}
	return result
}

const (
	n = iota
	ne
	e
	se
	s
	sw
	w
	nw
)

func getWord(i, j int, grid [][]string, d int) string {
	switch d {
	case n:
		if i >= 3 {
			return grid[i][j] + grid[i-1][j] + grid[i-2][j] + grid[i-3][j]
		}
	case ne:
		if i >= 3 && j < len(grid[i])-3 {
			return grid[i][j] + grid[i-1][j+1] + grid[i-2][j+2] + grid[i-3][j+3]
		}
	case e:
		if j < len(grid[i])-3 {
			return grid[i][j] + grid[i][j+1] + grid[i][j+2] + grid[i][j+3]
		}
	case se:
		if i < len(grid)-3 && j < len(grid[i])-3 {
			return grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2] + grid[i+3][j+3]
		}
	case s:
		if i < len(grid)-3 {
			return grid[i][j] + grid[i+1][j] + grid[i+2][j] + grid[i+3][j]
		}
	case sw:
		if i < len(grid)-3 && j >= 3 {
			return grid[i][j] + grid[i+1][j-1] + grid[i+2][j-2] + grid[i+3][j-3]
		}
	case w:
		if j >= 3 {
			return grid[i][j] + grid[i][j-1] + grid[i][j-2] + grid[i][j-3]
		}
	case nw:
		if i >= 3 && j >= 3 {
			return grid[i][j] + grid[i-1][j-1] + grid[i-2][j-2] + grid[i-3][j-3]
		}
	}
	return ""
}

func part2(filePath string) int {
	file, err := os.Open(filePath)
	parse.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	///////////////////////////////////
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	result := solve2(input)
	///////////////////////////////////
	parse.CheckError(scanner.Err())
	return result
}

func solve2(input []string) int {
	var result int
	var grid [][]string
	for _, line := range input {
		grid = append(grid, strings.Split(line, ""))
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "A" {
				if xmas(i, j, grid) {
					result++
				}
			}
		}
	}
	return result
}

func xmas(i int, j int, grid [][]string) bool {
	if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[0])-1 {
		return false
	}
	//SE
	SE := grid[i+1][j+1]
	//NW
	NW := grid[i-1][j-1]
	//SW
	SW := grid[i+1][j-1]
	//NE
	NE := grid[i-1][j+1]
	cross1 := (SE == "M" && NW == "S") || (SE == "S" && NW == "M")
	cross2 := (SW == "M" && NE == "S") || (SW == "S" && NE == "M")
	return cross1 && cross2
}
