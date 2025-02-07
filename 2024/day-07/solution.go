package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

type equation struct {
	target   int
	operands []int
}

type operation struct {
	applyFn func(x, y int) int
}

func (o *operation) apply(x, y int) int {
	return o.applyFn(x, y)
}

var add = &operation{applyFn: func(x, y int) int { return x + y }}
var mul = &operation{applyFn: func(x, y int) int { return x * y }}
var concat = &operation{applyFn: func(x, y int) int {
	res, err := strconv.Atoi(fmt.Sprintf("%d%d", x, y))
	if err != nil {
		panic(err)
	}
	return res
}}

func (e *equation) String() string {
	return fmt.Sprintf("{target: %d; operands: %v}", e.target, e.operands)
}

func parseInput(input string) ([]*equation, error) {
	lines := strings.Split(input, "\n")
	var err error
	equations := make([]*equation, len(lines))
	for i, line := range lines {
		eq := &equation{}
		parts := strings.Split(line, ":")
		eq.target, err = strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		nums := strings.Split(strings.TrimSpace(parts[1]), " ")
		eq.operands = make([]int, len(nums))
		for j, s := range nums {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			eq.operands[j] = n
		}
		equations[i] = eq
	}
	return equations, nil
}

func checkRec(eq *equation, curr int, idx int, op *operation, validOps []*operation) bool {
	if idx == len(eq.operands) {
		return curr == eq.target
	}

	val := op.apply(curr, eq.operands[idx])
	for _, o := range validOps {
		if checkRec(eq, val, idx+1, o, validOps) {
			return true
		}
	}
	return false
}

func part1(eqs []*equation) int {
	var sum int
	validOps := []*operation{add, mul}

	for _, eq := range eqs {
		if checkRec(eq, 0, 0, add, validOps) || checkRec(eq, 1, 0, mul, validOps) {
			sum += eq.target
		}
	}
	return sum
}

func part2(eqs []*equation) int {
	var sum int
	validOps := []*operation{add, mul, concat}

	for _, eq := range eqs {
		if checkRec(eq, 0, 0, add, validOps) || checkRec(eq, 1, 0, mul, validOps) || checkRec(eq, 0, 0, concat, validOps) {
			sum += eq.target
		}
	}
	return sum
}

func main() {
	input := util.GetInputStr("https://adventofcode.com/2024/day/7/input")

	eqs, err := parseInput(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %d\n", part1(eqs))
	fmt.Printf("Part 2: %d\n", part2(eqs))
}
