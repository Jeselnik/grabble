package board

import (
	"errors"
	"fmt"
)

const (
	outofBoundsMsg = `a value in co-ordinates x: %d y: %d is invalid
	due to size greater than max index of %d`

	notAdjacentMsg = `co-ordinates x: %d y: %d are not adjacent to an
	already placed tile`
)

var ErrNegativeCoord = errors.New(`co-ordinates cannot include a 
negative value`)

type ErrOutOfBounds struct {
	x, y int
}

func (e ErrOutOfBounds) Error() string {
	return fmt.Sprintf(outofBoundsMsg, e.x, e.y, size-1)
}

type ErrNotAdjacent struct {
	x, y int
}

func (e ErrNotAdjacent) Error() string {
	return fmt.Sprintf(notAdjacentMsg, e.x, e.y)
}
