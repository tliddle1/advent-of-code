package day15

import (
	"testing"
)

const (
	expectedPart1Sample = 10092
	expectedPart1       = 1527563
	expectedPart2Sample = 9021
	expectedPart2       = 1521635
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
