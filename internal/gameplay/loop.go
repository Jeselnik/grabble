package gameplay

import (
	"fmt"

	"github.com/jeselnik/grabble/internal/board"
)

func GameLoop(gs *GameState) {
	running := true

	for running {
		for p := 0; p < len(gs.players); p++ {

			for i := 0; i < len(gs.players); i++ {
				fmt.Printf("Player %s score is %d\n", gs.players[i].GetName(),
					gs.players[i].GetScore())
			}

			fmt.Printf("Player %s, it's your turn.\n", gs.players[p].GetName())
			board.Print(gs.board)
			gs.Turn(p)
		}

		running = false
	}
}
