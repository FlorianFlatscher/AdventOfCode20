package days

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"regexp"
	"strconv"
	"strings"
)

type Day2 struct {
}

type password struct {
	min, max int
	char     string
	password string
}

func (d Day2) Calc() string {
	input := input.ReadInputFile(2)
	var passwords []password

	for _, e := range strings.Split(input, "\r\n") {
		regex := regexp.MustCompile(`(\d*?)-(\d*?) ([a-z]): ([a-z]*)$`)
		values := regex.FindStringSubmatch(e)
		min, _ := strconv.Atoi(values[1])
		max, _ := strconv.Atoi(values[2])
		passwords = append(passwords, password{
			min,
			max,
			values[3],
			values[4],
		})
	}

	return fmt.Sprintf("%v\n%v", d.calc1(passwords), d.calc2(passwords))
}

func (d Day2) calc1(passwords []password) string {
	var sum = 0
	for _, password := range passwords {
		var count = strings.Count(password.password, password.char)
		if count >= password.min && count <= password.max {
			sum++
		}
	}
	return fmt.Sprintf("%d", sum)
}

func (d Day2) calc2(passwords []password) string {
	var sum = 0
	for _, password := range passwords {
		var occurrences = 0
		if password.min-1 < len(password.password) && password.password[password.min-1] == password.char[0] {
			occurrences++
		}
		if password.max-1 < len(password.password) && password.password[password.max-1] == password.char[0] {
			occurrences++
		}
		if occurrences == 1 {
			sum++
		}
	}
	return fmt.Sprintf("%d", sum)
}
