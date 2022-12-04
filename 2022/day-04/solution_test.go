package main

import "testing"

func TestPart1(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	want := 2
	if got := Part1(input); got != want {
		t.Errorf("unexpected result, got: %d, want: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	want := 4
	if got := Part2(input); got != want {
		t.Errorf("unexpected result, got: %d, want: %d", got, want)
	}
}
