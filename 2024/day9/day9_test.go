package day9

import (
	"slices"
	"strings"
	"testing"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

const (
	expectedPart1Sample = 1928
	expectedPart1       = 6385338159127
	expectedPart2Sample = 2858
	expectedPart2       = 6415163624282
)

func TestOriginalAllocation(t *testing.T) {
	expectedResult := strings.Split("00...111...2...333.44.5555.6666.777.888899", "")
	input := parse.GetInput("day9_sample.txt")[0]
	result := OriginalAllocation(input)

	if !slices.Equal(result, expectedResult) {
		t.Errorf("Expected %v, \n\t\t\t\t  got %v", expectedResult, result)
	}
}

func TestNewAllocation(t *testing.T) {
	expectedResult := strings.Split("0099811188827773336446555566", "")
	result := NewAllocation(strings.Split("00...111...2...333.44.5555.6666.777.888899", ""))
	if !slices.Equal(expectedResult, result) {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestCheckSum(t *testing.T) {
	str := "0099811188827773336446555566"
	result := checkSum(strings.Split(str, ""))
	if result != expectedPart1Sample {
		t.Errorf("Expected %d, got %d", expectedPart1, result)
	}
}

func TestPart1Sample(t *testing.T) {
	expectedResult := expectedPart1Sample
	result := part1("day9_sample.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart1(t *testing.T) {
	expectedResult := expectedPart1
	result := part1("day9.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2Sample(t *testing.T) {
	expectedResult := expectedPart2Sample
	result := part2("day9_sample.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}

func TestPart2(t *testing.T) {
	expectedResult := expectedPart2
	result := part2("day9.txt")
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}
