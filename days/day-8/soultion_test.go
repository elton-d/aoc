package main

import (
	"testing"
)

func TestFindConfig(t *testing.T) {
	inputStr := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |	cdfeb fcadb cdfeb cdbaf"
	i := newInput(inputStr)
	got := findConfig(i)
	want := map[rune]rune{
		'd': 'a',
		'e': 'b',
		'a': 'c',
		'f': 'd',
		'g': 'e',
		'b': 'f',
		'c': 'g',
	}

	for wk, wv := range want {
		gv, ok := got[wk]
		if !ok {
			t.Fatalf("didn't get corresponding mapping for %s", string(wk))
		}

		if wv != gv {
			t.Errorf("want %s -> %s got %s -> %s", string(wk), string(wv), string(wk), string(gv))
		}
	}

}

func TestGetValue(t *testing.T) {
	conf := map[rune]rune{
		'd': 'a',
		'e': 'b',
		'a': 'c',
		'f': 'd',
		'g': 'e',
		'b': 'f',
		'c': 'g',
	}

	want := map[string]int{
		"acedgfb": 8,
		"cdfbe":   5,
		"gcdfa":   2,
		"fbcad":   3,
		"dab":     7,
		"cefabd":  9,
		"cdfgeb":  6,
		"eafb":    4,
		"cagedb":  0,
		"ab":      1,
	}

	c := &Configuration{
		config: conf,
	}

	for k, v := range want {
		got := c.getValue(k)

		if got != v {
			t.Errorf("unexpected value for %v, want: %v, got:%v", k, v, got)
		}
	}

}

func TestDecodeOutput(t *testing.T) {
	conf := map[rune]rune{
		'd': 'a',
		'e': 'b',
		'a': 'c',
		'f': 'd',
		'g': 'e',
		'b': 'f',
		'c': 'g',
	}

	c := &Configuration{
		config: conf,
	}

	want := 5353

	s := []*Set{
		newSetFromStr("cdfeb"),
		newSetFromStr("fcadb"),
		newSetFromStr("cdfeb"),
		newSetFromStr("cdbaf"),
	}

	got := c.decodeOutput(s)

	if want != got {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
