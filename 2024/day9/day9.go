package day9

import (
	"strconv"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

func part1(filePath string) int {
	input := parse.GetInput(filePath)
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

func reallocate(allocation []string, curID string) []string {
	curIdx := getIdx(allocation, curID)
	curLen := getLen(allocation, curID, curIdx)
	i := 0
	for i < len(allocation) && i < curIdx {
		char := allocation[i]
		if char == "." {
			emptyLen := getLen(allocation, ".", i)
			if emptyLen >= curLen {
				//reallocate
				for j := range curLen {
					allocation[i+j] = curID
					allocation[curIdx+j] = "."
				}
				return allocation
			}
		}
		i++
	}
	return allocation
}

func getLen(allocation []string, id string, idx int) int {
	length := 0
	for _, char := range allocation[idx:] {
		if char == id {
			length++
		} else {
			return length
		}
	}
	return length
}

func getIdx(allocation []string, id string) int {
	for idx, char := range allocation {
		if char == id {
			return idx
		}
	}
	return -1
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	line := input[0]
	allocation := OriginalAllocation(line)
	for i := parse.StringToInt(allocation[len(allocation)-1]); i >= 0; i-- {
		curID := strconv.Itoa(i)
		allocation = reallocate(allocation, curID)
	}
	return checkSum(allocation)
}

func checkSum(str []string) int {
	var result int
	for i, numStr := range str {
		if numStr == "." {
			continue
		}
		result += i * parse.StringToInt(numStr)
	}
	return result
}
