package day9

import (
	"strconv"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	return solve1(input)
}

func solve1(input []string) int {
	line := input[0]
	originalAllocation := OriginalAllocation(line)
	newAllocation := NewAllocation(originalAllocation)
	return checkSum(newAllocation)
}

func NewAllocation(allocation []string) []string {
	var newAllocation []string
	var count int
	j := len(allocation) - 1
	for _, char := range allocation {
		for allocation[j] == "." {
			j--
		}
		if char != "." {
			count++
			newAllocation = append(newAllocation, char)
		} else {
			newAllocation = append(newAllocation, allocation[j])
			j--
		}
	}
	return newAllocation[:count]
}

func OriginalAllocation(line string) []string {
	var result []string
	idNumber := 0
	for i, char := range line {
		if i%2 == 0 {
			for _ = range parse.StringToInt(string(char)) {
				result = append(result, strconv.Itoa(idNumber))
			}
			idNumber++
		} else {
			for _ = range parse.StringToInt(string(char)) {
				result = append(result, ".")
			}
		}
	}
	return result
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	_ = input
	return 0
}

func checkSum(str []string) int {
	var result int
	for i, numStr := range str {
		result += i * parse.StringToInt(numStr)
	}
	return result
}
