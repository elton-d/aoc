package main

import "testing"

func TestSolution(t *testing.T) {
	inputStr := `2199943210
3987894921
9856789892
8767896789
9899965678`
	hm := &HeightMap{
		hm: heightMapFromStr(inputStr),
	}

	want := 1134
	got := hm.GetProduct()

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
