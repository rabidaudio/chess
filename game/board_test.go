package game_test

import (
	"encoding/hex"
	"testing"

	"github.com/rabidaudio/chess/game"
	"gotest.tools/v3/assert"
)

func TestRune(t *testing.T) {
	cases := []struct {
		color    game.Color
		expected string
	}{
		{color: game.White, expected: "♙♗♘♖♕♔"},
		{color: game.Black, expected: "♟♞♝♜♛♚"},
	}
	pieces := []game.Piece{game.Pawn, game.Bishop, game.Knight, game.Rook, game.Queen, game.King}
	for _, test := range cases {
		for i, p := range pieces {
			assert.Equal(t, []rune(test.expected)[i], p.Rune(test.color))
		}
	}
}

func TestNew(t *testing.T) {
	p := func(s string) game.Position {
		p, err := game.ParsePosition(s)
		if err != nil {
			t.Fatal(err)
		}
		return p
	}
	b := game.New(map[game.Position]struct {
		game.Piece
		game.Color
	}{
		p("a1"): {game.Rook, game.White},
		p("b1"): {game.Knight, game.White},
		p("c1"): {game.Bishop, game.White},
		p("d1"): {game.Queen, game.White},
		p("e1"): {game.King, game.White},
		p("f1"): {game.Bishop, game.White},
		p("g1"): {game.Knight, game.White},
		p("h1"): {game.Rook, game.White},

		p("a2"): {game.Pawn, game.White},
		p("b2"): {game.Pawn, game.White},
		p("c2"): {game.Pawn, game.White},
		p("d2"): {game.Pawn, game.White},
		p("e2"): {game.Pawn, game.White},
		p("f2"): {game.Pawn, game.White},
		p("g2"): {game.Pawn, game.White},
		p("h2"): {game.Pawn, game.White},

		p("a8"): {game.Rook, game.Black},
		p("b8"): {game.Knight, game.Black},
		p("c8"): {game.Bishop, game.Black},
		p("d8"): {game.Queen, game.Black},
		p("e8"): {game.King, game.Black},
		p("f8"): {game.Bishop, game.Black},
		p("g8"): {game.Knight, game.Black},
		p("h8"): {game.Rook, game.Black},

		p("a7"): {game.Pawn, game.Black},
		p("b7"): {game.Pawn, game.Black},
		p("c7"): {game.Pawn, game.Black},
		p("d7"): {game.Pawn, game.Black},
		p("e7"): {game.Pawn, game.Black},
		p("f7"): {game.Pawn, game.Black},
		p("g7"): {game.Pawn, game.Black},
		p("h7"): {game.Pawn, game.Black},
	})
	assert.Equal(t, b, game.InitialBoard())
	h := hex.EncodeToString(b[:])
	assert.Equal(t, h, "43256234111111110000000000000000000000000000000099999999cbadeabc")
}

func TestString(t *testing.T) {
	assert.Equal(t, game.InitialBoard().StringFromPerspective(game.White), `♜♝♞♛♚♞♝♜
♟♟♟♟♟♟♟♟
        
        
        
        
♙♙♙♙♙♙♙♙
♖♘♗♕♔♗♘♖`)

	assert.Equal(t, game.InitialBoard().StringFromPerspective(game.Black), `♖♘♗♔♕♗♘♖
♙♙♙♙♙♙♙♙
        
        
        
        
♟♟♟♟♟♟♟♟
♜♝♞♚♛♞♝♜`)
}
