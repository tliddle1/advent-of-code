package parse

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetInput(filePath string) []string {
	file, err := os.Open(filePath)
	CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	///////////////////////////////////
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	CheckError(scanner.Err())
	return input
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//////////////////////////////////// Ints ////////////////////////////////////

func IntPow(base int, exponent int) int {
	result := 1
	for range exponent {
		result *= base
	}
	return result
}

func IntAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

//////////////////////////////////// Strings and Ints ////////////////////////////////////

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	CheckError(err)
	return i
}

func IntToString(i int) string { return strconv.Itoa(i) }

func RemoveLeadingZeros(s string) string {
	for i := range s {
		if s[i] != '0' {
			return s[i:]
		}
	}
	return "0"
}

//////////////////////////////////// Int Slices ////////////////////////////////////

func AllBetweenAB(nums []int, a, b int) bool {
	for _, diff := range nums {
		if diff < a || diff > b {
			return false
		}
	}
	return true
}

func AllPositive(nums []int) bool {
	for _, diff := range nums {
		if diff < 1 {
			return false
		}
	}
	return true
}

func AllNegative(nums []int) bool {
	for _, diff := range nums {
		if diff > -1 {
			return false
		}
	}
	return true
}
