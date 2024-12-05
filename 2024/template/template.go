package template

import (
	"bufio"
	"os"

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
	///////////////////////////////////
	parse.CheckError(scanner.Err())
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
	///////////////////////////////////
	parse.CheckError(scanner.Err())
	return solve(input)
}
