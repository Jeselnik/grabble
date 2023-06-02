package gameplay

import (
	"bufio"
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jeselnik/grabble/internal/board"
	"github.com/jeselnik/grabble/internal/player"
	"github.com/jeselnik/grabble/internal/queue"
	"github.com/jeselnik/grabble/internal/tile"
)

const (
	playerCount     = 2
	tileLetterIndex = 0
	tileValueIndex  = 2
)

type GameState struct {
	players   [playerCount]*player.Player
	tileBag   queue.Queue
	board     board.Board
	highScore uint
}

//go:embed ScrabbleTiles.txt
var tiles []byte

func (gs *GameState) InitDefaultState() {
	for i := 0; i < len(gs.players); i++ {
		fmt.Printf("Enter name for Player %d: ", i+1)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()

		var playerName string
		if err != nil {
			playerName = fmt.Sprintf("Player %d", i+1)
		} else {
			playerName = scanner.Text()
		}

		gs.players[i] = player.New(playerName)
	}

	readTiles := []*tile.Tile{}
	tileLines := strings.Split(string(tiles), "\n")

	for i := 0; i < len(tileLines); i++ {
		if len(tileLines[i]) > 0 {
			letter := tileLines[i][tileLetterIndex]
			a, _ := strconv.Atoi(
				string(tileLines[i][tileValueIndex:len(tileLines[i])]))

			readTiles = append(readTiles,
				tile.New(tile.Letter(letter), tile.Value(a)))

		}
	}

	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(readTiles), func(i, j int) {
		readTiles[i], readTiles[j] = readTiles[j], readTiles[i]
	})

	for i := 0; i < len(readTiles); i++ {
		gs.tileBag.Push(readTiles[i])
	}

	for i := 0; i < len(gs.players); i++ {
		for x := 0; x < 7; x++ {
			t, _ := gs.tileBag.Pop()
			gs.players[i].Hand.Add(t)
		}
	}

}
