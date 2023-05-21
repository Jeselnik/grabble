package tile

type Letter rune
type Value uint

/*
Once created, the properties of a tile are immutable,
therefore its fields are private
*/
type Tile struct {
	letter Letter
	value  Value
}

func New(letter Letter, value Value) *Tile {
	return &Tile{letter, value}
}

func (t *Tile) GetLetter() string {
	return string(t.letter)
}

func (t *Tile) GetValue() Value {
	return t.value
}
