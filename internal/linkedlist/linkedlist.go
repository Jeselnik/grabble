package linkedlist

import (
	"fmt"

	"github.com/jeselnik/grabble/internal/tile"
)

type node struct {
	tile *tile.Tile
	next *node
}

type LinkedList struct {
	head *node
}

func (l *LinkedList) Add(t *tile.Tile) {
	if l.head == nil {
		l.head = &node{t, nil}
	} else {
		next := l.head
		l.head = &node{t, next}
	}
}

func (l *LinkedList) IndexOf(t tile.Letter) (int, error) {
	i := -1
	found := false
	n := l.head
	for n != nil {
		i++
		if n.tile.GetLetter() == string(t) {
			found = true
			break
		}
	}

	if !found {
		return i, newNotPresentError(t)
	}

	return i, nil
}

func (l *LinkedList) Remove(index int) error {
	if index > l.Len()-1 {
		return newOutOfBoundsError(index, l.Len())
	}
	var prev *node = nil
	curr := l.head
	var next *node = nil

	for i := 0; i < index; i++ {
		prev = curr
		curr = curr.next
	}

	next = curr.next

	if prev != nil {
		prev.next = next
	} else {
		l.head = next
	}

	return nil
}

func (l *LinkedList) Peek(index int) (*tile.Tile, error) {
	if index > l.Len()-1 {
		return &tile.Tile{}, newOutOfBoundsError(index, l.Len())
	}

	n := l.head
	for i := 0; i < index; i++ {
		n = n.next
	}
	return n.tile, nil
}

func (l *LinkedList) Pop(index int) (*tile.Tile, error) {
	if index > l.Len()-1 {
		return &tile.Tile{}, newOutOfBoundsError(index, l.Len())
	}

	t, err := l.Peek(index)
	if err != nil {
		return t, err
	}

	err = l.Remove(index)
	if err != nil {
		return t, err
	}

	return t, err
}

func (l *LinkedList) Len() int {
	len := 0
	n := l.head
	for n != nil {
		len++
		n = n.next
	}
	return len
}

func (l *LinkedList) Print() string {
	n := l.head
	s := ""
	for n != nil {
		tile := fmt.Sprintf("%s - %d, ", n.tile.GetLetter(), n.tile.GetValue())
		s += tile
		n = n.next
	}
	s += "\n"
	return s
}
