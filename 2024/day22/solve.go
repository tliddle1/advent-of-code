package day22

import (
	"github.com/tliddle1/advent-of-code/2024/parse"
)

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	var result int
	for i, line := range input {
		_ = i
		result += parse.StringToInt(line)
	}
	return result
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	var result int
	for i, line := range input {
		_ = i
		result += parse.StringToInt(line)
	}
	return result
}
