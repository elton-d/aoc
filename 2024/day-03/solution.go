package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

func main() {
	input := util.GetInputStr("https://adventofcode.com/2024/day/3/input")

	s, err := findSum(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %d\n", s)

	s, err = findSum2(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2: %d\n", s)
}

func findMatches(instrs string) ([]string, error) {
	matches := []string{}
	re, err := regexp.Compile(`mul\(\d{1,3}\,\d{1,3}\)`)
	if err != nil {
		return nil, err
	}
	for _, l := range strings.Split(instrs, "\n") {
		matches = slices.Concat(matches, re.FindAllString(l, -1))
	}
	return matches, nil
}

func findMatches2(instrs string) ([]string, error) {
	matches := []string{}
	re, err := regexp.Compile(`(mul\(\d{1,3}\,\d{1,3}\))|(do\(\))|(don\'t\(\))`)
	if err != nil {
		return nil, err
	}
	for _, l := range strings.Split(instrs, "\n") {
		matches = slices.Concat(matches, re.FindAllString(l, -1))
	}
	return matches, nil
}

func findSum2(instrs string) (int, error) {
	sum := 0
	keepCount := true
	matches, err := findMatches2(instrs)
	if err != nil {
		return -1, err
	}
	re := regexp.MustCompile(`\d+`)

	for _, m := range matches {
		if m == "do()" {
			keepCount = true
			continue
		} else if m == "don't()" {
			keepCount = false
			continue
		}
		res := 1
		if keepCount {
			for _, n := range re.FindAllString(m, -1) {
				if i, err := strconv.Atoi(n); err != nil {
					return -1, err
				} else {
					res *= i
				}
			}
			sum += res
		}
	}
	return sum, nil
}

func findSum(instrs string) (int, error) {
	sum := 0
	matches, err := findMatches(instrs)
	if err != nil {
		return -1, err
	}
	re := regexp.MustCompile(`\d+`)

	for _, m := range matches {
		res := 1
		for _, n := range re.FindAllString(m, -1) {
			if i, err := strconv.Atoi(n); err != nil {
				return -1, err
			} else {
				res *= i
			}
		}
		sum += res
	}
	return sum, nil
}
