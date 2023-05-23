package board

import "github.com/jeselnik/grabble/internal/tile"

const size = 15

type Board [size][size]*tile.Tile

func (b *Board) IsEmpty() bool {
	empty := true
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if b[row][col] != nil {
				empty = false
				break
			}
		}
	}

	return empty
}
