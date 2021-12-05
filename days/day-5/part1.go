package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

type LineSegment struct {
	x1, y1, x2, y2 int
}

func evaluateLine(space [][]int, l *LineSegment) [][]int {
	if l.x1 == l.x2 {
		if l.y2 < l.y1 {
			for i := l.y2; i <= l.y1; i++ {
				space[l.x1][i] += 1
			}
		} else {
			for i := l.y1; i <= l.y2; i++ {
				space[l.x1][i] += 1
			}
		}
	} else if l.y1 == l.y2 {
		if l.x2 < l.x1 {
			for i := l.x2; i <= l.x1; i++ {
				space[i][l.y1] += 1
			}
		} else {
			for i := l.x1; i <= l.x2; i++ {
				space[i][l.y1] += 1
			}
		}
	} else if slope := (l.y2 - l.y1) / (l.x2 - l.x1); slope == 1 || slope == -1 {

		if l.x2 < l.x1 {

			if l.y2 < l.y1 {
				for x, y := l.x2, l.y2; x <= l.x1; x++ {
					space[x][y] += 1
					y += 1
				}
			} else {
				for x, y := l.x2, l.y2; x <= l.x1; x++ {
					space[x][y] += 1
					y -= 1
				}
			}

		} else {

			if l.y2 < l.y1 {
				for x, y := l.x1, l.y1; x <= l.x2; x++ {
					space[x][y] += 1
					y -= 1
				}
			} else {
				for x, y := l.x1, l.y1; x <= l.x2; x++ {
					space[x][y] += 1
					y += 1
				}
			}

		}

	}
	return space
}

// part 1 solution
// func evaluateLine(space [][]int, l *LineSegment) [][]int {
// 	if l.x1 == l.x2 {
// 		if l.y2 < l.y1 {
// 			for i := l.y2; i <= l.y1; i++ {
// 				space[l.x1][i] += 1
// 			}
// 		} else {
// 			for i := l.y1; i <= l.y2; i++ {
// 				space[l.x1][i] += 1
// 			}
// 		}
// 	} else if l.y1 == l.y2 {
// 		if l.x2 < l.x1 {
// 			for i := l.x2; i <= l.x1; i++ {
// 				space[i][l.y1] += 1
// 			}
// 		} else {
// 			for i := l.x1; i <= l.x2; i++ {
// 				space[i][l.y1] += 1
// 			}
// 		}
// 	}
// 	return space
// }

func processInput() ([]*LineSegment, error) {
	var lines []*LineSegment
	b, err := util.GetInput("https://adventofcode.com/2021/day/5/input")
	if err != nil {
		return nil, err
	}
	inputStr := string(b)
	for _, line := range strings.Split(strings.TrimSpace(inputStr), "\n") {
		l := &LineSegment{}
		linesStr := strings.Split(line, " -> ")
		pointsInt := []int{}

		pointsStr := []string{}
		for _, lineStr := range linesStr {
			pointsStr = append(pointsStr, strings.Split(lineStr, ",")...)
		}

		for _, p := range pointsStr {
			i, err := strconv.Atoi(p)
			if err != nil {
				return nil, err
			}
			pointsInt = append(pointsInt, i)
		}

		l.x1, l.y1, l.x2, l.y2 = pointsInt[0], pointsInt[1], pointsInt[2], pointsInt[3]
		lines = append(lines, l)

	}
	return lines, nil

}

func overlapsCount(space [][]int) int {
	overlaps := 0
	for _, line := range space {
		for _, count := range line {
			if count > 1 {
				overlaps += 1
			}
		}
	}
	return overlaps
}

func evaluateLines(space [][]int, lines []*LineSegment) [][]int {
	for _, l := range lines {
		space = evaluateLine(space, l)
	}
	return space
}

func main() {
	space := make([][]int, 1000)
	for i := range space {
		space[i] = make([]int, 1000)
	}

	lines, err := processInput()
	if err != nil {
		panic(err)
	}

	space = evaluateLines(space, lines)
	fmt.Print(overlapsCount(space))
}
