package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/elton-d/aoc/util"
)

type HeightMap struct {
	hm [][]int
}

type MapCell struct {
	row, col, height int
}

func newMapCell(row, col, height int) *MapCell {
	return &MapCell{
		row:    row,
		col:    col,
		height: height,
	}
}

func (h *HeightMap) getNeighbors(m *MapCell) []*MapCell {
	neighbors := []*MapCell{}
	i, j := m.row, m.col

	if i > 0 {
		neighbors = append(neighbors, newMapCell(i-1, j, h.hm[i-1][j]))
	}

	if j > 0 {
		neighbors = append(neighbors, newMapCell(i, j-1, h.hm[i][j-1]))
	}

	if i < len(h.hm)-1 {
		neighbors = append(neighbors, newMapCell(i+1, j, h.hm[i+1][j]))
	}

	if j < len(h.hm[i])-1 {
		neighbors = append(neighbors, newMapCell(i, j+1, h.hm[i][j+1]))
	}
	return neighbors
}

func (h *HeightMap) isLowPoint(i, j int) bool {
	neighbors := h.getNeighbors(newMapCell(i, j, h.hm[i][j]))
	x := h.hm[i][j]

	for _, n := range neighbors {
		if n.height <= x {
			return false
		}
	}
	return true
}

func heightMapFromStr(s string) [][]int {
	hm := [][]int{}
	lines := strings.Split(strings.TrimSpace(s), "\n")

	for i, line := range lines {
		if len(hm) < i+1 {
			hm = append(hm, []int{})
		}

		for _, n := range line {
			hm[i] = append(hm[i], int(n-'0'))
		}
	}
	return hm
}

func getHeightMap() ([][]int, error) {
	b, err := util.GetInput("https://adventofcode.com/2021/day/9/input")
	if err != nil {
		return nil, err
	}

	return heightMapFromStr(string(b)), nil
}

func (h *HeightMap) getArea(m *MapCell, visited [][]bool) int {
	neighborsArea := 0
	visited[m.row][m.col] = true
	for _, n := range h.getNeighbors(m) {
		if !visited[n.row][n.col] && n.height != 9 {
			neighborsArea += h.getArea(n, visited)
		}
	}

	return neighborsArea + 1
}

func (h *HeightMap) GetProduct() int {
	lowPts := []*MapCell{}
	for i := range h.hm {
		for j := range h.hm[i] {
			if h.isLowPoint(i, j) {
				lowPts = append(lowPts, newMapCell(i, j, h.hm[i][j]))
			}
		}
	}

	areas := []int{}
	visited := make([][]bool, len(h.hm))
	for i := range visited {
		visited[i] = make([]bool, len(h.hm[i]))
	}

	for _, l := range lowPts {
		areas = append(areas, h.getArea(l, visited))
	}

	sort.Ints(areas)

	product := 1

	for _, a := range areas[len(areas)-3:] {
		product *= a
	}
	return product
}

func main() {
	hm, err := getHeightMap()
	if err != nil {
		panic(err)
	}

	h := &HeightMap{hm: hm}

	riskLevel := 0
	for i := range hm {
		for j := range hm[i] {
			if h.isLowPoint(i, j) {
				riskLevel += hm[i][j] + 1
			}
		}
	}
	fmt.Println(riskLevel)

	fmt.Println(h.GetProduct())
}
