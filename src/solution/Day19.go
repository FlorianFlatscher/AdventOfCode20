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

	fmt.Println(rules)
	fmt.Println(messages)

	//part1 := d.countRule0(rules, messages)
	rules[8] = "42 | 42 8"
	rules[11] = "42 31 | 42 11 31"
	part2 := d.countRule0(rules, messages)

	return fmt.Sprintf("1: %v\n2: %v\n", nil, part2)
}

//378 too high
//370 too high
//282 too low

func (d Day19) countRule0(rules map[int]rule, messages []message) int {
	var appliesRule func(message, int) (bool, int)
	appliesRule = func(m message, rIndex int) (bool, int) {
		r := rules[rIndex]
		switch r[0] {
		case '"':
			if len(m) == 0 {
				return false, 0
			}
			if m[0] == r[1] {
				return true, 1
			}
			return false, -1
		default:
			orGroups := strings.Split(string(r), " | ")
			highestMatch := -1
			for _, og := range orGroups {
				followingRules := strings.Split(og, " ")
				index := 0
				for i, fr := range followingRules {
					nextRule, err := strconv.Atoi(fr)
					if err != nil {
						panic(err)
					}
					//if nextRule == rIndex && i+2 == len(followingRules){
					//	tryNextRule, err := strconv.Atoi(followingRules[i+1])
					//	if err != nil {
					//		panic(err)
					//	}
					//	for tryStart := index+1; tryStart < len(m); tryStart++ {
					//		does, length := appliesRule(m[tryStart:], tryNextRule)
					//		doesThis, lengthThis := appliesRule(m[index+1:tryStart], nextRule)
					//		if  does && length + tryStart == len(m) && doesThis && index + lengthThis == tryStart + 1{
					//			return true, len(m)
					//		}
					//	}
					//}

					does, length := appliesRule(m[index:], nextRule)
					if does {
						index += length
						if i == len(followingRules)-1 {
							if index > highestMatch {
								highestMatch = index
							}
						}
					} else {
						break
					}
				}
			}
			if highestMatch > 0 {
				return true, highestMatch
			} else {
				return false, highestMatch
			}
		}
	}

	count := 0
	for _, m := range messages {
		if does, length := appliesRule(m, 0); does {
			fmt.Println(m, length)
			if length == len(m) {
				count++

			}
		}
	}
	return count
}
