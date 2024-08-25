package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/elton-d/aoc/util"
)

func main() {
	input := util.GetInputStr("https://adventofcode.com/2023/day/1/input")
	total := 0
	for _, i := range strings.Split(input, "\n") {
		val, err := strconv.Atoi(decodeCalibration(i))
		if err != nil {
			panic(err)
		}
		total += val
	}
	fmt.Printf("%d\n", total)

	total = 0
	for _, i := range strings.Split(input, "\n") {
		val, err := strconv.Atoi(decodeCalibration(replaceSpelledDigits(i)))
		if err != nil {
			panic(err)
		}
		total += val
	}
	fmt.Printf("%d\n", total)
}

func decodeCalibration(val string) string {
	start, end := "", ""

	for _, s := range val {
		if unicode.IsDigit(s) {
			if start == "" {
				start = string(s)
			}
			end = string(s)
		}
	}
	return fmt.Sprintf("%s%s", start, end)
}

func digitToInt(word string) int {
	switch word {
	case "zero":
		return 0
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return -1
	}
}

func replaceSpelledDigits(s string) string {
	res := []string{}
	start := 0
	for i, c := range s {
		if unicode.IsDigit(c) {
			res = append(res, string(c))
			start = i + 1
			continue
		}
		for j := start; j < i+1; j++ {
			substr := s[j : i+1]
			if val := digitToInt(substr); val != -1 {
				res = append(res, fmt.Sprintf("%d", val))
			}
		}
	}
	return strings.Join(res, "")
}
