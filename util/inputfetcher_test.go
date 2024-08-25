//go:build local
// +build local

package util

import (
	"testing"
)

func TestInputFetcher(t *testing.T) {
	input := GetInputStr("https://adventofcode.com/2023/day/1/input")
	if input == "" {
		t.Fatal("input cannot be empty")
	}
	t.Log(input)
}
