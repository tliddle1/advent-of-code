package template

import (
	"bufio"
	"os"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

type ParsedLine struct {
	line string
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
	return parse.StringToInt(this.line)
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
	return parse.StringToInt(this.line)
}

func parseLine(line string) ParsedLine {
	return ParsedLine{
		line: line,
	}
}
