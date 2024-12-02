package day1

import (
	"testing"
)

func TestPart1Sample(t *testing.T) {
	expectedResult := 11
	result := part1("day1_sample.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart1(t *testing.T) {
	expectedResult := 1765812
	result := part1("day1.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2Sample(t *testing.T) {
	expectedResult := 31
	result := part2("day1_sample.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2(t *testing.T) {
	expectedResult := 20520794
	result := part2("day1.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}
