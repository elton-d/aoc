package main

import "testing"

func TestSolution(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}
	tests := []struct {
		days int
		want int
	}{
		{
			days: 18,
			want: 26,
		},
		{
			days: 80,
			want: 5934,
		},
		{
			days: 256,
			want: 26984457539,
		},
	}
	for _, tc := range tests {
		got := fishCount(input, tc.days)
		if tc.want != got {
			t.Errorf("unexpected number of fish after %d days: want: %d, got: %d", tc.days, tc.want, got)
		}
	}
}
