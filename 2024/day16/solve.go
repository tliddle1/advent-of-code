package day16

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

type Orientation struct {
	X, Y, Dir int
}

type Maze [][]string

func (m Maze) getChar(char string) (int, int) {
	for i, line := range m {
		for j, c := range line {
			if c == char {
				return j, i
			}
		}
	}
	return -1, -1
}

const (
	north = 0
	east  = 1
	south = 2
	west  = 3
)

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	maze := getMaze(input)
	startX, startY := maze.getChar("S")
	cheapestCost, valid := getCheapestCost(maze, 0, startX, startY, east, 0, [][2]int{})
	if !valid {
		return -1
	}
	//crawlMaze(maze, 0, startX, startY, east, 0)
	//x, y := maze.getChar("E")
	//minCost := math.MaxInt32
	//for _, dir := range []int{north, east, south, west} {
	//	cost, ok := Cache[Orientation{x, y, dir}]
	//	if ok {
	//		minCost = int(math.Min(float64(minCost), float64(cost)))
	//	}
	//}
	return cheapestCost
}

func getCheapestCost(maze Maze, currentCost, currentX, currentY, currentDir, depth int, visited [][2]int) (int, bool) {
	if maze[currentY][currentX] == "#" {
		return -1, false
	}
	fmt.Println(currentCost, currentX, currentY, currentDir, depth)
	if depth > 10 {
		return -1, false
	}
	if slices.Contains(visited, [2]int{currentX, currentY}) {
		return -1, false
	} else {
		visited = append(visited, [2]int{currentX, currentY})
	}
	if maze[currentY][currentX] == "E" {
		return currentCost, true
	}
	if currentX < 0 || currentX >= len(maze[0]) || currentY < 0 || currentY >= len(maze) {
		return -1, false
	}
	var paths []Path
	for _, newDir := range []int{north, east, south, west} {
		// Don't turn 180 degrees
		if math.Abs(float64(newDir-currentDir)) == 2 {
			continue
		}
		var costInc int
		if newDir != currentDir {
			costInc = 1001
		} else {
			costInc = 1
		}
		xInc, yInc := getDirection(newDir)
		pathCost, pathValid := getCheapestCost(maze, currentCost+costInc, currentX+xInc, currentY+yInc, newDir, depth+1, visited)
		if pathValid {
			paths = append(paths, Path{pathCost, pathValid})
		}
	}
	minCost := math.MaxInt32
	for _, path := range paths {
		if path.valid {
			minCost = int(math.Min(float64(minCost), float64(path.cost)))
		}
	}
	if minCost == math.MaxInt32 {
		return -1, false
	}
	return minCost, true
}

//
//func crawlMaze(maze Maze, cost, x, y, dir, depth int) {
//	if maze[y][x] == "#" || depth > len(maze)*len(maze[0]) {
//		return
//	}
//	currentCost, ok := Cache[Orientation{x, y, dir}]
//	if !ok {
//		Cache[Orientation{x, y, dir}] = cost
//	} else if ok && currentCost < cost {
//		return
//	}
//	for _, newDir := range []int{north, east, south, west} {
//		// Don't turn 180 degrees
//		if math.Abs(float64(newDir-dir)) == 2 {
//			continue
//		}
//		var costInc int
//		if newDir != dir {
//			costInc = 1001
//		} else {
//			costInc = 1
//		}
//		xInc, yInc := getDirection(newDir)
//		crawlMaze(maze, cost+costInc, x+xInc, y+yInc, newDir, depth+1)
//	}
//}

type Path struct {
	cost  int
	valid bool
}

func getDirection(dir int) (int, int) {
	switch dir {
	case north:
		return 0, -1
	case east:
		return 1, 0
	case south:
		return 0, 1
	case west:
		return -1, 0
	}
	panic("Invalid direction")
}

func getMaze(input []string) Maze {
	var maze Maze
	for i, line := range input {
		_ = i
		maze = append(maze, strings.Split(line, ""))
	}
	return maze
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	var result int
	for i, line := range input {
		_ = i
		result += parse.StringToInt(line)
	}
	return result
}
