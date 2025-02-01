package main

import "testing"

func TestSafe(t *testing.T) {
	rep, err := getReports(`7 6 4 2 1`)
	if err != nil {
		t.Fatal(err)
	}
	if !safe(rep[0]) {
		t.Errorf("want safe")
	}

}

func TestSafeWithRemovals(t *testing.T) {
	rep, err := getReports(`9 7 6 2 1`)
	if err != nil {
		t.Fatal(err)
	}
	if safeWithRemovals(rep[0]) {
		t.Errorf("want unsafe")
	}

}
