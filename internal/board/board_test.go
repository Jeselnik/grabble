package board

import (
	"errors"
	"testing"

	"github.com/jeselnik/grabble/internal/tile"
)

func TestIsEmpty(t *testing.T) {
	var b Board

	t.Run("Empty", func(t *testing.T) {
		expected := true
		got := b.isEmpty()
		if expected != got {
			t.Errorf("Board should be empty!")
		}
	})

	t.Run("After Added Tile", func(t *testing.T) {
		b[0][5] = tile.New('A', 0)
		expected := false
		got := b.isEmpty()
		if expected != got {
			t.Errorf("IsEmpty should be false!")
		}
	})

}

func TestPlace(t *testing.T) {
	var b Board

	t.Run("Out of Bounds (Negative)", func(t *testing.T) {
		err := b.Place(tile.New('A', 0), -1, 0)

		if !errors.Is(err, ErrNegativeCoord) {
			t.Errorf("Should have gotten ErrNegativeCoord")
		}
	})

	t.Run("Out of Bounds ( > Size)", func(t *testing.T) {
		err := b.Place(tile.New('A', 0), 0, size)

		if !errors.Is(err, ErrInvalidPlacement{outofBoundsMsg, 0, size}) {
			t.Errorf("Should have got ErrInvalidPlacement")
		}
	})

	t.Run("Arbitrary placement on empty board", func(t *testing.T) {
		err := b.Place(tile.New('A', 0), 7, 7)

		if err != nil {
			t.Errorf("Board is empty, tile should have been placed")
		}
	})

	t.Run("Not adjacent placement", func(t *testing.T) {
		err := b.Place(tile.New('A', 0), 0, 0)

		if !errors.Is(err, ErrInvalidPlacement{notAdjacentMsg, 0, 0}) {
			t.Errorf("Should have gotten ErrInvalidPlacement")
		}
	})

	t.Run("Adjacent placement", func(t *testing.T) {
		err := b.Place(tile.New('A', 0), 8, 7)

		if err != nil {
			t.Errorf("err should have been nil")
		}
	})
}

func TestTranslateCoordinates(t *testing.T) {
	x := 'A'
	y := 1

	expX, expY := 0, 0

	gotX, gotY := TranslateCoordinates(x, y)

	if expX != gotX || expY != gotY {
		t.Errorf("x: exp %d got %d\ny: exp %d got %d",
			expX, gotX, expY, gotY)
	}

	t.Run("Lowercase", func(t *testing.T) {
		lowerX := 'b'
		y := 12

		expX, expY := 1, 11
		gotX, gotY := TranslateCoordinates(lowerX, y)

		if expX != gotX || expY != gotY {
			t.Errorf("x: exp %d got %d\ny: exp %d got %d",
				expX, gotX, expY, gotY)
		}
	})
}
