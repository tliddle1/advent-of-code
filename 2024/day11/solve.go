package day11

import (
	"strconv"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	stones := strings.Split(input[0], " ")
	return blinkStones(stones, 25)
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	stones := strings.Split(input[0], " ")
	return blinkStones(stones, 75)
}

func blinkStones(stones []string, numIterations int) int {
	stoneMap := make(map[string]int)
	for _, stone := range stones {
		stoneMap[stone]++
	}
	for _ = range numIterations {
		newStoneMap := make(map[string]int)
		for stone, count := range stoneMap {
			blinkResult := blink(stone)
			for _, newStone := range blinkResult {
				newStoneMap[newStone] += count
			}
		}
		stoneMap = newStoneMap
	}
	var result int
	for _, value := range stoneMap {
		result += value
	}
	return result
}

func blink(stone string) []string {
	if stone == "0" {
		return []string{"1"}
	}
	if len(stone)%2 == 0 {
		return []string{stone[:len(stone)/2], parse.RemoveLeadingZeros(stone[len(stone)/2:])}
	}
	return []string{strconv.Itoa(parse.StringToInt(stone) * 2024)}
}
