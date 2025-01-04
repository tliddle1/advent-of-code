package day17

import (
	"testing"
)

const (
	expectedPart1Sample = "4,6,3,5,6,3,5,2,1,0"
	expectedPart1       = "7,6,1,5,3,1,4,2,6"
	expectedPart2Sample = 0
	expectedPart2       = 0
)

func TestPart1Sample(t *testing.T) {
	expectedResult := expectedPart1Sample
	result := part1("sample_input.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart1(t *testing.T) {
	expectedResult := expectedPart1
	result := part1("input.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

// TODO
func TestPart2Sample(t *testing.T) {
	expectedResult := expectedPart2Sample
	result := part2("sample_input2.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2(t *testing.T) {
	expectedResult := expectedPart2
	result := part2("input.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}
