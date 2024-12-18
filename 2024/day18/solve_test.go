package day18

import (
	"testing"
)

const (
	expectedPart1Sample = 22
	expectedPart1       = 324
	expectedPart2Sample = "6,1"
	expectedPart2       = "46,23"
)

func TestPart1Sample(t *testing.T) {
	expectedResult := expectedPart1Sample
	result := part1("sample_input.txt", 7, 12)
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart1(t *testing.T) {
	expectedResult := expectedPart1
	result := part1("input.txt", 71, 1024)
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2Sample(t *testing.T) {
	expectedResult := expectedPart2Sample
	result := part2("sample_input.txt", 7, 12)
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2(t *testing.T) {
	expectedResult := expectedPart2
	result := part2("input.txt", 71, 1024)
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}
