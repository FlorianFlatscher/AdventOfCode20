package days

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/constants"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"strconv"
	"strings"
)

type Day5 struct{}

func (d Day5) Calc() string {
	puzzleInput := input.ReadInputFile(5)

	binarySeats := strings.Split(puzzleInput, constants.LineSeparator)
	return fmt.Sprintf("Last Seat: %vMy Seat: %v", findMax(binarySeats), getMySeat(binarySeats))
}

func getMySeat(binarySeats []string) int {
	var seats [128 * 8]bool

	for _, binarySeat := range binarySeats {
		seatNr, _ := decodeBinarySeat(binarySeat)
		seats[seatNr] = true
	}

	startSearch := false
	for seatNr, seat := range seats {
		if startSearch && !seat {
			return seatNr
		} else if seat {
			startSearch = true
		}
	}

	return -1
}

func findMax(binarySeats []string) int {
	max := -1
	for _, binarySeat := range binarySeats {
		seatNumber, err := decodeBinarySeat(binarySeat)
		if err != nil {
			panic(err)
		}
		if seatNumber > max {
			max = seatNumber
		}
	}

	return max
}

func decodeBinarySeat(binarySeat string) (int, error) {
	binarySeat = strings.ReplaceAll(binarySeat, "F", "0")
	binarySeat = strings.ReplaceAll(binarySeat, "L", "0")
	binarySeat = strings.ReplaceAll(binarySeat, "B", "1")
	binarySeat = strings.ReplaceAll(binarySeat, "R", "1")

	seatNumber, err := strconv.ParseInt(binarySeat, 2, 11)
	return int(seatNumber), err
}
