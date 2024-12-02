package day1

import (
	"bufio"
	"os"
	"sort"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

type ParsedLine struct {
	line       string
	num1, num2 int
}

func part1(filePath string) int {
	var result int
	var side1 []int
	var side2 []int

	file, err := os.Open(filePath)
	parse.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	///////////////////////////////////
	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := parseLine(line)
		side1 = append(side1, parsedLine.num1)
		side2 = append(side2, parsedLine.num2)
	}
	sort.Ints(side1)
	sort.Ints(side2)
	for i := range len(side1) {
		result += parse.Abs(side1[i] - side2[i])
	}
	///////////////////////////////////
	parse.CheckError(scanner.Err())
	return result
}

func (this *ParsedLine) operate1() int {
	return 0
}

func part2(filePath string) int {
	var result int
	frequencies := make(map[int]int) // number: count
	parsedLines := make([]ParsedLine, 0)

	file, err := os.Open(filePath)
	parse.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	///////////////////////////////////
	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := parseLine(line)
		frequencies[parsedLine.num2]++
		parsedLines = append(parsedLines, parsedLine)
	}
	for _, parsedLine := range parsedLines {
		result += parsedLine.operate2(frequencies)
	}
	///////////////////////////////////
	parse.CheckError(scanner.Err())
	return result
}

func (this *ParsedLine) operate2(freq2 map[int]int) int {
	return this.num1 * freq2[this.num1]
}

func parseLine(line string) ParsedLine {
	items := strings.Split(line, "   ")
	return ParsedLine{
		line: line,
		num1: parse.StringToInt(items[0]),
		num2: parse.StringToInt(items[1]),
	}
}
