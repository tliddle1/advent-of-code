package day19

import (
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

func getTowels(input []string) []string {
	return strings.Split(input[0], ", ")
}

func getRequests(input []string) []string {
	return input[2:]
}

var possible map[string]bool

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	towels := getTowels(input)
	requests := getRequests(input)
	var result int
	possible = make(map[string]bool)
	numPossible = make(map[string]int)
	for _, request := range requests {
		if requestPossible(request, towels) {
			result += 1
		}
	}
	return result
}

func requestPossible(request string, towels []string) bool {
	if isPossible, ok := possible[request]; ok {
		return isPossible
	}
	if request == "" {
		return true
	}
	for _, towel := range towels {
		if request == towel {
			return true
		}
		if strings.HasPrefix(request, towel) {
			if requestPossible(request[len(towel):], towels) {
				possible[request] = true
				return true
			}
		}
	}
	possible[request] = false
	return false
}

var numPossible map[string]int

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	towels := getTowels(input)
	requests := getRequests(input)
	var result int
	numPossible = make(map[string]int)
	for _, request := range requests {
		result += countPossible(request, towels)
	}
	return result
}

func countPossible(request string, towels []string) int {
	if num, ok := numPossible[request]; ok {
		return num
	}
	if request == "" {
		return 1
	}
	var total int
	for _, towel := range towels {
		if strings.HasPrefix(request, towel) {
			total += countPossible(request[len(towel):], towels)
		}
	}
	numPossible[request] = total
	return total
}
