package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

func getLists(in string) ([]int, []int, error) {
	var list1, list2 []int
	for _, l := range strings.Split(in, "\n") {
		parts := strings.Split(l, "   ")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid input: cannot split %q into two parts", l)
		}
		n1, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		n2, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	return list1, list2, nil
}

func totalDiff(l1, l2 []int) int {
	slices.Sort(l1)
	slices.Sort(l2)

	var diff int
	for i, n1 := range l1 {
		diff += int(math.Abs(float64(n1 - l2[i])))
	}
	return diff
}

func similarityScore(l1, l2 []int) int {
	var score int
	var counts = make(map[int]int)
	for _, n := range l2 {
		counts[n]++
	}
	for _, n := range l1 {
		score += n * counts[n]
	}
	return score
}

func main() {
	input := util.GetInputStr("https://adventofcode.com/2024/day/1/input")

	l1, l2, err := getLists(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %d, Part 2: %d", totalDiff(l1, l2), similarityScore(l1, l2))
}
