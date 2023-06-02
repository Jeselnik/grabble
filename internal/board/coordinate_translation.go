package board

const (
	ascii_A_VAL = 'A'
	ascii_Z_VAL = 'Z'
	ascii_a_VAL = 'a'
	ascii_z_VAL = 'z'
	diff        = ascii_a_VAL - ascii_A_VAL
)

func TranslateCoordinates(x rune, y int) (xConv, yConv int) {
	yConv = y - 1

	if x >= ascii_a_VAL && x <= ascii_z_VAL {
		x -= diff
	}

	if x > ascii_Z_VAL || x < ascii_A_VAL {
		xConv = -1
	} else {
		xConv = int(x) - ascii_A_VAL
	}

	return
}
