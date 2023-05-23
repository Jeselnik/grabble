package gameplay

import (
	"github.com/jeselnik/grabble/internal/board"
	"github.com/jeselnik/grabble/internal/player"
)

const playerCount = 2

type GameState struct {
	players   [playerCount]*player.Player
	board     board.Board
	highScore uint
}
