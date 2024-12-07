package day6

import (
	"testing"
)

const (
	expectedPart1Sample = 41
	expectedPart1       = 5145
	expectedPart2Sample = 6
	expectedPart2       = 1523
)

func TestPart1Sample(t *testing.T) {
	expectedResult := expectedPart1Sample
	result := Part1("day6_sample.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart1(t *testing.T) {
	expectedResult := expectedPart1
	result := Part1("day6.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2Sample(t *testing.T) {
	expectedResult := expectedPart2Sample
	result := Part2("day6_sample.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2(t *testing.T) {
	expectedResult := expectedPart2
	result := Part2("day6.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}
