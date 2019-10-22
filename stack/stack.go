package main

import "fmt"

type Item interface {
	fmt.Stringer
}

type Stack interface {
	Pop() Item
	Push(Item)
	IsFull() bool
	IsEmpty() bool
	Size() int
}

type S struct {
	Stack
	Items    []Item
	size     int
	capacity int
}

func New(s int) *S {
	return &S{
		Items:    make([]Item, s),
		capacity: s,
	}
}

func (s *S) Push(i Item) {
	if s.IsFull() {
		return
	}
	s.Items[s.size] = i
	s.size++
}

func (s *S) Pop() Item {
	var x Item
	if s.IsEmpty() {
		return x
	}
	x = s.Items[s.size-1]
	s.Items = s.Items[:(s.size - 1)]
	s.size--
	return x
}

func (s *S) IsEmpty() bool {
	return s.Size() == 0
}

func (s *S) IsFull() bool {
	return s.Size() == s.capacity
}
func (s *S) Size() int {
	return s.size
}
