package main

import "testing"

func TestCheckCorrupted(t *testing.T) {
	tests := []struct {
		in        string
		corrupted bool
		char      rune
	}{
		{
			in:        "{([(<{}[<>[]}>{[]{[(<()>",
			corrupted: true,
			char:      '}',
		},
		{
			in:        "[[<[([]))<([[{}[[()]]]",
			corrupted: true,
			char:      ')',
		},
		{
			in:        "[{[{({}]{}}([{[{{{}}([]",
			corrupted: true,
			char:      ']',
		},
		{
			in:        "[({(<(())[]>[[{[]{<()<>>",
			corrupted: false,
			char:      -1,
		},
	}

	for _, tc := range tests {
		isCorrupted, char := CheckCorrupted(tc.in)
		if tc.corrupted != isCorrupted {
			t.Errorf("unexpected value for corrupted, got: %v, want: %v", isCorrupted, tc.corrupted)
		}
		if tc.char != char {
			t.Errorf("unexpected char, got: %v, want: %v", char, tc.char)
		}
	}
}

func TestCheckIncomplete(t *testing.T) {
	tests := []struct {
		in         string
		incomplete bool
		comp       string
	}{
		{
			in:         "[({(<(())[]>[[{[]{<()<>>",
			incomplete: true,
			comp:       "}}]])})]",
		},
		{
			in:         "[(()[<>])]({[<{<<[]>>(",
			incomplete: true,
			comp:       ")}>]})",
		},
		{
			in:         "(((({<>}<{<{<>}{[]{[]{}",
			incomplete: true,
			comp:       "}}>}>))))",
		},
		{
			in:         "[{[{({}]{}}([{[{{{}}([]",
			incomplete: false,
			comp:       "",
		},
	}
	for _, tc := range tests {
		isIncomplete, comp := CheckIncomplete(tc.in)
		if tc.incomplete != isIncomplete {
			t.Errorf("unexpected value for incomplete, got: %v, want: %v", isIncomplete, tc.incomplete)
		}
		if tc.comp != comp {
			t.Errorf("unexpected comp, got: %v, want: %v", comp, tc.comp)
		}
	}
}
