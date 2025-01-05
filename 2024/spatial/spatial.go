package spatial

import (
	"errors"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

const (
	N Dir = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

type Dir uint8
type Pos [2]int

func (p Pos) Move(d Dir, distance int) Pos {
	rel, err := RelativePos(d)
	parse.CheckError(err)
	x, y := p[0], p[1]
	relX, relY := rel[0], rel[1]
	return Pos{x + relX*distance, y + relY*distance}
}

var ErrInvalidDir = errors.New("invalid dir")

func RelativePos(d Dir) (Pos, error) {
	switch d {
	case N:
		return [2]int{0, -1}, nil
	case NE:
		return [2]int{1, -1}, nil
	case E:
		return [2]int{1, 0}, nil
	case SE:
		return [2]int{1, 1}, nil
	case S:
		return [2]int{0, 1}, nil
	case SW:
		return [2]int{-1, 1}, nil
	case W:
		return [2]int{-1, 0}, nil
	case NW:
		return [2]int{-1, -1}, nil
	default:
		return [2]int{0, 0}, ErrInvalidDir
	}
}

type Grid[T any] [][]T

func New[T any](numCols, numRows int, defaultValue T) Grid[T] {
	var g [][]T
	for range numRows {
		var row []T
		for range numCols {
			row = append(row, defaultValue)
		}
		g = append(g, row)
	}
	return g
}

func (g Grid[T]) M() int {
	return len(g)
}

func (g Grid[T]) N() int {
	if len(g) == 0 {
		return 0
	}
	return len(g[0])
}

func (g Grid[T]) Set(pos Pos, value T) {
	x, y := pos[0], pos[1]
	g[y][x] = value
}

func (g Grid[T]) Get(pos Pos) T {
	x, y := pos[0], pos[1]
	return g[y][x]
}

func (g Grid[T]) Contains(pos Pos) bool {
	x, y := pos[0], pos[1]
	if x < 0 || y < 0 || x >= g.M() || y >= g.N() {
		return false
	}
	return true
}

func (g Grid[T]) OnEdge(pos Pos) bool {
	x, y := pos[0], pos[1]
	return y == 0 || x == 0 || y == g.N()-1 || x == g.M()-1
}
