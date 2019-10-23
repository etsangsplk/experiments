package main

import "github.com/etsangsplk/experiments/util"

type Queue interface {
	Enqueue(item util.Item)
	Dequeue() util.Item
	isFull() bool
	isEmpty() bool
	Size() int
}

type Q struct {
	Queue
	items       []util.Item
	currentSize int
	capacity    int
}

func New(size int) *Q {
	return &Q{
		items:    make([]util.Item, size),
		capacity: size,
	}
}

func (q *Q) isFull() bool {
	return q.currentSize == len(q.items)
}

func (q *Q) isEmpty() bool {
	return 0 == len(q.items)
}

func (q *Q) Enqueue(i util.Item) {
	q.items = append(q.items, i)
	q.currentSize++
}

func (q *Q) Dequeue() util.Item {
	x := q.items[0]
	q.items = q.items[1:]
	return x
}
