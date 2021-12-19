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

type Velocity struct {
	Vx, Vy int
}

type Probe struct {
	Trajectory      []Coordinate
	Vx, Vy, Step    int
	CurrPos         Coordinate
	HighestPoint    int
	InitialVelocity Velocity
}

func NewProbe(vx, vy int) *Probe {
	return &Probe{
		Trajectory:      []Coordinate{{x: 0, y: 0}},
		Vx:              vx,
		Vy:              vy,
		HighestPoint:    math.MinInt,
		InitialVelocity: Velocity{Vx: vx, Vy: vy},
	}
}

func (p *Probe) HitsTarget(t *TargetArea) bool {
	for {
		p.NextStep()
		if (t.xMax > 0 && p.CurrPos.x > t.xMax) || (t.xMin < 0 && p.CurrPos.x < t.xMin) || (p.CurrPos.y < t.yMin) {
			return false
		}
		if t.Contains(p.CurrPos) {
			return true
		}
	}
}

func (p *Probe) NextStep() {
	p.CurrPos.x += p.Vx
	p.CurrPos.y += p.Vy

	if p.Vx > 0 {
		p.Vx -= 1
	} else if p.Vx < 0 {
		p.Vx += 1
	}

	p.Vy -= 1
	p.Trajectory = append(p.Trajectory, p.CurrPos)
	if p.CurrPos.y > p.HighestPoint {
		p.HighestPoint = p.CurrPos.y
	}
}

type TargetArea struct {
	xMin, xMax, yMin, yMax int
}

func (t *TargetArea) Contains(c Coordinate) bool {
	return c.x >= t.xMin && c.x <= t.xMax && c.y >= t.yMin && c.y <= t.yMax
}

func (t *TargetArea) FindVelocityForMax() (int, int) {
	highest := math.MinInt
	hits := 0
	if t.xMax > 0 {
		for vx := 1; vx <= t.xMax; vx++ {
			for vy := -t.yMin - 1; vy >= t.yMin; vy-- {
				p := NewProbe(vx, vy)
				if p.HitsTarget(t) {
					hits += 1
					if p.HighestPoint > highest {
						highest = p.HighestPoint

					}
				}
			}
		}
	}
	return highest, hits
}

func NewTargetAreaFromStr(s string) (*TargetArea, error) {
	fields := strings.Split(strings.Split(s, ":")[1], ",")
	xSplit := strings.Split(strings.Split(fields[0], "=")[1], "..")
	ySplit := strings.Split(strings.Split(fields[1], "=")[1], "..")

	xMin, err := strconv.Atoi(xSplit[0])
	if err != nil {
		return nil, err
	}

	xMax, err := strconv.Atoi(xSplit[1])
	if err != nil {
		return nil, err
	}

	yMin, err := strconv.Atoi(ySplit[0])
	if err != nil {
		return nil, err
	}

	yMax, err := strconv.Atoi(ySplit[1])
	if err != nil {
		return nil, err
	}

	return &TargetArea{
		xMin: xMin,
		xMax: xMax,
		yMin: yMin,
		yMax: yMax,
	}, nil
}

func Part1(input string) {
	target, err := NewTargetAreaFromStr(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(target.FindVelocityForMax())
}

func main() {
	input := util.GetInputStr("https://adventofcode.com/2021/day/17/input")
	Part1(input)
}
