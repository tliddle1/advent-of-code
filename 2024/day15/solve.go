package day15

import (
	"fmt"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

type Room [][]string

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	room, robotY, robotX := getRoom(input)
	instructions := getInstructions(input)
	for _, dir := range instructions {
		if canMove(room, robotY, robotX, dir) {
			y, x := nextAvailable(room, robotY, robotX, dir)
			room[y][x] = "O"
			room[robotY][robotX] = "."
			direction := getDirection(dir)
			robotY += direction[0]
			robotX += direction[1]
			room[robotY][robotX] = "@"
		}
	}
	var result int
	for y, row := range room {
		for x, cell := range row {
			if cell == "O" {
				result += calculateGPS(y, x)
			}
		}
	}
	return result
}

func printRoom(room Room) {
	for _, row := range room {
		fmt.Println(row)
	}
}

func calculateGPS(y int, x int) int {
	return 100*y + x
}

func canMove(room Room, y int, x int, dir string) bool {
	nextX, nextY := nextAvailable(room, y, x, dir)
	return !(nextY == -1 || nextX == -1)
}

func nextAvailable(room Room, y int, x int, dir string) (int, int) {
	for room[y][x] != "#" {
		if room[y][x] == "." {
			return y, x
		}
		y += getDirection(dir)[0]
		x += getDirection(dir)[1]
	}
	return -1, -1
}

func getDirection(dir string) []int {
	switch dir {
	case "^":
		return []int{-1, 0}
	case ">":
		return []int{0, 1}
	case "v":
		return []int{1, 0}
	case "<":
		return []int{0, -1}
	}
	panic("invalid direction")
}

func (r Room) checkDir(dir string, y, x int) string {
	direction := getDirection(dir)
	return r[y+direction[0]][x+direction[1]]
}

func getInstructions(input []string) []string {
	instructions := ""
	reachedInstructions := false
	for _, line := range input {
		if line == "" {
			reachedInstructions = true
			continue
		}
		if reachedInstructions {
			instructions += line
		}
	}
	return strings.Split(instructions, "")
}

func getRoom(input []string) (Room, int, int) {
	var room Room
	robotY, robotX := 0, 0
	for i, line := range input {
		if line == "" {
			return room, robotY, robotX
		}
		x := strings.Index(line, "@")
		if x > 0 {
			robotY = i
			robotX = x
		}
		row := strings.Split(line, "")
		room = append(room, row)

	}
	return room, robotY, robotX
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	room, robotY, robotX := getRoom2(input)
	instructions := getInstructions(input)
	for _, dir := range instructions {
		//printRoom(room)
		//fmt.Println(dir)
		direction := getDirection(dir)
		if room.checkDir(dir, robotY, robotX) == "#" {
			continue
		} else if room.checkDir(dir, robotY, robotX) == "." {
			room[robotY][robotX] = "."
			robotY += direction[0]
			robotX += direction[1]
			room[robotY][robotX] = "@"
		} else {
			boxes, canPush := canPushBoxInFrontOfMe(room, robotY, robotX, dir)
			if canPush {
				boxes = parse.RemoveDuplicates[[2]int](boxes)
				for _, box := range boxes {
					room[box[0]+direction[0]][box[1]+direction[1]] = room[box[0]][box[1]]
					room[box[0]][box[1]] = "."
				}
				room[robotY][robotX] = "."
				robotY += direction[0]
				robotX += direction[1]
				room[robotY][robotX] = "@"
			}
		}
	}
	//printRoom(room)
	var result int
	for y, row := range room {
		for x, cell := range row {
			if cell == "[" {
				result += calculateGPS(y, x)
			}
		}
	}
	return result
}

func canPushBoxInFrontOfMe(room Room, y int, x int, dir string) ([][2]int, bool) {
	direction := getDirection(dir)
	nextY, nextX := y+direction[0], x+direction[1]
	if room[nextY][nextX] == "#" {
		return [][2]int{}, false
	} else if room[nextY][nextX] == "." {
		return [][2]int{}, true
	} else if room[nextY][nextX] == "[" || room[nextY][nextX] == "]" {
		if dir == ">" || dir == "<" {
			boxes, canPush := canPushBoxInFrontOfMe(room, nextY, nextX, dir)
			if canPush {
				return append(boxes, [2]int{nextY, nextX}), true
			} else {
				return [][2]int{}, false
			}
		}
		boxY, boxX := y+direction[0], x+direction[1]
		pairY, pairX := getPair(boxY, boxX, room)
		boxesA, canPushA := canPushBoxInFrontOfMe(room, boxY, boxX, dir)
		boxesB, canPushB := canPushBoxInFrontOfMe(room, pairY, pairX, dir)
		if canPushA && canPushB {
			return append(append(boxesA, boxesB...), [][2]int{{boxY, boxX}, {pairY, pairX}}...), true
		}
		return [][2]int{}, false
	}
	panic("invalid object")
}

func getPair(y int, x int, room Room) (int, int) {
	if room[y][x] == "[" {
		return y, x + 1
	} else if room[y][x] == "]" {
		return y, x - 1
	}
	panic("invalid object")
}

func getRoom2(input []string) (Room, int, int) {
	var room Room
	robotY, robotX := 0, 0
	for i, line := range input {
		if line == "" {
			return room, robotY, robotX
		}
		var row string
		splitLine := strings.Split(line, "")
		for _, obj := range splitLine {
			row += widerObj(obj)
		}
		x := strings.Index(row, "@")
		if x > 0 {
			robotX = x
			robotY = i
		}
		room = append(room, strings.Split(row, ""))
	}
	return room, robotX, robotY
}

func widerObj(obj string) string {
	switch obj {
	case "#":
		return "##"
	case ".":
		return ".."
	case "O":
		return "[]"
	case "@":
		return "@."
	}
	panic("invalid object")
}
