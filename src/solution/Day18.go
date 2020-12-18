package solution

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/constants"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"strconv"
	"strings"
)

type Day18 struct {
}

func (d Day18) Calc() string {
	lines := strings.Split(input.ReadInputFile(18), constants.LineSeparator)
	for i, line := range lines {
		lines[i] = strings.Replace(line, " ", "", -1)
	}

	return fmt.Sprintf("1: %v\n2: %v\n", d.sumAllExpressions(lines), d.sumAllExpressionsAdvancedMath(lines))
}

//5472527075827 too low

func (d Day18) sumAllExpressions(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += d.evaluateExpression(line)
	}
	return sum
}
func (d Day18) sumAllExpressionsAdvancedMath(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += d.evaluateAdvancedExpression(line)
	}
	return sum
}

func (d Day18) evaluateExpression(expression string) int {
	if len(expression) == 1 {
		value, err := strconv.Atoi(expression)
		if err != nil {
			panic(err)
		}
		return value
	}
	if expression[len(expression)-1] == ')' && d.findMatchingBracketReverse(expression, len(expression)-1) == 0 {
		return d.evaluateExpression(expression[1 : len(expression)-1])
	}
	switch expression[len(expression)-1] {
	case ')':
		closingBracketIndex := d.findMatchingBracketReverse(expression, len(expression)-1)
		leftSide := expression[0 : closingBracketIndex-1]
		operator := expression[closingBracketIndex-1]
		rightSide := expression[closingBracketIndex:]
		switch operator {
		case '*':
			return d.evaluateExpression(leftSide) * d.evaluateExpression(rightSide)
		case '+':
			return d.evaluateExpression(leftSide) + d.evaluateExpression(rightSide)
		}
	default:
		leftSide := expression[0 : len(expression)-2]
		operator := expression[len(expression)-2]
		rightSide := expression[len(expression)-1]
		switch operator {
		case '*':
			return d.evaluateExpression(leftSide) * d.evaluateExpression(string(rightSide))
		case '+':
			return d.evaluateExpression(leftSide) + d.evaluateExpression(string(rightSide))
		}
	}
	return -69
}

func (d Day18) evaluateAdvancedExpression(expression string) int {
	if len(expression) == 1 {
		value, err := strconv.Atoi(expression)
		if err != nil {
			panic(err)
		}
		return value
	}
	if expression[len(expression)-1] == ')' && d.findMatchingBracketReverse(expression, len(expression)-1) == 0 {
		return d.evaluateAdvancedExpression(expression[1 : len(expression)-1])
	}
	switch expression[len(expression)-1] {
	case ')':
		closingBracketIndex := d.findMatchingBracketReverse(expression, len(expression)-1)
		leftSide := expression[0 : closingBracketIndex-1]
		operator := expression[closingBracketIndex-1]
		rightSide := expression[closingBracketIndex:]
		switch operator {
		case '*':
			return d.evaluateAdvancedExpression(leftSide) * d.evaluateAdvancedExpression(rightSide)
		case '+':
			if leftSide[len(leftSide)-1] == ')' {
				leftOpening := d.findMatchingBracketReverse(leftSide, len(leftSide)-1)
				if leftOpening > 0 {
					sb := strings.Builder{}
					sb.WriteString(leftSide[:leftOpening])
					sb.WriteRune('(')
					sb.WriteString(leftSide[leftOpening:])
					sb.WriteByte(operator)
					sb.WriteString(rightSide)
					sb.WriteRune(')')

					return d.evaluateAdvancedExpression(sb.String())

				} else {
					return d.evaluateAdvancedExpression(leftSide) + d.evaluateAdvancedExpression(rightSide)
				}
			} else {
				if len(leftSide) > 1 {
					sb := strings.Builder{}
					sb.WriteString(leftSide[:len(leftSide)-1])
					sb.WriteRune('(')
					sb.WriteByte(leftSide[len(leftSide)-1])
					sb.WriteByte(operator)
					sb.WriteString(rightSide)
					sb.WriteRune(')')

					return d.evaluateAdvancedExpression(sb.String())
				} else {
					return d.evaluateAdvancedExpression(leftSide) + d.evaluateAdvancedExpression(string(rightSide))
				}
			}
		}
	default:
		leftSide := expression[0 : len(expression)-2]
		operator := expression[len(expression)-2]
		rightSide := expression[len(expression)-1]
		switch operator {
		case '*':
			return d.evaluateAdvancedExpression(leftSide) * d.evaluateAdvancedExpression(string(rightSide))
		case '+':
			if leftSide[len(leftSide)-1] == ')' {
				leftOpening := d.findMatchingBracketReverse(leftSide, len(leftSide)-1)
				if leftOpening > 0 {
					sb := strings.Builder{}
					sb.WriteString(leftSide[:leftOpening])
					sb.WriteRune('(')
					sb.WriteString(leftSide[leftOpening:])
					sb.WriteByte(operator)
					sb.WriteByte(rightSide)
					sb.WriteRune(')')

					return d.evaluateAdvancedExpression(sb.String())

				} else {
					return d.evaluateAdvancedExpression(leftSide) + d.evaluateAdvancedExpression(string(rightSide))
				}
			} else {
				if len(leftSide) > 1 {
					sb := strings.Builder{}
					sb.WriteString(leftSide[:len(leftSide)-1])
					sb.WriteRune('(')
					sb.WriteByte(leftSide[len(leftSide)-1])
					sb.WriteByte(operator)
					sb.WriteByte(rightSide)
					sb.WriteRune(')')

					return d.evaluateAdvancedExpression(sb.String())
				} else {
					return d.evaluateAdvancedExpression(leftSide) + d.evaluateAdvancedExpression(string(rightSide))
				}
			}
		}
	}
	return -69
}

func (d Day18) findMatchingBracketReverse(expression string, bracketIndex int) int {
	openBrackets := 1
	closingBracketIndex := -1
	for index := bracketIndex - 1; closingBracketIndex < 0; index-- {
		switch expression[index] {
		case ')':
			openBrackets++
		case '(':
			openBrackets--
			if openBrackets == 0 {
				closingBracketIndex = index
			}
		}
	}
	return closingBracketIndex
}

func (d Day18) findMatchingBracket(expression string, bracketIndex int) int {
	openBrackets := 1
	closingBracketIndex := -1
	for index := bracketIndex + 1; closingBracketIndex < 0; index++ {
		switch expression[index] {
		case '(':
			openBrackets++
		case ')':
			openBrackets--
			if openBrackets == 0 {
				closingBracketIndex = index
			}
		}
	}
	return closingBracketIndex
}
