package main

import "testing"

func TestSimilarityScore(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	want := 31

	l1, l2, err := getLists(input)
	if err != nil {
		t.Fatal(err)
	}
	if got := similarityScore(l1, l2); got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
