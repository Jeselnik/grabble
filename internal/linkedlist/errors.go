package linkedlist

import (
	"fmt"

	"github.com/jeselnik/grabble/internal/tile"
)

type NotPresentError struct {
	msg string
}

func (e NotPresentError) Error() string {
	return e.msg
}

func newNotPresentError(item tile.Letter) NotPresentError {
	return NotPresentError{fmt.Sprintf("Item '%s' not present \n",
		string(item))}
}

type OutOfBoundsError struct {
	msg string
}

func (e OutOfBoundsError) Error() string {
	return e.msg
}

func newOutOfBoundsError(index, len int) OutOfBoundsError {
	return OutOfBoundsError{fmt.Sprintf("Index %d > Len %d", index, len)}
}
