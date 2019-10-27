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
