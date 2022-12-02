package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getMaxCals(calList string) int {
	max := 0
	for _, elfCals := range strings.Split(calList, "\n\n") {
		calSum := 0
		for _, cal := range strings.Split(elfCals, "\n") {
			n, err := strconv.Atoi(cal)
			if err != nil {
				panic(fmt.Errorf("unexpected non-integer input %s: %w", cal, err))
			}
			calSum += n
			if calSum > max {
				max = calSum
			}
		}
	}
	return max
}

func getTopThreeCals(calList string) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, elfCals := range strings.Split(calList, "\n\n") {
		calSum := 0
		for _, cal := range strings.Split(elfCals, "\n") {
			n, err := strconv.Atoi(cal)
			if err != nil {
				panic(fmt.Errorf("unexpected non-integer input %s: %w", cal, err))
			}
			calSum += n
		}
		heap.Push(h, calSum)
	}
	total := 0
	for i := 0; i < 3; i++ {
		total += heap.Pop(h).(int)
	}
	return total
}

func main() {
	input := util.GetInputStr("https://adventofcode.com/2022/day/1/input")
	fmt.Println(getMaxCals(input))
	fmt.Println()
	fmt.Println(getTopThreeCals(input))
}
