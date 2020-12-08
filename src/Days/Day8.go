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
	expCommand = regexp.MustCompile("([a-z]{3}) [+]?([-]?[0-9]*)")
)

type Day8 struct {
	IDay
}

type command struct {
	command  string
	argument int
}

func (d Day8) Calc() string {
	input := strings.Split(Input.ReadInputFile(8), Constants.LineSeparator)
	var commands []command

	for _, e := range input {
		data := expCommand.FindStringSubmatch(e)
		argument, err := strconv.Atoi(data[2])
		if err != nil {
			panic(err)
		}
		commands = append(commands, command{data[1], argument})
	}

	return fmt.Sprintf("Acc at loop: %v\nAcc at end: %v", accAtLoop(commands), accAtEnd(commands))
}

func accAtLoop(commands []command) int {
	acc := 0
	commandHistory := make(map[int]bool)

	for i := 0; i < len(commands); i++ {
		command := commands[i]
		switch command.command {
		case "acc":
			acc += command.argument
		case "jmp":
			commandHistory[i] = true
			i = i + command.argument
			if commandHistory[i] {
				return acc
			}
			i--
			continue
		}

		commandHistory[i] = true
	}

	return -1
}

func accAtEnd(commands []command) int {
	commandHistory := make(map[int]bool)

	var execute func(int, bool) int
	execute = func(start int, maySkip bool) int {
		acc := 0
		for i := start; i < len(commands); i++ {
			if commandHistory[i] {
				return -1
			}
			commandHistory[i] = true

			command := commands[i]
			switch command.command {
			case "acc":
				acc += command.argument
			case "jmp":
				if command.argument != 0 {
					newIndex := i + command.argument
					if newIndex < len(commands) && newIndex > 0 {
						valueIfJump := execute(newIndex, maySkip)
						if valueIfJump >= 0 {
							return valueIfJump + acc
						} else {
							if !maySkip {
								return -1
							} else {
								maySkip = false
							}
						}
					}
				}
				//case "nop":
				//	if command.argument != 0 && maySkip{
				//		newIndex := i + command.argument
				//		if newIndex < len(commands) && newIndex > 0 {
				//			valueIfJump := execute(newIndex, false)
				//			if valueIfJump >= 0 {
				//				return valueIfJump + acc
				//			}
				//		}
				//	}
			}

		}
		return acc
	}

	return execute(0, true)
}
