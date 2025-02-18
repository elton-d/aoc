package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/elton-d/aoc/util"
)

func main() {
	input := util.GetInputStr("https://adventofcode.com/2024/day/8/input")

	fmt.Printf("Part 1: %d\n", part1(parseInput(input)))
	fmt.Printf("Part 2: %d\n", part2(parseInput(input)))
}

type location struct {
	row float64
	col float64
}

func parseInput(in string) [][]string {
	var grid [][]string
	lines := strings.Split(in, "\n")
	grid = make([][]string, len(lines))

	for i, l := range lines {
		grid[i] = make([]string, len(l))

		for j, c := range l {
			grid[i][j] = string(c)
		}
	}
	return grid
}

func part1(grid [][]string) int {
	var count int
	freqToLoc := make(map[string][]location)
	antinodes := make(map[location]struct{})

	for i, row := range grid {
		for j, val := range row {
			char := []rune(val)[0]
			if unicode.IsLetter(char) || unicode.IsDigit(char) {
				freqToLoc[val] = append(freqToLoc[val], location{float64(i), float64(j)})
			}
		}
	}

	for _, locations := range freqToLoc {
		for i, loc1 := range locations {
			if i == len(locations)-1 {
				break
			}
			for _, loc2 := range locations[i+1:] {
				an1, an2 := computeAntinodes(loc1, loc2)
				antinodes[an1] = struct{}{}
				antinodes[an2] = struct{}{}
			}
		}
	}

	for an := range antinodes {
		if an.row >= 0 && an.row < float64(len(grid)) && an.col >= 0 && an.col < float64(len(grid[0])) {
			count++
		}
	}
	return count
}

func part2(grid [][]string) int {
	var count int
	freqToLoc := make(map[string][]location)
	antinodes := make(map[location]struct{})

	for i, row := range grid {
		for j, val := range row {
			char := []rune(val)[0]
			if unicode.IsLetter(char) || unicode.IsDigit(char) {
				freqToLoc[val] = append(freqToLoc[val], location{float64(i), float64(j)})
			}
		}
	}

	for _, locations := range freqToLoc {
		for i, loc1 := range locations {
			if i == len(locations)-1 {
				break
			}
			for _, loc2 := range locations[i+1:] {
				for _, an := range computeAntinodes2(loc1, loc2, grid) {
					antinodes[an] = struct{}{}
				}
			}
		}
	}

	for an := range antinodes {
		if an.row >= 0 && an.row < float64(len(grid)) && an.col >= 0 && an.col < float64(len(grid[0])) {
			count++
		}
	}
	return count
}

func computeAntinodes(loc1, loc2 location) (location, location) {
	return computeAntiNode(loc1, loc2), computeAntiNode(loc2, loc1)
}

func computeAntiNode(loc1, loc2 location) location {
	x1 := float64(loc1.col)
	y1 := float64(loc1.row)

	x2 := float64(loc2.col)
	y2 := float64(loc2.row)

	x := x1 - (x2 - x1)
	y := y1 - (y2 - y1)

	return location{y, x}
}

func computeAntinodes2(loc1 location, loc2 location, grid [][]string) []location {
	var antinodes []location
	x1 := float64(loc1.col)
	y1 := float64(loc1.row)

	x2 := float64(loc2.col)
	y2 := float64(loc2.row)

	xDiff := x2 - x1
	yDiff := y2 - y1

	antinodes = append(antinodes, loc1, loc2)

	curr := loc1
	for {
		an := location{curr.row - yDiff, curr.col - xDiff}
		if an.row < 0 || an.row >= float64(len(grid)) || an.col < 0 || an.col >= float64(len(grid[0])) {
			break
		}
		antinodes = append(antinodes, an)
		curr = an
	}

	curr = loc2
	for {
		an := location{curr.row + yDiff, curr.col + xDiff}
		if an.row < 0 || an.row >= float64(len(grid)) || an.col < 0 || an.col >= float64(len(grid[0])) {
			break
		}
		antinodes = append(antinodes, an)
		curr = an
	}
	return antinodes
}
