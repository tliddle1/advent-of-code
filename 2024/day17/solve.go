package day17

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

func part1(filePath string) string {
	input := parse.GetInput(filePath)
	var program []string
	for i, line := range input {
		if i == 0 {
			A = parse.StringToInt(strings.Split(line, " ")[2])
		} else if i == 1 {
			B = parse.StringToInt(strings.Split(line, " ")[2])
		} else if i == 2 {
			C = parse.StringToInt(strings.Split(line, " ")[2])
		} else if i == 4 {
			program = strings.Split(strings.Split(line, " ")[1], ",")
		}
	}
	InstructionPointer = 0
	return strings.Join(strings.Split(execute(program), ""), ",")
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	_ = input
	return 0
}

func execute(program []string) string {
	var result string
	for InstructionPointer < len(program)-1 {
		opCode := program[InstructionPointer]
		operand := program[InstructionPointer+1]
		switch opCode {
		case "0":
			fmt.Println("adv")
			adv(operand)
		case "1":
			fmt.Println("bxl")
			bxl(operand)
		case "2":
			fmt.Println("bst")
			bst(operand)
		case "3":
			fmt.Println("jnz")
			jnz(operand)
		case "4":
			fmt.Println("bxc")
			bxc(operand)
		case "5":
			fmt.Println("out")
			result += strconv.Itoa(out(operand))
		case "6":
			fmt.Println("bdv")
			bdv(operand)
		case "7":
			fmt.Println("cdv")
			cdv(operand)
		default:
			panic("invalid instruction")
		}
	}
	return result
}

var (
	A                  = 0
	B                  = 0
	C                  = 0
	InstructionPointer = 0
)

func literalOperand(operand string) int {
	return parse.StringToInt(operand)
}

func comboOperand(operand string) int {
	switch operand {
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return A
	case "5":
		return B
	case "6":
		return C
	case "7":
		panic("reserved operand")
	}
	panic("invalid operand")
}

func incrementInstructionPointer() {
	InstructionPointer += 2
}

// The dv function performs division.
// The numerator is the value in the A register.
// The denominator is found by raising 2 to the power of the instruction's combo operand.
// (So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.)
func dv(operand string) int {
	return A / parse.IntPow(2, comboOperand(operand))
}

// The adv instruction (opcode 0) performs division.
// The numerator is the value in the A register.
// The denominator is found by raising 2 to the power of the instruction's combo operand.
// (So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.)
// The result of the division operation is truncated to an integer and then written to the A register.
func adv(operand string) {
	A = dv(operand)
	incrementInstructionPointer()
}

// The bxl instruction (opcode 1)
// calculates the bitwise XOR of register B and the instruction's literal operand,
// then stores the result in register B.
func bxl(operand string) {
	B = B ^ literalOperand(operand)
	incrementInstructionPointer()
}

// The bst instruction (opcode 2)
// calculates the value of its combo operand modulo 8
// (thereby keeping only its lowest 3 bits),
// then writes that value to the B register.
func bst(operand string) {
	B = comboOperand(operand) % 8
	incrementInstructionPointer()
}

// The jnz instruction (opcode 3)
// does nothing if the A register is 0. However,
// if the A register is not zero, it jumps by setting the instruction pointer to the value of its literal operand;
// if this instruction jumps, the instruction pointer is not increased by 2 after this instruction.
func jnz(operand string) {
	if A != 0 {
		InstructionPointer = literalOperand(operand)
	} else {
		incrementInstructionPointer()
	}
}

// The bxc instruction (opcode 4)
// calculates the bitwise XOR of register B and register C,
// then stores the result in register B.
// (For legacy reasons, this instruction reads an operand but ignores it.)
func bxc(operand string) {
	_ = operand
	B = B ^ C
	incrementInstructionPointer()
}

// The out instruction (opcode 5)
// calculates the value of its combo operand modulo 8,
// then outputs that value.
// (If a program outputs multiple values, they are separated by commas.)
func out(operand string) int {
	incrementInstructionPointer()
	return comboOperand(operand) % 8
}

// The bdv instruction (opcode 6)
// works exactly like the adv instruction except that the result is stored in the B register.
// (The numerator is still read from the A register.)
func bdv(operand string) {
	B = dv(operand)
	incrementInstructionPointer()
}

// The cdv instruction (opcode 7)
// works exactly like the adv instruction except that the result is stored in the C register.
// (The numerator is still read from the A register.)
func cdv(operand string) {
	C = dv(operand)
	incrementInstructionPointer()
}
