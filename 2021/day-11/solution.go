package main

import (
	"fmt"
	"strings"

	"github.com/elton-d/aoc/util"
)

type OctopusSimulator struct {
	octopuses         [][]*Octopus
	step              int
	totalFlashes      int
	currentlyFlashing []*Octopus
}

func (s *OctopusSimulator) String() string {
	var sb strings.Builder

	for _, row := range s.octopuses {
		for _, o := range row {
			sb.WriteString(fmt.Sprintf("%d", o.energyLevel))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (s *OctopusSimulator) Steps(n int) {
	gridSize := len(s.octopuses)
	for i := 0; i < n; i++ {
		s.currentlyFlashing = []*Octopus{}
		s.step += 1
		s.performOperations()
		if len(s.currentlyFlashing) == gridSize*gridSize {
			return
		}
		for _, o := range s.currentlyFlashing {
			o.flashing = false
		}
	}
}

func (s *OctopusSimulator) performOperations() {
	for _, row := range s.octopuses {
		for _, o := range row {
			o.energyLevel += 1

		}
	}

	for i, row := range s.octopuses {
		for j, o := range row {
			s.checkForFlash(o, i, j)
		}
	}

}

func (s *OctopusSimulator) checkForFlash(o *Octopus, row, col int) {
	if o.energyLevel > 9 {
		s.totalFlashes += 1
		o.energyLevel = 0
		o.flashing = true
		s.currentlyFlashing = append(s.currentlyFlashing, o)
		neighbors := s.getAdjacentOctopuses(row, col)
		for _, n := range neighbors {
			if !n.flashing {
				n.energyLevel += 1
				s.checkForFlash(n, n.row, n.col)
			}
		}
	}
}

func (s *OctopusSimulator) getAdjacentOctopuses(row, col int) []*Octopus {
	octopuses := []*Octopus{}

	gridSize := len(s.octopuses)

	if row > 0 {
		octopuses = append(octopuses, s.octopuses[row-1][col])
		if col > 0 {
			octopuses = append(octopuses, s.octopuses[row-1][col-1])
		}

		if col < gridSize-1 {
			octopuses = append(octopuses, s.octopuses[row-1][col+1])
		}
	}

	if col > 0 {
		octopuses = append(octopuses, s.octopuses[row][col-1])
		if row < gridSize-1 {
			octopuses = append(octopuses, s.octopuses[row+1][col-1])
		}
	}

	if row < len(s.octopuses)-1 {
		octopuses = append(octopuses, s.octopuses[row+1][col])
	}

	if col < gridSize-1 {
		octopuses = append(octopuses, s.octopuses[row][col+1])

		if row < gridSize-1 {
			octopuses = append(octopuses, s.octopuses[row+1][col+1])
		}
	}

	return octopuses
}

func newSimulatorFromStr(s string) (*OctopusSimulator, error) {
	octopuses := [][]*Octopus{}
	for i, line := range strings.Split(s, "\n") {
		octopuses = append(octopuses, []*Octopus{})
		for j, val := range line {
			octopuses[i] = append(octopuses[i], &Octopus{
				energyLevel: int(val - '0'),
				row:         i,
				col:         j,
			})

		}
	}
	return &OctopusSimulator{
		octopuses:         octopuses,
		currentlyFlashing: []*Octopus{},
	}, nil
}

type Octopus struct {
	energyLevel int
	flashing    bool
	row, col    int
}

func main() {
	s, err := newSimulatorFromStr(util.GetInputStr("https://adventofcode.com/2021/day/11/input"))
	if err != nil {
		panic(err)
	}
	s.Steps(100)
	fmt.Println(s.totalFlashes)

	s, err = newSimulatorFromStr(util.GetInputStr("https://adventofcode.com/2021/day/11/input"))
	if err != nil {
		panic(err)
	}

	s.Steps(10000)
	fmt.Println(s.step)
}
