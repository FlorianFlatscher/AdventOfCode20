package Days

import (
	"../Constants"
	"../Input"
	fmt "fmt"
	"strconv"
	"strings"
)

type Day9 struct {
}

func (d Day9) Calc() string {
	var input []int
	for _, line := range strings.Split(Input.ReadInputFile(9), Constants.LineSeparator) {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		input = append(input, value)
	}

	return fmt.Sprintf("First number not a sum of previous 25: %v\nEncryption weakness: %v\n", firstNotSum(input), encryptionWeakness(input))
}

func firstNotSum(input []int) int {
	for i, current := range input {
		past25 := make(map[int]bool)
		if i > 25 {
			for pastI := i - 25; pastI < i; pastI++ {
				if past25[current-input[pastI]] {
					break
				} else if pastI+1 == i {
					return current
				}
				past25[input[pastI]] = true
			}
		}

	}

	return -1
}

func encryptionWeakness(input []int) int {
	number := firstNotSum(input)

	for i := range input {
		sum := input[i]
		smallest, largest := input[i], input[i]
		for l := i + 1; l < len(input); l++ {
			newNumber := input[l]
			sum += newNumber
			if newNumber < smallest {
				smallest = newNumber
			}
			if newNumber > largest {
				largest = newNumber
			}
			if sum == number {
				return smallest + largest
			}
		}
		if sum > number {
			continue
		}
	}

	return -1
}
