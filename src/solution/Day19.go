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
	var appliesRule func(message, int) (bool, []int)
	appliesRule = func(m message, rIndex int) (bool, []int) {
		r := rules[rIndex]
		switch r[0] {
		case '"':
			if len(m) == 0 {
				return false, nil
			}
			if m[0] == r[1] {
				return true, []int{1}
			}
			return false, nil
		default:
			orGroups := strings.Split(string(r), " | ")
			var totalMatches []int
			for _, og := range orGroups {
				followingRules := strings.Split(og, " ")
				stringIndex := []int{0}
				for _, fr := range followingRules {
					nextRule, err := strconv.Atoi(fr)
					if err != nil {
						panic(err)
					}
					oldStLength := len(stringIndex)
					for stIndex := 0; stIndex < oldStLength; stIndex++ {
						st := stringIndex[stIndex]
						does, length := appliesRule(m[st:], nextRule)
						if does {
							stringIndex = append(stringIndex[:stIndex], stringIndex[stIndex+1:]...)
							stIndex--
							oldStLength--
							for _, newIndex := range length {
								if newIndex <= len(m) {
									stringIndex = append(stringIndex, st+newIndex)
								}
							}
						} else {
							stringIndex = append(stringIndex[:stIndex], stringIndex[stIndex+1:]...)
							stIndex--
							oldStLength--
						}
					}
				}
				totalMatches = append(totalMatches, stringIndex...)
			}
			if len(totalMatches) > 0 {
				return true, totalMatches
			} else {
				return false, nil
			}
		}
	}

	count := 0
	for _, m := range messages {
		if does, length := appliesRule(m, 0); does {
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
