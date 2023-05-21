package linkedlist

import (
	"testing"

	"github.com/jeselnik/grabble/internal/tile"
)

func TestLen(t *testing.T) {
	l := LinkedList{}
	t.Run("Empty List", func(t *testing.T) {
		expected := 0
		got := l.Len()
		if expected != got {
			t.Errorf("Got Len %d expected %d", got, expected)
		}
	})

	for i := tile.Value(0); i < 3; i++ {
		l.Add(tile.New('A', i))
	}

	t.Run("Iterate", func(t *testing.T) {
		expected := 3
		got := l.Len()
		if expected != got {
			t.Errorf("Got Len %d expected %d", got, expected)
		}
	})
}

func TestAdd(t *testing.T) {
	l := LinkedList{}
	firstAdded := tile.New('A', 0)

	t.Run("First Item", func(t *testing.T) {
		l.Add(firstAdded)
		expectedLen := 1
		gotLen := l.Len()
		expectedHead := firstAdded.GetLetter()
		gotHead := l.head.tile.GetLetter()

		if !(expectedLen == gotLen && expectedHead == gotHead) {
			t.Errorf("Expected item not at head")
		}
	})
}
