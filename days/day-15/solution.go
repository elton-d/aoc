package main

import (
	"container/heap"
	"fmt"
	"strings"

	"github.com/elton-d/aoc/util"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
)

type Cavern struct {
	riskLevels [][]int
}

type Coordinate struct {
	x, y int
}

type Node struct {
	dist  int
	prev  Coordinate
	curr  Coordinate
	index int
}

func (c *Cavern) PrintPath(sp map[Coordinate]*Node) {
	dest := Coordinate{x: len(c.riskLevels[0]) - 1, y: len(c.riskLevels) - 1}
	curr := sp[dest]

	path := make(map[Coordinate]bool)
	path[dest] = true
	stop := Coordinate{x: -1, y: -1}
	for curr.prev != stop {
		path[curr.prev] = true
		curr = sp[curr.prev]
	}
	for i, row := range c.riskLevels {
		for j, val := range row {
			if path[Coordinate{x: j, y: i}] {

				fmt.Print(colorRed, val, colorReset)
			} else {
				fmt.Print(val)
			}
		}
		fmt.Println()
	}
}

func (c *Cavern) FindShortestPath() int {
	shortestPaths := make(map[Coordinate]*Node)
	visited := make([][]bool, len(c.riskLevels))
	for i := range visited {
		visited[i] = make([]bool, len(c.riskLevels[i]))
	}
	visited[0][0] = true
	shortestPaths[Coordinate{x: 0, y: 0}] = &Node{
		dist: 0,
		prev: Coordinate{x: -1, y: -1},
	}

	queue := PriorityQueue{}
	heap.Push(&queue, &Node{curr: Coordinate{x: 0, y: 0}})
	enqueued := make(map[Coordinate]bool)
	for len(queue) > 0 {
		curr := heap.Pop(&queue).(*Node)
		delete(enqueued, curr.curr)
		visited[curr.curr.y][curr.curr.x] = true
		for _, n := range c.getNeighbors(curr.curr) {
			idx := len(queue)
			if !visited[n.y][n.x] {
				_, ok := enqueued[n]
				dist := shortestPaths[curr.curr].dist + c.RiskAt(n.x, n.y)
				val := Node{
					curr:  n,
					dist:  dist,
					prev:  curr.curr,
					index: idx,
				}
				if !ok {
					heap.Push(&queue, &val)
					enqueued[n] = true
				}

				_, ok = shortestPaths[n]
				if !ok {
					shortestPaths[n] = &val
				}

				if dist < shortestPaths[n].dist {
					shortestPaths[n] = &val
					queue.update(shortestPaths[n], curr.curr, dist)
				}
			}
		}
	}
	c.PrintPath(shortestPaths)
	return shortestPaths[Coordinate{x: len(c.riskLevels[0]) - 1, y: len(c.riskLevels) - 1}].dist
}

func (c *Cavern) RiskAt(x, y int) int {
	return c.riskLevels[y][x]
}

func NewCavernFromStr(s string) *Cavern {
	rl := [][]int{}
	for row, line := range strings.Split(s, "\n") {
		rl = append(rl, []int{})
		for _, val := range line {

			rl[row] = append(rl[row], int(val-'0'))
		}
	}
	return &Cavern{
		riskLevels: rl,
	}
}

func (c *Cavern) FiveX() {
	size := len(c.riskLevels)
	newRL := [][]int{}
	for i := 0; i < 5*len(c.riskLevels); i++ {
		newRL = append(newRL, []int{})
		for j := 0; j < 5*len(c.riskLevels[0]); j++ {
			idiff := i / size
			jdiff := j / size
			diff := idiff + jdiff
			newLevel := c.riskLevels[i-idiff*size][j-jdiff*size] + diff

			if newLevel > 9 {
				newLevel = newLevel % 9
			}

			newRL[i] = append(newRL[i], newLevel)
		}
	}
	c.riskLevels = newRL
}

func (c *Cavern) getNeighbors(pos Coordinate) []Coordinate {
	x, y := pos.x, pos.y
	neighbors := []Coordinate{}
	if x < len(c.riskLevels[0])-1 {
		neighbors = append(neighbors, Coordinate{x: x + 1, y: y})
	}
	if y < len(c.riskLevels)-1 {
		neighbors = append(neighbors, Coordinate{x: x, y: y + 1})
	}
	if x > 0 {
		neighbors = append(neighbors, Coordinate{x: x - 1, y: y})
	}
	if y > 0 {
		neighbors = append(neighbors, Coordinate{x: x, y: y - 1})
	}
	return neighbors
}

func Part1(input string) {
	c := NewCavernFromStr(input)
	fmt.Println(c.FindShortestPath())
}

func Part2(input string) {
	c := NewCavernFromStr(input)
	c.FiveX()
	fmt.Println(c.FindShortestPath())
}

func main() {
	input := util.GetInputStr("https://adventofcode.com/2021/day/15/input")
	Part1(input)
	Part2(input)
}
