package game

type Color uint8

const (
	White Color = 0x00
	Black Color = 0x08
)

type Piece uint8

const (
	Pawn Piece = iota + 1
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
