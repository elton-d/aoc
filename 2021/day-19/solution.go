package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

type Scanner struct {
	Name    string
	Beacons []Coordinate
}

type Pair struct {
	A, B Coordinate
}

func (s *Scanner) OverlapCount(other *Scanner) (int, Coordinate) {
	diffCount := map[Coordinate]int{}

	for i := 0; i < len(s.Beacons); i++ {
		for j := 0; j < len(other.Beacons); j++ {
			pair := Pair{A: s.Beacons[i], B: other.Beacons[j]}
			diff := Coordinate{
				X: pair.A.X - pair.B.X,
				Y: pair.A.Y - pair.B.Y,
				Z: pair.A.Z - pair.B.Z,
			}
			if _, ok := diffCount[diff]; !ok {
				diffCount[diff] = 0
			}
			diffCount[diff] += 1
		}
		for diff, count := range diffCount {
			if count >= 12 {
				return count, diff
			}
		}
	}
	return 0, Coordinate{}
}

func (s *Scanner) Copy() *Scanner {
	bc := make([]Coordinate, len(s.Beacons))
	copy(bc, s.Beacons)
	return &Scanner{
		Name:    s.Name,
		Beacons: bc,
	}
}

func (s *Scanner) Rotate(angle float64, axis Axis) *Scanner {
	ret := s.Copy()
	for i, b := range ret.Beacons {
		ret.Beacons[i] = b.Rotate(angle, axis)
	}
	return ret
}

func CountBeacons(scanners []*Scanner) int {
	q := []*Scanner{scanners[0]}

	unevaluated := map[string]*Scanner{}
	for _, s := range scanners[1:] {
		unevaluated[s.Name] = s
	}

	beacons := map[Coordinate]bool{}
	for _, b := range scanners[0].Beacons {
		beacons[b] = true
	}

	for len(unevaluated) > 0 {
		s1 := q[0]
		q = q[1:]
		delete(unevaluated, s1.Name)
		for _, s2 := range unevaluated {
			found := false
			for z := 0; z < 4 && !found; z++ {
				s2 = s2.Rotate(float64(z)*math.Pi/2, ZAxis)
				for y := 0; y < 4 && !found; y++ {
					s2 = s2.Rotate(float64(y)*math.Pi/2, YAxis)
					for x := 0; x < 4; x++ {
						s2 = s2.Rotate(float64(x)*math.Pi/2, XAxis)
						overlaps, diff := s1.OverlapCount(s2)
						if overlaps >= 12 {
							for i := range s2.Beacons {
								s2.Beacons[i].Translate(diff)
								beacons[s2.Beacons[i]] = true
							}
							q = append(q, s2)
							found = true
							break
						}
					}
				}
			}
		}
	}
	return len(beacons)
}

func LocateScanners(scanners []*Scanner) map[string]Coordinate {
	locations := map[string]Coordinate{
		scanners[0].Name: {0, 0, 0},
	}
	q := []*Scanner{scanners[0]}

	unevaluated := map[string]*Scanner{}
	for _, s := range scanners[1:] {
		unevaluated[s.Name] = s
	}

	for len(unevaluated) > 0 {
		s1 := q[0]
		q = q[1:]
		delete(unevaluated, s1.Name)
		for _, s2 := range unevaluated {
			found := false
			for z := 0; z < 4 && !found; z++ {
				s2 = s2.Rotate(float64(z)*math.Pi/2, ZAxis)
				for y := 0; y < 4 && !found; y++ {
					s2 = s2.Rotate(float64(y)*math.Pi/2, YAxis)
					for x := 0; x < 4; x++ {
						s2 = s2.Rotate(float64(x)*math.Pi/2, XAxis)
						overlaps, diff := s1.OverlapCount(s2)
						if overlaps >= 12 {
							locations[s2.Name] = diff
							for i := range s2.Beacons {
								s2.Beacons[i].Translate(diff)
							}
							q = append(q, s2)
							found = true
							break
						}
					}
				}
			}
		}
	}
	return locations
}

type Coordinate struct {
	X, Y, Z float64
}

type Axis string

const (
	XAxis Axis = "x"
	YAxis Axis = "y"
	ZAxis Axis = "z"
)

func (c *Coordinate) Rotate(angle float64, axis Axis) Coordinate {
	x, y, z := c.X, c.Y, c.Z
	if angle == 0 {
		return *c
	}

	sinTheta := math.Sin(angle)
	cosTheta := math.Cos(angle)
	if axis == ZAxis {
		x = c.X*cosTheta - c.Y*sinTheta
		y = c.Y*cosTheta + c.X*sinTheta
	} else if axis == XAxis {
		y = c.Y*cosTheta - c.Z*sinTheta
		z = c.Z*cosTheta + c.Y*sinTheta
	} else {
		z = c.Z*cosTheta - c.X*sinTheta
		x = c.X*cosTheta + c.Z*sinTheta
	}

	return Coordinate{X: math.Round(x), Y: math.Round(y), Z: math.Round(z)}
}

func (c *Coordinate) Equals(other Coordinate) bool {
	return c.X == other.X && c.Y == other.Y && c.Z == other.Z
}

func (c *Coordinate) Translate(diff Coordinate) {
	c.X += diff.X
	c.Y += diff.Y
	c.Z += diff.Z
}

func (c *Coordinate) ManhattanDistance(other Coordinate) float64 {
	dist := 0.0

	dist += math.Abs(c.X - other.X)
	dist += math.Abs(c.Y - other.Y)
	dist += math.Abs(c.Z - other.Z)

	return dist
}

func ScannerFromStr(s string) *Scanner {
	sc := &Scanner{}
	sc.Beacons = []Coordinate{}
	lines := strings.Split(s, "\n")
	sc.Name = strings.TrimSpace(strings.Replace(lines[0], "---", "", 2))
	for _, l := range lines[1:] {
		splits := strings.Split(l, ",")
		x, err := strconv.Atoi(splits[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(splits[1])
		if err != nil {
			panic(err)
		}
		z, err := strconv.Atoi(splits[2])
		if err != nil {
			panic(err)
		}
		sc.Beacons = append(sc.Beacons, Coordinate{X: float64(x), Y: float64(y), Z: float64(z)})
	}
	return sc
}

func ScannersFromStr(s string) []*Scanner {
	scanners := []*Scanner{}
	for _, scannerStr := range strings.Split(s, "\n\n") {
		scanners = append(scanners, ScannerFromStr(scannerStr))
	}
	return scanners
}

func Part1(input string) int {
	return CountBeacons(ScannersFromStr(input))
}

func Part2(input string) float64 {
	locationsMap := LocateScanners(ScannersFromStr(input))

	locations := []Coordinate{}

	for _, c := range locationsMap {
		locations = append(locations, c)
	}

	largestDist := -1.0
	for i := 0; i < len(locations)-1; i++ {
		for j := i; j < len(locations); j++ {
			dist := locations[i].ManhattanDistance(locations[j])
			if dist > largestDist {
				largestDist = dist
			}
		}
	}
	return largestDist
}

func main() {
	input := util.GetInputStr("https://adventofcode.com/2021/day/19/input")
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
