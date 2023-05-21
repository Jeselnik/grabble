package queue

import (
	"testing"

	"github.com/jeselnik/grabble/internal/tile"
)

func TestPush(t *testing.T) {
	q := Queue{}
	t.Run("Empty", func(t *testing.T) {
		expected := 0
		got := q.Len()
		if expected != got {
			t.Errorf("Expected len %d got %d", expected, got)
		}
	})

	for i := tile.Value(0); i < 2; i++ {
		q.Push(tile.New('A', i))
	}

	t.Run("Appended", func(t *testing.T) {
		expected := 2
		got := q.Len()
		if expected != got {
			t.Errorf("Expected len %d got %d", expected, got)
		}
	})

}

func TestFIFO(t *testing.T) {
	q := Queue{}

	for i := tile.Value(0); i < 3; i++ {
		q.Push(tile.New('A', i))
	}

	expectedHeadIndex := tile.Value(2)
	expectedTailIndex := tile.Value(0)

	gotTailIndex := q.Peek().GetValue()
	gotHeadIndex := q.head.tile.GetValue()

	if !(expectedHeadIndex == gotHeadIndex &&
		expectedTailIndex == gotTailIndex) {

		t.Errorf("HEAD: %d, expected: %d\n TAIL: %d, expected:%d", gotHeadIndex,
			expectedHeadIndex, gotTailIndex, expectedTailIndex)
	}
}

func TestPop(t *testing.T) {
	q := Queue{}

	intendedSize := 5

	for i := tile.Value(0); i < tile.Value(intendedSize); i++ {
		q.Push(tile.New('A', i))
	}

	pop := q.Pop()

	expectedPoppedIndex := tile.Value(0)
	gotPoppedIndex := pop.GetValue()

	expectedLen := intendedSize - 1
	gotLen := q.Len()

	if !(expectedLen == gotLen && expectedPoppedIndex == gotPoppedIndex) {
		t.Errorf("LEN: %d, expected %d\nIND: %d, expected %d",
			gotLen, expectedLen, gotPoppedIndex, expectedPoppedIndex)
	}
}
