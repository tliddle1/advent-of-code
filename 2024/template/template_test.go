package template

import (
	"testing"
)

const (
	expectedPart1Sample = 0
	expectedPart1       = 0
	expectedPart2Sample = 0
	expectedPart2       = 0
)

func TestPart1Sample(t *testing.T) {
	expectedResult := expectedPart1Sample
	result := part1("template_sample.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart1(t *testing.T) {
	expectedResult := expectedPart1
	result := part1("template.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2Sample(t *testing.T) {
	expectedResult := expectedPart2Sample
	result := part2("template_sample.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2(t *testing.T) {
	expectedResult := expectedPart2
	result := part2("template.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}
