package main

import "testing"

func TestCountPaths1(t *testing.T) {
	input := `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`
	cs := CavesFromStr(input)
	want := 226
	got := CountPaths(cs, 1)

	if got != want {
		t.Errorf("unexpected value: got: %d, want: %d", got, want)
	}
}

func TestCountPaths2(t *testing.T) {
	input := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`
	cs := CavesFromStr(input)
	want := 36
	got := CountPaths(cs, 2)

	if got != want {
		t.Errorf("unexpected value: got: %d, want: %d", got, want)
	}
}
