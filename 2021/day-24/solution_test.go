package main

import "testing"

func TestNegateInstr(t *testing.T) {
	program := `inp x
mul x -1`
	a := NewALU([]int{5})
	a.RunProgram(program)

	want := -5

	if got := a.Variables["x"]; got != want {
		t.Errorf("unexpected value for var x, got: %v, want: %v", got, want)
	}
}
