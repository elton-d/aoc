package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

func getInitialState() ([]int, error) {
	b, err := util.GetInput("https://adventofcode.com/2021/day/6/input")
	if err != nil {
		return nil, err
	}
	numsStr := strings.Split(strings.TrimSpace(string(b)), ",")
	var nums []int
	for _, s := range numsStr {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func fishCount(state []int, days int) int {
	counts := make(map[int]int)
	for i := 0; i < 9; i++ {
		counts[i] = 0
	}
	for _, i := range state {
		counts[i] += 1
	}
	for i := 0; i < days; i++ {
		newFish := counts[0]
		for j := 0; j < 8; j++ {
			counts[j] = counts[j+1]
		}
		counts[6] += newFish
		counts[8] = newFish
	}
	sum := 0
	for _, i := range counts {
		sum += i
	}
	return sum
}

func main() {
	state, err := getInitialState()
	if err != nil {
		panic(err)
	}
	fmt.Println(fishCount(state, 80))
	fmt.Println(fishCount(state, 256))
}
