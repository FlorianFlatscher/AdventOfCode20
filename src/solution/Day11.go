package solution

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/constants"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"strings"
)

type seatStatus int

const (
	FLOOR    = 0
	EMPTY    = 1
	OCCUPIED = 2
)

type Day11 struct{}

func (d Day11) Calc() string {
	var seats [][]seatStatus
	puzzleInput := input.ReadInputFile(11)
	for _, line := range strings.Split(puzzleInput, constants.LineSeparator) {
		var row []seatStatus
		for _, seatInput := range []rune(line) {
			switch seatInput {
			case '.':
				row = append(row, FLOOR)
			case 'L':
				row = append(row, EMPTY)
			case '#':
				row = append(row, OCCUPIED)
			}
		}
		seats = append(seats, row)
	}

	return fmt.Sprintf("Occupied seats: %v, Occupied with raycast: %v", clacOccupiedSeats(seats), clacOccupiedSeats2(seats))
}

func clacOccupiedSeats(seats [][]seatStatus) int {
	current := make([][]seatStatus, 0)
	for _, seat := range seats {
		clone := make([]seatStatus, len(seat))
		copy(clone, seat)
		current = append(current, clone)
	}

	countNeighbors := func(x, y int) (count int) {
		for xx := x - 1; xx <= x+1; xx++ {
			if xx >= 0 && xx < len(current) {
				for yy := y - 1; yy <= y+1; yy++ {
					if !(xx == x && yy == y) && yy >= 0 && yy < len(current[xx]) && current[xx][yy] == OCCUPIED {
						count++
					}
				}
			}
		}
		return
	}

	for {
		next := make([][]seatStatus, 0)
		hasChanged := false
		for x := range current {
			next = append(next, make([]seatStatus, len(current[x])))
			for y := range current[x] {
				if current[x][y] != FLOOR {
					neighbours := countNeighbors(x, y)
					switch {
					case neighbours == 0:
						next[x][y] = OCCUPIED
						hasChanged = current[x][y] != OCCUPIED || hasChanged
					case neighbours >= 4:
						next[x][y] = EMPTY
						hasChanged = current[x][y] != EMPTY || hasChanged
					default:
						next[x][y] = current[x][y]
					}
				}
			}
		}

		if !hasChanged {
			occupiedSeats := 0
			for x := range next {
				for y := range next[x] {
					if next[x][y] == OCCUPIED {
						occupiedSeats++
					}
				}
			}
			return occupiedSeats
		}
		current = next
	}
}

func clacOccupiedSeats2(seats [][]seatStatus) int {
	current := make([][]seatStatus, 0)
	for _, seat := range seats {
		clone := make([]seatStatus, len(seat))
		copy(clone, seat)
		current = append(current, clone)
	}

	countNeighbors := func(x, y int) (count int) {
		for xx := -1; xx <= 1; xx++ {
			for yy := -1; yy <= 1; yy++ {
				if !(xx == 0 && yy == 0) {
					foundSomething := false
					distance := 1
					for x+xx*distance >= 0 && x+xx*distance < len(current) && y+yy*distance >= 0 && y+yy*distance < len(current[x+xx*distance]) {
						switch current[x+xx*distance][y+yy*distance] {
						case OCCUPIED:
							count++
							foundSomething = true
						case EMPTY:
							foundSomething = true
						}
						if foundSomething {
							break
						}
						distance++
					}
				}
			}
		}
		return
	}

	for {
		next := make([][]seatStatus, 0)
		hasChanged := false
		for x := range current {
			next = append(next, make([]seatStatus, len(current[x])))
			for y := range current[x] {
				if current[x][y] != FLOOR {
					neighbours := countNeighbors(x, y)
					switch {
					case neighbours == 0:
						next[x][y] = OCCUPIED
						hasChanged = current[x][y] != OCCUPIED || hasChanged
					case neighbours >= 5:
						next[x][y] = EMPTY
						hasChanged = current[x][y] != EMPTY || hasChanged
					default:
						next[x][y] = current[x][y]
					}
				}
			}
		}

		if !hasChanged {
			occupiedSeats := 0
			for x := range next {
				for y := range next[x] {
					if next[x][y] == OCCUPIED {
						occupiedSeats++
					}
				}
			}
			return occupiedSeats
		}
		current = next
	}
}
