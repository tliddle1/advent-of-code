package day14

import (
	"fmt"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

func part1(filePath string, wide, tall int) int {
	input := parse.GetInput(filePath)
	var robots []*Robot
	iterations := 100
	for i, line := range input {
		_ = i
		robot := parseRobot(line)
		robots = append(robots, robot)
	}
	quadrantCounts := make(map[int]int)
	coordinates := make(map[Coordinate]int)
	for _, robot := range robots {
		quadrant, x, y := calculateQuadrant(*robot, wide, tall, iterations)
		coordinates[Coordinate{x, y}]++
		if quadrant != 0 {
			quadrantCounts[quadrant]++
		}
	}
	result := 1
	for _, quadrant := range quadrantCounts {
		result *= quadrant
	}
	return result
}

func printRoom(tall int, wide int, coordinates map[Coordinate]int) {
	for j := range tall {
		row := ""
		for i := range wide {
			if coordinates[Coordinate{i, j}] > 0 {
				row += "#"
			} else {
				row += "."
			}
		}
		fmt.Println(row)
	}
}

func (r *Robot) Move(wide, tall int) {
	r.P[0] += r.V[0]
	r.P[1] += r.V[1]
	r.P[0] = r.P[0] % wide
	if r.P[0] < 0 {
		r.P[0] += wide
	}
	r.P[1] = r.P[1] % tall
	if r.P[1] < 0 {
		r.P[1] += tall
	}
}

func calculateQuadrant(robot Robot, wide int, tall int, iterations int) (int, int, int) {
	x := robot.P[0]
	y := robot.P[1]
	x += iterations * robot.V[0]
	y += iterations * robot.V[1]
	x = x % wide
	if x < 0 {
		x += wide
	}
	y = y % tall
	if y < 0 {
		y += tall
	}
	return getQuadrant(x, y, wide, tall), x, y
}

func getQuadrant(x int, y int, wide int, tall int) int {
	midX := wide / 2
	midY := tall / 2
	if x == midX || y == midY {
		return 0
	}
	if x < midX {
		if y < midY {
			return 1
		} else {
			return 2
		}
	} else {
		if y < midY {
			return 3
		} else {
			return 4
		}
	}
}

type Coordinate [2]int
type Robot struct {
	P, V Coordinate
}

func parseRobot(line string) *Robot {
	pv := strings.Split(line, " ")
	p := strings.Split(pv[0][2:], ",")
	v := strings.Split(pv[1][2:], ",")
	return &Robot{
		P: Coordinate{parse.StringToInt(p[0]), parse.StringToInt(p[1])},
		V: Coordinate{parse.StringToInt(v[0]), parse.StringToInt(v[1])},
	}
}

func part2(filePath string, wide, tall int) int {
	input := parse.GetInput(filePath)
	var robots []*Robot
	for i, line := range input {
		_ = i
		robot := parseRobot(line)
		robots = append(robots, robot)
	}
	for iterations := 1; iterations <= 100_000; iterations++ {
		coordinates := make(map[Coordinate]int)
		for _, robot := range robots {
			robot.Move(wide, tall)
			coordinates[Coordinate{robot.P[0], robot.P[1]}]++

		}
		if hasTree(coordinates, wide, tall) {
			//printRoom(tall, wide, coordinates)
			return iterations
		}

	}
	return 0
}

func hasTree(coordinates map[Coordinate]int, wide int, tall int) bool {
	for j := range tall {
		row := ""
		for i := range wide {
			if coordinates[Coordinate{i, j}] > 0 {
				row += "#"
			} else {
				row += "."
			}
			//row += " "
		}
		if strings.Contains(row, "##########") {
			return true
		}
	}
	return false
}
