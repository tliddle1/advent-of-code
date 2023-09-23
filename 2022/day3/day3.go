package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("day3/day3.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	/////////////////////////////
	total := 0
	/////////////////////////////
	for {
		line, err := reader.ReadString('\n')
		line = trimSpace(line)
		/////////////////////////////
		midpoint := len(line) / 2

		firstHalf := line[:midpoint]
		secondHalf := line[midpoint:]
		c := commonCharacter(firstHalf, secondHalf)
		total += priority(c)
		/////////////////////////////
		if err != nil {
			break
		} //EOF
	}

	if err != nil {
		fmt.Println("Error:", err)
	}
	/////////////////////////////
	fmt.Println(total)
	/////////////////////////////
}

func trimSpace(s string) string {
	return string(bytes.TrimSpace([]byte(s)))
}

func commonCharacter(s1, s2 string) string {
	// Create a map to track characters in the second string.
	charMap := make(map[rune]bool)
	for _, char := range s2 {
		charMap[char] = true
	}
	for _, char := range s1 {
		if charMap[char] {
			return string(char)
		}
	}
	return ""
}

func priority(s string) int {
	num := int(rune(s[0]))
	if num <= 90 {
		return num - 38
	} else {
		return num - 96
	}
}
