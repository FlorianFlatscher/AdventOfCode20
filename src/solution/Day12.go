package solution

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/constants"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	regexParseIntruction = regexp.MustCompile("([A-Z])([0-9]*)")
)

type Day12 struct{}

type instruction struct {
	command   rune
	parameter int
}

type direction int

const (
	NORTH = 0
	EAST  = 1
	SOUTH = 2
	WEST  = 3
)

func (d Day12) Calc() string {
	puzzleInput := input.ReadInputFile(12)
	var instructions []instruction
	for _, line := range strings.Split(puzzleInput, constants.LineSeparator) {
		data := regexParseIntruction.FindStringSubmatch(line)
		parameter, err := strconv.Atoi(data[2])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instruction{
			command:   rune(data[1][0]),
			parameter: parameter,
		})
	}

	return fmt.Sprintf("Mannhattan: %v\nWaypoint: %v\n", calcManhattanDistance(instructions), calcManhattanDistanceWaypoint(instructions))
}

func calcManhattanDistance(instructions []instruction) int {
	var x int
	var y int
	d := direction(EAST)

	for _, instruction := range instructions {
		switch instruction.command {
		case 'N':
			y -= instruction.parameter
		case 'S':
			y += instruction.parameter
		case 'E':
			x += instruction.parameter
		case 'W':
			x -= instruction.parameter
		case 'L':
			d -= direction(instruction.parameter / 90)
			d %= 4
			if d < 0 {
				d = 4 + d
			}
		case 'R':
			d += direction(instruction.parameter / 90)
			d %= 4
		case 'F':
			switch d {
			case NORTH:
				y -= instruction.parameter
			case SOUTH:
				y += instruction.parameter
			case EAST:
				x += instruction.parameter
			case WEST:
				x -= instruction.parameter
			}
		default:
			panic("Not a valid command: " + fmt.Sprint(instruction))
		}
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func calcManhattanDistanceWaypoint(instructions []instruction) int {
	x := 10.
	y := -1.
	shipX := 0.
	shipY := 0.

	for _, instruction := range instructions {
		switch instruction.command {
		case 'N':
			y -= float64(instruction.parameter)
		case 'S':
			y += float64(instruction.parameter)
		case 'E':
			x += float64(instruction.parameter)
		case 'W':
			x -= float64(instruction.parameter)
		case 'L':
			z := math.Sqrt(x*x + y*y)
			currentAngle := math.Atan2(y, x)
			currentAngle -= float64(instruction.parameter) * ((2. * math.Pi) / 360.)
			x = z * math.Cos(currentAngle)
			y = z * math.Sin(currentAngle)
		case 'R':
			z := math.Sqrt(x*x + y*y)
			currentAngle := math.Atan2(y, x)
			currentAngle += float64(instruction.parameter) * ((2. * math.Pi) / 360.)
			x = z * math.Cos(currentAngle)
			y = z * math.Sin(currentAngle)
		case 'F':
			shipX += x * float64(instruction.parameter)
			shipY += y * float64(instruction.parameter)
		default:
			panic("Not a valid command: " + fmt.Sprint(instruction))
		}
	}

	return int(math.Round(math.Abs(shipX) + math.Abs(shipY)))
}

//17648 to low
//77366 to low
//76780 to low
//178986
