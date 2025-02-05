package main

import (
	"fmt"
	"strings"

	"github.com/elton-d/aoc/util"
)

type nextFn func(row, col int) (int, int)

func main() {
	input := util.GetInputStr("https://adventofcode.com/2024/day/4/input")
	target := "XMAS"
	grid := parseInput(input)

	fmt.Printf("Found %q %d times\n", target, wordSearch(grid, target))
	fmt.Printf("Found an X-MAS %d times\n", part2(grid))
}

func wordSearch(grid [][]string, target string) int {
	var total int
	for row := range grid {
		for col := range grid[row] {
			total += searchRec(row, col, grid, target, "", func(r, c int) (int, int) { return r - 1, c - 1 })
			total += searchRec(row, col, grid, target, "", func(r, c int) (int, int) { return r - 1, c })
			total += searchRec(row, col, grid, target, "", func(r, c int) (int, int) { return r - 1, c + 1 })
			total += searchRec(row, col, grid, target, "", func(r, c int) (int, int) { return r, c - 1 })
			total += searchRec(row, col, grid, target, "", func(r, c int) (int, int) { return r, c + 1 })
			total += searchRec(row, col, grid, target, "", func(r, c int) (int, int) { return r + 1, c - 1 })
			total += searchRec(row, col, grid, target, "", func(r, c int) (int, int) { return r + 1, c })
			total += searchRec(row, col, grid, target, "", func(r, c int) (int, int) { return r + 1, c + 1 })
		}
	}
	return total
}

func searchRec(row int, col int, grid [][]string, target string, curr string, next nextFn) int {
	curr += grid[row][col]
	if curr == target {
		return 1
	} else if len(curr) == len(target) {
		return 0
	}
	if grid[row][col] != string(target[len(curr)-1]) {
		return 0
	}
	nextRow, nextCol := next(row, col)
	if nextRow < 0 || nextCol < 0 || nextRow >= len(grid) || nextCol >= len(grid[row]) {
		return 0
	}
	return searchRec(nextRow, nextCol, grid, target, curr, next)
}

func part2(grid [][]string) int {
	var matches int
	for row := range grid {
		for col := range grid[row] {
			if row < 1 || row > len(grid)-2 || col < 1 || col > len(grid[row])-2 {
				continue
			}
			if grid[row][col] == "A" {
				if (grid[row-1][col-1] == "M" && grid[row-1][col+1] == "M" && grid[row+1][col-1] == "S" && grid[row+1][col+1] == "S") ||
					(grid[row-1][col-1] == "M" && grid[row-1][col+1] == "S" && grid[row+1][col-1] == "M" && grid[row+1][col+1] == "S") ||
					(grid[row-1][col-1] == "S" && grid[row-1][col+1] == "S" && grid[row+1][col-1] == "M" && grid[row+1][col+1] == "M") ||
					(grid[row-1][col-1] == "S" && grid[row-1][col+1] == "M" && grid[row+1][col-1] == "S" && grid[row+1][col+1] == "M") {
					matches++
				}
			}
		}
	}
	return matches
}

func parseInput(in string) [][]string {
	lines := strings.Split(in, "\n")
	grid := make([][]string, len(lines))
	for row, line := range lines {
		grid[row] = make([]string, len(line))
		for col, char := range line {
			grid[row][col] = string(char)
		}
	}
	return grid
}
