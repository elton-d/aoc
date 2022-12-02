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
