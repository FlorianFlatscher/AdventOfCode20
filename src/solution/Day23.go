package solution

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"strconv"
	"strings"
	"time"
)

type Day23 struct{}

type Dequeue []int

func (d *Dequeue) Pop() int {
	first := (*d)[0]
	*d = (*d)[1:]
	return first
}

func (d *Dequeue) Peek() int {
	return (*d)[0]
}

func (d *Dequeue) Append(x int) {
	*d = append(*d, x)
}

func newDequeue(x string) *Dequeue {
	a := Dequeue(make([]int, len(x)))
	for i, c := range x {
		a[i] = mustAtoi(string(c))
	}
	return &a
}

func (d *Dequeue) get(i int) int {
	return (*d)[i]
}

func (d *Dequeue) cut(min int, max int) []int {
	var slice = make([]int, 0)
	slice = append(slice, (*d)[min:max]...)
	*d = append((*d)[:min], (*d)[max:]...)
	return slice
}

func (d *Dequeue) find(x int) int {
	for i := 0; i < len(*d); i++ {
		if (*d)[i] == x {
			return i
		}
	}
	return -1
}

func (d *Dequeue) findFrom(x int, l int) int {
	if l < 0 {
		l = 0
	}
	for i := l; i < len(*d); i++ {
		if (*d)[i] == x {
			return i
		}
	}
	return -1
}

func (d *Dequeue) insert(i int, x []int) {
	*d = append(*d, make([]int, len(x))...)
	//*d = append((*d)[:i+len(x)], (*d)[i:]...)
	copy((*d)[i+len(x):], (*d)[i:])
	copy((*d)[i:], x)
}

func (day Day23) Calc() string {
	cycle := newDequeue(input.ReadInputFile(23))
	return fmt.Sprintf("1: %v\n2: %v\n", day.simulate(*cycle), day.simulateBig(*cycle))
}

func (day *Day23) simulate(dequeue Dequeue) string {
	var d = Dequeue(make([]int, len(dequeue)))
	copy(d, dequeue)

	for move := 0; move < 100; move++ {
		fmt.Println(move)
		currentValue := d.Pop()

		nextThree := []int{d.Pop(), d.Pop(), d.Pop()}

		isInNextThree := func(x int) bool {
			for _, n := range nextThree {
				if n == x {
					return true
				}
			}
			return false
		}

		next := currentValue - 1

		for isInNextThree(next) {
			next--
		}

		designationIndex := -1

		if next > 0 {
			designationIndex = d.find(next)
		} else {
			designationIndex = d.find(len(d) + next + 4)
		}
		d.insert(designationIndex+1, nextThree)

		d.Append(currentValue)
	}

	sb := strings.Builder{}
	for i := d.find(1) + 1; i < d.find(1)+len(d); i++ {
		sb.WriteString(strconv.Itoa(d.get(i % len(d))))
	}

	return sb.String()
}

func (day *Day23) simulateBig(cycle Dequeue) int {
	var d = Dequeue(make([]int, len(cycle), 1000000))
	copy(d, cycle)

	for i := len(cycle); len(d) < 1000000; i++ {
		d.Append(i + 1)
	}

	lastSeenAt := make(map[int]int)
	for i, v := range d {
		lastSeenAt[v] = i
	}

	offset := 0
	t := time.Now()

	for move := 0; move < 10000000; move++ {
		if move%10000 == 0 {
			fmt.Println(move, len(d), len(lastSeenAt))
			fmt.Println("Time: ", time.Since(t))
			t = time.Now()

		}
		currentValue := d.Pop()

		nextThree := []int{d.Pop(), d.Pop(), d.Pop()}

		offset += 4

		isInNextThree := func(x int) bool {
			for _, n := range nextThree {
				if n == x {
					return true
				}
			}
			return false
		}

		next := currentValue - 1

		for isInNextThree(next) {
			next--
		}

		designationIndex := -1

		if next > 0 {
			designationIndex = d.findFrom(next, lastSeenAt[next]-offset)
		} else {
			next = len(d) + next + 4
			designationIndex = d.findFrom(next, lastSeenAt[next]-offset)
		}
		if move%10000 == 0 {
			fmt.Println(next, offset, lastSeenAt[next], lastSeenAt[next]-offset, designationIndex)
		}
		d.insert(designationIndex+1, nextThree)
		d.Append(currentValue)
		lastSeenAt[currentValue] = len(d) - 1 + offset
		for i, v := range nextThree {
			lastSeenAt[v] = designationIndex + 1 + i + offset
		}

	}

	for d.Peek() != 1 {
		d.Pop()
	}
	d.Pop()

	return d.Pop() * d.Pop()
}
