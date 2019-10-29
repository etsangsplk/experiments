package main

import (
	"github.com/etsangsplk/experiment/util"
)

type Item util.Item

type PQ interface {
	Peek() Item
	Pop() Item
	Push(x Item, p int)
}

type PriorityQ struct {
	PQ
	Items []Item
}

func NewPQ() *PQ {
	return &PriorityQ{
		Items: []Item{},
	}
}

func (q *PriorityQ) Peek() Item {

}

func (q *PriorityQ) Push(x Item, p int) {

}

func (q *PriorityQ) Pop() Item {

}
