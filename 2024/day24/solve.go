package template

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

type Gate struct {
	input1    string
	input2    string
	output    string
	operation string
}

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	var wires map[string]int
	var gates []Gate
	wires = getStartingWires(input)
	gates = getGates(input)
	somethingUpdated := true
	for somethingUpdated {
		somethingUpdated = false
		for _, gate := range gates {
			_, ok := wires[gate.output]
			if !ok {
				value1, ok1 := wires[gate.input1]
				value2, ok2 := wires[gate.input2]
				if ok1 && ok2 {
					wires[gate.output] = logicGate(value1, value2, gate.operation)
					somethingUpdated = true
				}
			}
		}
	}
	var zWires []string
	var values []string
	for key, value := range wires {
		values = append(values, key+" "+strconv.Itoa(value))
		if strings.HasPrefix(key, "z") {
			zWires = append(zWires, key)
		}
	}
	sort.Strings(values)
	//fmt.Println(strings.Join(values, "\n"))
	slices.Sort(zWires)
	var result int
	for i, zWire := range zWires {
		result += parse.IntPow(2, i) * wires[zWire]
		//fmt.Println(zWire, wires[zWire], result)
	}
	return result
}

func getGates(input []string) []Gate {
	var result []Gate
	readingGates := false
	for _, line := range input {
		if readingGates {
			splitLine := strings.Split(line, " ")
			result = append(result, Gate{
				input1:    splitLine[0],
				input2:    splitLine[2],
				output:    splitLine[4],
				operation: splitLine[1],
			})
		}
		if line == "" {
			readingGates = true
		}

	}
	return result
}

func getStartingWires(input []string) map[string]int {
	result := make(map[string]int)
	for _, line := range input {
		if line == "" {
			return result
		}
		splitLine := strings.Split(line, ": ")
		result[splitLine[0]] = parse.StringToInt(splitLine[1])
	}
	return result
}

func logicGate(value1 int, value2 int, operation string) int {
	switch operation {
	case "AND":
		if value1 == 1 && value2 == 1 {
			return 1
		}
	case "OR":
		if value1 == 1 || value2 == 1 {
			return 1
		}
	case "XOR":
		if value1 != value2 {
			return 1
		}
	default:
		panic("invalid operation")
	}
	return 0
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	var gates []Gate
	outputs := make(map[string]string)
	gates = getGates(input)
	for _, gate := range gates {
		outputs[gate.output] = "(" + gate.input1 + " " + gate.operation + " " + gate.input2 + ")"
	}
	replacementOccurred := true
	for replacementOccurred {
		replacementOccurred = false
		for _, gate := range gates {
			for output, value := range outputs {
				if strings.Contains(outputs[gate.output], output) {
					replacementOccurred = true
					outputs[gate.output] = strings.Replace(outputs[gate.output], output, value, 1)
				}
			}
			outputs[gate.output] = "(" + gate.output + ")"
		}
	}
	for output, value := range outputs {
		fmt.Println(output, value)
	}
	return 0
}
