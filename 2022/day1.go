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
	// Open the text file for reading. Replace "input.txt" with the actual file path.
	file, err := os.Open("day1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a bufio.Reader to read from the file.
	reader := bufio.NewReader(file)

	// Read and print each line from the file.
	maxCal := 0
	curr := 0
	var calories []int
	for {
		line, err := reader.ReadString('\n') // Read until a newline character is encountered.
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
			break // Reached end of file
		}
	}

	if err != nil {
		fmt.Println("Error:", err)
	}

	// Sort the slice in ascending order.
	sort.Ints(calories)

	// Find and print the largest three integers.
	length := len(calories)
	largestThree := calories[length-3:]

	fmt.Print(maxCal, sum(largestThree))
}

func trimSpace(s string) string {
	// Remove leading and trailing whitespace.
	return string(bytes.TrimSpace([]byte(s)))
}

func sum(nums []int) int {
	total := 0
	for _, i := range nums {
		total += i
	}
	return total
}
