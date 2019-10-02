package main

type ItemType interface{}

type Queue interface {
	Enqueue(item ItemType)
	Dequeue() ItemType
	isFull() bool
	isEmpty() bool
	Size() int
}

type Q struct {
	Queue
	items       []ItemType
	currentSize int
	capacity    int
}

func New(size int) *Q {
	return &Q{
		items:    make([]ItemType, size),
		capacity: size,
	}
}

func (q *Q) isFull() bool {
	return q.currentSize == len(q.items)
}

func (q *Q) isEmpty() bool {
	return 0 == len(q.items)
}

func (q *Q) Enqueue(i ItemType) {
	q.items = append(q.items, i)
	q.currentSize++
}

func (q *Q) Dequeue() ItemType {
	x := q.items[0]
	q.items = q.items[1:]
	return x
}
