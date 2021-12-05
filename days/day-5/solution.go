package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

type Point struct{ X, Y int }

type LineSegment struct {
	Pt1, Pt2 *Point
}

func (l *LineSegment) Slope() int {
	return (l.Pt2.Y - l.Pt1.Y) / (l.Pt2.X - l.Pt1.X)
}

type filterFunc func(l *LineSegment) bool
type Graph struct {
	pointToWeight map[Point]int
	overlaps      int
	filter        filterFunc
}

func (g *Graph) increaseWeight(p Point) {
	if val, ok := g.pointToWeight[p]; ok {
		if val == 1 {
			// count it since weight will increase from 1 to 2
			g.overlaps += 1
		}
	} else {
		g.pointToWeight[p] = 0
	}
	g.pointToWeight[p] += 1
}

func newGraph(f filterFunc) *Graph {
	return &Graph{
		filter:        f,
		pointToWeight: make(map[Point]int),
	}
}

func (g *Graph) DrawLineSegs(lineSegs []*LineSegment) {
	for _, l := range lineSegs {
		g.Plot(l)
	}
}

func (g *Graph) Plot(l *LineSegment) {
	if g.filter(l) {
		x1, y1, x2, y2 := l.Pt1.X, l.Pt1.Y, l.Pt2.X, l.Pt2.Y
		if x2 != x1 {
			var left, right *Point
			m := l.Slope()
			c := y2 - m*x2
			if x1 < x2 {
				left = l.Pt1
				right = l.Pt2
			} else {
				left = l.Pt2
				right = l.Pt1
			}

			for x := left.X; x <= right.X; x++ {
				g.increaseWeight(Point{X: x, Y: m*x + c})
			}

		} else {
			var bottom, top *Point
			if y1 < y2 {
				bottom = l.Pt1
				top = l.Pt2
			} else {
				bottom = l.Pt2
				top = l.Pt1
			}

			for y := bottom.Y; y <= top.Y; y++ {
				g.increaseWeight(Point{X: x1, Y: y})
			}
		}
	}
}

func newLineFromStr(s string) (*LineSegment, error) {
	coords := strings.Split(s, " -> ")
	pointsInt := []int{}

	pointsStr := []string{}
	for _, lineStr := range coords {
		pointsStr = append(pointsStr, strings.Split(lineStr, ",")...)
	}
	for _, p := range pointsStr {
		i, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		pointsInt = append(pointsInt, i)
	}
	var pt1, pt2 Point

	pt1.X, pt1.Y, pt2.X, pt2.Y = pointsInt[0], pointsInt[1], pointsInt[2], pointsInt[3]
	return &LineSegment{Pt1: &pt1, Pt2: &pt2}, nil

}

func processInput() ([]*LineSegment, error) {
	var lines []*LineSegment
	b, err := util.GetInput("https://adventofcode.com/2021/day/5/input")
	if err != nil {
		return nil, err
	}
	inputStr := string(b)
	for _, line := range strings.Split(strings.TrimSpace(inputStr), "\n") {
		l, err := newLineFromStr(line)
		if err != nil {
			return nil, err
		}
		lines = append(lines, l)
	}
	return lines, nil

}

func main() {
	lineSegs, err := processInput()
	if err != nil {
		panic(err)
	}

	part1Filter := func(l *LineSegment) bool {
		return l.Pt1.X == l.Pt2.X || l.Pt1.Y == l.Pt2.Y
	}

	g := newGraph(part1Filter)
	g.DrawLineSegs(lineSegs)
	fmt.Println(g.overlaps)

	part2Filter := func(l *LineSegment) bool {
		if l.Pt1.X == l.Pt2.X {
			return true
		}
		m := l.Slope()

		return m == 0 || m == 1 || m == -1
	}
	g = newGraph(part2Filter)
	g.DrawLineSegs(lineSegs)
	fmt.Println(g.overlaps)
}
