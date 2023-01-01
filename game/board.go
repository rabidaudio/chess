package game

import (
	"fmt"
	"strings"
)

type Board [32]byte

var initial = Board([32]byte{
	0x43, 0x25, 0x62, 0x34, 0x11, 0x11, 0x11, 0x11,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x99, 0x99, 0x99, 0x99, 0xcb, 0xad, 0xea, 0xbc,
})

func InitialBoard() Board {
	return initial
}

type Position struct {
	Rank uint
	File uint
}

func ParsePosition(p string) (pos Position, err error) {
	if len(p) != 2 {
		return pos, fmt.Errorf("position must be 2 characters")
	}
	p = strings.ToLower(p)
	f := p[0]
	r := p[1]
	if f < 'a' || f > 'h' {
		return pos, fmt.Errorf("file must be a-f")
	}
	if r < '1' || r > '8' {
		return pos, fmt.Errorf("rank must be 1-8")
	}
	return Position{Rank: uint(r - '1'), File: uint(f - 'a')}, nil
}

func New(layout map[Position]struct {
	Piece
	Color
}) Board {
	b := Board{}
	for p, s := range layout {
		b.place(s.Piece, s.Color, p)
	}
	return b
}

func (p Position) index() (i uint, high bool) {
	// a1, b1, c1 .... f8, g8, h8
	i = p.Rank*8 + p.File
	high = i%2 == 0
	i /= 2
	return
}

func (b *Board) remove(pos Position) {
	i, h := pos.index()
	mask := byte(0xF0)
	if h {
		mask = byte(0x0F)
	}
	b[i] = b[i] & mask
}

func (b *Board) place(p Piece, c Color, pos Position) {
	b.remove(pos)
	i, h := pos.index()
	v := uint8(p) | uint8(c)
	if h {
		v = v << 4
	}
	b[i] = b[i] | v
}

func (b Board) At(pos Position) (Piece, Color) {
	i, h := pos.index()
	v := b[i]
	if h {
		v = (v & 0xF0) >> 4
	} else {
		v = v & 0x0F
	}
	return Piece(v & 0x07), Color(v & 0x08)
}

func (p Position) perspective(c Color) Position {
	if c == White {
		return p
	}
	return Position{Rank: 7 - p.Rank, File: 7 - p.File}
}

func (b Board) StringFromPerspective(color Color) string {
	// white: a-h,8-1
	// black: h-a,1-8
	data := make([]rune, 0, 8*8+7)
	for r := uint(0); r < 8; r++ {
		for f := uint(0); f < 8; f++ {
			pos := Position{7 - r, f}.perspective(color)
			p, c := b.At(pos)
			data = append(data, p.Rune(c))
		}
		if r < 7 {
			data = append(data, '\n')
		}
	}
	return string(data)
}

func (b Board) String() string {
	return b.StringFromPerspective(White)
}
