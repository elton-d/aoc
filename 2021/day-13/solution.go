package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

type Coordinate struct {
	x, y int
}

type Fold struct {
	axis string
	val  int
}

type Graph struct {
	points         map[Coordinate]bool
	boundX, boundY int
}

func (g *Graph) Plot(c *Coordinate) {
	g.points[*c] = true
}

func (g *Graph) FoldX(fold int) {
	prevBound := g.boundX
	g.boundX = fold - 1
	for x := 0; x <= g.boundX; x++ {
		for y := 0; y <= g.boundY; y++ {
			if g.IsMarked(x, y) || g.IsMarked(prevBound-x, y) {
				g.Plot(&Coordinate{x: x, y: y})
			}
		}
	}
}

func (g *Graph) FoldY(fold int) {
	prevBound := g.boundY
	g.boundY = fold - 1
	for x := 0; x <= g.boundX; x++ {
		for y := 0; y <= g.boundY; y++ {
			if g.IsMarked(x, y) || g.IsMarked(x, prevBound-y) {
				g.Plot(&Coordinate{x: x, y: y})
			}
		}
	}
}

func (g *Graph) String() string {
	sb := &strings.Builder{}
	for y := 0; y <= g.boundY; y++ {
		line := []string{}
		for x := 0; x <= g.boundX; x++ {
			if g.IsMarked(x, y) {
				line = append(line, "#")
			} else {
				line = append(line, ".")
			}
		}
		sb.WriteString(strings.Join(line, ""))
		sb.WriteString("\n")
	}
	return strings.TrimSpace(sb.String())
}

func (g *Graph) IsMarked(x, y int) bool {
	_, ok := g.points[Coordinate{x: x, y: y}]
	return ok
}

func (g *Graph) performFolds(folds []Fold) {
	for _, fold := range folds {
		if fold.axis == "x" {
			g.FoldX(fold.val)
		} else {
			g.FoldY(fold.val)
		}
	}
}

func (g *Graph) CountDots() int {
	count := 0
	for y := 0; y <= g.boundY; y++ {
		for x := 0; x <= g.boundX; x++ {
			if g.IsMarked(x, y) {
				count += 1
			}
		}
	}
	return count
}

func NewGraphFromStr(s string) (*Graph, error) {
	g := &Graph{
		points: make(map[Coordinate]bool),
	}
	g.boundX, g.boundY = math.MinInt, math.MinInt
	for _, line := range strings.Split(s, "\n") {
		coords := strings.Split(line, ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			return nil, err
		}
		if x > g.boundX {
			g.boundX = x
		}
		if y > g.boundY {
			g.boundY = y
		}
		g.points[Coordinate{x, y}] = true
	}
	return g, nil
}

func processFoldInstructions(foldsStr string) ([]Fold, error) {
	folds := []Fold{}
	for _, line := range strings.Split(foldsStr, "\n") {
		splits := strings.Split(strings.Fields(line)[2], "=")
		val, err := strconv.Atoi(splits[1])
		if err != nil {
			return nil, err
		}
		folds = append(folds, Fold{axis: splits[0], val: val})
	}
	return folds, nil
}

func Part1(pointsStr, foldsStr string) {
	g, err := NewGraphFromStr(pointsStr)
	if err != nil {
		panic(err)
	}
	folds, err := processFoldInstructions(foldsStr)
	if err != nil {
		panic(err)
	}
	g.performFolds(folds[:1])
	fmt.Println(g.CountDots())
}

func Part2(pointsStr, foldsStr string) {
	g, err := NewGraphFromStr(pointsStr)
	if err != nil {
		panic(err)
	}
	folds, err := processFoldInstructions(foldsStr)
	if err != nil {
		panic(err)
	}
	g.performFolds(folds)
	fmt.Println(g.String())
}
func main() {
	input := util.GetInputStr("https://adventofcode.com/2021/day/13/input")
	splits := strings.Split(input, "\n\n")
	pointsStr := splits[0]
	foldsStr := splits[1]
	Part1(pointsStr, foldsStr)
	Part2(pointsStr, foldsStr)
}
