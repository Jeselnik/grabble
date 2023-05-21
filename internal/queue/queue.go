/*
In this implementation the tile bag will have tiles inserted at the front
of a list and retrieved from the end in FIFO manner. The base of this queue
is a doubly linked list with pointers to the beginning and end of the queue,
allowing for insertion and retrieval in constant O(1) time.
*/
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

func (q *Queue) Peek() (*tile.Tile, error) {
	if q.tail == nil {
		return &tile.Tile{}, errEmptyQueue
	}
	return q.tail.tile, nil
}

func (q *Queue) Pop() (*tile.Tile, error) {
	elem, err := q.Peek()
	if err != nil {
		return &tile.Tile{}, err
	}
	if q.head != q.tail {
		newTail := q.tail.prev
		q.tail = newTail
		if q.tail != nil {
			q.tail.next = nil
		}
		/* When head and tail are equal and not pointing to nil, it means
		that there is only a single element present. As we are retrieving
		the final element, we need to empty the queue */
	} else {
		q.head = nil
		q.tail = nil
	}
	return elem, nil
}
