package main

import "github.com/etsangsplk/experiments/util"

type Item util.Item

type Queue2 interface {
	Enqueue(x Item)
	Dequeue()
	Size() int
	IsEmpty() bool
}

type Q2 struct {
	Queue2
	Items []Item
}

func NewQ2() *Q2 {
	return &Q2{
		Items: []Item{},
	}
}

func (q *Q2) Enqueue(x Item) {
	q.Items = append(q.Items, x)
}

func (q *Q2) Dequeue() Item {
	if q.IsEmpty() {
		return nil
	}
	x := q.Items[0]
	q.Items = q.Items[1:]
	return x
}

func (q *Q2) Size() int {
	return len(q.Items)
}

func (q *Q2) IsEmpty() bool {
	return q.Size() == 0
}
