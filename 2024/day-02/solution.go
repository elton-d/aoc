package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

func main() {
	input := util.GetInputStr("https://adventofcode.com/2024/day/2/input")

	reports, err := getReports(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(numSafe(reports, safe))
	fmt.Println(numSafe(reports, safeWithRemovals))

}

func getReports(input string) ([][]int, error) {
	var reports = [][]int{}
	for _, l := range strings.Split(input, "\n") {
		var report = []int{}
		for _, s := range strings.Split(l, " ") {
			if n, err := strconv.Atoi(s); err != nil {
				return nil, err
			} else {
				report = append(report, n)
			}
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func safe(report []int) bool {
	increasing := true
	decreasing := true
	for i, val := range report[1:] {
		diff := val - report[i]
		if diff == 0 || math.Abs(float64(diff)) > 3 {
			return false
		}
		decreasing = decreasing && (diff < 0)
		increasing = increasing && (diff > 0)

	}

	return increasing || decreasing
}

func safeWithRemovals(report []int) bool {
	if safe(report) {
		return true
	}

	for i := range report {
		var newSlice []int
		if i == len(report)-1 {
			newSlice = report[:i]
		} else {
			newSlice = slices.Concat(report[:i], report[i+1:])
		}
		if safe(newSlice) {
			return true
		}
	}
	return false
}

func numSafe(reports [][]int, checker func([]int) bool) int {
	var count int
	for _, r := range reports {
		if checker(r) {
			count++
		}
	}
	return count
}
