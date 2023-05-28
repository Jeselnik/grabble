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
		return ErrInvalidPlacement{outofBoundsMsg, x, y}
	}

	if !b.isEmpty() {
		xAdjs := [2]int{x - 1, x + 1}
		yAdjs := [2]int{y - 1, y + 1}

		invalid := true

		for _, xAdj := range xAdjs {
			if xAdj > -1 && xAdj < size {
				if b[y][xAdj] != nil {
					invalid = false
				}
			}
		}

		for _, yAdj := range yAdjs {
			if yAdj > -1 && yAdj < size {
				if b[yAdj][x] != nil {
					invalid = false
				}
			}
		}

		if invalid {
			return ErrInvalidPlacement{notAdjacentMsg, x, y}
		}

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
