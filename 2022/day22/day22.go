package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	day := 22
	fmt.Println(part1(day), part2(day))
}

func part1(day int) int {
	filename := fmt.Sprintf("day%d/day%d.txt", day, day)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	/////////////////////////////
	/////////////////////////////
	for {
		line, err := reader.ReadString('\n')
		line = trimSpace(line)
		/////////////////////////////
		/////////////////////////////
		if err != nil {
			break
		} //EOF
	}

	if err != nil {
		fmt.Println("Error:", err)
	}
	/////////////////////////////
	return 0
	/////////////////////////////
}

func part2(day int) int {
	filename := fmt.Sprintf("day%d/day%d.txt", day, day)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	/////////////////////////////
	/////////////////////////////
	for {
		line, err := reader.ReadString('\n')
		line = trimSpace(line)
		/////////////////////////////
		/////////////////////////////
		if err != nil {
			break
		} //EOF
	}

	if err != nil {
		fmt.Println("Error:", err)
	}
	/////////////////////////////
	return 0
	/////////////////////////////
}

func trimSpace(s string) string {
	return string(bytes.TrimSpace([]byte(s)))
}
