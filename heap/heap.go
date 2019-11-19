package main

import "github.com/etsangsplk/experiments/util"

type ComparableItem util.ComparableItem

type Heap interface {
	//heap.Interface
	Push(x ComparableItem, priority int)
	Pop() ComparableItem
	Swap(x, y *ComparableItem)
	Heapify()
}

type H1 struct {
	Heap
	Items []ComparableItem
}

func NewH1() *H1 {
	return &H1{
		Items: []ComparableItem{},
	}
}

func (h *H1) BuildHeap(x []ComparableItem) {

    n := len(x)
    for (i:= n/1; i >= ) {

    }
}

func (h *H1) MergeHeap(x *H1) {

}

func (h *H1) AddItems(x []ComparableItem) {

}

func (h *H1) Push(x ComparableItem, priority int) {

}

func (h *H1) Pop() ComparableItem {

}

func (h *H1) Swap(x, y *ComparableItem) {

}

func (h *H1) Heapify() {

}
