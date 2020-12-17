package solution

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/constants"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"regexp"
	"strconv"
	"strings"
)

type Day16 struct{}

type field struct {
	text string
	min1 int
	max1 int
	min2 int
	max2 int
}

var (
	regexField = regexp.MustCompile("^([a-z ]+): (\\d+)-(\\d+) or (\\d+)-(\\d+)")
)

func parseField(line string) field {
	data := regexField.FindStringSubmatch(line)
	text := data[1]
	min1, err := strconv.Atoi(data[2])
	max1, err := strconv.Atoi(data[3])
	min2, err := strconv.Atoi(data[4])
	max2, err := strconv.Atoi(data[5])
	if err != nil {
		panic(err)
	}
	return field{
		text,
		min1,
		max1,
		min2,
		max2,
	}
}

type ticket []int

func parseTicket(line string) ticket {
	var result ticket
	for _, data := range strings.Split(line, ",") {
		num, err := strconv.Atoi(data)
		if err != nil {
			panic(err)
		}
		result = append(result, num)
	}
	return result
}

func (d Day16) Calc() string {
	pInput := strings.Split(input.ReadInputFile(16), strings.Repeat(constants.LineSeparator, 2))
	var fields []field
	for _, fieldLine := range strings.Split(pInput[0], constants.LineSeparator) {
		fields = append(fields, parseField(fieldLine))
	}

	myTicket := parseTicket(strings.Split(pInput[1], constants.LineSeparator)[1])
	var tickets []ticket
	for _, ticketLine := range strings.Split(pInput[2], constants.LineSeparator)[1:] {
		tickets = append(tickets, parseTicket(ticketLine))
	}

	return fmt.Sprintf("1: %v\n2: %v\n", d.sumInvalidTickets(fields, tickets), d.sumDeparture(fields, myTicket, tickets))
}

func (d *Day16) sumInvalidTickets(fields []field, tickets []ticket) int {
	sumInvalid := 0

	for _, ticket := range tickets {
		for _, value := range ticket {
			invalid := true
			for _, field := range fields {
				if value >= field.min1 && value <= field.max1 ||
					value >= field.min2 && value <= field.max2 {
					invalid = false
					break
				}
			}
			if invalid {
				sumInvalid += value
			}
		}
	}
	return sumInvalid
}

func (d *Day16) sumDeparture(fields []field, myTicket ticket, tickets []ticket) int {
	validTickets := d.getValidTickets(fields, tickets)
	fieldOrder := d.getFieldOrder(nil, fields, validTickets)
	product := 1
	regexDeparture := regexp.MustCompile("departure")
	for i, field := range fieldOrder {
		if regexDeparture.MatchString(field.text) {
			product *= myTicket[i]
		}
	}
	return product
}

//1213906609901 too low
//1213906609901

func (d *Day16) getFieldOrder(order []field, fields []field, tickets []ticket) []field {
	if len(fields) == 0 && d.validOrderUntilNow(order, tickets) {
		return order
	}
	for i, f := range fields {
		newOrder := make([]field, len(order), len(order)+1)
		copy(newOrder, order)
		newOrder = append(newOrder, f)
		if d.validOrderUntilNow(newOrder, tickets) {
			fieldsLeft := RemoveFieldAtIndex(fields, i)
			newOrder = d.getFieldOrder(newOrder, fieldsLeft, tickets)
			if newOrder != nil {
				return newOrder
			}
		}
	}
	return nil
}

func RemoveFieldAtIndex(s []field, index int) []field {
	ret := make([]field, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func (d *Day16) validOrderUntilNow(order []field, tickets []ticket) bool {
	//not working for 606 at index 0
	for _, ticket := range tickets {
		for i, field := range order {
			if !(ticket[i] >= field.min1 && ticket[i] <= field.max1) &&
				!(ticket[i] >= field.min2 && ticket[i] <= field.max2) {
				return false
			}
		}
	}

	return true
}

func (d *Day16) getValidTickets(fields []field, tickets []ticket) []ticket {
	var validTickets []ticket

	for _, ticket := range tickets {
		valid := true
		for _, value := range ticket {
			valueValid := false
			for _, field := range fields {
				if value >= field.min1 && value <= field.max1 ||
					value >= field.min2 && value <= field.max2 {
					valueValid = true
					break
				}
			}
			if !valueValid {
				valid = false
				break
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets
}
