package main

import "testing"

func TestPart1(t *testing.T) {
	input := `A Y
B X
C Z`
	want := 15
	got := totalScore(parseInput(input))

	if got != want {
		t.Errorf("unexpected total score, got: %d, want: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `A Y
B X
C Z`
	want := 12
	got := totalScore2(parseInput2(input))

	if got != want {
		t.Errorf("unexpected total score, got: %d, want: %d", got, want)
	}
}
