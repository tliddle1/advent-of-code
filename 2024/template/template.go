package template

import (
	"github.com/tliddle1/advent-of-code/2024/parse"
)

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	return solve(input)
}

func solve(input []string) int {
	var result int
	for i, line := range input {
		_ = i
		result += calculate(line)
	}
	return result
}

func calculate(line string) int {
	return parse.StringToInt(line)
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	return solve(input)
}
