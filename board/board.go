package board

type Color uint8

const (
	White Color = 0x00
	Black Color = 0x08
)

type Piece uint8

const (
	None Piece = 0
	Pawn Piece = iota
	Bishop
	Knight
	Rook
	Queen
	King
)

var pieces = []rune{
	' ', '♙', '♗', '♘', '♖', '♕', '♔',
	' ', '♟', '♞', '♝', '♜', '♛', '♚',
}

func (p Piece) Rune(c Color) rune {
	i := ((uint8(c) >> 3) * 7) + uint8(p)
	return pieces[i]
}

type Board [32]byte

var initial = Board([32]byte{
	0x43, 0x25, 0x62, 0x34, 0x11, 0x11, 0x11, 0x11,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x99, 0x99, 0x99, 0x99, 0xcb, 0xad, 0xea, 0xbc,
})

func Initial() Board {
	return initial
}

type Position struct {
	Piece
	Color
	Rank int
	File int
}

func New(pos ...Position) Board {
	b := Board{}
	for _, p := range pos {
		b.place(p)
	}
	return b
}

func index(rank, file int) (i int, high bool) {
	// a1, b1, c1 .... f8, g8, h8
	i = rank*8 + file
	high = i%2 == 0
	i /= 2
	return
}

func (b *Board) remove(rank, file int) {
	i, h := index(rank, file)
	mask := byte(0xF0)
	if h {
		mask = byte(0x0F)
	}
	b[i] = b[i] & mask
}

func (p Position) value() uint8 {
	return uint8(p.Piece) | uint8(p.Color)
}

func (b *Board) place(p Position) {
	b.remove(p.Rank, p.File)
	if p.Piece == None {
		return
	}
	i, h := index(p.Rank, p.File)
	v := p.value()
	if h {
		v = v << 4
	}
	b[i] = b[i] | v
}

func (b Board) At(rank, file int) (Piece, Color) {
	i, h := index(rank, file)
	v := b[i]
	if h {
		v = (v & 0xF0) >> 4
	} else {
		v = v & 0x0F
	}
	return Piece(v & 0x07), Color(v & 0x08)
}

func (b Board) StringFromPerspective(color Color) string {
	// white: a-h,8-1
	// black: h-a,1-8
	data := make([]rune, 0, 8*9)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			r := 7 - i
			f := j
			if color == Black {
				r = i
				f = 7 - j
			}
			p, c := b.At(r, f)
			data = append(data, p.Rune(c))
		}
		if i < 7 {
			data = append(data, '\n')
		}
	}
	return string(data)
}

func (b Board) String() string {
	return b.StringFromPerspective(White)
}
