package solution

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/constants"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"regexp"
	"strconv"
	"strings"
)

type Day19 struct {
}

type rule string

type message string

var (
	regexRule = regexp.MustCompile("(\\d+): (.*?)$")
)

func (d Day19) Calc() string {
	inputParts := strings.Split(input.ReadInputFile(19), strings.Repeat(constants.LineSeparator, 2))

	rules := make(map[int]rule)
	for _, ruleInput := range strings.Split(inputParts[0], constants.LineSeparator) {
		result := regexRule.FindStringSubmatch(ruleInput)
		newRuleIndex, err := strconv.Atoi(result[1])
		if err != nil {
			panic(err)
		}
		newRule := rule(result[2])
		rules[newRuleIndex] = newRule
	}

	var messages []message
	for _, m := range strings.Split(inputParts[1], constants.LineSeparator) {
		messages = append(messages, message(m))
	}

	part1 := d.countRule0(rules, messages)
	rules[8] = "42 | 42 8"
	rules[11] = "42 31 | 42 11 31"
	part2 := d.countRule0(rules, messages)

	return fmt.Sprintf("1: %v\n2: %v\n", part1, part2)
}

func (d Day19) countRule0(rules map[int]rule, messages []message) int {
	var checkMessageForRule func(message, int) (bool, []int)
	checkMessageForRule = func(m message, ruleNumber int) (bool, []int) {
		rule := rules[ruleNumber]
		switch rule[0] {
		case '"':
			if len(m) == 0 {
				return false, nil
			}
			if m[0] == rule[1] {
				return true, []int{1}
			}
			return false, nil
		default:
			orGroups := strings.Split(string(rule), " | ")
			var totalMatchesByAllOrGroups []int
			for _, orGroup := range orGroups {
				rulesInOrgroup := strings.Split(orGroup, " ")
				allCurrentIndex := []int{0}
				for _, fr := range rulesInOrgroup {
					nextRule, err := strconv.Atoi(fr)
					if err != nil {
						panic(err)
					}
					oldIndexLength := len(allCurrentIndex)
					for possibilityIndex := 0; possibilityIndex < oldIndexLength; possibilityIndex++ {
						currentIndex := allCurrentIndex[possibilityIndex]
						ruleApplies, length := checkMessageForRule(m[currentIndex:], nextRule)
						if ruleApplies {
							for _, newIndex := range length {
								if newIndex <= len(m) {
									allCurrentIndex = append(allCurrentIndex, currentIndex+newIndex)
								}
							}
						}
						allCurrentIndex = append(allCurrentIndex[:possibilityIndex], allCurrentIndex[possibilityIndex+1:]...)
						possibilityIndex--
						oldIndexLength--
					}
				}
				totalMatchesByAllOrGroups = append(totalMatchesByAllOrGroups, allCurrentIndex...)
			}
			if len(totalMatchesByAllOrGroups) > 0 {
				return true, totalMatchesByAllOrGroups
			} else {
				return false, nil
			}
		}
	}

	count := 0
	for _, m := range messages {
		if does, length := checkMessageForRule(m, 0); does {
			for _, l := range length {
				if l == len(m) {
					count++
					break
				}
			}

		}
	}
	return count
}
