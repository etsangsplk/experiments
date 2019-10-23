package main

import "github.com/etsangsplk/experiments/util"

type Stack interface {
	Pop() util.Item
	Push(util.Item)
	IsFull() bool
	IsEmpty() bool
	Size() int
}

type S struct {
	Stack
	Items    []util.Item
	size     int
	capacity int
}

func New(s int) *S {
	return &S{
		Items:    make([]util.Item, s),
		capacity: 0,
	}
}

func (s *S) Push(i util.Item) {
	if s.IsFull() {
		return
	}
	s.Items[s.size] = i
	s.size++
}

func (s *S) Pop() util.Item {
	var x util.Item
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
	return s.Size() == cap(s.Items)
}
func (s *S) Size() int {
	return s.size
}
