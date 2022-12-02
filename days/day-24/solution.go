package main

import (
	"strconv"
	"strings"
)

type ALU struct {
	Variables map[string]int
	Inputs    []int
	inputPos  int
}

func (a *ALU) RunInstruction(i string) {
	parts := strings.Split(i, " ")
	var second int
	if len(parts) > 2 {
		val, err := strconv.Atoi(parts[2])
		if err != nil {
			second = a.Variables[parts[2]]
		} else {
			second = val
		}
	}
	switch parts[0] {
	case "inp":
		a.Variables[parts[1]] = a.Inputs[a.inputPos]
		a.inputPos += 1
	case "add":
		a.Variables[parts[1]] += second
	case "mul":
		a.Variables[parts[1]] *= second
	case "div":
		a.Variables[parts[1]] /= second
	case "mod":
		a.Variables[parts[1]] %= second
	case "eql":
		if a.Variables[parts[1]] == second {
			a.Variables[parts[1]] = 1
		} else {
			a.Variables[parts[1]] = 0
		}
	}
}

func (a *ALU) RunProgram(p string) {
	for _, i := range strings.Split(p, "\n") {
		a.RunInstruction(i)
	}
}

func NewALU(in []int) *ALU {
	return &ALU{
		Variables: map[string]int{
			"w": 0,
			"x": 0,
			"y": 0,
			"z": 0,
		},
		Inputs: in,
	}
}

func main() {

}
