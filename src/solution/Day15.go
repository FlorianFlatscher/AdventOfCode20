package solution

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"strconv"
	"strings"
)

type Day15 struct {
}

func (d Day15) Calc() string {
	pInput := strings.Split(input.ReadInputFile(15), ",")
	var numbers []int
	for _, n := range pInput {
		parseN, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, parseN)
	}

	return fmt.Sprintf("2020th number: %v\n30000000th number: %v\n", d.calcNumberAt2020(numbers), d.calcNumberAt30000000(numbers))
}

func (d Day15) calcNumberAt2020(numbers []int) int {
	lastNumber := numbers[len(numbers)-1]
	store := make(map[int]int)
	for i, number := range numbers[:len(numbers)-1] {
		store[number] = i + 1
	}

	turn := len(numbers)

	for turn < 2020 {
		if mostRecent, exists := store[lastNumber]; !exists {
			store[lastNumber] = turn
			lastNumber = 0
		} else {
			newNumber := turn - mostRecent
			store[lastNumber] = turn
			lastNumber = newNumber
		}
		turn++
	}

	return lastNumber
}

func (d Day15) calcNumberAt30000000(numbers []int) int {
	lastNumber := numbers[len(numbers)-1]
	store := make(map[int]int)
	for i, number := range numbers[:len(numbers)-1] {
		store[number] = i + 1
	}

	turn := len(numbers)

	for turn < 30000000 {
		if mostRecent, exists := store[lastNumber]; !exists {
			store[lastNumber] = turn
			lastNumber = 0
		} else {
			newNumber := turn - mostRecent
			store[lastNumber] = turn
			lastNumber = newNumber
		}
		turn++
	}

	return lastNumber
}
