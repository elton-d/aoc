package main

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/elton-d/aoc/util"
)

var (
	ErrEndOfStream = errors.New("EOS")
)

func hexToBin(s string) string {
	mapping := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}
	sb := strings.Builder{}

	for _, c := range s {
		sb.WriteString(mapping[c])
	}
	return sb.String()
}

type Packet struct {
	Hdr        *Header
	SubPackets []*Packet
	Value      int
}

func (p *Packet) evaluateExpression() {
	switch p.Hdr.TypeID {
	case 0:
		for _, sp := range p.SubPackets {
			p.Value += sp.Value
		}
	case 1:
		p.Value = 1
		for _, sp := range p.SubPackets {
			p.Value *= sp.Value
		}
	case 2:
		p.Value = math.MaxInt
		for _, sp := range p.SubPackets {
			if sp.Value < p.Value {
				p.Value = sp.Value
			}
		}
	case 3:
		p.Value = math.MinInt
		for _, sp := range p.SubPackets {
			if sp.Value > p.Value {
				p.Value = sp.Value
			}
		}
	case 5:
		if p.SubPackets[0].Value > p.SubPackets[1].Value {
			p.Value = 1
		} else {
			p.Value = 0
		}
	case 6:
		if p.SubPackets[0].Value < p.SubPackets[1].Value {
			p.Value = 1
		} else {
			p.Value = 0
		}
	case 7:
		if p.SubPackets[0].Value == p.SubPackets[1].Value {
			p.Value = 1
		} else {
			p.Value = 0
		}
	}

}

type Header struct {
	Version int
	TypeID  int
}

type PacketReader struct {
	pos int
	str string
}

func (pr *PacketReader) NextN(n int) string {
	ret := pr.str[pr.pos : pr.pos+n]
	pr.pos += n
	return ret
}

func (pr *PacketReader) parseLiteralValue() int {
	sb := strings.Builder{}
	bits := pr.NextN(5)
	for bits[0] != '0' {
		sb.WriteString(bits[1:])
		bits = pr.NextN(5)
	}
	sb.WriteString(bits[1:])

	return util.BinToDec(sb.String())
}

func (pr *PacketReader) NextPacket() (*Packet, error) {
	if len(pr.str)-pr.pos < 6 {
		return nil, ErrEndOfStream
	}
	p := &Packet{}
	p.Hdr = HeaderFromString(pr.NextN(6))
	switch p.Hdr.TypeID {
	case 4:
		p.Value = pr.parseLiteralValue()
	default:
		lengthTypeID := pr.NextN(1)
		if p.SubPackets == nil {
			p.SubPackets = []*Packet{}
		}
		switch lengthTypeID {
		case "0":
			length := util.BinToDec(pr.NextN(15))
			start := pr.pos
			for pr.pos < start+length {
				sp, err := pr.NextPacket()
				if err != nil {
					return nil, err
				}
				p.SubPackets = append(p.SubPackets, sp)
			}
		case "1":
			numSubPackets := util.BinToDec(pr.NextN(11))
			for len(p.SubPackets) < numSubPackets {
				sp, err := pr.NextPacket()
				if err != nil {
					return nil, err
				}
				p.SubPackets = append(p.SubPackets, sp)
			}
		}
		p.evaluateExpression()
	}
	return p, nil
}

func PacketReaderFromStr(s string) *PacketReader {
	return &PacketReader{
		str: hexToBin(s),
	}
}

func HeaderFromString(s string) *Header {
	return &Header{
		Version: util.BinToDec(s[:3]),
		TypeID:  util.BinToDec(s[3:]),
	}
}

func recursiveVersionSum(p *Packet) int {
	sum := p.Hdr.Version
	for _, sp := range p.SubPackets {
		sum += recursiveVersionSum(sp)
	}
	return sum
}

func Part2(input string) {
	packets := []*Packet{}

	pr := PacketReaderFromStr(input)
	var (
		p   *Packet
		err error
	)
	for p, err = pr.NextPacket(); err == nil; p, err = pr.NextPacket() {
		packets = append(packets, p)

	}
	if err != nil {
		if !errors.Is(err, ErrEndOfStream) {
			panic(err)
		}
	}

	fmt.Println(packets[0].Value)
}

func Part1(input string) {
	packets := []*Packet{}

	pr := PacketReaderFromStr(input)
	var (
		p   *Packet
		err error
	)
	for p, err = pr.NextPacket(); err == nil; p, err = pr.NextPacket() {
		packets = append(packets, p)

	}
	if err != nil {
		if !errors.Is(err, ErrEndOfStream) {
			panic(err)
		}
	}

	sum := 0
	for _, p := range packets {
		sum += recursiveVersionSum(p)
	}
	fmt.Println(sum)
}

func main() {
	input := util.GetInputStr("https://adventofcode.com/2021/day/16/input")
	Part1(input)
	Part2(input)
}
