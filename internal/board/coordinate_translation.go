package board

const (
	ascii_A_VAL = 65
	ascii_Z_VAL = 90
	ascii_a_VAL = 97
	ascii_z_VAL = 122
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
