package main

import "testing"

func TestSolution1(t *testing.T) {
	pos := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	want := 37
	got := getFuelUsedForBestPos(pos)
	if got != want {
		t.Errorf("unexpected value for fuel used; want: %v, got: %v", want, got)
	}
}

func TestSolution2(t *testing.T) {
	pos := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	want := 168
	got := getFuelUsedForBestPos2(pos)
	if got != want {
		t.Errorf("unexpected value for fuel used; want: %v, got: %v", want, got)
	}
}
