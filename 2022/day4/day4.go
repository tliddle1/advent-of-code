package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1(), part2())
}

func part1() int {
	file, err := os.Open("day4/day4.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	/////////////////////////////
	count := 0
	/////////////////////////////
	for {
		line, err := reader.ReadString('\n')
		line = trimSpace(line)
		/////////////////////////////
		assignments := strings.Split(line, ",")
		a1 := strings.Split(assignments[0], "-")
		a2 := strings.Split(assignments[1], "-")
		x1, _ := strconv.Atoi(a1[0])
		y1, _ := strconv.Atoi(a1[1])
		x2, _ := strconv.Atoi(a2[0])
		y2, _ := strconv.Atoi(a2[1])
		if (x1 <= x2 && y1 >= y2) || (x2 <= x1 && y2 >= y1) {
			count += 1
		}
		/////////////////////////////
		if err != nil {
			break
		} //EOF
	}

	if err != nil {
		fmt.Println("Error:", err)
	}
	/////////////////////////////
	return count
	/////////////////////////////
}

func part2() int {
	file, err := os.Open("day4/day4.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	/////////////////////////////
	count := 0
	/////////////////////////////
	for {
		line, err := reader.ReadString('\n')
		line = trimSpace(line)
		/////////////////////////////
		assignments := strings.Split(line, ",")
		a1 := strings.Split(assignments[0], "-")
		a2 := strings.Split(assignments[1], "-")
		x1, _ := strconv.Atoi(a1[0])
		y1, _ := strconv.Atoi(a1[1])
		x2, _ := strconv.Atoi(a2[0])
		y2, _ := strconv.Atoi(a2[1])
		if x1 <= y2 && x1 >= x2 || y1 <= y2 && y1 >= x2 || x2 <= y1 && x2 >= x1 || y2 <= y1 && y2 >= x1 {
			count += 1
		}
		/////////////////////////////
		if err != nil {
			break
		} //EOF
	}

	if err != nil {
		fmt.Println("Error:", err)
	}
	/////////////////////////////
	return count
	/////////////////////////////
}

func trimSpace(s string) string {
	return string(bytes.TrimSpace([]byte(s)))
}
