package queue

import (
	"errors"
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

	gotTail, _ := q.Peek()
	gotTailIndex := gotTail.GetValue()
	gotHeadIndex := q.head.tile.GetValue()

	if !(expectedHeadIndex == gotHeadIndex &&
		expectedTailIndex == gotTailIndex) {

		t.Errorf("HEAD: %d, expected: %d\n TAIL: %d, expected:%d", gotHeadIndex,
			expectedHeadIndex, gotTailIndex, expectedTailIndex)
	}
}

func TestPop(t *testing.T) {
	t.Run("Basic Operation", func(t *testing.T) {
		q := Queue{}

		intendedSize := 5

		for i := tile.Value(0); i < tile.Value(intendedSize); i++ {
			q.Push(tile.New('A', i))
		}

		pop, _ := q.Pop()

		expectedPoppedIndex := tile.Value(0)
		gotPoppedIndex := pop.GetValue()

		expectedLen := intendedSize - 1
		gotLen := q.Len()

		if !(expectedLen == gotLen && expectedPoppedIndex == gotPoppedIndex) {
			t.Errorf("LEN: %d, expected %d\nIND: %d, expected %d",
				gotLen, expectedLen, gotPoppedIndex, expectedPoppedIndex)
		}
	})

	t.Run("One Element", func(t *testing.T) {
		q := Queue{}
		q.Push(tile.New('A', 0))
		q.Pop()

		var expHead, expTail *node = nil, nil

		if !(expHead == q.head && expTail == q.tail) {
			t.Log(q.head)
			t.Log(q.tail)
			t.Error("Queue Head & Tail Not Empty")
		}

		t.Run("Error on Empty", func(t *testing.T) {
			_, err := q.Pop()
			if !(errors.Is(err, errEmptyQueue)) {
				t.Errorf("Expected error %s got %s", errEmptyQueue, err)
			}
		})

	})
}
