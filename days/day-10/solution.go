package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/elton-d/aoc/util"
)

func getLines() []string {
	return strings.Split(util.GetInputStr("https://adventofcode.com/2021/day/10/input"), "\n")
}

var openingBraces = map[rune]bool{
	'[': true,
	'(': true,
	'<': true,
	'{': true,
}

func isOpeningBrace(c rune) bool {
	_, ok := openingBraces[c]
	return ok
}

func matches(open, close rune) bool {
	switch open {
	case '(':
		return close == ')'
	case '{':
		return close == '}'
	case '[':
		return close == ']'
	case '<':
		return close == '>'
	}
	return false
}

func CheckCorrupted(s string) (bool, rune) {
	stack := []rune{}
	for _, c := range s {
		if isOpeningBrace(c) {
			stack = append(stack, c)
		} else {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if !matches(popped, c) {
				return true, c
			}
		}
	}
	return false, -1
}

func getComplement(s []rune) string {
	match := map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
		'<': '>',
	}
	c := []rune{}
	for i := len(s) - 1; i >= 0; i-- {
		c = append(c, match[s[i]])
	}
	return string(c)
}

func CheckIncomplete(s string) (bool, string) {
	stack := []rune{}
	for _, c := range s {
		if isOpeningBrace(c) {
			stack = append(stack, c)
		} else {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if !matches(popped, c) {
				return false, ""
			}
		}
	}
	return true, getComplement(stack)
}

func getCompletionScore(s string) int {
	pts := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	score := 0
	for _, c := range s {
		score *= 5
		score += pts[c]
	}
	return score
}

func getPoints(c rune) int {
	return map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}[c]
}

func main() {
	lines := getLines()

	score := 0

	for _, l := range lines {
		corr, c := CheckCorrupted(l)
		if corr {
			score += getPoints(c)
		}
	}
	fmt.Println(score)

	completionStrs := []string{}

	for _, l := range lines {
		incomplete, comp := CheckIncomplete(l)
		if incomplete {
			completionStrs = append(completionStrs, comp)
		}
	}

	completionScores := []int{}

	for _, comp := range completionStrs {
		completionScores = append(completionScores, getCompletionScore(comp))
	}

	sort.Ints(completionScores)
	fmt.Println(completionScores[len(completionScores)/2])

}
