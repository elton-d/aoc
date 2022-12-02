package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/elton-d/aoc/util"
)

var (
	canonicals = map[string]int{
		"abcefg":  0,
		"cf":      1,
		"acdeg":   2,
		"acdfg":   3,
		"bcdf":    4,
		"abdfg":   5,
		"abdefg":  6,
		"acf":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}
)

type Configuration struct {
	config map[rune]rune
}

func (c *Configuration) getValue(in string) int {
	s := []string{}

	for _, i := range in {
		s = append(s, string(c.config[i]))
	}
	sort.Strings(s)

	return canonicals[strings.Join(s, "")]
}

func (c *Configuration) decodeOutput(s []*Set) int {
	n := 0
	digits := []int{}
	for _, i := range s {
		digits = append(digits, c.getValue(i.String()))
	}

	for i, d := range digits {
		n += d * int(math.Pow10(len(digits)-i-1))
	}

	return n
}

type Input struct {
	patterns []*Set
	outputs  []*Set
}

func newInput(s string) *Input {
	splits := strings.Split(s, "|")
	p := strings.Fields(splits[0])
	o := strings.Fields(splits[1])

	patterns := []*Set{}
	outputs := []*Set{}

	for _, i := range p {
		patterns = append(patterns, newSetFromStr(i))
	}
	for _, i := range o {
		outputs = append(outputs, newSetFromStr(i))
	}
	return &Input{
		patterns: patterns,
		outputs:  outputs,
	}

}

type Set struct {
	set map[rune]bool
}

func newSet() *Set {
	return &Set{
		set: make(map[rune]bool),
	}
}

func newSetWith(elems ...rune) *Set {
	s := newSet()
	for _, r := range elems {
		s.Add(r)
	}
	return s
}

func (s *Set) Contains(x rune) bool {
	_, ok := s.set[x]
	return ok
}

func (s *Set) Add(x ...rune) {
	for _, i := range x {
		s.set[i] = true
	}
}

func (s *Set) Intersection(other *Set) *Set {
	i := newSet()
	for j := range other.set {
		if s.Contains(j) {
			i.Add(j)
		}
	}
	return i
}

func (s *Set) Size() int {
	return len(s.set)
}

func (s *Set) List() []rune {
	keys := []rune{}
	for k := range s.set {
		keys = append(keys, k)
	}
	return keys
}

func (s *Set) String() string {
	return string(s.List())
}

func (s *Set) Difference(other *Set) *Set {
	d := newSet()
	intersection := s.Intersection(other)
	for i := range s.set {
		if !intersection.Contains(i) {
			d.Add(i)
		}
	}
	return d
}

func (s *Set) Equal(other *Set) bool {
	if s.Size() != other.Size() {
		return false
	}
	for i := range s.set {
		if !other.Contains(i) {
			return false
		}
	}
	return true
}

func newSetFromStr(in string) *Set {
	s := newSet()
	for _, i := range in {
		s.Add(i)
	}
	return s
}

func part1() (int, error) {
	b, err := util.GetInput("https://adventofcode.com/2021/day/8/input")
	if err != nil {
		return -1, err
	}

	var outputs []string
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")

	for _, line := range lines {
		o := strings.Split(line, "|")[1]
		outputs = append(outputs, strings.Fields(o)...)
	}

	count := 0

	for _, o := range outputs {
		switch len(o) {
		case 2, 4, 3, 7:
			count += 1
		}
	}
	return count, nil

}

func part2Input() ([]*Input, error) {
	inputs := []*Input{}
	b, err := util.GetInput("https://adventofcode.com/2021/day/8/input")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	for _, l := range lines {
		inputs = append(inputs, newInput(l))
	}
	return inputs, nil
}

func findConfig(i *Input) map[rune]rune {

	var one, four, seven, eight *Set

	for _, j := range i.patterns {
		switch j.Size() {
		case 2:
			one = j
		case 4:
			four = j
		case 3:
			seven = j
		case 7:
			eight = j
		}
	}

	a := seven.Difference(one).List()[0]
	bd := four.Difference(one)
	eg := eight.Difference(four).Difference(newSetWith(a))

	var b, f rune

	for _, j := range i.patterns {
		if j.Size() == 6 && j.Difference(bd).Difference(eg).Difference(newSetWith(a)).Size() == 1 {
			f = j.Difference(bd).Difference(eg).Difference(newSetWith(a)).List()[0]
		}
	}
	c := one.Difference(newSetWith(f)).List()[0]

	for _, j := range i.patterns {
		if j.Size() == 6 && j.Difference(eg).Difference(newSetWith(a, c, f)).Size() == 1 {
			b = j.Difference(eg).Difference(newSetWith(a, c, f)).List()[0]
		}
	}

	d := bd.Difference(newSetWith(b)).List()[0]

	var e, g rune
	for _, j := range i.patterns {
		if j.Size() == 5 && j.Intersection(eg).Size() == 1 {
			k := j.Intersection(eg)
			g = k.List()[0]
			e = eg.Difference(k).List()[0]
		}
	}

	return map[rune]rune{
		a: 'a',
		b: 'b',
		c: 'c',
		d: 'd',
		e: 'e',
		f: 'f',
		g: 'g',
	}

}

func main() {
	p1, err := part1()
	if err != nil {
		panic(err)
	}

	fmt.Println(p1)

	inputs, err := part2Input()
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, in := range inputs {
		c := findConfig(in)
		sum += (&Configuration{config: c}).decodeOutput(in.outputs)
	}

	fmt.Println(sum)
}
