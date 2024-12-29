package parse

import (
	"bufio"
	"log"
	"os"
	"slices"
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

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	CheckError(err)
	return i
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

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

func Remove(numbers []int, idx int) []int {
	var newNumbers []int
	for i, n := range numbers {
		if i != idx {
			newNumbers = append(newNumbers, n)
		}
	}
	return newNumbers
}

func RemoveLeadingZeros(s string) string {
	for i := range s {
		if s[i] != '0' {
			return s[i:]
		}
	}
	return "0"
}

func IntPow(base int, exponent int) int {
	result := 1
	for range exponent {
		result *= base
	}
	return result
}

func RemoveDuplicates[T comparable](objs []T) []T {
	var newObjs []T
	for _, obj := range objs {
		if slices.Contains(newObjs, obj) {
			continue
		} else {
			newObjs = append(newObjs, obj)
		}
	}
	return newObjs
}
