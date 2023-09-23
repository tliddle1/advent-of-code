package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("day2/day2.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	scores1 := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}
	scores2 := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}
	totalScore1 := 0
	totalScore2 := 0
	for {
		line, err := reader.ReadString('\n')
		line = trimSpace(line)
		/////////////////////////////
		totalScore1 += scores1[line]
		totalScore2 += scores2[line]
		/////////////////////////////
		if err != nil {
			break
		} //EOF
	}

	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(totalScore1, totalScore2)
}

func trimSpace(s string) string {
	return string(bytes.TrimSpace([]byte(s)))
}
