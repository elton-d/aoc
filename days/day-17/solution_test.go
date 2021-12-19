package main

import (
	"testing"
)

func TestHitsTarget(t *testing.T) {
	tests := []struct {
		vx, vy    int
		targetStr string
		want      bool
	}{
		{
			vx:        7,
			vy:        2,
			targetStr: "target area: x=20..30, y=-10..-5",
			want:      true,
		},
		{
			vx:        6,
			vy:        3,
			targetStr: "target area: x=20..30, y=-10..-5",
			want:      true,
		},
		{
			vx:        9,
			vy:        0,
			targetStr: "target area: x=20..30, y=-10..-5",
			want:      true,
		},
		{
			vx:        17,
			vy:        -4,
			targetStr: "target area: x=20..30, y=-10..-5",
			want:      false,
		},
	}

	for _, tc := range tests {
		p := NewProbe(tc.vx, tc.vy)
		target, err := NewTargetAreaFromStr(tc.targetStr)
		if err != nil {
			t.Fatal(err)
		}
		got := p.HitsTarget(target)

		if got != tc.want {
			t.Errorf("unexpected value, want: %v, got: %v", tc.want, got)
		}
	}
}

func TestFindVelocityForMax(t *testing.T) {
	target, err := NewTargetAreaFromStr("target area: x=20..30, y=-10..-5")
	if err != nil {
		t.Fatal(err)
	}
	wantH := 45
	wantV := 112

	gotH, gotV := target.FindVelocityForMax()
	if gotH != wantH {
		t.Errorf("unexpected value for max height: got: %v, want: %v", gotH, wantH)
	}
	if gotV != wantV {
		t.Errorf("unexpected value for initial velocities: got: %v, want: %v", gotV, wantV)
	}

}
