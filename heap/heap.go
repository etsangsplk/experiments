package main

import "github.com/etsangsplk/experiments/util"

type Item util.Item

type Heap interface {
	//heap.Interface
	Push(x Item, priority int)
	Pop() Item
	Swap(x, y *Item)
	Heapify()
}

type H1 struct {
	Heap
	Items []Item
}

func NewH1() *H1 {
	return &H1{
		Items: []Item{},
	}
}

func (h *H1) Push(x Item, priority int) {

}

func (h *H1) Pop() Item {

}

func (h *H1) Swap(x, y *Item) {

}

func (h *H1) Heapify() {

}
