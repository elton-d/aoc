package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/elton-d/aoc/util"
)

func main() {
	input := util.GetInputStr("https://adventofcode.com/2023/day/3/input")

	mapping, err := mapCoordinateToPart(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(sumIncludedParts(mapping, input))
}

type partInstance struct {
	id      int
	partNum int
}

type output struct {
	sumParts       int
	sumGearRations int
}

func mapCoordinateToPart(input string) ([][]partInstance, error) {
	lines := strings.Split(input, "\n")
	mapping := make([][]partInstance, len(lines))
	id := 1

	for j, line := range lines {
		mapping[j] = make([]partInstance, len(line))
		var l, r int
		var char rune
		var digitChars []rune
		for i := 0; i < len(line); {
			char = rune(line[i])
			if unicode.IsDigit(char) {
				digitChars = []rune{}
				l = i
				for r = l; r < len(line); r++ {
					char = rune(line[r])
					if !unicode.IsDigit(char) {
						num, err := strconv.Atoi(string(digitChars))
						if err != nil {
							return nil, err
						}
						for k := l; k < r; k++ {
							mapping[j][k] = partInstance{id: id, partNum: num}
						}
						id += 1
						break
					}
					digitChars = append(digitChars, char)
					i += 1
				}
			} else {
				i += 1
			}
		}
		if len(digitChars) > 0 {
			num, err := strconv.Atoi(string(digitChars))
			if err != nil {
				return nil, err
			}
			for k := l; k < r; k++ {
				mapping[j][k] = partInstance{id: id, partNum: num}
			}
			id += 1
		}
	}
	return mapping, nil
}

func sumIncludedParts(mapping [][]partInstance, input string) output {
	sum := 0
	partsSet := make(map[partInstance]bool)
	sumGearRatios := 0
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		for j, char := range line {
			if !unicode.IsDigit(char) && char != '.' {
				adjacentParts := map[partInstance]bool{}
				if j > 0 {
					left := mapping[i][j-1]
					if left.partNum != 0 {
						adjacentParts[left] = true
						partsSet[left] = true
					}
				}
				if j < len(line)-1 {
					right := mapping[i][j+1]
					if right.partNum != 0 {
						adjacentParts[right] = true
						partsSet[right] = true
					}
				}
				if i > 0 {
					top := mapping[i-1][j]
					if top.partNum != 0 {
						adjacentParts[top] = true
						partsSet[top] = true
					}
				}
				if i < len(lines)-1 {
					bottom := mapping[i+1][j]
					if bottom.partNum != 0 {
						adjacentParts[bottom] = true
						partsSet[bottom] = true
					}
				}
				if i > 0 && j > 0 {
					topLeft := mapping[i-1][j-1]
					if topLeft.partNum != 0 {
						adjacentParts[topLeft] = true
						partsSet[topLeft] = true
					}
				}
				if i > 0 && j < len(line)-1 {
					topRight := mapping[i-1][j+1]
					if topRight.partNum != 0 {
						adjacentParts[topRight] = true
						partsSet[topRight] = true
					}
				}
				if i < len(lines)-1 && j > 0 {
					bottomLeft := mapping[i+1][j-1]
					if bottomLeft.partNum != 0 {
						adjacentParts[bottomLeft] = true
						partsSet[bottomLeft] = true
					}
				}
				if i < len(lines)-1 && j < len(line)-1 {
					bottomRight := mapping[i+1][j+1]
					if bottomRight.partNum != 0 {
						adjacentParts[bottomRight] = true
						partsSet[bottomRight] = true
					}
				}
				if char == '*' && len(adjacentParts) == 2 {
					ratio := 1
					for p := range adjacentParts {
						ratio *= p.partNum
					}
					sumGearRatios += ratio
				}
			}
		}
	}

	for pi := range partsSet {
		sum += pi.partNum
	}

	return output{sumParts: sum, sumGearRations: sumGearRatios}
}
