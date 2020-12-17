package solution

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/constants"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"strings"
)

type Day17 struct{}

type cubeState int

const (
	INACTIVE cubeState = iota
	ACTIVE   cubeState = iota
)

func (c cubeState) String() string {
	switch c {
	case ACTIVE:
		return "#"
	case INACTIVE:
		return "."
	}
	return "?"
}

func (d Day17) Calc() string {
	grid := d.parseInput(input.ReadInputFile(17))
	return fmt.Sprintf("1: %v\n", d.conwayCubes3rdDimension(grid))
}

func (d Day17) conwayCubes4thDimension(grid [][][][]cubeState) int {
	active := 0
	for cycle := 0; cycle < 6; cycle++ {
		nextCycle := make([][][][]cubeState, len(grid)+2)
		active = 0
		for w := range nextCycle {
			nextCycle[w] = make([][][]cubeState, len(grid[0])+2)
			for z := range nextCycle[w] {
				nextCycle[w][z] = make([][]cubeState, len(grid[0][0])+2)
				for y := range nextCycle[z] {
					nextCycle[w][z][y] = make([]cubeState, len(grid[0][0][0])+2)
					for x := range nextCycle[w][z][y] {
						state := INACTIVE
						if w > 0 && z > 0 && y > 0 && x > 0 && w <= len(grid) && z <= len(grid[0]) && y <= len(grid[0][0]) && x <= len(grid[0][0][0]) {
							state = grid[w-1][z-1][y-1][x-1]
						}
						neighbours := d.countActiveNeighbors4thDimension(grid, w-1, z-1, y-1, x-1)
						switch state {
						case ACTIVE:
							if neighbours == 2 || neighbours == 3 {
								nextCycle[w][z][y][x] = ACTIVE
								active++
							} else {
								nextCycle[w][z][y][x] = INACTIVE
							}
						case INACTIVE:
							if neighbours == 3 {
								nextCycle[w][z][y][x] = ACTIVE
								active++
							} else {
								nextCycle[w][z][y][x] = INACTIVE
							}
						}
					}
				}
			}
		}
		grid = nextCycle
	}

	return active
}

func (d Day17) conwayCubes3rdDimension(grid [][][]cubeState) int {
	active := 0
	for cycle := 0; cycle < 6; cycle++ {
		nextCycle := make([][][]cubeState, len(grid)+2)
		active = 0
		for z := range nextCycle {
			nextCycle[z] = make([][]cubeState, len(grid[0])+2)
			for y := range nextCycle[z] {
				nextCycle[z][y] = make([]cubeState, len(grid[0][0])+2)
				for x := range nextCycle[z][y] {
					state := INACTIVE
					if z > 0 && y > 0 && x > 0 && z <= len(grid) && y <= len(grid[0]) && x <= len(grid[0][0]) {
						state = grid[z-1][y-1][x-1]
					}
					neighbours := d.countActiveNeighbors3rdDimension(grid, z-1, y-1, x-1)
					switch state {
					case ACTIVE:
						if neighbours == 2 || neighbours == 3 {
							nextCycle[z][y][x] = ACTIVE
							active++
						} else {
							nextCycle[z][y][x] = INACTIVE
						}
					case INACTIVE:
						if neighbours == 3 {
							nextCycle[z][y][x] = ACTIVE
							active++
						} else {
							nextCycle[z][y][x] = INACTIVE
						}
					}
				}
			}
		}
		grid = nextCycle
	}

	return active
}

func (_ Day17) countActiveNeighbors3rdDimension(grid [][][]cubeState, z, y, x int) int {
	neighbors := 0
	for zz := -1; zz < 2; zz++ {
		for yy := -1; yy < 2; yy++ {
			for xx := -1; xx < 2; xx++ {
				if z+zz >= 0 && z+zz < len(grid) &&
					y+yy >= 0 && y+yy < len(grid[z+zz]) &&
					x+xx >= 0 && x+xx < len(grid[z+zz][y+yy]) &&
					!(zz == 0 && yy == 0 && xx == 0) {
					if grid[z+zz][y+yy][x+xx] == ACTIVE {
						neighbors++
					}
				}
			}
		}
	}
	return neighbors
}

func (_ Day17) countActiveNeighbors4thDimension(grid [][][][]cubeState, w, z, y, x int) int {
	neighbors := 0
	for ww := -1; ww < 2; ww++ {
		for zz := -1; zz < 2; zz++ {
			for yy := -1; yy < 2; yy++ {
				for xx := -1; xx < 2; xx++ {
					if w+ww >= 0 && w+ww < len(grid) &&
						z+zz >= 0 && z+zz < len(grid[w+ww]) &&
						y+yy >= 0 && y+yy < len(grid[w+ww][z+zz]) &&
						x+xx >= 0 && x+xx < len(grid[w+ww][z+zz][y+yy]) &&
						!(ww == 0 && zz == 0 && yy == 0 && xx == 0) {
						if grid[w+ww][z+zz][y+yy][x+xx] == ACTIVE {
							neighbors++
						}
					}
				}
			}
		}
	}
	return neighbors
}

func (d *Day17) parseInput(input string) [][][]cubeState {
	grid := make([][][]cubeState, 1)
	lines := strings.Split(input, constants.LineSeparator)
	grid[0] = make([][]cubeState, len(lines))
	for i, line := range lines {
		chars := strings.Split(line, "")
		grid[0][i] = make([]cubeState, len(chars))
		for l, char := range chars {
			switch char[0] {
			case '.':
				grid[0][i][l] = INACTIVE
			case '#':
				grid[0][i][l] = ACTIVE
			}
		}
	}
	return grid
}
