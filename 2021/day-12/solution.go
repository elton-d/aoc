package main

import (
	"fmt"
	"strings"

	"github.com/elton-d/aoc/util"
)

type Cave struct {
	val   string
	Edges map[string]*Cave
}

func (c *Cave) IsLargeCave() bool {
	return strings.ToUpper(c.val) == c.val
}

func (c *Cave) IsStart() bool {
	return c.val == "start"
}

func (c *Cave) IsEnd() bool {
	return c.val == "end"
}

func CavesFromStr(s string) map[string]*Cave {

	caves := make(map[string]*Cave)
	for _, edge := range strings.Split(s, "\n") {
		nodes := strings.Split(edge, "-")
		cave1Label := nodes[0]
		cave2Label := nodes[1]

		var c1, c2 *Cave

		if _, ok := caves[cave1Label]; !ok {
			caves[cave1Label] = &Cave{
				val:   cave1Label,
				Edges: make(map[string]*Cave),
			}
		}

		if _, ok := caves[cave2Label]; !ok {
			caves[cave2Label] = &Cave{
				val:   cave2Label,
				Edges: make(map[string]*Cave),
			}
		}

		c1, c2 = caves[cave1Label], caves[cave2Label]
		c1.Edges[cave2Label] = c2
		c2.Edges[cave1Label] = c1
	}
	return caves
}

func CountPaths(cs map[string]*Cave, part int) int {
	count := 0
	if part == 1 {
		count = recursiveSearch1(cs["start"], cs["end"], map[string]int{"start": 1})
	}
	if part == 2 {
		count = recursiveSearch(cs["start"], cs["end"], map[string]int{"start": 1}, false)
	}
	return count
}

func recursiveSearch(from, to *Cave, visitCount map[string]int, doubleSmall bool) int {
	if from == to {
		return 1
	}
	if !from.IsStart() && !from.IsLargeCave() && visitCount[from.val] > 1 {
		if doubleSmall {
			return 0
		}
	}
	count := 0
	doubleSmall = doubleSmall || (!from.IsLargeCave() && visitCount[from.val] > 1)

	for i, c := range from.Edges {
		if _, ok := visitCount[i]; !ok {
			visitCount[i] = 0
		}
		if c.IsStart() {
			continue
		}
		visitCount[i] += 1
		count += recursiveSearch(c, to, visitCount, doubleSmall)
		visitCount[i] -= 1
	}

	return count
}

func recursiveSearch1(from, to *Cave, visitCount map[string]int) int {
	if from == to {
		return 1
	}
	if !from.IsStart() && !from.IsLargeCave() && visitCount[from.val] > 1 {
		return 0
	}
	count := 0

	for i, c := range from.Edges {
		if _, ok := visitCount[i]; !ok {
			visitCount[i] = 0
		}
		if c.IsStart() {
			continue
		}
		visitCount[i] += 1
		count += recursiveSearch1(c, to, visitCount)
		visitCount[i] -= 1
	}

	return count
}

func Part1() {
	cs := CavesFromStr(util.GetInputStr("https://adventofcode.com/2021/day/12/input"))
	fmt.Println(CountPaths(cs, 1))
}

func Part2() {
	cs := CavesFromStr(util.GetInputStr("https://adventofcode.com/2021/day/12/input"))
	fmt.Println(CountPaths(cs, 2))
}

func main() {
	Part1()
	Part2()
}
