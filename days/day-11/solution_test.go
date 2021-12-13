package main

import (
	"strings"
	"testing"
)

func TestSimulation(t *testing.T) {
	orig := `11111
19991
19191
19991
11111`
	tests := []struct {
		step int
		want string
	}{
		{
			step: 1,
			want: `34543
40004
50005
40004
34543`,
		},

		{
			step: 2,
			want: `45654
51115
61116
51115
45654`,
		},
	}

	for _, tc := range tests {
		s, err := newSimulatorFromStr(orig)
		if err != nil {
			t.Fatal(err)
		}
		s.Steps(tc.step)
		if got := s.String(); strings.TrimSpace(got) != tc.want {
			t.Errorf("unexpected result at step %d, \ngot:\n%s\nwant:\n%s", tc.step, got, tc.want)
		}
	}
}

func TestFlashCount(t *testing.T) {
	s, err := newSimulatorFromStr(`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`)

	if err != nil {
		t.Fatal(err)
	}

	s.Steps(100)
	want := 1656

	if got := s.totalFlashes; got != want {
		t.Errorf("unexpected number of flashes: got: %d, want: %d", got, want)
	}
}
