package template

import (
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	keys, locks := getKeysAndLocks(input)
	var result int
	for _, key := range keys {
		for _, lock := range locks {
			valid := true
			for i := range 5 {
				if key[i]+lock[i] > 5 {
					valid = false
					break
				}
			}
			if valid {
				result++
			}
		}
	}
	return result
}

type Key [5]int
type Lock [5]int

func getKeysAndLocks(input []string) ([]Key, []Lock) {
	var keys []Key
	var locks []Lock
	objects := strings.Split(strings.Join(input, "\n"), "\n\n")
	for _, obj := range objects {
		lines := strings.Split(obj, "\n")
		curType := lines[0][0]
		curObject := [5]int{}
		for _, line := range lines[1 : len(lines)-1] {
			for i, char := range line {
				if char == '#' {
					curObject[i]++
				}
			}
		}
		if curType == '#' {
			locks = append(locks, curObject)
		} else if curType == '.' {
			keys = append(keys, curObject)
		}
	}
	return keys, locks
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	_ = input
	var result int
	return result
}
