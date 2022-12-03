package main

import "testing"

func TestSumOverlapPriorities(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	want := 157

	got := sumOverlapPriorities(parseInput(input))

	if got != want {
		t.Errorf("unexpected sum: got: %d, want: %d", got, want)
	}
}

func TestSumBadgePriorities(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	want := 70

	got := sumBadgePriorities(parseInput(input))

	if got != want {
		t.Errorf("unexpected sum: got: %d, want: %d", got, want)
	}
}
