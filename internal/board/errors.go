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

type ErrInvalidPlacement struct {
	msg  string
	x, y int
}

func (e ErrInvalidPlacement) Error() string {
	return fmt.Sprintf(e.msg, e.x, e.y, size-1)
}
