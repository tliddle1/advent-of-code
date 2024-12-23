package template

import (
	"slices"
	"sort"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	connections := make(map[string][]string)
	for _, line := range input {
		computers := strings.Split(line, "-")
		a, b := computers[0], computers[1]
		connections[a] = append(connections[a], b)
		connections[b] = append(connections[b], a)
	}
	var result []string
	for comp, connectedComp := range connections {
		if strings.HasPrefix(comp, "t") {
			result = append(result, numLANs(comp, connectedComp, connections)...)
		}
	}
	return len(removeDuplicates(result))
}

func removeDuplicates(objs []string) []string {
	var newObjs []string
	for _, obj := range objs {
		if slices.Contains(newObjs, obj) {
			continue
		} else {
			newObjs = append(newObjs, obj)
		}
	}
	return newObjs
}

func numLANs(comp string, connectedComps []string, connections map[string][]string) []string {
	var result []string
	for i := 0; i < len(connectedComps)-1; i++ {
		a := connectedComps[i]
		for j := i + 1; j < len(connectedComps); j++ {
			b := connectedComps[j]
			if slices.Contains(connections[a], b) {
				arr := []string{comp, a, b}
				sort.Strings(arr)
				result = append(result, strings.Join(arr, ","))
			}
		}
	}
	return result
}

func part2(filePath string) string {
	input := parse.GetInput(filePath)
	_ = input
	return ""
}
