package gameplay

import (
	"fmt"

	"github.com/jeselnik/grabble/internal/board"
)

func GamePlay(gs GameState) {
	running := true

	for running {
		for p := 0; p < len(gs.players); p++ {
			fmt.Printf("Player %s, it's your turn.\n", gs.players[p].GetName())
			board.Print(gs.board)
			fmt.Printf("Your hand is %s\n", gs.players[p].Hand.Print())
		}

		running = false
	}
}
