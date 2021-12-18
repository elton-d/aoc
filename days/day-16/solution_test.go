package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHexToBin(t *testing.T) {
	tests := []struct {
		hex string
		bin string
	}{
		{
			hex: "D2FE28",
			bin: "110100101111111000101000",
		},
		{
			hex: "38006F45291200",
			bin: "00111000000000000110111101000101001010010001001000000000",
		},
		{
			hex: "EE00D40C823060",
			bin: "11101110000000001101010000001100100000100011000001100000",
		},
	}

	for _, tc := range tests {
		got := hexToBin(tc.hex)
		if got != tc.bin {
			t.Errorf("unexpected value, got: %v, want:%v", got, tc.bin)
		}
	}
}

func TestHeaderFromString(t *testing.T) {
	tests := []struct {
		str  string
		want *Header
	}{
		{
			str: "110100",
			want: &Header{
				Version: 6,
				TypeID:  4,
			},
		},
	}

	for _, tc := range tests {
		got := HeaderFromString(tc.str)
		if diff := cmp.Diff(got, tc.want); diff != "" {
			t.Errorf("unexpected value, diff:%v", diff)
		}
	}
}

func TestNextPacket(t *testing.T) {
	tests := []struct {
		str  string
		want *Packet
	}{
		{
			str: "D2FE28",
			want: &Packet{
				Hdr: &Header{
					Version: 6,
					TypeID:  4,
				},
				Value: 2021,
			},
		},
		{
			str: "38006F45291200",
			want: &Packet{
				Hdr: &Header{
					Version: 1,
					TypeID:  6,
				},
				Value: 1,
				SubPackets: []*Packet{
					{
						Hdr: &Header{
							Version: 6,
							TypeID:  4,
						},
						Value: 10,
					},
					{
						Hdr: &Header{
							Version: 2,
							TypeID:  4,
						},
						Value: 20,
					},
				},
			},
		},
		{
			str: "EE00D40C823060",
			want: &Packet{
				Hdr: &Header{
					Version: 7,
					TypeID:  3,
				},
				Value: 3,
				SubPackets: []*Packet{
					{
						Hdr: &Header{
							Version: 2,
							TypeID:  4,
						},
						Value: 1,
					},
					{
						Hdr: &Header{
							Version: 4,
							TypeID:  4,
						},
						Value: 2,
					},
					{
						Hdr: &Header{
							Version: 1,
							TypeID:  4,
						},
						Value: 3,
					},
				},
			},
		},
		{
			str: "C200B40A82",
			want: &Packet{
				Hdr: &Header{
					Version: 6,
					TypeID:  0,
				},
				Value: 3,
				SubPackets: []*Packet{
					{
						Hdr: &Header{
							Version: 6,
							TypeID:  4,
						},
						Value: 1,
					},
					{
						Hdr: &Header{
							Version: 2,
							TypeID:  4,
						},
						Value: 2,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		pr := PacketReaderFromStr(tc.str)
		got, err := pr.NextPacket()
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(got, tc.want); diff != "" {
			t.Errorf("unexpected value, diff:%v", diff)
		}
	}
}
