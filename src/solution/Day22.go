package solution

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/constants"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"math"
	"strconv"
	"strings"
)

type Day22 struct{}

type Deck []int

func newDeck(input string) *Deck {
	var cards Deck
	for _, line := range strings.Split(input, constants.LineSeparator)[1:] {
		cards = append(cards, mustAtoi(line))
	}
	return &cards
}

func (d *Deck) Pop() int {
	first := (*d)[0]
	*d = (*d)[1:]
	return first
}

func (d *Deck) Peek() int {
	return (*d)[0]
}

func (d *Deck) Append(x int) {
	*d = append(*d, x)
}

func mustAtoi(x string) int {
	res, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}
	return res
}

func (d Day22) Calc() string {
	par := strings.Split(input.ReadInputFile(22), strings.Repeat(constants.LineSeparator, 2))
	player1 := newDeck(par[0])
	player2 := newDeck(par[1])
	recPlayer1 := newDeck(par[0])
	recPlayer2 := newDeck(par[1])

	return fmt.Sprintf("1: %v\n2:%v\n", d.calculateScore(*player1, *player2), d.recursiveCombat(*recPlayer1, *recPlayer2))
}

func (d Day22) calculateScore(player1, player2 Deck) int {
	for len(player1) > 0 && len(player2) > 0 {
		cardPlayer1 := player1.Pop()
		cardPlayer2 := player2.Pop()
		if cardPlayer1 > cardPlayer2 {
			player1.Append(cardPlayer1)
			player1.Append(cardPlayer2)
		} else {
			player2.Append(cardPlayer2)
			player2.Append(cardPlayer1)
		}
	}

	var winner Deck
	if len(player1) > 0 {
		winner = player1
	} else {
		winner = player2
	}

	var score int
	for len(winner) > 0 {
		mult := len(winner)
		score += mult * winner.Pop()
	}

	return score
}

func (d Day22) recursiveCombat(player1, player2 Deck) int {
	var rec func(player1, player2 Deck) (int, int)
	rec = func(player1, player2 Deck) (int, int) {
		history := make(map[string]bool)

		for len(player1) > 0 && len(player2) > 0 {
			if _, ok := history[fmt.Sprint(player1, player2)]; ok {
				return 1, 0
			}
			history[fmt.Sprint(player1, player2)] = true

			cardPlayer1 := player1.Pop()
			cardPlayer2 := player2.Pop()

			if cardPlayer1 <= len(player1) && cardPlayer2 <= len(player2) {
				copyPlayer1 := Deck(make([]int, cardPlayer1))
				copyPlayer2 := Deck(make([]int, cardPlayer2))
				copy(copyPlayer1, player1)
				copy(copyPlayer2, player2)
				score1, score2 := rec(copyPlayer1, copyPlayer2)
				if score1 > score2 {
					player1.Append(cardPlayer1)
					player1.Append(cardPlayer2)
				} else {
					player2.Append(cardPlayer2)
					player2.Append(cardPlayer1)
				}

			} else {
				if cardPlayer1 > cardPlayer2 {
					player1.Append(cardPlayer1)
					player1.Append(cardPlayer2)
				} else {
					player2.Append(cardPlayer2)
					player2.Append(cardPlayer1)
				}
			}
		}

		var score1 int
		var score2 int

		for len(player1) > 0 {
			mult := len(player1)
			score1 += mult * player1.Pop()
		}

		for len(player2) > 0 {
			mult := len(player2)
			score2 += mult * player2.Pop()
		}

		return score1, score2
	}

	res1, res2 := rec(player1, player2)
	return int(math.Round(math.Max(float64(res1), float64(res2))))
}
