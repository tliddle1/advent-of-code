package day5

import (
	"testing"
)

const (
	expected1sample = 143
	expected1       = 4996
	expected2sample = 123
	expected2       = 6311
)

func TestPart1Sample(t *testing.T) {
	expectedResult := expected1sample
	result := part1("day5_sample.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart1(t *testing.T) {
	expectedResult := expected1
	result := part1("day5.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2Sample(t *testing.T) {
	expectedResult := expected2sample
	result := part2("day5_sample.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2(t *testing.T) {
	expectedResult := expected2
	result := part2("day5.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}
