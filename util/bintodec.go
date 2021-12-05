package util

import (
	"math"
)

func BinToDec(bin string) int {
	var dec int
	for i, b := range bin {

		dec += int(math.Pow(2.0, float64(len(bin)-i-1))) * int(b-'0')
	}
	return dec
}
