package day2

import (
	"bufio"
	"os"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
	"github.com/tliddle1/advent-of-code/2024/slice"
)

type ParsedLine struct {
	line    string
	numbers []int
}

func parseLine(line string) ParsedLine {
	var numbers []int
	nums := strings.Split(line, " ")
	for _, num := range nums {
		numbers = append(numbers, parse.StringToInt(num))
	}
	return ParsedLine{
		line:    line,
		numbers: numbers,
	}
}

func part1(filePath string) int {
	var result int

	file, err := os.Open(filePath)
	parse.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	///////////////////////////////////
	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := parseLine(line)
		result += parsedLine.operate1()
	}
	///////////////////////////////////
	parse.CheckError(scanner.Err())
	return result
}

func (this *ParsedLine) operate1() int {
	if safe(this.numbers) {
		return 1
	}
	return 0
}

func safe(numbers []int) bool {

	var diffs []int
	for i := range len(numbers) - 1 {
		diff := numbers[i] - numbers[i+1]
		diffs = append(diffs, diff)
	}
	result := verifySafe(diffs)
	return result
}

func verifySafe(diffs []int) bool {
	return (parse.AllNegative(diffs) || parse.AllPositive(diffs)) && (parse.AllBetweenAB(diffs, 1, 3) || parse.AllBetweenAB(diffs, -3, -1))
}

func part2(filePath string) int {
	var result int

	file, err := os.Open(filePath)
	parse.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	///////////////////////////////////
	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := parseLine(line)
		result += parsedLine.operate2()
	}
	///////////////////////////////////
	parse.CheckError(scanner.Err())
	return result
}

func (this *ParsedLine) operate2() int {
	if safe(this.numbers) {
		return 1
	}
	for i := range this.numbers {
		if safe(slice.Remove(this.numbers, i)) {
			return 1
		}
	}
	return 0
}
