package parse

import (
	"log"
	"strconv"
)

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
