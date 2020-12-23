package solution

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"strconv"
	"strings"
)

type Day23 struct{}

type Cycle []int

func newCycle(x string) *Cycle {
	a := Cycle(make([]int, len(x)))
	for i, c := range x {
		a[i] = mustAtoi(string(c))
	}
	return &a
}

func (c *Cycle) get(i int) int {
	return (*c)[i%len(*c)]
}

func (c *Cycle) cut(min int, max int) []int {
	max = max % len(*c)
	min = min % len(*c)
	var slice = make([]int, 0, max)
	if max <= min {
		slice = append(slice, (*c)[min:]...)
		slice = append(slice, (*c)[:max]...)
		*c = (*c)[:min]
		*c = (*c)[max:]
	} else {
		slice = append(slice, (*c)[min:max]...)
		*c = append((*c)[:min], (*c)[max:]...)
	}
	return slice
}

func (c *Cycle) find(x int) int {
	for i := len(*c) - 1; i >= 0; i-- {
		if (*c)[i] == x {
			return i
		}
	}
	return -1
}

func (c *Cycle) insert(i int, x []int) {
	*c = append(*c, make([]int, len(x))...)
	//*d = append((*d)[:i+len(x)], (*d)[i:]...)
	copy((*c)[i+len(x):], (*c)[i:])
	copy((*c)[i:], x)
}

func (d Day23) Calc() string {
	cycle := newCycle(input.ReadInputFile(23))
	return fmt.Sprintf("1: %v\n2: %v\n", d.simulate(*cycle), d.simulateBig(*cycle))
}

func (d *Day23) simulate(cycle Cycle) string {
	var c = Cycle(make([]int, len(cycle)))
	copy(c, cycle)

	var current = 0
	for move := 0; move < 100; move++ {
		currentValue := c.get(current)
		fmt.Println("--", move+1, "--")
		fmt.Println("cups:", c)
		nextThree := c.cut(current+1, current+4)
		fmt.Println("pick up:", nextThree)
		fmt.Println("current", currentValue)

		designationIndex := -1
		for next := currentValue - 1; designationIndex < 0; next-- {
			if next > 0 {
				designationIndex = c.find(next)
			} else {
				designationIndex = c.find(len(c) + next + 3)
			}
		}
		fmt.Println("destination:", c.get(designationIndex))

		c.insert(designationIndex+1, nextThree)
		if designationIndex < current {
			oldCurrent := current
			current = current + 4
			current %= len(c)

			if len(c)-oldCurrent <= 3 {
				current -= 4 - (len(c) - oldCurrent)
			}
		} else {
			current = current + 1
			current %= len(c)
		}
	}

	sb := strings.Builder{}
	for i := c.find(1) + 1; i < c.find(1)+len(c); i++ {
		sb.WriteString(strconv.Itoa(c.get(i)))
	}

	return sb.String()
}

func (d *Day23) simulateBig(cycle Cycle) int {
	return 0
	var c = Cycle(make([]int, 1000000))
	copy(c, cycle)

	for i := len(cycle); i < 1000000; i++ {
		c[i] = i + 1
	}

	var current = 0
	for move := 0; move < 10000000; move++ {
		fmt.Println(move)
		//currentValue := c.get(current)
		nextThree := c.cut(current+1, current+4)

		designationIndex := -1

		//for next := currentValue - 1; designationIndex < 0; next-- {
		//	if _, ok := nextMap[next]; !ok {
		//		if next > 0 {
		//			designationIndex = c.find(next)
		//		} else {
		//			designationIndex = c.find(len(c) + next + 3)
		//		}
		//	}
		//}

		c.insert(designationIndex+1, nextThree)

		if designationIndex < current {
			oldCurrent := current
			current = current + 4
			current %= len(c)

			if len(c)-oldCurrent <= 3 {
				current -= 4 - (len(c) - oldCurrent)
			}
		} else {
			current = current + 1
			current %= len(c)
		}
	}

	pro := 1
	for i := c.find(1) + 1; i < c.find(1)+3; i++ {
		pro *= c[i]
	}

	return pro
}
