package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Cube struct {
	X, Y, Z int
	// On      bool
}

// func (c *Cube) SwitchOn() {
// 	c.On = true
// }

// func (c *Cube) SwitchOff() {
// 	c.On = false
// }

func Less(a, b Cube) bool {
	return a.X < b.X && a.Y < b.Y && a.Z < b.Z
}

type Reactor struct {
	Cubes map[Cube]bool
	Steps []*RebootStep
}

func (r *Reactor) RunSteps() {
	for _, step := range r.Steps {
		cubes := r.getCubesInCuboid(step.cuboid)

		for _, cube := range cubes {
			if step.state == OnState {
				r.Cubes[cube] = true
			} else {
				delete(r.Cubes, cube)
			}
		}
	}
}

func (r *Reactor) SwitchedOnCubes() int {
	return len(r.Cubes)
}

func NewReactor(input string) *Reactor {
	steps := RebootStepsFromStr(input)
	return &Reactor{
		Steps: steps,
		Cubes: make(map[Cube]bool),
	}
}

type RebootStep struct {
	state  State
	cuboid *Cuboid
}

type Cuboid struct {
	xmin, xmax, ymin, ymax, zmin, zmax int
}

type State string

const (
	OnState  State = "on"
	OffState State = "off"
)

func CuboidFromStr(s string) *Cuboid {
	splits := strings.Split(s, ",")

	xBounds := strings.Split(strings.Split(splits[0], "=")[1], "..")
	yBounds := strings.Split(strings.Split(splits[1], "=")[1], "..")
	zBounds := strings.Split(strings.Split(splits[2], "=")[1], "..")

	xMin, _ := strconv.Atoi(xBounds[0])
	xMax, _ := strconv.Atoi(xBounds[1])

	yMin, _ := strconv.Atoi(yBounds[0])
	yMax, _ := strconv.Atoi(yBounds[1])

	zMin, _ := strconv.Atoi(zBounds[0])
	zMax, _ := strconv.Atoi(zBounds[1])

	return &Cuboid{
		xmin: xMin,
		xmax: xMax,
		ymin: yMin,
		ymax: yMax,
		zmin: zMin,
		zmax: zMax,
	}
}

func RebootStepFromStr(s string) *RebootStep {
	splits := strings.Split(s, " ")

	return &RebootStep{
		state:  State(splits[0]),
		cuboid: CuboidFromStr(splits[1]),
	}
}

func RebootStepsFromStr(s string) []*RebootStep {
	steps := []*RebootStep{}
	for _, line := range strings.Split(s, "\n") {
		steps = append(steps, RebootStepFromStr(line))
	}
	return steps
}

func Part1() int {
	input := `on x=-16..31,y=1..46,z=-4..43
on x=-25..20,y=-25..28,z=-48..-3
on x=-1..48,y=-8..36,z=-12..41
on x=-14..38,y=-46..0,z=-33..15
on x=-36..14,y=-35..14,z=-24..28
on x=-13..37,y=-36..8,z=-24..21
on x=-43..11,y=-48..6,z=-12..39
on x=-20..29,y=-45..5,z=-42..11
on x=-32..21,y=-9..45,z=-26..18
on x=-3..43,y=-38..14,z=-10..36
off x=-13..1,y=9..26,z=18..33
on x=-5..43,y=-2..49,z=-41..9
off x=-15..-1,y=-46..-32,z=-43..-29
on x=-3..46,y=-10..34,z=-35..19
off x=-44..-25,y=10..27,z=-46..-29
on x=-33..20,y=-35..14,z=-32..18
off x=-10..4,y=-33..-21,z=16..35
on x=-6..44,y=-11..39,z=-9..43
off x=-40..-30,y=-35..-25,z=-42..-27
on x=-45..7,y=-12..41,z=-19..35`
	r := NewReactor(input)

	r.RunSteps()
	return r.SwitchedOnCubes()
}

func (r *Reactor) getCubesInCuboid(c *Cuboid) []Cube {
	xmin, xmax, ymin, ymax, zmin, zmax := c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax
	cubes := []Cube{}
	for x := xmin; x <= xmax; x++ {
		for y := ymin; y <= ymax; y++ {
			for z := zmin; z <= zmax; z++ {
				cubes = append(cubes, Cube{
					X: x,
					Y: y,
					Z: z,
				})
			}
		}
	}
	return cubes
}

func main() {
	fmt.Println(Part1())
}
