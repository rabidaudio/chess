package main

import (
	"fmt"

	"github.com/rabidaudio/chess/board"
)

func main() {
	fmt.Print(board.Initial().StringFromPerspective(board.Black))
}
