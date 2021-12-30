package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/elton-d/aoc/util"
)

type Player struct {
	Score int
	Name  string
	Pos   int
}

func (p *Player) Move(n int) {
	newPos := (p.Pos + n) % 10
	if newPos == 0 {
		p.Pos = 10
	} else {
		p.Pos = newPos
	}

	p.Score += p.Pos
}

func (p *Player) Copy() *Player {
	return &Player{
		Score: p.Score,
		Pos:   p.Pos,
		Name:  p.Name,
	}
}

func PlayerFromStr(s string) *Player {
	playerName := strings.Split(s, " starting")[0]
	positionStr := strings.TrimSpace(strings.Split(s, ":")[1])

	pos, err := strconv.Atoi(positionStr)
	if err != nil {
		panic(err)
	}
	return &Player{
		Name: playerName,
		Pos:  pos,
	}
}

func PlayersFromStr(s string) []*Player {
	p := []*Player{}
	for _, line := range strings.Split(s, "\n") {
		p = append(p, PlayerFromStr(line))
	}
	return p
}

type Game struct {
	Players       []*Player
	Dice          Dice
	Turn          int
	TerminalScore int
	CurrRollSum   int
}

type GameState struct {
	playerPos    string
	playerScores string
	diceVal      int
	diceRolls    int
	currRollSum  int
	turn         int
}

func (g *Game) Play() int {
	for {
		for i, p := range g.Players {
			g.Turn = i
			rollsSum := 0
			rollsSum += g.Dice.Roll()
			rollsSum += g.Dice.Roll()
			rollsSum += g.Dice.Roll()
			p.Move(rollsSum)
			if p.Score >= g.TerminalScore {
				lowest := math.MaxInt
				for _, p := range g.Players {
					if p.Score < lowest {
						lowest = p.Score
					}
				}
				return lowest * g.Dice.RollCount()
			}
		}
	}
}

func (g *Game) Roll(rollValue int) bool {
	currPlayer := g.Players[g.Turn]
	g.Dice.RollWithValue(rollValue)
	g.CurrRollSum += rollValue

	if g.Dice.RollCount()%3 == 0 {
		currPlayer.Move(g.CurrRollSum)
		if g.Complete() {
			return true
		}
		g.CurrRollSum = 0
		g.Turn = (g.Turn + 1) % len(g.Players)
	}
	return false
}

func (g *Game) Winner() string {
	if g.Complete() {
		return g.Players[g.Turn].Name
	}
	return ""
}

func (g *Game) Complete() bool {
	return g.Players[g.Turn].Score >= g.TerminalScore
}

func (g *Game) Copy() *Game {
	players := []*Player{}
	for _, p := range g.Players {
		players = append(players, p.Copy())
	}
	return &Game{
		Players:       players,
		Dice:          g.Dice.Copy(),
		Turn:          g.Turn,
		TerminalScore: g.TerminalScore,
		CurrRollSum:   g.CurrRollSum,
	}
}

func (g *Game) State() GameState {
	positions := []string{}
	scores := []string{}

	for _, p := range g.Players {
		positions = append(positions, fmt.Sprintf("%d", p.Pos))
		scores = append(scores, fmt.Sprintf("%d", p.Score))
	}
	return GameState{
		playerPos:    strings.Join(positions, ","),
		playerScores: strings.Join(scores, ","),
		diceVal:      g.Dice.LastValue(),
		currRollSum:  g.CurrRollSum,
		diceRolls:    g.Dice.RollCount(),
		turn:         g.Turn,
	}
}

type Dice interface {
	Roll() int
	RollCount() int
	Copy() Dice
	RollWithValue(i int) int
	LastValue() int
}

type DeterministicDice struct {
	Value         int
	Rolls         int
	ResetInterval int
}

func (d *DeterministicDice) Copy() Dice {
	return &DeterministicDice{
		Value:         d.Value,
		Rolls:         d.Rolls,
		ResetInterval: d.ResetInterval,
	}
}

func (d *DeterministicDice) Roll() int {
	if d.Value == d.ResetInterval {
		d.Value = 1
	} else {
		d.Value += 1
	}
	d.Rolls += 1
	return d.Value
}

func (d *DeterministicDice) RollCount() int {
	return d.Rolls
}

func (d *DeterministicDice) RollWithValue(i int) int {
	d.Value = i
	d.Rolls += 1
	return d.Value
}

func (d *DeterministicDice) LastValue() int {
	return d.Value
}

func Part1(input string) int {
	game := &Game{
		Players:       PlayersFromStr(input),
		Dice:          &DeterministicDice{ResetInterval: 100},
		TerminalScore: 1000,
	}
	return game.Play()
}

type Result map[string]int

func (g *Game) PlayQuantum(cache *map[GameState]Result) Result {
	if g.Complete() {
		return map[string]int{
			g.Winner(): 1,
		}
	}
	children := []*Game{g.Copy(), g.Copy(), g.Copy()}

	for i, game := range children {
		game.Roll(i + 1)
	}
	state := g.State()
	if val, ok := (*cache)[state]; ok {
		return val
	}

	combined := combineMaps(
		children[0].PlayQuantum(cache),
		children[1].PlayQuantum(cache),
		children[2].PlayQuantum(cache),
	)

	(*cache)[state] = combined
	return combined
}

func combineMaps(maps ...map[string]int) map[string]int {
	combined := map[string]int{}

	for _, m := range maps {
		for k, v := range m {
			if _, ok := combined[k]; !ok {
				combined[k] = 0
			}
			combined[k] += v
		}
	}
	return combined
}

func Part2(input string) int {
	g := &Game{
		Players:       PlayersFromStr(input),
		Dice:          &DeterministicDice{ResetInterval: 3},
		TerminalScore: 21,
	}
	res := g.PlayQuantum(&map[GameState]Result{})
	max := math.MinInt

	for _, count := range res {
		if count > max {
			max = count
		}
	}
	return max
}

func main() {
	input := util.GetInputStr("https://adventofcode.com/2021/day/21/input")
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
