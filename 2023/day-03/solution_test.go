package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		input string
		want  output
	}{
		{
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			want: output{sumParts: 4361, sumGearRations: 467835},
		},
		{
			input: `.1.
1*1
111`,
			want: output{sumParts: 114, sumGearRations: 0},
		},
	}
	for _, tc := range tests {
		mapping, err := mapCoordinateToPart(tc.input)
		if err != nil {
			t.Fatal(err)
		}

		got := sumIncludedParts(mapping, tc.input)

		if tc.want.sumParts != got.sumParts {
			t.Errorf("want: %d, got: %d", tc.want.sumParts, got.sumParts)
		}
		if tc.want.sumGearRations != got.sumGearRations {
			t.Errorf("want: %d, got: %d", tc.want.sumGearRations, got.sumGearRations)
		}
	}

}
