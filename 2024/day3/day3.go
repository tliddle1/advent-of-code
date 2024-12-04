package day3

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

type ParsedInput struct {
	lines []string
}

func part1(filePath string) int {
	var result int

	file, err := os.Open(filePath)
	parse.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	parsedInput := ParsedInput{}
	///////////////////////////////////
	for scanner.Scan() {
		line := scanner.Text()
		parsedInput.lines = append(parsedInput.lines, line)
	}
	result += parsedInput.operate1()
	///////////////////////////////////
	parse.CheckError(scanner.Err())
	return result
}

func (this *ParsedInput) operate1() int {
	var result int
	for _, line := range this.lines {
		result += getResult(line)
	}
	return result
}

func mul(s string) int {
	nums := strings.Split(s[4:len(s)-1], ",")
	a := parse.StringToInt(nums[0])
	b := parse.StringToInt(nums[1])
	return a * b
}

func part2(filePath string) int {
	var result int

	file, err := os.Open(filePath)
	parse.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	parsedInput := ParsedInput{}
	///////////////////////////////////
	for scanner.Scan() {
		line := scanner.Text()
		parsedInput.lines = append(parsedInput.lines, line)
	}
	result += parsedInput.operate2()
	///////////////////////////////////
	parse.CheckError(scanner.Err())
	return result
}

func (this *ParsedInput) operate2() int {
	var input string
	for _, line := range this.lines {
		input += line
	}
	return getResult2(input, true, 0)
}

func getResult2(input string, do bool, result int) int {
	if do {
		str := "don't()"
		idx := strings.Index(input, str)
		if idx == -1 {
			return result + getResult(input)
		}
		result += getResult(input[:idx])
		return getResult2(input[idx:], !do, result)
	} else {
		str := "do()"
		idx := strings.Index(input, str)
		if idx == -1 {
			return result
		}
		return getResult2(input[idx:], !do, result)
	}
}

func getResult(line string) int {
	var result int
	r, err := regexp.Compile("mul\\([0-9][0-9]?[0-9]?,[0-9][0-9]?[0-9]?\\)")
	parse.CheckError(err)
	mulStrings := r.FindAll([]byte(line), -1)
	for _, mulString := range mulStrings {
		result += mul(string(mulString))
	}
	return result
}
