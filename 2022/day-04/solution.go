package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

type Range struct {
	start int
	end   int
}

func NewRange(r string) *Range {
	parts := strings.Split(r, "-")
	s, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	e, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return &Range{start: s, end: e}
}

func fullyContains(r1, r2 *Range) bool {
	return (r2.start >= r1.start && r2.end <= r1.end) || (r1.start >= r2.start && r1.end <= r2.end)
}

func overlaps(r1, r2 *Range) bool {
	return r1.end >= r2.start && r1.end <= r2.end || r2.end >= r1.start && r2.end <= r1.end
}

func Part1(input string) int {
	count := 0
	for _, l := range strings.Split(input, "\n") {
		parts := strings.Split(l, ",")
		r1 := NewRange(parts[0])
		r2 := NewRange(parts[1])
		if fullyContains(r1, r2) {
			count += 1
		}
	}
	return count
}

func Part2(input string) int {
	count := 0
	for _, l := range strings.Split(input, "\n") {
		parts := strings.Split(l, ",")
		r1 := NewRange(parts[0])
		r2 := NewRange(parts[1])
		if overlaps(r1, r2) {
			count += 1
		}
	}
	return count
}

func main() {
	input := util.GetInputStr("https://adventofcode.com/2022/day/4/input")
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
