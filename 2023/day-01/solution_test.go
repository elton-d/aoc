package main

import (
	"fmt"
	"strings"
	"testing"
)

var testInput2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestDecodeCalibration(t *testing.T) {
	inputs := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	want := []int{12, 38, 15, 77}
	for i, input := range strings.Split(inputs, "\n") {
		if got, want := decodeCalibration(input), fmt.Sprintf("%d", want[i]); got != want {
			t.Fatalf("got: %s; want: %s", got, want)
		}
	}
}

func TestReplaceSpelledDigits(t *testing.T) {
	want := []string{"219", "823", "123", "2134", "49872", "18234", "76"}
	for i, input := range strings.Split(testInput2, "\n") {
		got := replaceSpelledDigits(input)
		if got != want[i] {
			t.Fatalf("got : %s, want: %s", got, want[i])
		}
	}
}
