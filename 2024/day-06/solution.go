package main

import (
	"fmt"
	"strings"

	"github.com/elton-d/aoc/util"
)

const (
	up    = "^"
	down  = "v"
	right = ">"
	left  = "<"
)

type movement struct {
	rowChange int
	colChange int
	direction string
}

func (m *movement) rotate() {
	switch {
	case m.direction == up:
		m.direction = right
	case m.direction == right:
		m.direction = down
	case m.direction == down:
		m.direction = left
	case m.direction == left:
		m.direction = up
	}
	m.setDisplacements()
}

func (m *movement) setDisplacements() {
	switch m.direction {
	case up:
		m.rowChange = -1
		m.colChange = 0
	case right:
		m.rowChange = 0
		m.colChange = 1
	case down:
		m.rowChange = 1
		m.colChange = 0
	case left:
		m.rowChange = 0
		m.colChange = -1
	}
}

func (m *movement) forward(row, col int) (int, int) {
	return row + m.rowChange, col + m.colChange
}

func (m *movement) backward(row, col int) (int, int) {
	return row - m.rowChange, col - m.colChange
}

func (m *movement) copy() *movement {
	return &movement{
		rowChange: m.rowChange,
		colChange: m.colChange,
		direction: m.direction,
	}
}

func newMovement(dir string) *movement {
	m := &movement{direction: dir}
	m.setDisplacements()
	return m
}

func parseInput(input string) [][]string {
	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))
	for i, line := range lines {
		linesSlice := make([]string, len(line))
		for j, c := range line {
			linesSlice[j] = string(c)
		}
		grid[i] = linesSlice
	}
	return grid
}

func part1(grid [][]string) int {
	var count int
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	for i, row := range grid {
		for j, val := range row {
			if val == up || val == down || val == left || val == right {
				m := newMovement(val)
				search(i, j, grid, visited, m)
				break
			}
		}
	}
	for _, row := range visited {
		for _, val := range row {
			if val {
				count++
			}
		}
	}
	return count
}

func part2(grid [][]string) int {
	var startRow, startCol, obstructions int
	var m *movement
	for i, row := range grid {
		for j, val := range row {
			if val == up || val == down || val == left || val == right {
				m = newMovement(val)
				startRow = i
				startCol = j
			}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "." {
				grid[i][j] = "#"
				if stuck(startRow, startCol, grid, m.copy(), 0) {
					obstructions++
				}
				grid[i][j] = "."
			}
		}
	}
	return obstructions
}

func search(row int, col int, grid [][]string, visited [][]bool, m *movement) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return
	}

	switch grid[row][col] {
	case ".", up, down, left, right:
		visited[row][col] = true
		nextRow, nextCol := m.forward(row, col)
		search(nextRow, nextCol, grid, visited, m)
	case "#":
		prevRow, prevCol := m.backward(row, col)
		m.rotate()
		nextRow, nextCol := m.forward(prevRow, prevCol)
		search(nextRow, nextCol, grid, visited, m)
	}
}

func stuck(row int, col int, grid [][]string, m *movement, count int) bool {
	count++
	if count > len(grid)*len(grid[0]) {
		return true
	}
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return false
	}

	switch grid[row][col] {
	case ".", up, down, left, right:
		nextRow, nextCol := m.forward(row, col)
		return stuck(nextRow, nextCol, grid, m, count)
	case "#":
		prevRow, prevCol := m.backward(row, col)
		m.rotate()
		nextRow, nextCol := m.forward(prevRow, prevCol)
		return stuck(nextRow, nextCol, grid, m, count)
	}
	// should never reach here
	return false
}

func main() {
	input := util.GetInputStr("https://adventofcode.com/2024/day/6/input")

	fmt.Printf("Part 1: %d\n", part1(parseInput(input)))
	fmt.Printf("Part 2: %d\n", part2(parseInput(input)))
}
