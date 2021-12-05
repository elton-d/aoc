package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewLineFromStr(t *testing.T) {
	tests := []struct {
		str  string
		want *LineSegment
	}{
		{
			str: "427,523 -> 427,790",
			want: &LineSegment{
				Pt1: &Point{
					X: 427,
					Y: 523,
				},
				Pt2: &Point{
					X: 427,
					Y: 790,
				},
			},
		},
		{
			str: "94,639 -> 94,951",
			want: &LineSegment{
				Pt1: &Point{
					X: 94,
					Y: 639,
				},
				Pt2: &Point{
					X: 94,
					Y: 951,
				},
			},
		},
	}

	for _, tc := range tests {
		got, err := newLineFromStr(tc.str)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("diff(+got-want): %s", diff)
		}
	}
}
