package main

import (
	"fmt"

	"github.com/rabidaudio/chess/game"
)

func main() {
	fmt.Print(game.InitialBoard().StringFromPerspective(game.Black))
}
