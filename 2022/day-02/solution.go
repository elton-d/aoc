package main

import (
	"fmt"
	"strings"

	"github.com/elton-d/aoc/util"
)

type ShapeType string

type Shape struct {
	Type  ShapeType
	Value int
}

type Round struct {
	other   *Shape
	yours   *Shape
	outcome string
}

const (
	Rock     ShapeType = "rock"
	Paper    ShapeType = "paper"
	Scissors ShapeType = "scissors"

	Win  = 6
	Draw = 3
	Loss = 0

	RockSymbol     = "A"
	PaperSymbol    = "B"
	ScissorsSymbol = "C"

	LossSymbol = "X"
	DrawSymbol = "Y"
	WinSymbol  = "Z"
)

func outcome(r *Round) int {
	if r.yours.Type == Rock {
		switch r.other.Type {
		case Rock:
			return Draw
		case Paper:
			return Loss
		case Scissors:
			return Win
		}
	}
	if r.yours.Type == Paper {
		switch r.other.Type {
		case Rock:
			return Win
		case Paper:
			return Draw
		case Scissors:
			return Loss
		}
	}
	if r.yours.Type == Scissors {
		switch r.other.Type {
		case Rock:
			return Loss
		case Paper:
			return Win
		case Scissors:
			return Draw
		}
	}
	panic("invalid shape")
}

func score(r *Round) int {
	return r.yours.Value + outcome(r)
}

func totalScore(rounds []*Round) int {
	total := 0
	for _, r := range rounds {
		total += score(r)
	}
	return total
}

func parseInput(input string) []*Round {
	rounds := []*Round{}
	for _, l := range strings.Split(input, "\n") {
		parts := strings.Split(l, " ")
		if len(parts) != 2 {
			panic("invalid input")
		}
		rounds = append(rounds, &Round{
			other: shapeFactory(parts[0]),
			yours: shapeFactory(parts[1]),
		})
	}
	return rounds
}

func shapeFactory(symbol string) *Shape {
	switch symbol {
	case "A", "X":
		return &Shape{
			Type:  Rock,
			Value: 1,
		}
	case "B", "Y":
		return &Shape{
			Type:  Paper,
			Value: 2,
		}
	case "C", "Z":
		return &Shape{
			Type:  Scissors,
			Value: 3,
		}
	default:
		panic("invalid symbol")
	}
}

func parseInput2(input string) []*Round {
	rounds := []*Round{}
	for _, l := range strings.Split(input, "\n") {
		parts := strings.Split(l, " ")
		if len(parts) != 2 {
			panic("invalid input")
		}
		rounds = append(rounds, &Round{
			other:   shapeFactory(parts[0]),
			outcome: parts[1],
		})
	}
	return rounds
}

func pickShapeForOutcome(r *Round) {
	switch r.other.Type {
	case Rock:
		{
			switch r.outcome {
			case LossSymbol:
				r.yours = shapeFactory(ScissorsSymbol)
			case DrawSymbol:
				r.yours = shapeFactory(RockSymbol)
			case WinSymbol:
				r.yours = shapeFactory(PaperSymbol)
			}
		}
	case Paper:
		{
			switch r.outcome {
			case LossSymbol:
				r.yours = shapeFactory(RockSymbol)
			case DrawSymbol:
				r.yours = shapeFactory(PaperSymbol)
			case WinSymbol:
				r.yours = shapeFactory(ScissorsSymbol)
			}
		}
	case Scissors:
		{
			switch r.outcome {
			case LossSymbol:
				r.yours = shapeFactory(PaperSymbol)
			case DrawSymbol:
				r.yours = shapeFactory(ScissorsSymbol)
			case WinSymbol:
				r.yours = shapeFactory(RockSymbol)
			}
		}
	}
}

func score2(r *Round) int {
	pickShapeForOutcome(r)
	var outcomeScore int
	switch r.outcome {
	case "X":
		outcomeScore = 0
	case "Y":
		outcomeScore = 3
	case "Z":
		outcomeScore = 6
	}
	return r.yours.Value + outcomeScore
}

func totalScore2(rounds []*Round) int {
	total := 0
	for _, r := range rounds {
		total += score2(r)
	}
	return total
}

func main() {
	input := util.GetInputStr("https://adventofcode.com/2022/day/2/input")
	fmt.Println(totalScore(parseInput(input)))
	fmt.Println(totalScore2(parseInput2(input)))
}
