package Days

import (
	"../Input"
	"fmt"
	"strconv"
	"strings"
)

type Day1 struct {
}

func (d Day1) Calc() string {
	input := Input.ReadInputFile(1)

	lines := strings.Split(input, "\r\n")
	var arr []int

	for _, i := range lines {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		arr = append(arr, j)
	}

	return fmt.Sprintf("Part1: %v\nPart2: %v", d.calc1(arr), d.calc2(arr))
}

func (d Day1) calc1(arr []int) string {
	set := make(map[int]bool)

	for _, number := range arr {
		if set[2020-number] {
			return strconv.Itoa(number * (2020 - number))
		}
		set[number] = true
	}

	return ""
}

func (d Day1) calc2(arr []int) string {
	set := make(map[int]bool)

	result := make(chan string)

	for _, newNumber := range arr {

		for number2, _ := range set {
			number3 := 2020 - newNumber - number2
			if set[number3] {
				go func() {
					result <- strconv.Itoa(newNumber * number2 * number3)
				}()
			}
		}
		set[newNumber] = true
	}

	output := <-result
	return output
}
