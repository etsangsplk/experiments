package main

import "github.com/etsangsplk/experiments/util"

type Item util.Item

type Heap interface {
	//heap.Interface
	Push(x Item)
	Pop() Item
	Swap(x, y *Item)
	Heapify()
}

type H1 struct {
	Heap
	Items []Item
}

func (h *H1) Push(x Item) {

}

func (h *H1) Pop() Item {

}

func (h *H1) Swap(x, y *Item) {

}

func (h *H1) Heapify() {

}
