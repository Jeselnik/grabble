package board

import "github.com/jeselnik/grabble/internal/tile"

const size = 15

type Board [size][size]*tile.Tile

func (b *Board) isEmpty() bool {
	empty := true
	row := 0
	for row < size && empty {
		for col := 0; col < size; col++ {
			if b[row][col] != nil {
				empty = false
				break
			}
		}
		row++
	}

	return empty
}

func validPlacement(b *Board, x, y int) error {
	if x < 0 || y < 0 {
		return ErrNegativeCoord
	}

	if x > size-1 || y > size-1 {
		return ErrOutOfBounds{x, y}
	}

	/* todo - If board not empty ensure desired place is adjacent
	to an existing tile */
	if !b.isEmpty() {
	}

	return nil
}

func (b *Board) Place(t *tile.Tile, x, y int) error {
	err := validPlacement(b, x, y)

	if err == nil {
		b[y][x] = t
	}

	return err
}
