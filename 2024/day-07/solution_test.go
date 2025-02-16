package main

import "testing"

func TestPart1(t *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
	eqs, err := parseInput(input)
	if err != nil {
		t.Fatal(err)
	}

	want := 3749

	got := part1(eqs)

	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
	eqs, err := parseInput(input)
	if err != nil {
		t.Fatal(err)
	}

	want := 11387

	got := part2(eqs)

	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
