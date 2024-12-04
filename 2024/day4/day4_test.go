package day4

import (
	"testing"
)

const (
	expected1sample = 18
	expected1       = 2530
	expected2sample = 9
	expected2       = 1921
)

func TestPart1Sample(t *testing.T) {
	expectedResult := expected1sample
	result := part1("day4_sample.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart1(t *testing.T) {
	expectedResult := expected1
	result := part1("day4.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2Sample(t *testing.T) {
	expectedResult := expected2sample
	result := part2("day4_sample.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2(t *testing.T) {
	expectedResult := expected2
	result := part2("day4.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}
