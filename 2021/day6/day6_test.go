package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func NumFish(input []int, numDays int) int {
	fish := map[int]int{}
	for i := 0; i <= 8; i++ {
		fish[i] = 0
	}
	for _, timer := range input {
		fish[timer] += 1
	}

	for d := 0; d < numDays; d++ {
		fish = iterate(fish)
	}

	result := 0
	for _, num := range fish {
		result += num
	}
	return result
}

func iterate(fish map[int]int) map[int]int {
	numNewFish := fish[0]
	for i := 0; i < 8; i++ {
		fish[i] = fish[i+1]
	}
	fish[6] += numNewFish
	fish[8] = numNewFish
	return fish
}

func TestAddTable(t *testing.T) {
	TestInput := []int{3, 4, 3, 1, 2}

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	RealInput := csvToIntSlice(records)

	testCases := []struct {
		name     string
		input    []int
		days     int
		expected int
	}{
		{"Test Input, 18 days", TestInput, 18, 26},
		{"Test Input, 80 days", TestInput, 80, 5934},
		{"Real Input, 80 days", RealInput, 80, 353274},
		{"Test Input, 256 days", TestInput, 256, 26984457539},
		{"Real Input, 256 days", RealInput, 256, 1609314870967},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := NumFish(tc.input, tc.days)
			if result != tc.expected {
				t.Errorf("After %d days expected %d fish but got %d", tc.days, tc.expected, result)
			}
		})
	}
}

func csvToIntSlice(records [][]string) []int {
	var intSlice []int

	for _, record := range records {
		for _, value := range record {
			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error converting to int:", err)
				continue
			}
			intSlice = append(intSlice, num)
		}
	}
	return intSlice
}
