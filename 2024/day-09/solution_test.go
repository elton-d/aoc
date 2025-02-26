package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "2333133121414131402",
			want:  "00...111...2...333.44.5555.6666.777.888899",
		},
		{
			input: "12345",
			want:  "0..111....22222",
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			got := strings.Join(expand(tc.input), "")
			if got != tc.want {
				t.Errorf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}

func TestCompact2(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "00...111...2...333.44.5555.6666.777.888899",
			want:  "00992111777.44.333....5555.6666.....8888..",
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			got := strings.Join(compact2(strings.Split(tc.input, "")), "")
			if got != tc.want {
				t.Errorf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}
func TestCompact(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "00...111...2...333.44.5555.6666.777.888899",
			want:  "0099811188827773336446555566..............",
		},
		{

			input: "0..111....22222",
			want:  "022111222......",
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			got := strings.Join(compact(strings.Split(tc.input, "")), "")
			if got != tc.want {
				t.Errorf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}

func TestChecksum(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "2333133121414131402",
			want:  1928,
		},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			got := checksum(compact(expand(tc.input)))
			if got != tc.want {
				t.Errorf("got: %d, want: %d", got, tc.want)
			}
		})
	}

}
