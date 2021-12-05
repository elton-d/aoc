package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
)

type BingoCard struct {
	lookupTable                  map[int]*BingoCoordinate
	rowMarkCounts, colMarkCounts []int
	strRep                       string
	complete                     bool
}

type BingoCoordinate struct {
	row, col int
	marked   bool
}

func (b *BingoCard) Check(i int) bool {
	if b.complete {
		return false
	}
	if coords, ok := b.lookupTable[i]; ok {
		coords.marked = true
		size := b.size()

		b.rowMarkCounts[coords.row] += 1
		b.colMarkCounts[coords.col] += 1
		b.complete = b.rowMarkCounts[coords.row] == size || b.colMarkCounts[coords.col] == size
		return b.complete
	}
	return false
}

func (b *BingoCard) size() int {
	return len(b.rowMarkCounts)
}

func (b *BingoCard) Score() int {
	score := 0
	for val, coord := range b.lookupTable {
		if !coord.marked {
			score += val
		}
	}
	return score
}

func (b *BingoCard) Print() {
	for _, row := range strings.Split(b.strRep, "\n") {
		for _, val := range strings.Fields(row) {
			num, _ := strconv.Atoi(val)
			valToPrint := fmt.Sprintf("%2d ", num)
			if coords, ok := b.lookupTable[num]; ok {
				if coords.marked {
					fmt.Print(string(colorRed), valToPrint, string(colorReset))
				} else {
					fmt.Print(valToPrint)
				}
			}

		}
		fmt.Print("\n")
	}
}

type BingoGame struct {
	cards []*BingoCard
	draws []int
}

func (g *BingoGame) PlayToWin() int {
	for _, num := range g.draws {
		for _, c := range g.cards {
			if c.Check(num) {
				c.Print()
				fmt.Printf("Wins on %d, board score: %d \n", num, c.Score())
				return num * c.Score()
			}
		}
	}
	return 0
}

func (g *BingoGame) PlayToLose() int {
	var boardsCompleted int

	for _, num := range g.draws {
		for _, c := range g.cards {
			if c.Check(num) {
				boardsCompleted += 1
				if len(g.cards) == boardsCompleted {
					c.Print()
					fmt.Printf("Wins on %d, board score: %d \n", num, c.Score())
					return num * c.Score()
				}

			}
		}
	}
	return 0
}

func newBingoCard(strRep string) (*BingoCard, error) {
	b := &BingoCard{
		strRep: strRep,
	}
	b.lookupTable = make(map[int]*BingoCoordinate)
	var i int
	var row string
	for i, row = range strings.Split(strRep, "\n") {
		for j, val := range strings.Fields(row) {
			num, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			b.lookupTable[num] = &BingoCoordinate{row: i, col: j}
		}
	}
	b.rowMarkCounts = make([]int, i+1)
	b.colMarkCounts = make([]int, i+1)
	return b, nil
}

func newBingoGame() (*BingoGame, error) {
	b, err := util.GetInput("https://adventofcode.com/2021/day/4/input")
	if err != nil {
		return nil, err
	}

	splits := strings.SplitN(strings.TrimSpace(string(b)), "\n\n", 2)
	var draws []int
	for _, d := range strings.Split(splits[0], ",") {
		num, err := strconv.Atoi(d)
		if err != nil {
			return nil, err
		}
		draws = append(draws, num)
	}

	var cards []*BingoCard
	for _, boardStr := range strings.Split(splits[1], "\n\n") {
		card, err := newBingoCard(boardStr)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return &BingoGame{
		cards: cards,
		draws: draws,
	}, nil
}

func main() {
	game, err := newBingoGame()
	if err != nil {
		panic(err)
	}
	fmt.Println(game.PlayToWin())
	fmt.Println("======================================")
	game, err = newBingoGame()
	if err != nil {
		panic(err)
	}
	fmt.Println(game.PlayToLose())
}
