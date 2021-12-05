package main

import (
	"strings"
	"testing"
)

var testInput = strings.Split(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`, "\n")

func TestCalcO2(t *testing.T) {
	want := 23
	got := calcO2Rating(testInput)
	if got != want {
		t.Errorf("Unexpected O2 rating, got: %v, want: %v", got, want)
	}
}

func TestCalcCO2(t *testing.T) {
	want := 10
	got := calcCO2Rating(testInput)
	if got != want {
		t.Errorf("Unexpected CO2 rating, got: %v, want: %v", got, want)
	}
}
