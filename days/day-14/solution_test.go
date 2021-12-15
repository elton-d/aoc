package main

import "testing"

func TestSteps(t *testing.T) {
	template := "NNCB"
	rulesStr := `CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`
	tests := []struct {
		steps int
		want  string
	}{
		{steps: 1, want: "NCNBCHB"},
		{steps: 2, want: "NBCCNBBBCBHCB"},
		{steps: 3, want: "NBBBCNCCNBBNBNBBCHBHHBCHB"},
		{steps: 4, want: "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB"},
	}

	for _, tc := range tests {
		p := &Polymer{
			state: template,
			rules: RulesFromStr(rulesStr),
			start: NodeFromStr(template),
		}
		p.RunSteps(tc.steps)
		if got := p.String(); got != tc.want {
			t.Errorf("unexpected value, got: %v, want: %v", got, tc.want)
		}
	}
}

func TestMaxMinusMin(t *testing.T) {
	tests := []struct {
		steps int
		want  int
	}{
		{
			steps: 10,
			want:  1588,
		},
	}

	template := "NNCB"
	rulesStr := `CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

	for _, tc := range tests {
		p := &Polymer{
			state: template,
			rules: RulesFromStr(rulesStr),
			start: NodeFromStr(template),
		}
		p.RunSteps(tc.steps)
		if got := p.MaxMinusMin(); got != tc.want {
			t.Errorf("unexpected value, got: %v, want: %v", got, tc.want)
		}
	}
}

func TestMaxMinusMin2(t *testing.T) {
	tests := []struct {
		steps int
		want  int
	}{
		{
			steps: 10,
			want:  1588,
		},
		{
			steps: 40,
			want:  2188189693529,
		},
	}

	template := "NNCB"
	rulesStr := `CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

	for _, tc := range tests {
		p := &Polymer{
			state: template,
			rules: RulesFromStr(rulesStr),
			start: NodeFromStr(template),
		}

		if got := p.MaxMinusMin2(tc.steps, "B"); got != tc.want {
			t.Errorf("unexpected value, got: %v, want: %v", got, tc.want)
		}
	}
}
