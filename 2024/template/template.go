package template

import (
	"github.com/tliddle1/advent-of-code/2024/parse"
)

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	return solve1(input)
}

func solve1(input []string) int {
	var result int
	for i, line := range input {
		_ = i
		result += calculate1(line)
	}
	return result
}

func calculate1(line string) int {
	return parse.StringToInt(line)
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	return solve2(input)
}

func solve2(input []string) int {
	var result int
	for i, line := range input {
		_ = i
		result += calculate2(line)
	}
	return result
}

func calculate2(line string) int {
	return parse.StringToInt(line)
}
