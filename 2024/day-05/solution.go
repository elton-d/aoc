package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

func main() {
	input := util.GetInputStr("https://adventofcode.com/2024/day/5/input")

	sum, err := part1(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %d\n", sum)
	sum, err = part2(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2: %d\n", sum)
}

type rule struct {
	l int
	r int
}

func (r *rule) eval(pos map[int]int) bool {
	var lIdx, rIdx int
	var ok bool
	if lIdx, ok = pos[r.l]; !ok {
		return true
	}
	if rIdx, ok = pos[r.r]; !ok {
		return true
	}

	return lIdx < rIdx
}

func NewRule(s string) (*rule, error) {
	parts := strings.Split(s, "|")
	if len(parts) != 2 {
		return nil, fmt.Errorf("bad string %q", s)
	}
	var l, r int
	var err error
	if l, err = strconv.Atoi(parts[0]); err != nil {
		return nil, err
	}
	if r, err = strconv.Atoi(parts[1]); err != nil {
		return nil, err
	}
	return &rule{l: l, r: r}, nil
}

func parseInput(in string) ([]*rule, [][]int, error) {
	var rules []*rule
	var updates [][]int
	parts := strings.Split(in, "\n\n")
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("input processing error")
	}
	for _, line := range strings.Split(parts[0], "\n") {
		rule, err := NewRule(line)
		if err != nil {
			return nil, nil, err
		}
		rules = append(rules, rule)
	}

	for _, line := range strings.Split(parts[1], "\n") {
		update := []int{}
		for _, s := range strings.Split(line, ",") {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, nil, err
			}
			update = append(update, n)
		}
		updates = append(updates, update)
	}
	return rules, updates, nil
}

func part1(input string) (int, error) {
	var sum int
	rules, updates, err := parseInput(input)
	if err != nil {
		return -1, err
	}
	for _, update := range updates {
		valid := true
		pos := make(map[int]int)
		for i, n := range update {
			pos[n] = i
		}
		for _, r := range rules {
			if !r.eval(pos) {
				valid = false
				break
			}
		}
		if valid {
			sum += update[len(update)/2]
		}
	}
	return sum, nil
}

func part2(input string) (int, error) {
	var sum int
	rules, updates, err := parseInput(input)
	if err != nil {
		return -1, err
	}

	numToRules := make(map[int][]*rule)
	for _, rule := range rules {
		numToRules[rule.l] = append(numToRules[rule.l], rule)
		numToRules[rule.r] = append(numToRules[rule.r], rule)
	}

	for _, update := range updates {
		pos := make(map[int]int)
		for i, n := range update {
			pos[n] = i
		}
		for _, r := range rules {
			if !r.eval(pos) {
				numToIdx := make(map[int]int)
				for _, r := range getApplicableRules(update, numToRules) {
					numToIdx[r.r] += 1
				}
				mid := len(update) / 2
				for n, idx := range numToIdx {
					if idx == mid {
						sum += n
						break
					}
				}
				break
			}
		}

	}
	return sum, nil
}

func getApplicableRules(update []int, numToRule map[int][]*rule) []*rule {
	allRules := []*rule{}
	updateMap := make(map[int]struct{})
	for _, n := range update {
		updateMap[n] = struct{}{}
	}

	for _, n := range update {
		rules := numToRule[n]
		for _, r := range rules {
			if _, ok := updateMap[r.l]; !ok {
				continue
			}
			if _, ok := updateMap[r.r]; !ok {
				continue
			}
			allRules = append(allRules, r)
		}
	}
	m := make(map[rule]struct{})
	for _, r := range allRules {
		m[*r] = struct{}{}
	}
	var deduped []*rule

	for r := range m {
		deduped = append(deduped, &r)
	}
	return deduped
}
