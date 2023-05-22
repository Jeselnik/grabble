package player

import "github.com/jeselnik/grabble/internal/linkedlist"

type Player struct {
	Hand               linkedlist.LinkedList
	name               string
	score, turnsPassed uint
}

func New(name string) Player {
	p := Player{}
	p.name = name
	return p
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) PassTurn() {
	p.turnsPassed++
}

func (p *Player) GetPasses() uint {
	return p.turnsPassed
}

func (p *Player) GetScore() uint {
	return p.score
}

func (p *Player) AddToScore(a uint) {
	p.score += a
}

func (p *Player) SetScore(s uint) {
	p.score = s
}
