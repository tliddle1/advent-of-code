package main

import (
	"bufio"
	"fmt"
	"os"
)

func getFileData(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var data []string
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func solve(grid []string) (int, []Coordinate) {
	var s Coordinate
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'S' {
				s = Coordinate{x, y}
			}
		}
	}
	coords := []Coordinate{s}
	curX, curY := s.X, s.Y

	d := map[rune][]rune{
		'F': {'S', 'E'},
		'J': {'N', 'W'},
		'7': {'S', 'W'},
		'L': {'N', 'E'},
		'-': {'W', 'E'},
		'|': {'S', 'N'},
	}

	var comingFrom rune

	if 'S' == d[rune(grid[curY-1][curX])][0] {
		comingFrom = 'S'
		curY--
	} else if 'N' == d[rune(grid[curY+1][curX])][0] {
		comingFrom = 'N'
		curY++
	} else if 'E' == d[rune(grid[curY][curX-1])][0] {
		comingFrom = 'E'
		curX--
	}

	steps := 1

	for grid[curY][curX] != 'S' {
		cur := rune(grid[curY][curX])
		for k, v := range d {
			if k == cur {
				for _, dir := range v {
					if dir != comingFrom {
						coords = append(coords, Coordinate{curX, curY})
						switch dir {
						case 'N':
							curY--
							comingFrom = 'S'
						case 'S':
							curY++
							comingFrom = 'N'
						case 'E':
							curX++
							comingFrom = 'W'
						case 'W':
							curX--
							comingFrom = 'E'
						}
						break
					}
				}
				break
			}
		}
		steps++
	}

	return steps / 2, coords
}

type Coordinate struct {
	X, Y int
}

func main() {
	lines, err := getFileData("/Users/thomas/go/src/personal/advent-of-code/2023/day10/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	n := len(lines)
	m := len(lines[0])

	part1, coords := solve(lines)
	fmt.Println(part1)

	var interiorPoints, exteriorPoints []Coordinate

	for y := 0; y < n; y++ {
		var numBordersCrossed int
		var lastVerticalEdge rune

		for x := 0; x < m; x++ {
			if containsCoordinate(Coordinate{x, y}, coords) {
				if lines[y][x] == '|' || lines[y][x] == 'S' {
					numBordersCrossed++
				} else if lines[y][x] == '7' && lastVerticalEdge == 'L' {
					numBordersCrossed++
				} else if lines[y][x] == 'J' && lastVerticalEdge == 'F' {
					numBordersCrossed++
				}

				if lines[y][x] != '-' {
					lastVerticalEdge = rune(lines[y][x])
				}
			} else {
				if numBordersCrossed%2 == 1 {
					interiorPoints = append(interiorPoints, Coordinate{x, y})
				} else {
					exteriorPoints = append(exteriorPoints, Coordinate{x, y})
				}
			}
		}
	}

	fmt.Println(len(interiorPoints))
}

func containsCoordinate(c Coordinate, coords []Coordinate) bool {
	for _, coordinate := range coords {
		if coordinate == c {
			return true
		}
	}
	return false
}
