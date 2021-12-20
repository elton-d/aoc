package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

type Node struct {
	Left, Right, Parent *Node
	Value               int
	Pos                 int
}

func (n *Node) Root() *Node {
	curr := n
	for curr.Parent != nil {
		curr = curr.Parent
	}
	return curr
}

func (n *Node) String() string {
	if n.IsRegularNum() {
		return fmt.Sprintf("%d", n.Value)
	}
	var left, right string
	if n.Left != nil {
		left = n.Left.String()
	}
	if n.Right != nil {
		right = n.Right.String()
	}
	return fmt.Sprintf("[%s,%s]", left, right)
}

func (n *Node) getOrderedNodes() []*Node {
	ordered := []*Node{}
	if n.IsRegularNum() {
		return []*Node{n}
	}

	if n.Left != nil {
		ordered = append(ordered, n.Left.getOrderedNodes()...)
	}
	if n.Right != nil {
		ordered = append(ordered, n.Right.getOrderedNodes()...)
	}
	return ordered
}

func (n *Node) RefreshNodePositions() []*Node {
	ord := n.getOrderedNodes()
	for i, node := range ord {
		node.Pos = i
	}
	return ord
}

func (n *Node) IsRegularNum() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node) Add(other *Node) *Node {
	parent := &Node{
		Left:  n,
		Right: other,
	}
	n.Parent = parent
	other.Parent = parent
	parent.Reduce()
	return parent
}

func (n *Node) Reduce() {
	reduced := false
	var toExp, toSplit *Node
	for !reduced {
		if toExp = n.GetNodeToExplode(0); toExp != nil {
			toExp.Explode()
			continue
		}

		if toSplit = n.GetNodeToSplit(); toSplit != nil {
			toSplit.Split()
		}

		reduced = toExp == nil && toSplit == nil
	}
}

func (n *Node) GetNodeToExplode(depth int) *Node {
	if n.IsRegularNum() {
		return nil
	}
	if depth == 4 {
		return n
	}
	if n.Left != nil {
		if cand := n.Left.GetNodeToExplode(depth + 1); cand != nil {
			return cand
		}
	}
	if n.Right != nil {
		if cand := n.Right.GetNodeToExplode(depth + 1); cand != nil {
			return cand
		}
	}
	return nil
}

func (n *Node) GetNodeToSplit() *Node {
	if n.Left != nil {
		if cand := n.Left.GetNodeToSplit(); cand != nil {
			return cand
		}
	}
	if n.Right != nil {
		if cand := n.Right.GetNodeToSplit(); cand != nil {
			return cand
		}
	}
	if n.Value >= 10 {
		return n
	}
	return nil
}

func (n *Node) Explode() {
	// fmt.Printf("will explode: %s\n", n.String())
	if left := n.GetNextRegularNumToLeft(); left != nil {
		left.Value += n.Left.Value
	}

	if right := n.GetNextRegularNumToRight(); right != nil {
		right.Value += n.Right.Value
	}
	newNode := &Node{
		Value:  0,
		Parent: n.Parent,
	}

	n.Value = math.MinInt
	if n.Parent.Left != nil && n.Parent.Left.Value == n.Value {
		n.Parent.Left = newNode
	} else {
		n.Parent.Right = newNode
	}
	n.Parent = nil
	// fmt.Printf("Explode: %s\n", newNode.Root().String())

}

func (n *Node) GetNextRegularNumToLeft() *Node {
	root := n.Root()
	positions := root.RefreshNodePositions()
	if n.Left.Pos > 0 {
		return positions[n.Left.Pos-1]
	}
	return nil
}

func (n *Node) GetNextRegularNumToRight() *Node {
	root := n.Root()
	positions := root.RefreshNodePositions()
	if n.Right.Pos < len(positions)-1 {
		return positions[n.Right.Pos+1]
	}
	return nil
}

func (n *Node) Split() {
	// fmt.Printf("will split: %s\n", n.String())
	leftVal := int(math.Floor(float64(n.Value) / 2.0))
	rightVal := int(math.Ceil(float64(n.Value) / 2.0))

	newNode := &Node{
		Left: &Node{
			Value: leftVal,
		},
		Right: &Node{
			Value: rightVal,
		},
		Parent: n.Parent,
	}
	newNode.Left.Parent = newNode
	newNode.Right.Parent = newNode
	n.Value = math.MinInt
	if n.Parent != nil {
		if n.Parent.Left != nil && n.Parent.Left.Value == n.Value {
			n.Parent.Left = newNode
		} else {
			n.Parent.Right = newNode
		}
	}
	n.Parent = nil
	// fmt.Printf("Split: %s\n", newNode.Root().String())
}

func (n *Node) Magnitude() int {
	mag := 0
	if n.IsRegularNum() {
		return n.Value
	}
	if n.Left != nil {
		mag += 3 * n.Left.Magnitude()
	}
	if n.Right != nil {
		mag += 2 * n.Right.Magnitude()
	}
	return mag
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func NewTreeFromStr(s string) *Node {
	stack := []*Node{}

	var curr *Node
	i := 0
	for i < len(s) {
		c := string(s[i])
		if isDigit(c) {
			numStr := c
			i += 1
			for isDigit(string(s[i])) {
				numStr += c
				i += 1
			}
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			curr = &Node{
				Value: num,
			}
			i -= 1
		} else if c == "," {
			stack = append(stack, curr)
		} else if c == "]" {
			if len(stack) > 0 {
				parent := &Node{}
				popped := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				curr.Parent = parent
				popped.Parent = parent
				parent.Left = popped
				parent.Right = curr
				curr = parent
			}
		}
		i += 1
	}
	return curr
}

func Part1(numbers string) int {
	return addNums(numbers).Magnitude()
}

func Part2(numbers string) int {
	numStrs := strings.Split(numbers, "\n")
	max := math.MinInt

	for i, num1Str := range numStrs[:len(numStrs)-1] {
		for j := i + 1; j < len(numStrs); j++ {
			num1 := NewTreeFromStr(num1Str)
			num2 := NewTreeFromStr(numStrs[j])

			sum1 := num1.Add(num2)

			num1 = NewTreeFromStr(num1Str)
			num2 = NewTreeFromStr(numStrs[j])

			sum2 := num2.Add(num1)

			if mag := sum1.Magnitude(); mag > max {
				max = mag
			}

			if mag := sum2.Magnitude(); mag > max {
				max = mag
			}
		}
	}
	return max
}

func addNums(numbers string) *Node {
	var tree *Node
	for _, num := range strings.Split(numbers, "\n") {
		if tree == nil {
			tree = NewTreeFromStr(num)
		} else {
			other := NewTreeFromStr(num)
			tree = tree.Add(other)
		}
	}
	return tree
}

func main() {
	input := util.GetInputStr("https://adventofcode.com/2021/day/18/input")
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
