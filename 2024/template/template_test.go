package template

import (
	"testing"
)

func TestPart1Sample(t *testing.T) {
	expectedResult := 0
	result := part1("template_sample.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart1(t *testing.T) {
	expectedResult := 0
	result := part1("template.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2Sample(t *testing.T) {
	expectedResult := 0
	result := part2("template_sample.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2(t *testing.T) {
	expectedResult := 0
	result := part2("template.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}
