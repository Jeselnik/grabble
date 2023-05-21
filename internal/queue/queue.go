package queue

import "github.com/jeselnik/grabble/internal/tile"

type node struct {
	tile *tile.Tile
	next *node
	prev *node
}

type Queue struct {
	head *node
	tail *node
}

func (q *Queue) Len() int {
	len := 0
	n := q.head
	for n != nil {
		len++
		n = n.next
	}
	return len
}

func (q *Queue) Push(t *tile.Tile) {
	h := q.head
	if h == nil {
		n := &node{t, nil, nil}
		q.head = n
		q.tail = n
	} else {
		n := &node{t, q.head, nil}
		q.head.prev = n
		q.head = n
	}
}

func (q *Queue) Peek() *tile.Tile {
	return q.tail.tile
}

func (q *Queue) Pop() *tile.Tile {
	elem := q.Peek()
	newTail := q.tail.prev
	q.tail = newTail
	q.tail.next = nil
	return elem
}
