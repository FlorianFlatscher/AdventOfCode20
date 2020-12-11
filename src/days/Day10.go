package days

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/constants"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day10 struct {
}

func (d Day10) Calc() string {
	puzzleInput := input.ReadInputFile(10)
	var adapters []int
	for _, a := range strings.Split(puzzleInput, constants.LineSeparator) {
		joltage, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}
		adapters = append(adapters, joltage)
	}
	sort.Ints(adapters)
	return fmt.Sprintf("1: %v\n2: %v\n", chainAdapters(adapters), numberDistinctArrangements(adapters))
}

func chainAdapters(adapters []int) int {
	differencesOf1 := 0
	differencesOf3 := 1
	joltage := 0
	for _, a := range adapters {
		difference := a - joltage
		switch difference {
		case 1:
			differencesOf1++
		case 3:
			differencesOf3++
		}
		joltage += difference
	}
	return differencesOf3 * differencesOf1
}

func numberDistinctArrangements(adapters []int) int {
	adapters = append([]int{0}, adapters...)
	var chain func(adapterIndex int) int
	memo := make(map[int]int)

	chain = func(adapterIndex int) int {
		possibilities := 0
		if memo[adapterIndex] == 0 {
			for i := adapterIndex + 1; i < len(adapters) && adapters[i]-adapters[adapterIndex] <= 3; i++ {
				possibilities += chain(i)
			}
			memo[adapterIndex] = int(math.Max(float64(possibilities), 1))
		} else {
			possibilities += memo[adapterIndex]
		}

		if possibilities > 0 {
			return possibilities
		} else {
			return 1
		}
	}
	return chain(0)
}
