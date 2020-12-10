package days

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"strings"
)

type Day3 struct {
}

func (d Day3) Calc() string {
	input := input.ReadInputFile(3)
	lines := strings.Split(input, "\r\n")
	var data [][]rune

	for _, e := range lines {
		data = append(data, []rune(e))
	}

	return fmt.Sprintf("1: %d\n2: %d\n", d.calc1(data, 3, 1), d.calc2(data))
}

func (d Day3) calc1(data [][]rune, xSpeed int, ySpeed int) int {
	x, y, count := 0, 0, 0

	for {
		x += xSpeed
		y += ySpeed

		if y >= len(data) {
			return count
		}

		x %= len(data[y])

		if data[y][x] == '#' {
			count++
		}
	}
}

func (d Day3) calc2(data [][]rune) int {
	return d.calc1(data, 1, 1) *
		d.calc1(data, 3, 1) *
		d.calc1(data, 5, 1) *
		d.calc1(data, 7, 1) *
		d.calc1(data, 1, 2)
}
