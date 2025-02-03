package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindMatches(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{
			s:    "mul(12,33)}I{SDFISDF{SPDF{Pmul(12,34)",
			want: []string{"mul(12,33)", "mul(12,34)"},
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("%s-%d", t.Name(), i), func(t *testing.T) {
			got, err := findMatches(tc.s)
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(got, tc.want) {
				t.Errorf("got: %q, want: %q", got, tc.want)
			}
		})
	}
}

func TestFindMatches2(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{
			s:    "mul(12,33)don't()}I{SDFISDF{SPDF{Pmul(12,34)do()",
			want: []string{"mul(12,33)", "don't()", "mul(12,34)", "do()"},
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("%s-%d", t.Name(), i), func(t *testing.T) {
			got, err := findMatches2(tc.s)
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(got, tc.want) {
				t.Errorf("got: %q, want: %q", got, tc.want)
			}
		})
	}
}
