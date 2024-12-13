package day12

import (
	"slices"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

const (
	north = iota
	east
	south
	west
)

type Grid [][]string

func (g Grid) exploreRegion(coordinate Coordinate, region *Region) {
	for _, dir := range []int{north, east, south, west} {
		neighbor, inGrid := coordinate.GetNeighbor(dir, g)
		if inGrid {
			if !region.Contains(neighbor) && g.at(neighbor) == region.type_ {
				region.Add(neighbor)
				g.exploreRegion(neighbor, region)
			}
			if g.at(neighbor) != region.type_ {
				region.crossings = append(region.crossings, Crossing{
					dir:        dir,
					coordinate: coordinate,
				})
				region.perimeter++
			}
		} else {
			region.crossings = append(region.crossings, Crossing{
				dir:        dir,
				coordinate: coordinate,
			})
			region.perimeter++
		}
	}
}

func (g Grid) at(coordinate Coordinate) string {
	return g[coordinate[0]][coordinate[1]]
}

type Coordinate [2]int

func (this Coordinate) GetNeighbor(dir int, grid Grid) (coordinate Coordinate, ok bool) {
	x, y := this[0], this[1]
	var newX, newY int
	switch dir {
	case north:
		newX, newY = x, y-1
	case east:
		newX, newY = x+1, y
	case south:
		newX, newY = x, y+1
	case west:
		newX, newY = x-1, y
	}
	coordinate = Coordinate{newX, newY}
	ok = newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0])
	return coordinate, ok
}

func (this Coordinate) IsInterior(r Region, g Grid) bool {
	for _, dir := range []int{north, east, south, west} {
		neighbor, _ := this.GetNeighbor(dir, g)
		if !slices.Contains(r.plots, neighbor) {
			return false
		}
	}
	return true
}

func (this Coordinate) neighborInRegion(dir int, r *Region, g Grid) bool {
	neighbor, _ := this.GetNeighbor(dir, g)
	return slices.Contains(r.plots, neighbor)
}

type Crossing struct {
	dir        int
	coordinate Coordinate
}

type Region struct {
	type_     string
	plots     []Coordinate
	perimeter int
	crossings []Crossing
}

func (r *Region) Contains(neighbor Coordinate) bool {
	return slices.Contains(r.plots, neighbor)
}

func (r *Region) Add(neighbor Coordinate) {
	r.plots = append(r.plots, neighbor)
}

func (r *Region) getPrice1() int {
	p := r.getPerimeter()
	a := r.getArea()
	return p * a
}

func (r *Region) getPerimeter() int {
	return r.perimeter
}

func (r *Region) getArea() int {
	return len(r.plots)
}

type Corner Coordinate

func (r *Region) numSides() int {
	var result int
	corners := make(map[Corner]struct{})
	for _, plot := range r.plots {
		corners[Corner{plot[0], plot[1]}] = struct{}{}
		corners[Corner{plot[0], plot[1] + 1}] = struct{}{}
		corners[Corner{plot[0] + 1, plot[1]}] = struct{}{}
		corners[Corner{plot[0] + 1, plot[1] + 1}] = struct{}{}
	}
	for corner := range corners {
		result += numRegionCorners(corner, r.plots)
	}
	return result
}

func numRegionCorners(corner Corner, plots []Coordinate) int {
	fourPlots := []Coordinate{
		{corner[0], corner[1]},         // NW
		{corner[0], corner[1] - 1},     // NE
		{corner[0] - 1, corner[1]},     // SW
		{corner[0] - 1, corner[1] - 1}, // SE
	}
	var count int
	var containedPlots []Coordinate
	for _, plot := range fourPlots {
		if slices.Contains(plots, plot) {
			containedPlots = append(containedPlots, plot)
			count++
		}
	}
	if count == 1 || count == 3 {
		return 1
	} else {
		if count == 2 {
			if !(containedPlots[0][0] == containedPlots[1][0] || containedPlots[0][1] == containedPlots[1][1]) {
				return 2
			}
		}
		return 0
	}
}

func (r *Region) getPrice2() int {
	p := r.numSides()
	a := r.getArea()
	return p * a
}

func part1(filePath string) int {
	input := parse.GetInput(filePath)
	grid := newGrid(input)
	var visited []Coordinate
	var result int
	for i := range grid {
		for j := range grid[i] {
			cur := Coordinate{i, j}
			if !slices.Contains(visited, cur) {
				region := &Region{
					type_: grid.at(cur),
					plots: []Coordinate{cur},
				}
				grid.exploreRegion(cur, region)
				visited = append(visited, region.plots...)
				result += region.getPrice1()
			}
		}
	}
	return result
}

func newGrid(input []string) Grid {
	var grid Grid
	for _, line := range input {
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}

func part2(filePath string) int {
	input := parse.GetInput(filePath)
	grid := newGrid(input)
	var visited []Coordinate
	var result int
	for i := range grid {
		for j := range grid[i] {
			cur := Coordinate{i, j}
			if !slices.Contains(visited, cur) {
				region := &Region{
					type_: grid.at(cur),
					plots: []Coordinate{cur},
				}
				grid.exploreRegion(cur, region)
				visited = append(visited, region.plots...)
				result += region.getPrice2()
			}
		}
	}
	return result
}
