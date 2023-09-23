package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("day1/day1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	maxCal := 0
	curr := 0
	var calories []int
	for {
		line, err := reader.ReadString('\n')
		line = trimSpace(line)
		if num, err := strconv.Atoi(line); err == nil {
			curr += num
		} else if line == "" {
			maxCal = max(maxCal, curr)
			calories = append(calories, curr)
			curr = 0
		} else {
			fmt.Printf("Not an integer: %v\n", line)
		}
		if err != nil {
			break
		}
	}

	if err != nil {
		fmt.Println("Error:", err)
	}

	sort.Ints(calories)
	length := len(calories)
	largestThree := calories[length-3:]

	fmt.Print(maxCal, sum(largestThree))
}

func trimSpace(s string) string {
	return string(bytes.TrimSpace([]byte(s)))
}

func sum(nums []int) int {
	total := 0
	for _, i := range nums {
		total += i
	}
	return total
}
