package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

type Equation struct {
	result   int
	operands []int
}

// operators
const (
	add = iota
	multiply
)

func isPossible1(result int, operands []int) bool {
	if len(operands) == 1 {
		return result == operands[0]
	}
	i := len(operands) - 1
	num := operands[i]
	if result%num == 0 && isPossible1(result/num, operands[:i]) {
		return true
	} else {
		return isPossible1(result-num, operands[:i])
	}
}

func isPossible2(result int, operands []int) bool {
	if result < 0 {
		return false
	}
	if len(operands) == 1 {
		return result == operands[0]
	}
	i := len(operands) - 1
	num := operands[i]
	if result%num == 0 && isPossible2(result/num, operands[:i]) {
		return true
	} else if isPossible2(result-num, operands[:i]) {
		return true
	} else {
		resultStr := strconv.Itoa(result)
		numStr := strconv.Itoa(num)
		if strings.HasSuffix(resultStr, numStr) {
			newNum := resultStr[:len(resultStr)-len(numStr)]
			if newNum == "-" {
				return false
			}
			if newNum == "" {
				newNum = "0"
			}
			newResult := parse.StringToInt(newNum)
			return isPossible2(newResult, operands[:i])
		}
		return false
	}
}

func (e Equation) calculate(operators []int) int {
	result := e.operands[0]
	for i := 1; i < len(e.operands); i++ {
		result = operate(result, e.operands[i], operators[i-1])
	}
	return result
}

func operate(a, b int, operator any) int {
	switch operator {
	case add:
		return a + b
	case multiply:
		return a * b
	}
	panic("Bad operator")
}

func newEquation(line string) Equation {
	var operands []int

	parts := strings.Split(line, ":")
	operandsString := strings.TrimSpace(parts[1])
	operandsStrs := strings.Split(operandsString, " ")
	for i := range operandsStrs {
		operands = append(operands, parse.StringToInt(operandsStrs[i]))
	}
	return Equation{
		result:   parse.StringToInt(parts[0]),
		operands: operands,
	}
}

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	var equations []Equation
	for _, line := range input {
		equations = append(equations, newEquation(line))
	}
	return solve1(equations)
}

func solve1(input []Equation) int {
	var result int
	for _, equation := range input {
		if isPossible1(equation.result, equation.operands) {
			result += equation.result
		}
	}
	return result
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	var equations []Equation
	for _, line := range input {
		equations = append(equations, newEquation(line))
	}
	return solve2(equations)
}

func solve2(input []Equation) int {
	var result int
	for i, equation := range input {
		if i == 233 {
			fmt.Println("")
		}
		if isPossible2(equation.result, equation.operands) {
			result += equation.result
		}
	}
	return result
}
