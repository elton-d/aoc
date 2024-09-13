package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

func main() {
	input := util.GetInputStr("https://adventofcode.com/2023/day/4/input")
	total := 0.0

	cards := []*Card{}

	for _, line := range strings.Split(input, "\n") {
		c, err := cardFromStr(line)
		if err != nil {
			panic(err)
		}
		total += c.Points()
		cards = append(cards, c)
	}
	fmt.Println(total)
	fmt.Println(totalCards(cards))
}

type Card struct {
	cardNo     int
	winningNos []int
	numbers    []int
	matches    int
}

func (c *Card) Matches() int {
	if c.matches != -1 {
		return c.matches
	}
	m := make(map[int]bool)
	for _, n := range c.numbers {
		m[n] = true
	}

	matches := 0
	for _, w := range c.winningNos {
		if _, ok := m[w]; ok {
			matches++
		}
	}
	c.matches = matches
	return matches
}

func (c *Card) Points() float64 {
	matches := c.Matches()
	if matches > 0 {
		return math.Pow(2, float64(matches-1))
	}
	return 0
}

func cardFromStr(s string) (*Card, error) {
	c := &Card{winningNos: []int{}, numbers: []int{}, matches: -1}
	cardNo, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(strings.Split(s, ":")[0], "Card ")))
	if err != nil {
		return nil, err
	}
	c.cardNo = cardNo
	parts := strings.Split(strings.Split(s, ":")[1], "|")
	for _, nStr := range strings.Split(strings.TrimSpace(parts[0]), " ") {
		if nStr == "" {
			continue
		}
		n, err := strconv.Atoi(nStr)
		if err != nil {
			return nil, err
		}
		c.winningNos = append(c.winningNos, n)
	}
	for _, nStr := range strings.Split(strings.TrimSpace(parts[1]), " ") {
		if nStr == "" {
			continue
		}
		n, err := strconv.Atoi(nStr)
		if err != nil {
			return nil, err
		}
		c.numbers = append(c.numbers, n)
	}
	return c, nil
}

func totalCards(cards []*Card) int {
	queue := make([]*Card, len(cards))
	copy(queue, cards)
	initialCards := make(map[int]*Card)
	for _, c := range queue {
		initialCards[c.cardNo] = c
	}
	maxCardNo := queue[len(queue)-1].cardNo
	pos := 0

	for pos < len(queue) {
		c := queue[pos]
		matches := c.Matches()
		for i := c.cardNo + 1; i <= min(maxCardNo, c.cardNo+matches); i++ {
			queue = append(queue, initialCards[i])
		}
		pos++
	}
	return len(queue)
}
