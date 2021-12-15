package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/elton-d/aoc/util"
)

type Node struct {
	val  string
	next *Node
}

func (n *Node) InsertAfter(val string) {
	var tmp *Node
	if n.next != nil {
		tmp = n.next
	}

	n.next = &Node{val: val, next: tmp}
}

func NodeFromStr(s string) *Node {
	start := &Node{val: string(s[0])}
	curr := start
	for _, c := range s[1:] {
		curr.next = &Node{val: string(c)}
		curr = curr.next
	}
	return start
}

type Polymer struct {
	start *Node
	state string
	rules Rules
}

func (p *Polymer) String() string {
	curr := p.start
	sb := strings.Builder{}
	for curr != nil {
		sb.WriteString(curr.val)
		curr = curr.next
	}
	return sb.String()
}
func (p *Polymer) RunSteps(n int) {
	for k := 0; k < n; k++ {
		curr := p.start
		for curr.next != nil {
			ruleKey := strings.Join([]string{curr.val, curr.next.val}, "")

			if x, ok := p.rules[ruleKey]; ok {
				curr.InsertAfter(x)
			}
			curr = curr.next.next
		}
	}
}

func (p *Polymer) RunSteps2(n int) map[string]int {
	pairCounts := make(map[string]int)
	elemCounts := make(map[string]int)
	for i := 0; i < len(p.state)-1; i++ {
		char := string(p.state[i])
		if _, ok := elemCounts[char]; !ok {
			elemCounts[char] = 0
		}
		elemCounts[char] += 1
	}

	for i := 0; i < len(p.state)-1; i++ {
		j := i + 1
		pair := p.state[i : j+1]
		if _, ok := pairCounts[pair]; !ok {
			pairCounts[pair] = 0
		}
		pairCounts[pair] += 1
	}
	for s := 0; s < n; s++ {
		newPairCounts := make(map[string]int)
		for pair, count := range pairCounts {
			res := p.rules[pair]
			k1 := strings.Join([]string{string(pair[0]), res}, "")
			k2 := strings.Join([]string{res, string(pair[1])}, "")
			if _, ok := newPairCounts[k1]; !ok {
				newPairCounts[k1] = 0
			}
			if _, ok := newPairCounts[k2]; !ok {
				newPairCounts[k2] = 0
			}
			newPairCounts[k1] += count
			newPairCounts[k2] += count

		}

		pairCounts = newPairCounts
	}
	return pairCounts
}

func (p *Polymer) MaxMinusMin2(steps int, last string) int {
	elemCounts := make(map[string]int)
	for pair, count := range p.RunSteps2(steps) {
		elem := string(pair[0])
		if _, ok := elemCounts[elem]; !ok {
			elemCounts[elem] = 0
		}
		elemCounts[elem] += count
	}

	elemCounts[last] += 1

	max := math.MinInt
	min := math.MaxInt

	for _, count := range elemCounts {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}
	return max - min
}

func (p *Polymer) ElementCounts() map[string]int {
	counts := make(map[string]int)
	curr := p.start
	for curr != nil {
		if _, ok := counts[curr.val]; !ok {
			counts[curr.val] = 0
		}
		counts[curr.val] += 1
		curr = curr.next
	}
	return counts
}

func (p *Polymer) MaxMinusMin() int {
	max := 0
	min := math.MaxInt
	for _, i := range p.ElementCounts() {
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}
	return max - min
}

func RulesFromStr(rulesStr string) Rules {
	rules := make(map[string]string)
	for _, r := range strings.Split(rulesStr, "\n") {
		splits := strings.Split(r, " -> ")
		k, v := splits[0], splits[1]
		rules[k] = v
	}
	return rules
}

type Rules map[string]string

func Part1(template, rulesStr string, steps int) {
	p := &Polymer{
		rules: RulesFromStr(rulesStr),
		state: template,
		start: NodeFromStr(template),
	}
	p.RunSteps(steps)
	fmt.Println(p.MaxMinusMin())
}

func Part2(template, rulesStr string, steps int) {
	p := &Polymer{
		rules: RulesFromStr(rulesStr),
		state: template,
		start: NodeFromStr(template),
	}
	fmt.Println(p.MaxMinusMin2(steps, string(template[len(template)-1])))
}

func main() {
	input := util.GetInputStr("https://adventofcode.com/2021/day/14/input")
	splits := strings.Split(input, "\n\n")
	template, rulesStr := splits[0], splits[1]
	Part1(template, rulesStr, 10)
	Part2(template, rulesStr, 40)
}
