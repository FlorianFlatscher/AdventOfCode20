package solution

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/constants"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"regexp"
	"strconv"
	"strings"
)

type Day14 struct {
}

type ValueCommand struct {
	address int
	value   int
	maskTo0 int
	maskTo1 int
}

type AddressCommand struct {
	address      int
	value        int
	orMask       int
	floatingBits []int
}

func (d Day14) Calc() string {
	lines := strings.Split(input.ReadInputFile(14), constants.LineSeparator)
	valueMaskTo0, valueMaskTo1 := d.parseValueMask(lines[0])
	addressOrMask, addressFloatingBits := d.parseAddressMask(lines[0])
	var valueCommands []ValueCommand
	var addressCommands []AddressCommand
	for _, commandLine := range lines[1:] {
		if commandLine[1] == 'e' {
			data := regexp.MustCompile("\\[([0-9]+)\\] = (.*)$").FindStringSubmatch(commandLine)
			address, err := strconv.Atoi(data[1])
			value, err := strconv.Atoi(data[2])
			if err != nil {
				panic(err)
			}
			valueCommands = append(valueCommands, ValueCommand{
				address: address,
				value:   value,
				maskTo0: valueMaskTo0,
				maskTo1: valueMaskTo1,
			})
			addressCommands = append(addressCommands, AddressCommand{
				address:      address,
				value:        value,
				orMask:       addressOrMask,
				floatingBits: addressFloatingBits,
			})
		} else {
			valueMaskTo0, valueMaskTo1 = d.parseValueMask(commandLine)
			addressOrMask, addressFloatingBits = d.parseAddressMask(commandLine)
		}
	}

	return fmt.Sprintf("1: %v\n2: %v\n", d.executeValueMasks(valueCommands), d.executeAddressMasks(addressCommands))
}

func (d Day14) parseValueMask(mask string) (int, int) {
	mask = regexp.MustCompile("= (.*)$").FindStringSubmatch(mask)[1]
	var maskTo1 int
	var maskTo0 int
	for i, char := range mask {
		switch char {
		case '0':
			maskTo0 |= 1 << (len(mask) - i - 1)
		case '1':
			maskTo1 |= 1 << (len(mask) - i - 1)
		}
	}
	return maskTo0, maskTo1
}

func (d Day14) parseAddressMask(mask string) (int, []int) {
	mask = regexp.MustCompile("= (.*)$").FindStringSubmatch(mask)[1]
	var orMask int
	var floatingBits []int
	for i, char := range mask {
		switch char {
		case '1':
			orMask |= 1 << (len(mask) - i - 1)
		case 'X':
			floatingBits = append(floatingBits, len(mask)-1-i)
		}
	}
	return orMask, floatingBits
}

func (d Day14) executeValueMasks(commands []ValueCommand) int {
	memory := make(map[int]int)

	for _, command := range commands {
		value := command.value
		value |= command.maskTo1
		value &= ^command.maskTo0
		memory[command.address] = value
	}

	sum := 0
	for _, value := range memory {
		sum += value
	}
	return sum
}

func (d Day14) executeAddressMasks(commands []AddressCommand) int {
	memory := make(map[int]int)

	for _, command := range commands {
		address := command.address
		address |= command.orMask
		var writeAddresses func(int, int)
		writeAddresses = func(address int, index int) {
			if index < len(command.floatingBits) {
				targetMemoryAddress := address
				targetMemoryAddress &= ^(1 << command.floatingBits[index])
				writeAddresses(targetMemoryAddress, index+1)
				targetMemoryAddress |= 1 << command.floatingBits[index]
				writeAddresses(targetMemoryAddress, index+1)
			} else {
				memory[address] = command.value
			}
		}
		writeAddresses(address, 0)
	}

	sum := 0
	for _, value := range memory {
		sum += value
	}
	return sum
}
