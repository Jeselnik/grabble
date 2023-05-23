package board

import (
	"testing"

	"github.com/jeselnik/grabble/internal/tile"
)

func TestIsEmpty(t *testing.T) {
	var b Board

	t.Run("Empty", func(t *testing.T) {
		expected := true
		got := b.IsEmpty()
		if expected != got {
			t.Errorf("Board should be empty!")
		}
	})

	t.Run("After Added Tile", func(t *testing.T) {
		b[0][5] = tile.New('A', 0)
		expected := false
		got := b.IsEmpty()
		if expected != got {
			t.Errorf("IsEmpty should be false!")
		}
	})

}
