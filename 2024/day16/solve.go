package day16

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
	return 0
}

type Path struct {
	cost  int
	valid bool
}

func part2(filePath string) int {
	return 0
}
