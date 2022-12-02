package main

import (
	"testing"

	"github.com/elton-d/aoc/util"
)

func TestGame(t *testing.T) {
	input := `Player 1 starting position: 4
Player 2 starting position: 8`
	game := &Game{
		Players:       PlayersFromStr(input),
		Dice:          &DeterministicDice{ResetInterval: 100},
		TerminalScore: 1000,
	}
	want := 739785

	got := game.Play()
	if want != got {
		t.Errorf("unexpected value, got: %d, want: %d", got, want)
	}
}

func TestPart1(t *testing.T) {
	want := 855624
	got := Part1(util.GetInputStr("https://adventofcode.com/2021/day/21/input"))
	if want != got {
		t.Errorf("unexpected value, got: %d, want: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 187451244607486
	got := Part2(util.GetInputStr("https://adventofcode.com/2021/day/21/input"))
	if want != got {
		t.Errorf("unexpected value, got: %d, want: %d", got, want)
	}
}
