package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

func getCandidatePositions(pos []int) []int {
	candidates := []int{}
	sort.Ints(pos)
	n := len(pos)

	if n%2 == 1 {
		candidates = append(candidates, pos[(n-1)/2])
	} else {
		candidates = append(candidates, pos[n/2], pos[(n/2)-1])
	}

	return candidates
}

func calculateFuelUsed(pos []int, alignIdx int) int {
	fuel := 0
	for _, p := range pos {
		if alignIdx < p {
			fuel += p - alignIdx
		} else {
			fuel += alignIdx - p
		}
	}
	return fuel
}

func getFuelUsedForBestPos(pos []int) int {
	c := getCandidatePositions(pos)
	fuelUsed := math.MaxInt

	for _, i := range c {
		f := calculateFuelUsed(pos, i)
		if f < fuelUsed {
			fuelUsed = f
		}
	}

	return fuelUsed
}

func getInput() ([]int, error) {
	nums := []int{}
	b, err := util.GetInput("https://adventofcode.com/2021/day/7/input")
	if err != nil {
		return nil, err
	}

	in := strings.Split(strings.TrimSpace(string(b)), ",")
	for _, s := range in {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func mean(nums []int) int {
	sum := 0

	for _, i := range nums {
		sum += i
	}
	return sum / len(nums)
}

func summation(n int) int {
	return (n * (n + 1)) / 2
}

func getFuelUsedForBestPos2(pos []int) int {
	alignIdx := mean(pos)

	fuel := 0

	for _, i := range pos {
		if i < alignIdx {
			fuel += summation(alignIdx - i)
		} else {
			fuel += summation(i - alignIdx)
		}
	}

	return fuel
}

func main() {
	n, err := getInput()
	if err != nil {
		panic(err)
	}
	fmt.Println(getFuelUsedForBestPos(n))
	fmt.Println(getFuelUsedForBestPos2(n))
}
