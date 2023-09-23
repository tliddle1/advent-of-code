package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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
	for {
		line, err := reader.ReadString('\n') // Read until a newline character is encountered.
		line = trimSpace(line)
		if num, err := strconv.Atoi(line); err == nil {
			curr += num
		} else if line == "" {
			maxCal = max(maxCal, curr)
			curr = 0
		} else {
			break
			fmt.Printf("Not an integer: %v\n", line)
		}
		if err != nil {
			break // Reached end of file
		}
	}

	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Print(maxCal)
}

func trimSpace(s string) string {
	// Remove leading and trailing whitespace.
	return string(bytes.TrimSpace([]byte(s)))
}
