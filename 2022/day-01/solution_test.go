package main

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, 2)
	heap.Push(h, 1)
	heap.Push(h, 5)
	heap.Push(h, 3)
	ordered := []int{5, 3, 2, 1}

	for i, want := range ordered {
		if got := heap.Pop(h); got != want {
			t.Fatalf("unexpected value returned by pop, called: %d, got: %d, want: %d", i, got, want)
		}
	}
}

func TestPart1(t *testing.T) {
	want := 24000
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	if got := getMaxCals(input); got != want {
		t.Errorf("unexpected answer for part 1, got: %d, want: %d", got, want)
	}

}

func TestPart2(t *testing.T) {
	want := 45000
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	if got := getTopThreeCals(input); got != want {
		t.Errorf("unexpected answer for part 1, got: %d, want: %d", got, want)
	}
}
