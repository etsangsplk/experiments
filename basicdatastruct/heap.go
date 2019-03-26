package main

import (
	"fmt"

	"../basicsort/helper"
)

type Heap struct {
	Data  []int
	order int
}

func MakeHeap(d []int) *Heap {
	return &Heap{
		Data:  make([]int, len(d)),
		order: 1,
	}
}

func (h *Heap) Insert(d int) {

	return
}

func (h *Heap) Max() int {
	return 0
}

func (h *Heap) RemoveMax() int {
	return 0
}

func (h *Heap) ReplaceValueAt(index int, targetValue int) {
	return
}

func (h *Heap) MergeHeap(r *Heap) {
	return
}

func main() {
	for i := 0; i < 10; i++ {
		a := helper.GenerateSlice(i)
		h := MakeHeap(a)
		fmt.Printf("make heap: %v\n", h.Data)
	}
}
