package gameplay

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	SEPARATOR    = " "
	OPTION_INDEX = 0

	/* Place Syntax:
	place {letter} at {x}{y}
	*/
	TILE_INDEX     = 1
	COORD_INDEX    = 3
	X_COORD_SUBSTR = 0
	Y_COORD_SUBSTR = 1
)

func (gs *GameState) Turn(playerIndex int) {
	finished := false
	placed := false
	for !finished {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		var userInput string
		if scanner.Err() == nil {
			userInput = strings.ToLower(scanner.Text())
		}

		s := strings.Split(userInput, SEPARATOR)

		switch s[OPTION_INDEX] {
		case "place":
			fmt.Println("aaa")
			placed = true
		case "pass":
			gs.players[playerIndex].PassTurn()
			finished = true
		case "done":
			if placed {
				finished = true
			}
		default:
			fmt.Println("Invalid Input")
		}
	}
}
