package main

import "testing"

func TestGraphString(t *testing.T) {
	in := `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0`
	g, err := NewGraphFromStr(in)
	if err != nil {
		t.Fatal(err)
	}
	want := `...#..#..#.
....#......
...........
#..........
...#....#.#
...........
...........
...........
...........
...........
.#....#.##.
....#......
......#...#
#..........
#.#........`
	got := g.String()

	if got != want {
		t.Errorf("unexpected value,\ngot:\n%v\n,\nwant:\n%v", got, want)
	}
}

func TestFoldY(t *testing.T) {
	in := `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0`
	g, err := NewGraphFromStr(in)
	if err != nil {
		t.Fatal(err)
	}
	want := `#.##..#..#.
#...#......
......#...#
#...#......
.#.#..#.###
...........
...........`
	g.FoldY(7)
	got := g.String()

	if got != want {
		t.Errorf("unexpected value,\ngot:\n%v\n,\nwant:\n%v", got, want)
	}
}

func TestFoldX(t *testing.T) {
	in := `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0`
	g, err := NewGraphFromStr(in)
	if err != nil {
		t.Fatal(err)
	}
	want := `#####
#...#
#...#
#...#
#####
.....
.....`
	g.FoldY(7)
	g.FoldX(5)
	got := g.String()

	if got != want {
		t.Errorf("unexpected value,\ngot:\n%v\n,\nwant:\n%v", got, want)
	}
}
