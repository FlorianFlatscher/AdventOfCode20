package solution

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/constants"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"math"
	"strconv"
	"strings"
)

type Day13 struct {
}

type BusLine struct {
	lineId int
	offset int
}

func (d Day13) Calc() string {
	puzzleInput := strings.Split(input.ReadInputFile(13), constants.LineSeparator)
	timeStamp, _ := strconv.Atoi(puzzleInput[0])
	var busLines []BusLine
	for i, e := range strings.Split(puzzleInput[1], ",") {
		newBusLine, err := strconv.Atoi(e)
		if err == nil {
			busLines = append(busLines, BusLine{
				newBusLine,
				i,
			})
		}
	}
	return fmt.Sprintf("1: %v\n2: %v\n", d.calcEarliestBus(timeStamp, busLines), d.calcOffsetMeetingBusAdvanced(busLines))
}

func (d *Day13) calcEarliestBus(timestamp int, busLines []BusLine) int {
	earliestBus, waitingTime := 0, math.MaxInt64

	for _, line := range busLines {
		linesWaitingTime := line.lineId - timestamp%line.lineId
		if linesWaitingTime < waitingTime {
			earliestBus, waitingTime = line.lineId, linesWaitingTime
		}
	}

	return earliestBus * waitingTime
}

func (d *Day13) calcOffsetMeetingBusAdvanced(busLines []BusLine) int {
	step := 1
	for timeStamp := 0; true; timeStamp += step {
		step = 1
		rightOne := true
		for _, line := range busLines {
			if (timeStamp+line.offset)%line.lineId == 0 {
				step *= line.lineId
			} else {
				rightOne = false
			}
		}
		if rightOne {
			return timeStamp
		}
	}
	return -1
}
