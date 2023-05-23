package board

import "fmt"

type ErrOutOfBounds struct {
	msg string
}

func (e ErrOutOfBounds) Error() string {
	return e.msg
}

func newErrOutOfBounds(row, col int) ErrOutOfBounds {
	return ErrOutOfBounds{
		fmt.Sprintf(`A value in co-ordinates x: %d y: %d is invalid 
		due to size greater than max index of %d.`, col, row, size-1)}
}
