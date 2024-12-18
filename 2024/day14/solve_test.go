package day14

import (
	"testing"
)

const (
	expectedPart1Sample = 12
	expectedPart1       = 218619324
	expectedPart2       = 6446
)

func TestPart1Sample(t *testing.T) {
	expectedResult := expectedPart1Sample
	wide := 11
	tall := 7
	result := part1("sample_input.txt", wide, tall)
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart1(t *testing.T) {
	expectedResult := expectedPart1
	wide := 101
	tall := 103
	result := part1("input.txt", wide, tall)
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2(t *testing.T) {
	expectedResult := expectedPart2
	wide := 101
	tall := 103
	result := part2("input.txt", wide, tall)
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}
