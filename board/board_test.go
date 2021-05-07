package board_test

import (
	"encoding/hex"
	"testing"

	"github.com/rabidaudio/chess/board"
	"gotest.tools/v3/assert"
)

func TestRune(t *testing.T) {
	cases := []struct {
		color    board.Color
		expected string
	}{
		{color: board.White, expected: " ♙♗♘♖♕♔"},
		{color: board.Black, expected: " ♟♞♝♜♛♚"},
	}
	pieces := []board.Piece{board.None, board.Pawn, board.Bishop, board.Knight, board.Rook, board.Queen, board.King}
	for _, test := range cases {
		for i, p := range pieces {
			assert.Equal(t, []rune(test.expected)[i], p.Rune(test.color))
		}
	}
}

func TestNew(t *testing.T) {
	b := board.New(
		board.Position{Piece: board.Rook, Color: board.White, Rank: 0, File: 0},
		board.Position{Piece: board.Knight, Color: board.White, Rank: 0, File: 1},
		board.Position{Piece: board.Bishop, Color: board.White, Rank: 0, File: 2},
		board.Position{Piece: board.Queen, Color: board.White, Rank: 0, File: 3},
		board.Position{Piece: board.King, Color: board.White, Rank: 0, File: 4},
		board.Position{Piece: board.Bishop, Color: board.White, Rank: 0, File: 5},
		board.Position{Piece: board.Knight, Color: board.White, Rank: 0, File: 6},
		board.Position{Piece: board.Rook, Color: board.White, Rank: 0, File: 7},

		board.Position{Piece: board.Pawn, Color: board.White, Rank: 1, File: 0},
		board.Position{Piece: board.Pawn, Color: board.White, Rank: 1, File: 1},
		board.Position{Piece: board.Pawn, Color: board.White, Rank: 1, File: 2},
		board.Position{Piece: board.Pawn, Color: board.White, Rank: 1, File: 3},
		board.Position{Piece: board.Pawn, Color: board.White, Rank: 1, File: 4},
		board.Position{Piece: board.Pawn, Color: board.White, Rank: 1, File: 5},
		board.Position{Piece: board.Pawn, Color: board.White, Rank: 1, File: 6},
		board.Position{Piece: board.Pawn, Color: board.White, Rank: 1, File: 7},

		board.Position{Piece: board.Rook, Color: board.Black, Rank: 7, File: 0},
		board.Position{Piece: board.Knight, Color: board.Black, Rank: 7, File: 1},
		board.Position{Piece: board.Bishop, Color: board.Black, Rank: 7, File: 2},
		board.Position{Piece: board.Queen, Color: board.Black, Rank: 7, File: 3},
		board.Position{Piece: board.King, Color: board.Black, Rank: 7, File: 4},
		board.Position{Piece: board.Bishop, Color: board.Black, Rank: 7, File: 5},
		board.Position{Piece: board.Knight, Color: board.Black, Rank: 7, File: 6},
		board.Position{Piece: board.Rook, Color: board.Black, Rank: 7, File: 7},

		board.Position{Piece: board.Pawn, Color: board.Black, Rank: 6, File: 0},
		board.Position{Piece: board.Pawn, Color: board.Black, Rank: 6, File: 1},
		board.Position{Piece: board.Pawn, Color: board.Black, Rank: 6, File: 2},
		board.Position{Piece: board.Pawn, Color: board.Black, Rank: 6, File: 3},
		board.Position{Piece: board.Pawn, Color: board.Black, Rank: 6, File: 4},
		board.Position{Piece: board.Pawn, Color: board.Black, Rank: 6, File: 5},
		board.Position{Piece: board.Pawn, Color: board.Black, Rank: 6, File: 6},
		board.Position{Piece: board.Pawn, Color: board.Black, Rank: 6, File: 7},
	)
	assert.Equal(t, b, board.Initial())
	h := hex.EncodeToString(b[:])
	assert.Equal(t, h, "43256234111111110000000000000000000000000000000099999999cbadeabc")
}

func TestString(t *testing.T) {
	assert.Equal(t, board.Initial().StringFromPerspective(board.White), `♜♝♞♛♚♞♝♜
♟♟♟♟♟♟♟♟
        
        
        
        
♙♙♙♙♙♙♙♙
♖♘♗♕♔♗♘♖`)

	assert.Equal(t, board.Initial().StringFromPerspective(board.Black), `♖♘♗♔♕♗♘♖
♙♙♙♙♙♙♙♙
        
        
        
        
♟♟♟♟♟♟♟♟
♜♝♞♚♛♞♝♜`)
}
