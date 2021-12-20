package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewTreeFromStr(t *testing.T) {
	tests := []struct {
		str  string
		want *Node
	}{
		{
			str: "[1,2]",
			want: &Node{
				Left: &Node{
					Value: 1,
				},
				Right: &Node{
					Value: 2,
				},
			},
		},
		{
			str: "[[1,2],3]",
			want: &Node{
				Left: &Node{
					Left: &Node{
						Value: 1,
					},
					Right: &Node{
						Value: 2,
					},
				},
				Right: &Node{
					Value: 3,
				},
			},
		},
		{
			str: "[[1,1],[2,2]]",
			want: &Node{
				Left: &Node{
					Left: &Node{
						Value: 1,
					},
					Right: &Node{
						Value: 1,
					},
				},
				Right: &Node{
					Left: &Node{
						Value: 2,
					},
					Right: &Node{
						Value: 2,
					},
				},
			},
		},
	}
	ignoreOpt := cmpopts.IgnoreFields(Node{}, "Parent")
	for _, tc := range tests {
		got := NewTreeFromStr(tc.str)

		if diff := cmp.Diff(tc.want, got, ignoreOpt); diff != "" {
			t.Error(diff)
		}
	}
}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		treeStr string
		mag     int
	}{
		{treeStr: "[9,1]", mag: 29},
		{treeStr: "[[9,1],[1,9]]", mag: 129},
		{treeStr: "[[1,2],[[3,4],5]]", mag: 143},
		{treeStr: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", mag: 1384},
		{treeStr: "[[[[1,1],[2,2]],[3,3]],[4,4]]", mag: 445},
		{treeStr: "[[[[3,0],[5,3]],[4,4]],[5,5]]", mag: 791},
		{treeStr: "[[[[5,0],[7,4]],[5,5]],[6,6]]", mag: 1137},
		{treeStr: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", mag: 3488},
	}

	for _, tc := range tests {
		tree := NewTreeFromStr(tc.treeStr)
		if got := tree.Magnitude(); got != tc.mag {
			t.Errorf("unexpected magnitude, got: %d, want: %d", got, tc.mag)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		nums string
		want string
	}{
		{
			nums: `[1,1]
[2,2]
[3,3]
[4,4]`,
			want: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			nums: `[1,1]
[2,2]
[3,3]
[4,4]
[5,5]`,
			want: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			nums: `[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
[6,6]`,
			want: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			nums: `[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]`,
			want: "[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]",
		},
	}

	for _, tc := range tests {
		res := addNums(tc.nums)
		m := res.Magnitude()
		fmt.Println(m)
		if diff := cmp.Diff(tc.want, res.String()); diff != "" {
			t.Errorf(diff)
		}
	}
}

func TestPart2(t *testing.T) {
	input := `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`

	want := 3993
	got := Part2(input)

	if got != want {
		t.Errorf("unexpected value for magnitude, got: %d, want: %d", got, want)
	}
}
