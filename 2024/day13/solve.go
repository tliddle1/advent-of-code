package day13

import (
	"math"
	"strings"

	"github.com/tliddle1/advent-of-code/2024/parse"
)

type Button struct {
	cost       int
	xInc, yInc int
}
type ClawMachine struct {
	buttons        []Button
	prizeX, prizeY int
}

func (m ClawMachine) getCost() (cost int, won bool) {
	ax, ay := float64(m.buttons[0].xInc), float64(m.buttons[0].yInc)
	bx, by := float64(m.buttons[1].xInc), float64(m.buttons[1].yInc)
	prizeX, prizeY := float64(m.prizeX), float64(m.prizeY)
	b := ((ay * prizeX) - (ax * prizeY)) / ((ay * bx) - (ax * by))
	a := ((by * prizeX) - (bx * prizeY)) / ((by * ax) - (bx * ay))
	if math.Floor(a) == a && math.Floor(b) == b {
		return int(a*3 + b*1), true
	}
	return 0, false
}

func part1(filePath string) int {
	var result int
	input := parse.GetInput(filePath)
	clawMachines := parseClawMachines(input, false)
	for _, m := range clawMachines {
		cost, won := m.getCost()
		if won {
			result += cost
		}
	}
	return result
}

func parseClawMachines(input []string, part2 bool) []ClawMachine {
	var clawMachines []ClawMachine
	var buttonA = Button{cost: 3}
	var buttonB = Button{cost: 1}
	var prizeX, prizeY int
	for i, line := range input {
		switch i % 4 {
		case 0:
			//get Button A
			parts := strings.Split(line, " ")
			xIncStr := parts[2][2 : len(parts[2])-1]
			yIncStr := parts[3][2:len(parts[3])]
			buttonA.xInc = parse.StringToInt(xIncStr)
			buttonA.yInc = parse.StringToInt(yIncStr)
		case 1:
			//get Button B
			parts := strings.Split(line, " ")
			xIncStr := parts[2][2 : len(parts[2])-1]
			yIncStr := parts[3][2:len(parts[3])]
			buttonB.xInc = parse.StringToInt(xIncStr)
			buttonB.yInc = parse.StringToInt(yIncStr)
		case 2:
			//get Prize
			parts := strings.Split(line, " ")
			prizeXStr := parts[1][2 : len(parts[1])-1]
			prizeYStr := parts[2][2:len(parts[2])]
			prizeX = parse.StringToInt(prizeXStr)
			prizeY = parse.StringToInt(prizeYStr)
			if part2 {
				prizeX += 10000000000000
				prizeY += 10000000000000
			}
		default:
			clawMachines = append(clawMachines, ClawMachine{
				buttons: []Button{buttonA, buttonB},
				prizeX:  prizeX,
				prizeY:  prizeY,
			})
		}
	}
	clawMachines = append(clawMachines, ClawMachine{
		buttons: []Button{buttonA, buttonB},
		prizeX:  prizeX,
		prizeY:  prizeY,
	})
	return clawMachines
}

func part2(filePath string) int {
	var result int
	input := parse.GetInput(filePath)
	clawMachines := parseClawMachines(input, true)
	for _, m := range clawMachines {
		cost, won := m.getCost()
		if won {
			result += cost
		}
	}
	return result
}
