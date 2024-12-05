package day5

import (
	"bufio"
	"os"
	"slices"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

type Rule []string

type Update struct {
	pages []string
}

func (this *Update) getMiddle() string {
	return this.pages[len(this.pages)/2]
}

type Input struct {
	rules   []Rule
	updates []Update
}

func part1(filePath string) int {
	file, err := os.Open(filePath)
	parse.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input Input
	ruleMode := true
	///////////////////////////////////
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			ruleMode = false
			continue
		}
		if ruleMode {
			input.rules = append(input.rules, strings.Split(line, "|"))
		} else {
			input.updates = append(input.updates, Update{strings.Split(line, ",")})
		}
	}
	///////////////////////////////////
	parse.CheckError(scanner.Err())
	return solve(input)
}

func solve(input Input) int {
	var result int
	for _, update := range input.updates {
		curFails := false
		for _, rule := range input.rules {
			if !fitsRule(rule, update) {
				curFails = true
			}
		}
		if curFails {
			continue
		}
		result += parse.StringToInt(update.getMiddle())
	}
	return result
}

func fitsRule(rule Rule, update Update) bool {
	if !slices.Contains(update.pages, rule[0]) || !slices.Contains(update.pages, rule[1]) {
		return true
	}
	for _, page := range update.pages {
		if page == rule[0] {
			return true
		}
		if page == rule[1] {
			return false
		}
	}
	return true
}

func part2(filePath string) int {
	file, err := os.Open(filePath)
	parse.CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input Input
	ruleMode := true
	///////////////////////////////////
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			ruleMode = false
			continue
		}
		if ruleMode {
			input.rules = append(input.rules, strings.Split(line, "|"))
		} else {
			input.updates = append(input.updates, Update{strings.Split(line, ",")})
		}
	}
	///////////////////////////////////
	parse.CheckError(scanner.Err())
	return solve2(input)
}

func solve2(input Input) int {
	var result int
	for _, update := range input.updates {
		curFails := false
		for _, rule := range input.rules {
			if !fitsRule(rule, update) {
				curFails = true
			}
		}
		if curFails {
			update.reorder(input.rules)
			result += parse.StringToInt(update.getMiddle())
		}
	}
	return result
}

func (this *Update) reorder(rules []Rule) {
	var ordered bool
	for !ordered {
		ordered = true
		for _, rule := range rules {
			if !fitsRule(rule, *this) {
				ordered = false
				a := rule[0]
				b := rule[1]
				idxA := slices.Index(this.pages, a)
				idxB := slices.Index(this.pages, b)
				this.pages[idxA], this.pages[idxB] = this.pages[idxB], this.pages[idxA]
			}
		}
	}
}
