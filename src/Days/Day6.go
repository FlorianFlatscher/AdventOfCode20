package Days

import (
	"../Constants"
	"../Input"
	"fmt"
	"strings"
)

type Day6 struct{}

func (d Day6) Calc() string {
	puzzleInput := Input.ReadInputFile(6)

	answeredQuestions := strings.Split(puzzleInput, strings.Repeat(Constants.LineSeparator, 2))

	return fmt.Sprintf(
		"Awnsered questions by anyone: %v\nAwnsered questions by everyone: %v", countAnsweredQuestionsAnyone(answeredQuestions), countAnsweredQuestionsEveryone(answeredQuestions))
}

func countAnsweredQuestionsAnyone(answeredQuestions []string) int {
	sum := 0

	for _, group := range answeredQuestions {
		collected := make(map[rune]bool)

		for _, question := range strings.ReplaceAll(group, "\n", "") {
			collected[question] = true
		}

		sum += len(collected)
	}

	return sum
}

func countAnsweredQuestionsEveryone(answeredQuestions []string) int {
	sum := 0

	for _, group := range answeredQuestions {
		collected := make(map[rune]int)

		for _, question := range strings.ReplaceAll(group, "\n", "") {
			collected[question] = collected[question] + 1
		}

		for _, value := range collected {
			if value == strings.Count(group, "\n")+1 {
				sum++
			}
		}
	}

	return sum
}
