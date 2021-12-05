package main

import (
	"fmt"
	"strings"

	"github.com/elton-d/aoc/util"
)

func getInput() ([]string, error) {
	b, err := util.GetInput("https://adventofcode.com/2021/day/3/input")
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(b)), "\n"), nil
}

type FrequencyTable struct {
	counts [][]int
}

func newFrequencyTable(in []string) *FrequencyTable {
	counts := make([][]int, len(in[0]))
	for i := range counts {
		counts[i] = make([]int, 2)
	}
	f := &FrequencyTable{
		counts: counts,
	}

	for _, i := range in {
		f.Add(i)
	}
	return f
}

func (f *FrequencyTable) Add(num string) {
	for i, r := range num {
		f.counts[i][int(r-'0')] += 1
	}
}

func (f *FrequencyTable) Del(num string) {
	for i, r := range num {
		f.counts[i][int(r-'0')] -= 1
	}
}

func (f *FrequencyTable) MoreFrequent(pos int) int {
	if f.counts[pos][0] > f.counts[pos][1] {
		return 0
	}
	return 1
}

func (f *FrequencyTable) LessFrequent(pos int) int {
	if f.counts[pos][0] <= f.counts[pos][1] {
		return 0
	}
	return 1
}

func calcO2Rating(in []string) (int, error) {
	freq := newFrequencyTable(in)
	candidates := in[:]
	var newCandidates []string

	for i := 0; i < len(candidates[0]); i++ {
		moreFreq := freq.MoreFrequent(i)
		for _, c := range candidates {
			if int(c[i]-'0') == moreFreq {
				newCandidates = append(newCandidates, c)
			} else {
				freq.Del(c)
			}
		}
		candidates = newCandidates
		newCandidates = []string{}
		if len(candidates) == 1 {
			break
		}
	}

	return util.BinToDec(candidates[0])

}

func calcCO2Rating(in []string) (int, error) {
	freq := newFrequencyTable(in)
	candidates := in[:]
	var newCandidates []string

	for i := 0; i < len(candidates[0]); i++ {
		lessFreq := freq.LessFrequent(i)
		for _, c := range candidates {
			if int(c[i]-'0') == lessFreq {
				newCandidates = append(newCandidates, c)
			} else {
				freq.Del(c)
			}
		}
		candidates = newCandidates
		newCandidates = []string{}
		if len(candidates) == 1 {
			break
		}
	}
	return util.BinToDec(candidates[0])
}

func main() {
	in, err := getInput()
	if err != nil {
		panic(err)
	}
	o2, err := calcO2Rating(in)
	if err != nil {
		panic(err)
	}
	co2, err := calcCO2Rating(in)
	if err != nil {
		panic(err)
	}
	fmt.Println(o2 * co2)
}
