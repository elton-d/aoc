package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

type game struct {
	rounds    []*round
	redSeen   int
	greenSeen int
	blueSeen  int
}

type round struct {
	red   int
	green int
	blue  int
}

func parseRound(str string) (*round, error) {
	t := &round{}

	for _, p := range strings.Split(str, ",") {
		countColor := strings.Split(strings.TrimSpace(p), " ")
		count, err := strconv.Atoi(countColor[0])
		if err != nil {
			return nil, err
		}
		switch countColor[1] {
		case "red":
			t.red = count
		case "blue":
			t.blue = count
		case "green":
			t.green = count
		}

	}
	return t, nil
}

func (g *game) isPossible(loadedRed, loadedGreen, loadedBlue int) bool {
	return g.blueSeen <= loadedBlue && g.greenSeen <= loadedGreen && g.redSeen <= loadedRed
}

func (g *game) playRound(r *round) {
	g.rounds = append(g.rounds, r)
	g.blueSeen = max(g.blueSeen, r.blue)
	g.greenSeen = max(g.greenSeen, r.green)
	g.redSeen = max(g.redSeen, r.red)
}

func (g *game) power() int {
	return g.redSeen * g.blueSeen * g.greenSeen
}

func main() {
	input := util.GetInputStr("https://adventofcode.com/2023/day/2/input")

	gameNoTotal := 0
	powerTotal := 0

	for _, l := range strings.Split(input, "\n") {
		parts := strings.Split(l, ":")
		gameNo, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
		if err != nil {
			panic(err)
		}
		g := &game{}
		for _, rstr := range strings.Split(parts[1], ";") {
			r, err := parseRound(rstr)
			if err != nil {
				panic(err)
			}
			g.playRound(r)
		}

		if g.isPossible(12, 13, 14) {
			gameNoTotal += gameNo
		}

		powerTotal += g.power()
	}

	fmt.Println(gameNoTotal) // want: 2913
	fmt.Println(powerTotal)  // want: 55593
}
