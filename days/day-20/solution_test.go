package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEnhance(t *testing.T) {
	enhancer := &ImageEnhancer{
		algo: "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#",
	}
	img := ImageFromStr(`#..#.
#....
##..#
..#..
..###`)
	want := `.##.##.
#..#.#.
##.#..#
####..#
.#..##.
..##..#
...#.#.`

	got := enhancer.Enhance(img, ".")

	if g := got.String(); g != want {
		t.Errorf("incorrect image, diff: %s", cmp.Diff(want, g))
	}

	got = enhancer.Enhance(got, ".")

	wantCount := 35

	if c := got.GetLitPixelsCount(); c != wantCount {
		t.Errorf("incorrect count: got: %d, want: %d", c, wantCount)
	}

	en2 := enhancer.EnhanceN(img, 2)

	if got.String() != en2.String() {
		t.Errorf("unexpected diff: %v", cmp.Diff(en2.String(), got.String()))
	}
}

func TestGetCode(t *testing.T) {
	img := ImageFromStr(`#..#.
#....
##..#
..#..
..###`)
	want := 34
	got := img.getPixelCode(2, 2, ".")

	if got != want {
		t.Errorf("incorrect code: got %d, want: %d", got, want)
	}

}

func TestGetLitPixelCount(t *testing.T) {
	img := ImageFromStr(`...............
...............
...............
...............
...............
.....#..#......
.....#.........
.....##..#.....
.......#.......
.......###.....
...............
...............
...............
...............
...............`)
	want := 10
	got := img.GetLitPixelsCount()

	if want != got {
		t.Errorf("unexpected value, got: %d, want: %d", got, want)
	}
}
