package main

import (
	"fmt"
	"strings"

	"github.com/elton-d/aoc/util"
)

type Rucksack struct {
	c1      string
	c2      string
	content string
}

func (r *Rucksack) findOverlap() rune {
	charMap := map[rune]bool{}
	for _, c := range r.c1 {
		charMap[c] = true
	}

	for _, c := range r.c2 {
		if _, ok := charMap[c]; ok {
			return c
		}
	}
	panic("no overlap found")
}

func (r *Rucksack) overlapPriority() int {
	overlap := r.findOverlap()
	return priority(overlap)
}

func priority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	}
	return int(item - 'A' + 27)
}

func parseInput(input string) []*Rucksack {
	rucksacks := []*Rucksack{}
	for _, l := range strings.Split(input, "\n") {
		size := len(l)
		rucksacks = append(rucksacks, &Rucksack{c1: l[:size/2], c2: l[size/2:], content: l})
	}
	return rucksacks
}

func sumOverlapPriorities(sacks []*Rucksack) int {
	sum := 0
	for _, s := range sacks {
		sum += s.overlapPriority()
	}
	return sum
}

func findBadge(sacks []*Rucksack) rune {
	maps := []map[rune]bool{}
	for _, s := range sacks {
		m := map[rune]bool{}
		for _, c := range s.content {
			m[c] = true
		}
		maps = append(maps, m)
	}

	for _, c := range sacks[0].content {
		count := 0
		for _, m := range maps {
			if _, ok := m[c]; !ok {
				break
			}
			count += 1
		}
		if count == 3 {
			return c
		}
	}
	panic("no common item across group rucksacks")
}

func sumBadgePriorities(sacks []*Rucksack) int {
	sum := 0
	for i := 0; i <= len(sacks)-3; i += 3 {
		sum += priority(findBadge(sacks[i : i+3]))
	}
	return sum
}

func main() {
	sacks := parseInput(util.GetInputStr("https://adventofcode.com/2022/day/3/input"))
	fmt.Println(sumOverlapPriorities(sacks))
	fmt.Println(sumBadgePriorities(sacks))
}
