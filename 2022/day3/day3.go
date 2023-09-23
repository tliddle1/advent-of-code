package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println(part1(), part2())
}

func part1() int {
	file, err := os.Open("day3/day3.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return 0
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
	return total
	/////////////////////////////
}

func part2() int {
	file, err := os.Open("day3/day3.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	/////////////////////////////
	total := 0
	/////////////////////////////
	for {
		line1, err := reader.ReadString('\n')
		line1 = trimSpace(line1)
		line2, err := reader.ReadString('\n')
		line2 = trimSpace(line2)
		line3, err := reader.ReadString('\n')
		line3 = trimSpace(line3)
		/////////////////////////////

		c := findCommonCharacters(line1, line2, line3)[0]
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
	return total
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

func findCommonCharacters(str1, str2, str3 string) []string {
	commonChars := []string{}

	// Convert the strings to sets of characters
	set1 := make(map[rune]struct{})
	set2 := make(map[rune]struct{})
	set3 := make(map[rune]struct{})

	for _, char := range str1 {
		set1[char] = struct{}{}
	}
	for _, char := range str2 {
		set2[char] = struct{}{}
	}
	for _, char := range str3 {
		set3[char] = struct{}{}
	}

	// Find common characters
	for char := range set1 {
		if _, exists2 := set2[char]; exists2 {
			if _, exists3 := set3[char]; exists3 {
				commonChars = append(commonChars, string(char))
			}
		}
	}

	return commonChars
}
