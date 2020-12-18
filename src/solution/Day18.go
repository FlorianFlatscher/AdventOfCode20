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

func (d Day18) convertExpressionToAdvancedMath(expression string) string {
	if len(expression) == 1 {
		return expression
	}
	if expression[len(expression)-1] == ')' && d.findMatchingBracketReverse(expression, len(expression)-1) == 0 {
		var sb strings.Builder
		sb.WriteRune('(')
		sb.WriteString(d.convertExpressionToAdvancedMath(expression[1 : len(expression)-1]))
		sb.WriteRune(')')
		return sb.String()
	}
	switch expression[0] {
	case '(':
		closingBracketIndex := d.findMatchingBracket(expression, 0)
		leftSide := expression[0 : closingBracketIndex+1]
		operator := expression[closingBracketIndex+1]
		rightSide := expression[closingBracketIndex+2:]

		var sb strings.Builder

		if operator == '+' {
			sb.WriteRune('(')
			sb.WriteString(leftSide)
			sb.WriteByte(operator)
			sb.WriteByte(rightSide[0])
			sb.WriteRune(')')
			if len(rightSide) > 1 {
				sb.WriteString(rightSide[1:])
			}
			return sb.String()
		}
		sb.WriteString(leftSide)
		sb.WriteByte(operator)
		sb.WriteString(d.convertExpressionToAdvancedMath(rightSide))
		return sb.String()

	default:
		leftSide := expression[0]
		operator := expression[1]
		rightSide := expression[2:]
		var sb strings.Builder

		if operator == '+' {
			sb.WriteRune('(')
			sb.WriteByte(leftSide)
			sb.WriteByte(operator)
			if rightSide[0] == '(' {
				sb.WriteString(d.convertExpressionToAdvancedMath(rightSide))
				sb.WriteRune(')')
				return sb.String()
			} else {
				sb.WriteByte(rightSide[0])
				sb.WriteRune(')')
				if len(rightSide) > 1 {
					sb.WriteString(rightSide[1:])
				}
				return d.convertExpressionToAdvancedMath(sb.String())
			}
		}
		sb.WriteByte(leftSide)
		sb.WriteByte(operator)
		sb.WriteString(d.convertExpressionToAdvancedMath(rightSide))
		return sb.String()
	}
}

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
		advancedLine := d.convertExpressionToAdvancedMath(line)
		fmt.Println(advancedLine)
		sum += d.evaluateExpression(advancedLine)
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
