package main

import "testing"

func TestPoints(t *testing.T) {
	tests := []struct {
		cardStr string
		want    float64
	}{
		{
			cardStr: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			want:    8,
		},
	}

	for _, tc := range tests {
		c, err := cardFromStr(tc.cardStr)
		if err != nil {
			t.Fatal(err)
		}
		if got := c.Points(); got != tc.want {
			t.Errorf("%s, got: %v, want: %v", tc.cardStr, got, tc.want)
		}
	}
}
