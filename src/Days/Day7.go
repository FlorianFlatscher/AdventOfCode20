package Days

import (
	"../Constants"
	"../Input"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	expGetParent        = regexp.MustCompile("(.+?) bag")
	expGetChildren      = regexp.MustCompile("[0-9]+ (.+?) bag[s]?[,.]?")
	expGetChildrenCount = regexp.MustCompile("([0-9]+) (.+?) bag[s]?[,.]?")
)

type Day7 struct {
	IDay
}

func (d Day7) Calc() string {
	input := strings.Split(Input.ReadInputFile(7), Constants.LineSeparator)

	return fmt.Sprintf("Number Containing Shiny Gold: %v\nBags in Shiny Gold: %v\n", numContainingShinyGold(input), bagsInShinyGold(input))
}

func numContainingShinyGold(input []string) int {
	parents := make(map[string][]string)

	for _, line := range input {
		parent := expGetParent.FindStringSubmatch(line)[1]
		children := expGetChildren.FindAllStringSubmatch(line, -1)

		for _, v := range children {
			parents[v[1]] = append(parents[v[1]], parent)
		}
	}

	containingGold := make(map[string]bool)

	//stack := []string{"shiny gold"}

	//for {
	//	child := stack[len(stack)-1]
	//	stack := stack[:len(stack)-1]
	//	for _, parent := range parents[child] {
	//		containingGold[parent] = true
	//		stack = append(stack, parent)
	//	}
	//
	//
	//	if len(stack) == 0 {
	//	}
	//}
	var reconstruct func(child string)
	reconstruct = func(child string) {
		parentList := parents[child]
		for _, parent := range parentList {
			containingGold[parent] = true
			reconstruct(parent)
		}
	}
	reconstruct("shiny gold")

	return len(containingGold)
}

func bagsInShinyGold(input []string) int {
	type bag struct {
		color string
		count int
	}

	children := make(map[string][]bag)

	for _, line := range input {
		parent := expGetParent.FindStringSubmatch(line)[1]
		childBags := expGetChildrenCount.FindAllStringSubmatch(line, -1)

		for _, v := range childBags {
			count, err := strconv.Atoi(v[1])
			if err != nil {
				panic(err)
			}
			children[parent] = append(children[parent], bag{color: v[2], count: count})
		}
	}

	var reconstruct func(child string) int
	reconstruct = func(child string) int {
		childrenList := children[child]
		sum := 0
		for _, children := range childrenList {
			sum += children.count * reconstruct(children.color)
		}
		return sum + 1
	}

	return reconstruct("shiny gold") - 1
}
